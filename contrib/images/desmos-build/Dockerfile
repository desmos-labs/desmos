FROM ubuntu:18.04

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc python3 wget
RUN apt-get update
RUN apt-get install --yes $PACKAGES

RUN wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz
RUN tar -xvf go1.16.4.linux-amd64.tar.gz
RUN mv go /usr/local
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:/usr/local/go/bin

ARG UID=1000
ARG GID=1000
USER ${UID}:${GID}

# Fixes build errors. Required since Go 1.12 (https://github.com/golang/go/issues/26280#issuecomment-445294378)
ENV GOPATH /tmp/go
ENV GOCACHE /tmp/.cache

VOLUME [ "/desmos" ]
WORKDIR /desmos

CMD ["sh"]