# Install and Running Desmos Fullnode

## Setup Your Environment

In order to run a fullnode, you need to build `desmosd` and `desmoscli` which require `Go`, `git`, `gcc` and `make` installed.

This process depends on your working environment.

:::: tabs

::: tab Linux
The following example is based on **Ubuntu(Debian)** and assume you are using terminal environment by default. Please run the equivalent commands if you are running other Linux distributions.

``` bash
# Install git, gcc and make
sudo apt install build-essential

# Install Go with Snap
sudo snap install go -classic

# Export environment variables
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

:::

::: tab MacOS
To install the required build tools, simple [install Xcode from Mac App Store](https://apps.apple.com/hk/app/xcode/id497799835?l=en&mt=12).

To install `Go` on __MacOS__, the best option is to install with [__Homebrew__](https://brew.sh/). Open `Terminal` on __MacOS__ and run the following command to install __Homebrew__. If you don't know how to open a `Terminal`, you can search it by typing `terminal` in `Spotlight`.

When you see the prompt, run this.

``` bash
# Install Homebrew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

After __Homebrew__ is installed, run

``` bash
# Install Go using Homebrew
brew install go

# Install Git using Homebrew
brew install git

# Export environment variables
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

:::

::: tab Windows
The software has not been tested on __Windows__. If you are currently running a __Windows__ PC, there are options you can consider.

1. Switch to a __Mac__ 👨‍💻
2. Swipe your harddrive and install a __Linux__ system on your PC
3. Install a separate __Linux__ system using [VirutalBox](https://www.virtualbox.org/wiki/Downloads)
4. Run a __Linux__ instance on a cloud provider

The software is still possible to be built and run on __Windows__ but it may give unexpected results and it requires some settings on __Windows__. If you insist to build and run the software on __Windows__, the best bet would be installing the [Chocolatey](https://chocolatey.org/) package manager.
:::

::::

## Build the software

The following operations will all be done in the terminal environment under your home directory.

``` bash{2,6,9}
# Clone the Desmos software
git clone https://github.com/desmos-labs/desmos.git && cd desmos

# Checkout correct commit
# We are using v0.2.0 for morpheus-1001 testnet
git checkout v0.2.0

# Build the software
make install
```

If the software is built successfully, `desmosd`, `desmoscli` and `desmoskeyutil` will be located at `/go/bin` of your home directory. If you have setup your ennivroment variables correctly in the previous step, you should be able to access them correctly. Try to check the version of the software.

``` bash
desmosd version --long
```

## Initialize the Desmos working directory

Configuration files and chain data will be stored in the `.desmosd` directory under your home directory by default. It will be created when you initialize the environment.

``` bash
# Initialize the working envinorment for Desmos
desmosd init <your_moniker>
```

You can choose any moniker your like. It will be saved in the `config.toml` under the `.desmosd` working directory.

## Get the genesis file

To connect or start a network, a genesis file is required. The file contains all the settings how the genesis block of the network look like. To connect to the `Morpheus` testnets, you will need the corresponding genesis file of each testnet. Visit the [testnet repo](https://github.com/desmos-labs/morpheus) and download the correct genesis file by running the following command.

``` bash
# First, remove the newly created genesis file during the initialization
rm $HOME/.desmosd/config/genesis.json

# Download the existing genesis file for the testnet
# Assuming you are getting the genesis file for morpheus-1000
curl https://raw.githubusercontent.com/desmos-labs/morpheus/master/1000/genesis.json -o $HOME/.desmosd/genesis.json
```

## Connect to persistent peer

You need to connect your node to other full nodes running with the same software and genesis file. This can be configured in the `config.toml` under the `.desmosd` working directory.

``` bash
# Open the config.toml file using text editor
nano $HOME/.desmosd/config/config.toml
```

Locate `persistent_peers = ""` at line 164. Update the value of it to a node address of a peer. The format of a node address is in `<node_id>@<node_ip_address>:<port>`

``` bash
persistent_peers = "8307c16191e249d6d3871ce764262d40d9cf249f@34.74.131.47:26656"
```

Save the file and exit the text editor.

## Start the Desmos node

Now you are good to run the full node.

``` bash
# Run Desmos full node
desmosd start
```

The full node will connect to the peers and start syncing. You can check the status of the node.

``` bash
# Check status of the node
desmoscli status
```

You should see the output like this.

``` json{24}
{
  "node_info": {
    "protocol_version": {
      "p2p": "7",
      "block": "10",
      "app": "0"
    },
    "id": "8307c16191e249d6d3871ce764262d40d9cf249f",
    "listen_addr": "tcp://0.0.0.0:26656",
    "network": "morpheus-1001",
    "version": "0.32.7",
    "channels": "4020212223303800",
    "moniker": "Morpheus",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://0.0.0.0:26657"
    }
  },
  "sync_info": {
    "latest_block_hash": "AAB066E5B020C325E5AE41CFACFB95DAC83B261D0C4A24439A66A2977A03B7EC",
    "latest_app_hash": "5B4CE89D3C1EFA1AE8E16710103EEAE1FDF9D13BE44F5847F5376810E8F39DAE",
    "latest_block_height": "480950",
    "latest_block_time": "2020-01-13T06:35:29.237599298Z",
    "catching_up": false
  },
  "validator_info": {
    "address": "25AD49347EC88C1922F18B317D12EA59DB0EC8D6",
    "pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "WPZGfRMuyd8X4B4vAx79yfyqH+nmEboaML8YlJKT/uE="
    },
    "voting_power": "0"
  }
}
```

If you see `catching_up` is `false` in the `sync_info`, it means that you are fully synced. If it is `true`, it is still syncing. 

After your node is fully synced, you can consider running your full node as a [validator node](/validators/validator-setup.html#create-your-validator).
