package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	scCs "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset"
	catalogInformers "github.com/kubernetes-sigs/service-catalog/pkg/client/informers_generated/externalversions"
	appCli "github.com/kyma-project/kyma/components/application-operator/pkg/client/clientset/versioned"
	appInformer "github.com/kyma-project/kyma/components/application-operator/pkg/client/informers/externalversions"
	"github.com/sirupsen/logrus"
	securityclientv1beta1 "istio.io/client-go/pkg/clientset/versioned/typed/security/v1beta1"
	v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"

	"github.com/kyma-project/kyma/components/application-broker/internal/access"
	"github.com/kyma-project/kyma/components/application-broker/internal/broker"
	"github.com/kyma-project/kyma/components/application-broker/internal/config"
	"github.com/kyma-project/kyma/components/application-broker/internal/mapping"
	"github.com/kyma-project/kyma/components/application-broker/internal/nsbroker"
	"github.com/kyma-project/kyma/components/application-broker/internal/servicecatalog"
	"github.com/kyma-project/kyma/components/application-broker/internal/storage"
	"github.com/kyma-project/kyma/components/application-broker/internal/storage/populator"
	"github.com/kyma-project/kyma/components/application-broker/internal/syncer"
	mappingCli "github.com/kyma-project/kyma/components/application-broker/pkg/client/clientset/versioned"
	mappingInformer "github.com/kyma-project/kyma/components/application-broker/pkg/client/informers/externalversions"
	"github.com/kyma-project/kyma/components/application-broker/platform/logger"
)

// informerResyncPeriod defines how often informer will execute relist action. Setting to zero disable resync.
// BEWARE: too short period time will increase the CPU load.
const (
	informerResyncPeriod = 30 * time.Minute
	Verbose              = "verbose"
)

func main() {
	verbose := flag.Bool(Verbose, false, "specify if log verbosely loading configuration")
	flag.Parse()
	cfg, err := config.Load(*verbose)
	fatalOnError(err)

	log := logger.New(&cfg.Logger)

	// setup graceful shutdown signals
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	stopCh := make(chan struct{})
	cancelOnInterrupt(ctx, stopCh, cancelFunc)

	k8sConfig, err := restclient.InClusterConfig()
	fatalOnError(err)

	appClient, err := appCli.NewForConfig(k8sConfig)
	fatalOnError(err)
	mClient, err := mappingCli.NewForConfig(k8sConfig)
	fatalOnError(err)
	scClientSet, err := scCs.NewForConfig(k8sConfig)
	fatalOnError(err)
	k8sClient, err := kubernetes.NewForConfig(k8sConfig)
	fatalOnError(err)
	istioClient, err := securityclientv1beta1.NewForConfig(k8sConfig)
	fatalOnError(err)

	livenessCheckStatus := broker.LivenessCheckStatus{Succeeded: false}
	srv := SetupServerAndRunControllers(cfg, log, stopCh, k8sClient, scClientSet, appClient, mClient,
		istioClient, &livenessCheckStatus)

	fatalOnError(srv.Run(ctx, fmt.Sprintf(":%d", cfg.Port)))
}

