FROM golang:1.17-alpine

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

ARG UID=1000
ARG GID=1000
USER ${UID}:${GID}

# Fixes build errors. Required since Go 1.12 (https://github.com/golang/go/issues/26280#issuecomment-445294378)
ENV GOPATH /tmp/go
ENV GOCACHE /tmp/.cache

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta5/libwasmvm_muslc.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep d16a2cab22c75dbe8af32265b9346c6266070bdcf9ed5aa9b7b39a7e32e25fe0

VOLUME [ "/desmos" ]
WORKDIR /desmos

CMD ["sh"]