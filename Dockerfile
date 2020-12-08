# To build the Desmos image, just run:
# > docker build -t desmos .
#
# Simple usage with a mounted data directory:
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.desmosd:/root/.desmosd desmos desmosd init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.desmosd:/root/.desmosd desmos desmosd start
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -p 46657:46657 -p 46656:46656 -v ~/.desmosd:/root/.desmosd --name desmos desmos
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it desmos bash
#
# To exit the bash, just execute
# > exit
FROM golang:alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

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
COPY --from=build-env /go/src/github.com/desmos-labs/desmos/build/desmosd /usr/bin/desmosd

EXPOSE 26656 26657 1317 9090

# Run desmosd by default, omit entrypoint to ease using container with desmoscli
CMD ["desmosd"]
