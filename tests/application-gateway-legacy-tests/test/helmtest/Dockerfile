FROM golang:1.13-alpine as builder

ENV SRC_DIR=/go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests

WORKDIR $SRC_DIR
COPY . $SRC_DIR

RUN CGO_ENABLED=0 go test -o proxyhelmtests.test -c ./test/helmtest/proxy/tests

FROM eu.gcr.io/kyma-project/external/alpine:3.14.2

LABEL source=git@github.com:kyma-project/kyma.git

WORKDIR /app

COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests/scripts/helm-test-entrypoint.sh .
COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests/proxyhelmtests.test .
COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests/licenses ./licenses

ARG TEST_EXECUTOR_IMAGE
ENV TEST_EXECUTOR_IMAGE=$TEST_EXECUTOR_IMAGE

ENTRYPOINT ["./helm-test-entrypoint.sh"]
