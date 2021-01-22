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

# Checkout the correct commit
# Please check on https://github.com/desmos-labs/morpheus to get
# the tag to use based on the current network version
git checkout tags/<version>

# Build the software
# If you want to use the default database backend run 
make install

# If you want to use RocksDB run instead
make install DB_BACKEND=rocksdb 
```

If the software is built successfully, the `desmos` executable will be located inside your `$GOBIN` path. If you setup
your environment variables correctly in the previous step, you should also be able to run it properly. To check this,
try running:

```bash
desmos version --long
```

## 2. Initialize the Desmos working directory

Configuration files and chain data will be stored inside the `$HOME/.desmos` directory by default.

We can create this folder and all the necessary data by initializing a new fullnode. To do this, run:

```bash
# Initialize the working environment for Desmos
desmos init <your_moniker>
```

You can choose any moniker your like. It will be saved in the `config.toml` under the `.desmos` working directory.

### Recovering a previous node
Starting from `v0.15.0`, you are now able to provide a custom seed when initializing your node. This will be
particularly useful because, in the case that you want to reset your node, you will be able to re-generate the same
private node key instead of having to create a new node.

In order to provide a custom seed to your private key, you can do as follows:

1. Get a new random seed by running
   ```shell
   desmos keys add node --dry-run
   ```
   This will create a new key **without** adding it to your keystore, and output the underlying seed.

2. Copy the above provided seed, and then pass it to the `init` command using the `--recover` flag:
   ```shell
   desmos init <your_moniker> --recover <your_seed>
   ```

:::tip Recovering a node If you already have a seed, you can directly use the `--recover` flag without generating a new
one. This will recover the private key associated to that seed.
:::

## 3. Get the genesis file

To connect to an existing network, or start a new one, a genesis file is required. The file contains all the settings
telling how the genesis block of the network should look like. To connect to the `morpheus` testnets, you will need the
corresponding genesis file of each testnet. Visit the [testnet repo](https://github.com/desmos-labs/morpheus) and
download the correct genesis file by running the following command.

```bash
# Download the existing genesis file for the testnet
# Replace <chain-id> with the id of the testnet you would like to join
curl https://raw.githubusercontent.com/desmos-labs/morpheus/master/<chain-id>/genesis.json > $HOME/.desmos/config/genesis.json
```

## 4. Connect to seed nodes

To properly run your node, you will need to connect it to other full nodes running with the same software and genesis
file. This can be done configuring the `seeds` value inside the `config.toml` file localed under the `.desmos` working
directory.

```bash
# Open the config.toml file using text editor
nano $HOME/.desmos/config/config.toml
```

:::tip Where to get seeds  
Each testnet has their own seed nodes. You can get the ones for the testnet you would like to connect with inside
our [testnet repo](https://github.com/desmos-labs/morpheus), inside the specific testnet folder.
:::

Once you have a list of seeds to use, locate the `seeds = ""` text and update its value to a list of node addresses. The
format of a node address must be `<node_id>@<node_ip_address>:<port>` and multiple addresses must be separated by a
comma (`,`):

```bash
# Example
seeds = "cd4612957461881d5f62367c589aaa0fdf933bd8@seed-1.morpheus.desmos.network:26656,fc4714d15629e3b016847c45d5648230a30a50f1@seed-2.morpheus.desmos.network:26656"
``` 

Save the file and exit the text editor.

### Using persistent peers instead of seed nodes

Sometimes, it might happen that the seed nodes you have inserted do not work. If this happens, your node won't be able
to connect to the networks and start syncing. If this happens, what you can do is use the `persistent_peers` field
instead.

Persistent peers are other nodes to which your fullnode will persist the connection. To use this feature, open
the `config.toml` file and locate the `persistent_peers` line. The, fill its value with a list of node ids.

:::tip Where to get seeds  
Each testnet has their own persistent peers. You can get the ones for the testnet you would like to connect with inside
our [testnet repo](https://github.com/desmos-labs/morpheus), inside the specific testnet folder.
:::

## (Optional) Change your database backend

If you would like to run your node using [Facebook's RocksDB](https://github.com/facebook/rocksdb) as the database
backend, and you have correctly built the Desmos binaries to work with it following the instructions
at [point 1](#1-build-the-software), there is one more thing you need to do.

In order to tell Tendermint to use RocksDB as its database backend engine, you are required to change the following like
inside the `config.toml` file:

```toml
db_backend = "goleveldb"
```

To become

```toml
db_backend = "rocksdb"
```

## 5. Setup state sync

Starting from Desmos `v0.15.0`, we've added the support for Tendermint'
s [state sync](https://docs.tendermint.com/master/tendermint-core/state-sync.html). This feature allow new nodes to sync
with the chain extremely fast, by downloading snapshots created by other full nodes.

In order to use this feature, you need to edit your `$HOME/.desmos/config/config.toml` changing a couple of things under
the `statesync` section.

First of all, enable state sync setting `enable = true`.

Then, set the RPC addresses from where to get the snapshosts to `172.105.247.238:26657,139.162.131.107:26657`. These are
two of our fullnodes that are set up to create periodic snapshots every 600 blocks.

Thirdly, get a trusted chain height and the associated block hash. You can do so by running:

```bash
curl -s http://172.105.247.238:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

