FROM alpine:3.8

RUN apk add --update bash curl

COPY gardener-extension-os-coreos /gardener-extension-os-coreos

WORKDIR /

ENTRYPOINT ["/gardener-extension-os-coreos"]
