#!/bin/bash

DIR=".thirdparty/relayer"
[ ! -d "$DIR" ] && echo "Repositry for relayer does not exist at $DIR. Try running 'make get-relayer'..." && exit 1

cd $DIR
echo "Building relayer binary at branch($(git branch --show-current)) tag($(git describe --tags)) commit($(git rev-parse HEAD))"
make install &> /dev/null
rly version