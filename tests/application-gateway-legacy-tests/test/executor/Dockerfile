FROM golang:1.15-alpine as builder

ENV SRC_DIR=/go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests

WORKDIR $SRC_DIR
COPY . $SRC_DIR

RUN CGO_ENABLED=0 go test -o proxytestsexecutor.test -c ./test/executor/proxy/tests

FROM alpine:3.13

LABEL source=git@github.com:kyma-project/kyma.git

WORKDIR /app

COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests/spec.json .
COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests/scripts/executor-entrypoint.sh .
COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests/proxytestsexecutor.test .
COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/application-gateway-legacy-tests/licenses ./licenses

ENTRYPOINT ["./executor-entrypoint.sh"]
