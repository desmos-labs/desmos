#!/usr/bin/env bash

NODES=$1
GENESIS_VERSION=$2
GENESIS_URL=$3
UPGRADE_NAME=$4

BUILDDIR=$(pwd)/build
CONTRIBFOLDER=$(pwd)/contrib
TESTNETDIR=$CONTRIBFOLDER/upgrade_testnet

# Remove all build files
echo "===> Removing build folder"
rm -r -f $BUILDDIR && mkdir $BUILDDIR

# Create the 4 nodes folders with the correct denom
echo "===> Creating $NODES nodes localnet"
docker run --rm --name desmos-tesnet --user $UID:$GID \
  -v $BUILDDIR:/workerplace/build:Z --workdir /workerplace \
  desmoslabs/desmos:$GENESIS_VERSION \
    desmos testnet \
      --home ./build \
      -o ./build \
      --starting-ip-address 192.168.255.2 \
      --keyring-backend=test \
	    --v=$NODES \
	    --gentx-coin-denom="udaric" \
	    --minimum-gas-prices="0.000001udaric"

# Run the Python script to setup the genesis
echo "===> Setting up the genesis file"
docker run --rm --name desmos-python --user $UID:$GID \
  -v $TESTNETDIR:/usr/src/app -v $BUILDDIR:/desmos:Z \
  desmoslabs/desmos-python \
    python setup_genesis.py /desmos $NODES $GENESIS_URL > /dev/null

# Build the new Desmos-Cosmovisor image
echo "===> Building the new Desmos-Cosmovisor image"
make -C $CONTRIBFOLDER/images desmos-cosmovisor DESMOS_VERSION=$GENESIS_VERSION > /dev/null

# Set the correct Desmos image version inside the docker compose file
echo "===> Setting up the Docker compose file"
sed -i "s|image: \".*\"|image: \"desmoslabs/desmos-cosmovisor:$GENESIS_VERSION\"|g" $TESTNETDIR/docker-compose.yml

# Build the current code using Alpine to make sure it's later compatible with the devnet
echo "===> Building Desmos"
make build-alpine > /dev/null

# Copy the Desmos binary into the proper folders
UPGRADE_FOLDER="$BUILDDIR/node0/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin"
if [ ! -d "$UPGRADE_FOLDER" ]; then
  echo "===> Setting up upgrade binary"

  for ((i = 0; i < $NODES; i++)); do
    echo "====> Node $i"
    mkdir -p "$BUILDDIR/node$i/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin"
    cp "$BUILDDIR/desmos" "$BUILDDIR/node$i/desmos/cosmovisor/upgrades/$UPGRADE_NAME/bin/desmos"
  done
fi

# Start the devnet
echo "===> Starting the devnet"
docker-compose -f $TESTNETDIR/docker-compose.yml up -d
