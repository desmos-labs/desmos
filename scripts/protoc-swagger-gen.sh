#!/usr/bin/env bash
SWAGGER_DIR=./swagger-proto

set -eo pipefail

# prepare swagger generation
mkdir -p "$SWAGGER_DIR/proto"
printf "version: v1\ndirectories:\n  - proto\n  - third_party" > "$SWAGGER_DIR/buf.work.yaml"
cp ./proto/buf.gen.swagger.yaml "$SWAGGER_DIR/buf.gen.swagger.yaml"

# copy existing proto files
cp -r ./proto/desmos "$SWAGGER_DIR/proto"
cp -r ./proto/oracle "$SWAGGER_DIR/proto"
cp -r ./proto/ethermint "$SWAGGER_DIR/proto"

# create temporary folder to store intermediate results from `buf generate`
mkdir -p ./tmp-swagger-gen

# step into swagger folder
cd "$SWAGGER_DIR"

proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cd ..

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen
