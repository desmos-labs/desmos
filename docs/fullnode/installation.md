# Install and Running Desmos Fullnode

## Setup Your Environment
In order to run a fullnode, you need to build `desmosd` and `desmoscli` which require `Go`, `git`, `gcc` and `make` installed.

This process depends on your working environment.

:::: tabs

::: tab Linux
The following example is based on **Ubuntu (Debian)** and assumes you are using a terminal environment by default. Please run the equivalent commands if you are running other Linux distributions.

```bash
# Install git, gcc and make
sudo apt install build-essential --yes

# Install Go with Snap
sudo snap install go --classic

# Export environment variables
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

:::

::: tab MacOS
To install the required build tools, simple [install Xcode from the Mac App Store](https://apps.apple.com/hk/app/xcode/id497799835?l=en&mt=12).

To install `Go` on __MacOS__, the best option is to install with [__Homebrew__](https://brew.sh/). To do so, open the `Terminal` application and run the following command: 

```bash
# Install Homebrew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```


> If you don't know how to open a `Terminal`, you can search it by typing `terminal` in `Spotlight`. 

After __Homebrew__ is installed, run

```bash
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
The software has not been tested on __Windows__. If you are currently running a __Windows__ PC, the following options should be considered:

1. Switch to a __Mac__ 👨‍💻. 
2. Wipe your hard drive and install a __Linux__ system on your PC.
3. Install a separate __Linux__ system using [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
4. Run a __Linux__ instance on a cloud provider.

Note that is still possible to build and run the software on __Windows__ but it may give you unexpected results and it may require additional setup to be done. If you insist to build and run the software on __Windows__, the best bet would be installing the [Chocolatey](https://chocolatey.org/) package manager.

:::

::::

## Build the software
The following operations will all be done in the terminal environment under your home directory.

```bash
# Clone the Desmos software
git clone https://github.com/desmos-labs/desmos.git && cd desmos

# If you do not have git installed, you can install it by running
# sudo apt install git-all --yes

# Checkout the correct commit
# We are using v0.2.0 for morpheus-1001 testnet
git checkout v0.2.0

# Build the software
make install
```

If the software is built successfully, `desmosd` and `desmoscli` will be located at `/go/bin` of your home directory. If you have setup your environment variables correctly in the previous step, you should be able to access them correctly. Try to check the version of the software.

```bash
desmosd version --long
```

## Initialize the Desmos working directory
Configuration files and chain data will be stored inside the `.desmosd` directory under your home directory by default. It will be created when you initialize the environment.

```bash
# Initialize the working envinorment for Desmos
desmosd init <your_moniker>
```

You can choose any moniker your like. It will be saved in the `config.toml` under the `.desmosd` working directory.

## Get the genesis file
To connect to or start a new network, a genesis file is required. The file contains all the settings telling how the genesis block of the network should look like. To connect to the `morpheus` testnets, you will need the corresponding genesis file of each testnet. Visit the [testnet repo](https://github.com/desmos-labs/morpheus) and download the correct genesis file by running the following command.

```bash
# First, remove the newly created genesis file during the initialization
rm $HOME/.desmosd/config/genesis.json

# Download the existing genesis file for the testnet
# Assuming you are getting the genesis file for the latest testnet
curl https://raw.githubusercontent.com/desmos-labs/morpheus/master/genesis.json -o $HOME/.desmosd/config/genesis.json
```

## Connect to persistent peer
To properly run your node, you will need to connect it to other full nodes running with the same software and genesis file. This can be done configuring the `persisten_peers` value inside the `config.toml` file localed under the `.desmosd` working directory.

```bash
# Open the config.toml file using text editor
nano $HOME/.desmosd/config/config.toml
```

Locate the `persistent_peers = ""` text at line 164. Update its value to a node address of a peer. The format of a node address must be `<node_id>@<node_ip_address>:<port>`

```bash
# Example
persistent_peers = "89f913e84b58da594eb449fca7b0fcb540e52d05@35.240.254.97:26656,5ebcf26295e966c4d705ce5f42e78203c94ad98d@34.76.79.154:26656"
```

Save the file and exit the text editor.

## Start the Desmos node
Now you are good to run the full node. To do so, run:

```bash
# Run Desmos full node
desmosd start
```

The full node will connect to the peers and start syncing. You can check the status of the node by executing: 

```bash
# Check status of the node
desmoscli status
```

You should see an output like the following one:

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

If you see that the `catching_up` value is `false` under the `sync_info`, it means that you are fully synced. If it is `true`, it means your node is still syncing. 

After your node is fully synced, you can consider running your full node as a [validator node](../validators/validator-setup.md#create-your-validator).

## Configure the service
To allow your `desmosd` instance to run in the background as a service you need to execute the following command

```bash
tee /etc/systemd/system/desmosd.service > /dev/null <<EOF  
[Unit]
Description=Desmosd Node
After=network-online.target

[Service]
User=root
ExecStart=/root/go/bin/desmosd start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
```

:::warning  
If you are logged as a user which is not `root`, make sure to edit the `User` and `ExecStart` values accordingly  
::: 

Once you have successfully created the service, you need to first enable it. You can do so by running 

```bash
systemctl enable desmosd
```

After this, you can run it by executing

```bash
systemctl start desmosd
```

### Service operations 
#### Check the service status
If you want to see if the service is running properly, you can execute

```bash
systmctl status desmosd
``` 

If everything is running smoothly you should see something like 

```bash
$ systemctl status desmosd
● desmosd.service - Desmosd Node
   Loaded: loaded (/etc/systemd/system/desmosd.service; enabled; vendor preset: 
   Active: active (running) since Fri 2020-01-17 10:23:12 CET; 2min 3s ago
 Main PID: 11318 (desmosd)
    Tasks: 10 (limit: 4419)
   CGroup: /system.slice/desmosd.service
           └─11318 /root/go/bin/desmosd start
```

#### Check the node status
If you want to see the current status of the node, you can do so by running

```bash
tail -100f /var/log/syslog
```

This should return something like 

```
Jan 17 09:24:55 <your-moniker> desmosd[11318]: I[2020-01-17|10:24:55.212] Executed block                               module=state height=10183 validTxs=0 invalidTxs=0
Jan 17 09:24:55 <your-moniker> desmosd[11318]: I[2020-01-17|10:24:55.237] Committed state                              module=state height=10183 txs=0 appHash=0D8BEBCAC81A7B8DA1FBBF93FA6E921E7815AE3EBF53B78DB66CD8437DFD70C8
Jan 17 09:24:55 <your-moniker> desmosd[11318]: I[2020-01-17|10:24:55.252] Executed block                               module=state height=10184 validTxs=0 invalidTxs=0
Jan 17 09:24:55 <your-moniker> desmosd[11318]: I[2020-01-17|10:24:55.261] Committed state                              module=state height=10184 txs=0 appHash=459F68E6C5BF31EA5E58FB959829A587BB09B9F4DCA9C31CB754E5F26125FCD5
```

#### Stopping the service
If you wish to stop the service from running, you can do so by running

```bash
systemctl stop desmosd
```

To check the successful stop, execute `systemctl status desmosd`. This should return

```bash
$ systemctl status desmosd
● desmosd.service - Desmosd Node
   Loaded: loaded (/etc/systemd/system/desmosd.service; enabled; vendor preset: enabled)
   Active: failed (Result: exit-code) since Fri 2020-01-17 10:28:04 CET; 3s ago
  Process: 11318 ExecStart=/root/go/bin/desmosd start (code=exited, status=143)
 Main PID: 11318 (code=exited, status=143)
```