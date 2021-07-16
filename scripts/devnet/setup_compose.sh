DESMOS_VERSION=$1
sed -i "s|image: \".*\"|image: \"desmoslabs/desmos:$DESMOS_VERSION\"|g" scripts/devnet/docker-compose.yml