// SetupServerAndRunControllers setups the application - create and start informers, create all services and HTTP server.
func SetupServerAndRunControllers(cfg *config.Config, log *logrus.Entry, stopCh chan struct{},
	k8sClient kubernetes.Interface,
	scClientSet scCs.Interface,
	appClient appCli.Interface,
	mClient mappingCli.Interface,
	istioClient securityclientv1beta1.SecurityV1beta1Interface,
	livenessCheckStatus *broker.LivenessCheckStatus,
) *broker.Server {

	// create storage factory
	storageConfig := storage.ConfigList(cfg.Storage)
	sFact, err := storage.NewFactory(&storageConfig)
	fatalOnError(err)

	// k8s
	nsInformer := v1.NewNamespaceInformer(k8sClient, informerResyncPeriod, cache.Indexers{})

	// ServiceCatalog

	scInformerFactory := catalogInformers.NewSharedInformerFactory(scClientSet, informerResyncPeriod)
	scInformersGroup := scInformerFactory.Servicecatalog().V1beta1()

	// Applications
	appInformerFactory := appInformer.NewSharedInformerFactory(appClient, informerResyncPeriod)
	appInformersGroup := appInformerFactory.Applicationconnector().V1alpha1()

	// Mapping
	mInformerFactory := mappingInformer.NewSharedInformerFactory(mClient, informerResyncPeriod)
	mInformersGroup := mInformerFactory.Applicationconnector().V1alpha1()

	// internal services
	nsBrokerSyncer := syncer.NewServiceBrokerSyncer(scClientSet.ServicecatalogV1beta1())
	relistRequester := syncer.NewRelistRequester(nsBrokerSyncer, cfg.BrokerRelistDurationWindow, log)
	siFacade := servicecatalog.NewFacade(scInformersGroup.ServiceInstances().Informer(), scInformersGroup.ServiceClasses().Informer())

	accessChecker := access.New(sFact.Application(), mClient.ApplicationconnectorV1alpha1(), sFact.Instance(), cfg.APIPackagesSupport)

	appSyncCtrl := syncer.New(appInformersGroup.Applications(), sFact.Application(), sFact.Application(), relistRequester, log, cfg.APIPackagesSupport)

	brokerService, err := broker.NewNsBrokerService()
	fatalOnError(err)

	nsBrokerFacade := nsbroker.NewFacade(scClientSet.ServicecatalogV1beta1(), k8sClient.CoreV1(), nsBrokerSyncer, cfg.Namespace,
		cfg.UniqueSelectorLabelKey, cfg.UniqueSelectorLabelValue, cfg.ServiceName, int32(cfg.Port), log)

	mappingCtrl := mapping.New(mInformersGroup.ApplicationMappings().Informer(), scInformersGroup.ServiceInstances().Informer(),
		nsBrokerFacade, nsBrokerSyncer, siFacade, log, livenessCheckStatus)

	// create ApplicationServiceID selector
	idSelector := broker.NewIDSelector(cfg.APIPackagesSupport)

	// create broker
	srv := broker.New(sFact.Application(), sFact.Instance(), sFact.InstanceOperation(), accessChecker,
		mClient.ApplicationconnectorV1alpha1(),
		mInformersGroup.ApplicationMappings().Lister(), brokerService,
		&mClient, &istioClient, log, livenessCheckStatus,
		cfg.APIPackagesSupport, cfg.Director.Service, cfg.Director.ProxyURL,
		scInformersGroup.ServiceBindings().Informer(), cfg.GatewayBaseURLFormat, idSelector)

	// wait for api server
	err = wait.PollImmediate(time.Second, time.Minute, func() (bool, error) {
		_, err := k8sClient.CoreV1().Namespaces().List(metav1.ListOptions{})
		if err != nil {
			log.Errorf("while waiting for api server: %s", err)
			return false, nil
		}
		return true, nil
	})
	fatalOnError(err)

	// start informers
	scInformerFactory.Start(stopCh)
	appInformerFactory.Start(stopCh)
	mInformerFactory.Start(stopCh)
	go nsInformer.Run(stopCh)

	// wait for cache sync
	scInformerFactory.WaitForCacheSync(stopCh)
	appInformerFactory.WaitForCacheSync(stopCh)
	mInformerFactory.WaitForCacheSync(stopCh)
	cache.WaitForCacheSync(stopCh, nsInformer.HasSynced)

	// start services & ctrl
	go appSyncCtrl.Run(stopCh)
	go mappingCtrl.Run(stopCh)
	go relistRequester.Run(stopCh)

	// instance populator
	instancePopulator := populator.NewInstances(scClientSet, sFact.Instance(), &populator.Converter{}, sFact.InstanceOperation(), srv, idSelector, log)
	err = instancePopulator.Do()
	fatalOnError(err)

	return srv
}

func fatalOnError(err error) {
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

// cancelOnInterrupt closes given channel and also calls cancel func when os.Interrupt or SIGTERM is received
func cancelOnInterrupt(ctx context.Context, ch chan<- struct{}, cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-ctx.Done():
			close(ch)
		case <-c:
			close(ch)
			cancel()
		}
	}()
}
