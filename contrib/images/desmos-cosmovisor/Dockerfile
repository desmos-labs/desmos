ARG DESMOS_VERSION="v0.17.2"
FROM golang:1.15-alpine AS build-env

ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

ENV GOBIN /go/bin

# Set working directory for the build
WORKDIR /go/src/github.com/cosmos/cosmos-sdk
RUN git clone https://github.com/cosmos/cosmos-sdk.git /go/src/github.com/cosmos/cosmos-sdk
RUN cd /go/src/github.com/cosmos/cosmos-sdk
RUN make cosmovisor

# Desmos image
FROM desmoslabs/desmos:$DESMOS_VERSION as desmos

# Final image
FROM ubuntu:18.04

# Copy over binaries from the build-env
COPY --from=desmos /usr/bin/desmos /usr/bin/desmos
COPY --from=build-env /go/src/github.com/cosmos/cosmos-sdk/cosmovisor/cosmovisor /usr/bin/cosmovisor

ARG UID=1000
ARG GID=1000
USER ${UID}:${GID}

EXPOSE 26656 26657 1317 9090

COPY wrapper.sh /usr/bin/wrapper.sh
ENTRYPOINT ["bash", "/usr/bin/wrapper.sh"]

# Run cosmovisor by default, omit entrypoint to ease using container with Desmos
CMD ["cosmovisor"]
STOPSIGNAL SIGTERM