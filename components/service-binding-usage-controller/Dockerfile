FROM eu.gcr.io/kyma-project/external/golang:1.17.6-alpine3.15 as builder

ENV SRC_DIR=/go/src/github.com/kyma-project/kyma/components/service-binding-usage-controller
COPY . $SRC_DIR

# creates a non-root user to give him write permissions to tmp folder
# needs for logger which saves logs under tmp dir
RUN mkdir /user && \
    echo 'appuser:x:2000:2000:appuser:/:' > /user/passwd && \
    echo 'appuser:x:2000:' > /user/group
RUN mkdir -p tmp

WORKDIR $SRC_DIR

RUN CGO_ENABLED=0 go build -o /binding-usage-controller ./cmd/controller

# hadolint ignore=DL3007
FROM alpine:latest as certs
RUN apk --no-cache add ca-certificates

FROM scratch

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder ./binding-usage-controller /root/binding-usage-controller

USER appuser:appuser
# appuser must be an owner of the tmp dir to write there
COPY --from=builder --chown=appuser /tmp /tmp

LABEL source=git@github.com:kyma-project/kyma.git

ENTRYPOINT ["/root/binding-usage-controller"]
