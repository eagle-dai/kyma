FROM eu.gcr.io/kyma-project/external/golang:1.14.8-alpine as builder

ENV SRC_DIR=/go/src/github.com/kyma-project/kyma/tests/compass-runtime-agent

RUN apk add -U --no-cache ca-certificates

WORKDIR $SRC_DIR
COPY . $SRC_DIR

RUN CGO_ENABLED=0 GOOS=linux go test -c ./test/runtimeagent/test

FROM eu.gcr.io/kyma-project/external/alpine:3.12.0

LABEL source=git@github.com:kyma-project/kyma.git
RUN apk add --no-cache curl

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/compass-runtime-agent/scripts/entrypoint.sh .
COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/compass-runtime-agent/test.test .
COPY --from=builder /go/src/github.com/kyma-project/kyma/tests/compass-runtime-agent/licenses ./licenses

ENTRYPOINT ["./entrypoint.sh"]
