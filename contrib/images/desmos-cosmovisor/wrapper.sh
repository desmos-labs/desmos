#!/usr/bin/env sh

BINARY=cosmovisor
ID=${ID:-0}
LOG=${LOG:-desmos.log}

export DESMOSDHOME="/desmos/node${ID}/desmos"

# Set environment variables
export DAEMON_NAME=desmos
export DAEMON_HOME="$DESMOSDHOME"
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_RESTART_AFTER_UPGRADE=true

# Setup Cosmovisor
COSMOVISOR_GENESIS="$DESMOSDHOME/cosmovisor/genesis/bin"
if [ ! -d "$COSMOVISOR_GENESIS" ]; then
  mkdir -p $COSMOVISOR_GENESIS
  cp $(which desmos) "$COSMOVISOR_GENESIS/desmos"
fi

for file in $(find $DESMOSDHOME -type f -name "desmos"); do
  ldd $file
  file $file
done

# Run the command
if [ -d "$(dirname "${DESMOSDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${DESMOSDHOME}" "$@" | tee "${DESMOSDHOME}/${LOG}"
else
  "${BINARY}" --home "${DESMOSDHOME}" "$@"
fi