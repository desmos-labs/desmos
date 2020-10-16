#!/bin/bash
shopt -s expand_aliases

####################################
## Variables
####################################
MONIKER=$1
if [ -z "$MONIKER" ]; then
  echo "Validator moniker not given. Please specify it as the first argument"
  exit
fi

USER=$(id -u -n)


####################################
## Setup environmental variables
####################################
echo "===> Setting up environmental variables"

if [ -z "$GOPATH" ]; then
  echo "GOPATH environmental variable not set"
  exit
fi

if [ -z "$GOBIN" ]; then
  echo "export GOBIN=$GOPATH/bin" >> ~/.profile
  source ~/.profile
fi

if [ -z "$DAEMON_NAME" ]; then
  {
    echo " "
    echo "# Setup Cosmovisor"
    echo "export DAEMON_NAME=desmosd"
    echo "export DAEMON_HOME=$HOME/.desmosd"
    echo "export DAEMON_RESTART_AFTER_UPGRADE=on"
  } >> ~/.profile
  source ~/.profile
fi

echo "===> Completed environmental variables setup"
echo ""

####################################
## Setup Cosmovisor
####################################

COSMOVISOR_FILE="$GOBIN/cosmovisor"
if [ ! -f "$COSMOVISOR_FILE" ]; then
  echo "===> Installing Cosmovisor"

  {
    git clone https://github.com/cosmos/cosmos-sdk.git ~/cosmos
    cd ~/cosmos/cosmovisor
    make cosmovisor
    mkdir -p "$GOBIN" && cp cosmovisor --target-directory="$GOBIN"
    cd ~
  } &> /dev/null

  echo "===> Cosmovisor installed"
  echo ""
fi


####################################
## Setup Desmos
####################################
echo "===> Setting up Desmos"

# Backup the priv validator key
VALIDATOR_PRIV_KEY="$HOME/.desmosd/config/priv_validator_key.json"
BACKUP_FILE="$HOME/priv_validator_key.json"
if [ -f "$VALIDATOR_PRIV_KEY" ]; then
  echo "====> Backing up the private validator key"
  cp "$VALIDATOR_PRIV_KEY" "$BACKUP_FILE"
fi

# Delete the old ~/.desmosd folder
DESMOSD_FOLDER="$HOME/.desmosd"
if [ -d "$DESMOSD_FOLDER" ]; then
  echo "====> Removing existing desmosd folder"
  sudo rm -r ~/.desmosd
fi

# Delete the old ~/.desmoscli folder
DESMOSCLI_FOLDER="$HOME/.desmoscli"
if [ -d "$DESMOSCLI_FOLDER" ]; then
  echo "====> Removing existing desmoscli folder"
  sudo rm -r ~/.desmoscli
fi

# Clone Desmos
echo "====> Downloading Desmos"
{
  DESMOS_FOLDER=~/desmos
  if [ ! -d "$DESMOS_FOLDER" ]; then
    git clone https://github.com/desmos-labs/desmos.git ~/desmos
  fi

  cd ~/desmos || exit
  git fetch -a
  git checkout tags/v0.12.2
  make build

  mkdir -p ~/.desmosd/cosmovisor/genesis/bin
  mkdir -p ~/.desmosd/cosmovisor/upgrades
  mv build/desmos* ~/.desmosd/cosmovisor/genesis/bin

  alias desmosd=~/.desmosd/cosmovisor/current/bin/desmosd
  alias desmoscli=~/.desmosd/cosmovisor/current/bin/desmoscli
} &> /dev/null

# Initialize the chain
echo "====> Initializing a new chain"
{
  cosmovisor unsafe-reset-all
  cosmovisor init "$MONIKER"
} &> /dev/null

# Restore the priv validator key
if [ -f "$BACKUP_FILE" ]; then
  echo "====> Restoring private validator key"
  cp "$BACKUP_FILE" "$VALIDATOR_PRIV_KEY"
  rm "$BACKUP_FILE"
fi

# Download the genesis file
echo "====> Downloading the genesis file"
{
  curl https://raw.githubusercontent.com/desmos-labs/morpheus/master/genesis.json -o $HOME/.desmosd/config/genesis.json
} &> /dev/null

# Setup the persistent peers
echo "====> Setting persistent peers"
{
  sed -i -e 's/persistent_peers = ""/persistent_peers = "7fed5624ca577eb0333d3631b5e4f16ba1736979@54.180.98.75:26656"/g' ~/.desmosd/config/config.toml
} &> /dev/null

echo "===> Completed Desmos setup"
echo ""

####################################
## Setup the service
####################################
echo "===> Setting up Desmos service"
FILE=/etc/systemd/system/desmosd.service
{
  sudo tee $FILE > /dev/null <<EOF
[Unit]
Description=Desmos full node watched by Cosmovisor
After=network-online.target
[Service]
User=$USER
ExecStart=$GOBIN/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=4096
Environment="DAEMON_NAME=desmosd"
Environment="DAEMON_HOME=$HOME/.desmosd"
Environment="DAEMON_RESTART_AFTER_UPGRADE=on"
[Install]
WantedBy=multi-user.target
EOF
} &> /dev/null

echo "====> Starting Desmos service"
echo ""
{
  sudo systemctl daemon-reload
  sudo systemctl enable desmosd
  sudo systemctl restart desmosd
} &> /dev/null

tail -100f /var/log/syslog
