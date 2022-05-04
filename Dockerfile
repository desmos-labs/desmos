# To build the Desmos image, just run:
# > docker build -t desmos .
#
# Simple usage with a mounted data directory:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos desmos desmos init
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos desmos desmos start
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -p 26657:26657 -p 26656:26656 -v ~/.desmos:/root/.desmos --name desmos desmos
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it desmos bash
#
# To exit the bash, just execute
# > exit
FROM golang:1.17.3-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3 ca-certificates build-base
RUN set -eux; apk add --no-cache $PACKAGES;

# Set working directory for the build
WORKDIR /go/src/github.com/desmos-labs/desmos

# Add source files
COPY . .

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta10/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep 2f44efa9c6c1cda138bd1f46d8d53c5ebfe1f4a53cf3457b01db86472c4917ac

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc GOOS=linux GOARCH=amd64 LEDGER_ENABLED=true LINK_STATICALLY=true make build

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Install bash
RUN apk add --no-cache bash

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/build/desmos /usr/bin/desmos

EXPOSE 26656 26657 1317 9090

# Run desmos by default, omit entrypoint to ease using container with desmos
CMD ["desmos"]
