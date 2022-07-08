FROM ubuntu:20.04

# Copy over binaries from the build-env
COPY --from=desmoslabs/builder:latest /code/build/desmos /usr/bin/desmos

EXPOSE 26656 26657 1317 9090

# Run desmos by default, omit entrypoint to ease using container with desmos
CMD ["desmos"]