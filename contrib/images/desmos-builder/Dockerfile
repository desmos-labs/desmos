FROM golang:1.19-alpine
ARG arch=x86_64

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3 ca-certificates build-base
RUN set -eux; apk add --no-cache $PACKAGES;

# Set working directory for the build
WORKDIR /code

# Add source files
COPY . /code/

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 86bc5fdc0f01201481c36e17cd3dfed6e9650d22e1c5c8983a5b78c231789ee0

ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep a00700aa19f5bfe0f46290ddf69bf51eb03a6dfcd88b905e1081af2e42dbbafc

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /usr/local/lib/libwasmvm_muslc.a

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN BUILD_TAGS=muslc GOOS=linux GOARCH=amd64 LEDGER_ENABLED=true LINK_STATICALLY=true make build
RUN echo "Ensuring binary is statically linked ..." \
  && (file ./build/desmos | grep "statically linked")