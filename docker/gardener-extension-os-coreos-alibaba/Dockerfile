FROM alpine:3.8

RUN apk add --update bash curl

COPY gardener-extension-os-coreos-alibaba /gardener-extension-os-coreos-alibaba

WORKDIR /

ENTRYPOINT ["/gardener-extension-os-coreos"]
