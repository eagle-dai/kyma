FROM eu.gcr.io/kyma-project/external/golang:1.17.8-alpine3.15 as builder

ENV SRC_DIR=/go/src/github.com/kyma-project/kyma/components/application-broker
COPY . $SRC_DIR

# creates a non-root user to give him write permissions to tmp folder
# needs for logger which saves logs under tmp dir
RUN mkdir /user && \
    echo 'appuser:x:2000:2000:appuser:/:' > /user/passwd && \
    echo 'appuser:x:2000:' > /user/group
RUN mkdir -p tmp

WORKDIR $SRC_DIR

RUN CGO_ENABLED=0 go build -o /application-broker ./cmd/broker

FROM eu.gcr.io/kyma-project/external/alpine:3.12.4 as certs
RUN apk --no-cache add ca-certificates

FROM scratch

COPY --from=builder /user/group /user/passwd /etc/

USER appuser:appuser

# appuser must be an owner of the tmp dir to write there
COPY --from=builder --chown=appuser /tmp /tmp

COPY --from=builder ./application-broker /root/application-broker
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./licenses/ /root/licenses

LABEL source=git@github.com:kyma-project/kyma.git

ENTRYPOINT ["/root/application-broker"]
