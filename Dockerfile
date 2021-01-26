# To build the Desmos image, just run:
# > docker build -t desmos .
#
# Simple usage with a mounted data directory:
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.desmos:/root/.desmos desmos desmos init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.desmos:/root/.desmos desmos desmos start
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -p 46657:46657 -p 46656:46656 -v ~/.desmos:/root/.desmos --name desmos desmos
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it desmos bash
#
# To exit the bash, just execute
# > exit
FROM golang:1.15-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3 ca-certificates wget
RUN apk add --no-cache $PACKAGES

ADD https://github.com/CosmWasm/wasmvm/releases/download/v0.13.0/libwasmvm_muslc.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep 39dc389cc6b556280cbeaebeda2b62cf884993137b83f90d1398ac47d09d3900

RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.28-r0/glibc-2.28-r0.apk
RUN apk add glibc-2.28-r0.apk

# Set working directory for the build
WORKDIR /go/src/github.com/desmos-labs/desmos

# Add source files
COPY . .

# Install Desmos, remove packages
RUN make build-linux

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
