# this script is for generating protobuf files for the new google.golang.org/protobuf API

set -eo pipefail

protoc_install_gopulsar() {
  go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
}

protoc_install_gopulsar

echo "Cleaning API directory"
(cd api; find ./ -type f \( -iname \*.pulsar.go -o -iname \*.pb.go -o -iname \*.cosmos_orm.go -o -iname \*.pb.gw.go \) -delete; find . -empty -type d -delete; cd ..)

echo "Generating API"
cd proto

module_dirs=$(find ./desmos -path -prune -o -name 'module.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $module_dirs; do
  buf generate --template buf.gen.pulsar.yaml --path $dir
done

chmod 755 ../api -R
