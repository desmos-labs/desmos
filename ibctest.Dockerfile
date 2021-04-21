FROM golang:1.15-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/desmos-labs/desmos

# Add source files
COPY . .

# Install Desmos, remove packages
RUN make build-linux

# Install relayer
RUN make get-relayer
RUN make build-relayer

RUN make setup-ibctest

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Install bash
RUN apk add --no-cache bash

# Copy over binaries and chain configs from the build-env
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/build/desmos /usr/bin/desmos
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/.thirdparty/build/rly /usr/bin/rly
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/build/ibc /ibc