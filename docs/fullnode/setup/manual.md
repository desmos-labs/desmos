# Manual full node setup
Following you will find the instructions on how to manually setup your Desmos full node.

:::warning Requirements  
Before starting, make sure you read the [setup overview](overview.md) to make sure your hardware meets the needed requirements.  
:::

## 1. Build the software
:::tip Choose your DB backend  
Before installing the software, a consideration must be done. 

By default, Desmos uses [LevelDB](https://github.com/google/leveldb) as its database backend engine. However, since version `v0.6.0` we've also added the possibility of optionally using [Facebook's RocksDB](https://github.com/facebook/rocksdb), which, although still being experimental, is know to be faster and could lead to lower syncing times. If you want to try out RocksDB (which we suggest you to do) you can take a look at our [RocksDB installation guide](../rocksdb-installation.md) before proceeding further.  
:::

The following operations will all be done in the terminal environment under your home directory.

```bash
# Clone the Desmos software
git clone https://github.com/desmos-labs/desmos.git && cd desmos

# If you do not have git installed, you can install it by running
# sudo apt install git-all --yes

# Checkout the correct commit
# Please check on https://github.com/desmos-labs/morpheus to get the proper 
# the tag to use based on the current network version
git checkout tags/<version>

# Build the software
# If you want to use the default database backend run 
make install

# If you want to use RocksDB run instead
make install DB_BACKEND=rocksdb 
```

If the software is built successfully, `desmosd` and `desmoscli` will be located at `/go/bin` of your home directory. If you have setup your environment variables correctly in the previous step, you should be able to access them correctly. Try to check the version of the software.

```bash
desmosd version --long
```

## 2. Initialize the Desmos working directory
Configuration files and chain data will be stored inside the `.desmosd` directory under your home directory by default. It will be created when you initialize the environment.

```bash
# Initialize the working envinorment for Desmos
desmosd init <your_moniker>
```

You can choose any moniker your like. It will be saved in the `config.toml` under the `.desmosd` working directory.

## 3. Get the genesis file
To connect to or start a new network, a genesis file is required. The file contains all the settings telling how the genesis block of the network should look like. To connect to the `morpheus` testnets, you will need the corresponding genesis file of each testnet. Visit the [testnet repo](https://github.com/desmos-labs/morpheus) and download the correct genesis file by running the following command.

```bash
# First, remove the newly created genesis file during the initialization
rm $HOME/.desmosd/config/genesis.json

# Download the existing genesis file for the testnet
# Assuming you are getting the genesis file for the latest testnet
curl https://raw.githubusercontent.com/desmos-labs/morpheus/master/genesis.json -o $HOME/.desmosd/config/genesis.json
```

## 4. Connect to seed nodes
To properly run your node, you will need to connect it to other full nodes running with the same software and genesis file. This can be done configuring the `seeds` value inside the `config.toml` file localed under the `.desmosd` working directory.

```bash
# Open the config.toml file using text editor
nano $HOME/.desmosd/config/config.toml
```

Locate the `seeds = ""` text at line 164. Update its value to a node address of a seed node. The format of a node address must be `<node_id>@<node_ip_address>:<port>`

```bash
# Example
seeds = "cd4612957461881d5f62367c589aaa0fdf933bd8@seed-1.morpheus.desmos.network:26656,fc4714d15629e3b016847c45d5648230a30a50f1@seed-2.morpheus.desmos.network:26656"
```

Save the file and exit the text editor.

## (Optional) Change your database backend
If you would like to run your node using [Facebook's RocksDB](https://github.com/facebook/rocksdb) as the database backend, and you have correctly built the Desmos binaries to work with it following the instructions at [point 1](#1-build-the-software), there is one more thing you need to do. 

In order to tell Tendermint to use RocksDB as its database backend engine, you are required to change the following like inside the `config.toml` file: 

```toml
db_backend="goleveldb"
```

To become

```toml
db_backend="rocksdb"
```

Once you have done so, you can go ahead with [point 5](#5-start-the-desmos-node).

## 5. Start the Desmos node
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

```
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

After your node is fully synced, you can consider running your full node as a [validator node](../../validators/setup.md).

## (Optional) Configure the service
To allow your `desmosd` instance to run in the background as a service you need to execute the following command

```bash
tee /etc/systemd/system/desmosd.service > /dev/null <<EOF  
[Unit]
Description=Desmosd Full Node
After=network-online.target

[Service]
User=ubuntu
ExecStart=/home/ubuntu/go/bin/desmosd start
Restart=always
RestartSec=3
LimitNOFILE=4096 # To compensate the "Too many open files" issue.

[Install]
WantedBy=multi-user.target
EOF
```

:::warning  
If you are logged as a user which is not `ubuntu`, make sure to edit the `User` and `ExecStart` values accordingly  
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
systemctl status desmosd
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
