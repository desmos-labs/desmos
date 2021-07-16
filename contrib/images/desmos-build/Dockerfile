FROM golang:1.15-alpine

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

ARG UID=1000
ARG GID=1000
USER ${UID}:${GID}

# Fixes build errors. Required since Go 1.12 (https://github.com/golang/go/issues/26280#issuecomment-445294378)
ENV GOPATH /tmp/go
ENV GOCACHE /tmp/.cache

VOLUME [ "/desmos" ]
WORKDIR /desmos

CMD ["sh"]