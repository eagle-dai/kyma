APP_NAME = rafter-test
APP_PATH = tests/rafter
BUILDPACK = eu.gcr.io/kyma-project/test-infra/buildpack-golang-toolbox:v20210601-19326f24
SCRIPTS_DIR = $(realpath $(shell pwd)/../..)/common/makefiles

include $(SCRIPTS_DIR)/generic-make-go.mk

VERIFY_IGNORE := /vendor\|/mocks

release:
	$(MAKE) gomod-release

test-local:
	@echo "Tests are not to be run at this stage."

.PHONY: path-to-referenced-charts
path-to-referenced-charts:
	@echo "resources/rafter"