Once you have those values, use them as the `trust_height` and `trust_hash` values.

Here is an example of what it should look like in the end:

```toml
enable = true

rpc_servers = "172.105.247.238:26657,139.162.131.107:26657"
trust_height = 16962
trust_hash = "E8ED7A890A64986246EEB02D7D8C4A6D497E3B60C0CAFDDE30F2EE385204C314"
trust_period = "168h0m0s"
```

:::tip Make the snapshot download faster  
If you want to make the snapshot discovery faster, we suggest you setting the following `persistent_peers` inside
the `$HOME/.desmos/config/config.toml` file:

```toml
persistent_peers = "1441bc29cd8ce4b91d64a4bfa8138360d022dee7@139.162.131.107:26656,84cc13d6acf22c32c209f4205d2693f70f458dde@172.105.247.238:26656"
```

These are the nodes associated with the RPC servers that will be used to download the snapshots. Setting them as
persistent peers will avoid having to wait until they are added as peers later, resulting in an extremely fast download
experience.
:::

## 6. Open the proper ports

Now that everything is in place to start the node, the last thing to do is to open up the proper ports.

Your node uses vary different ports to interact with the rest of the chain. Particularly, it relies on:

- port `26656` to listen for incoming connections from other nodes
- port `26657` to expose the RPC service to clients

A part from those, it also uses:

- port `9090` to expose the [gRPC](https://grpc.io/) service that allows clients to query the chain state
- port `1317` to expose the REST APIs service

Most of the validators do not have interest in opening any port. However, it might be beneficial to the whole chain if
you can open port `26656`. This would allow new nodes to connect to you as a peer, making their sync faster and the
connection more reliable. For this reason, we suggest you opening that port by using `ufw`:

```bash
# Accept connections to port 26656 from any address
sudo ufw allow from any to any port 26656 proto tcp
```

If you also want to run a gRPC server, RPC node or the REST APIs, you also need to remember to open the related ports as
well.

## 7. Start the Desmos node

After setting up the binary and opening up ports, you are now finally ready to start your node:

```bash
# Run Desmos full node
desmos start
```

The full node will connect to the peers and start syncing. You can check the status of the node by executing:

```bash
# Check status of the node
desmos status
```

You should see an output like the following one:

```json
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
To allow your `desmos` instance to run in the background as a service you need to execute the following command

```bash
tee /etc/systemd/system/desmosd.service > /dev/null <<EOF  
[Unit]
Description=Desmos Full Node
After=network-online.target

[Service]
User=ubuntu
ExecStart=$GOBIN/desmos start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
```

:::warning  
If you are logged as a user which is not `ubuntu`, make sure to edit the `User` value accordingly  
:::

Once you have successfully created the service, you need to enable it. You can do so by running

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
● desmos.service - Desmos Node
   Loaded: loaded (/etc/systemd/system/desmosd.service; enabled; vendor preset: 
   Active: active (running) since Fri 2020-01-17 10:23:12 CET; 2min 3s ago
 Main PID: 11318 (desmos)
    Tasks: 10 (limit: 4419)
   CGroup: /system.slice/desmosd.service
           └─11318 /root/go/bin/desmos start
```

#### Check the node status
If you want to see the current status of the node, you can do so by running

```bash
journalctl -u desmosd -f
```

#### Stopping the service
If you wish to stop the service from running, you can do so by running

```bash
systemctl stop desmosd
```

To check the successful stop, execute `systemctl status desmos`. This should return

```bash
$ systemctl status desmosd
● desmos.service - Desmos Node
   Loaded: loaded (/etc/systemd/system/desmosd.service; enabled; vendor preset: enabled)
   Active: failed (Result: exit-code) since Fri 2020-01-17 10:28:04 CET; 3s ago
  Process: 11318 ExecStart=/root/go/bin/desmos start (code=exited, status=143)
 Main PID: 11318 (code=exited, status=143)
```
