# Full node setup
Following you will find the instructions on how to manually setup your Desmos full node.

:::warning Requirements
Before starting, make sure you read the [overview](overview.md) to make sure your hardware meets the needed
requirements.
:::

## 1. Build the software

:::tip Choose your DB backend
Before installing the software, a consideration must be done.

By default, Desmos uses [LevelDB](https://github.com/google/leveldb) as its database backend engine. However, since
version `v0.6.0` we've also added the possibility of optionally
using [Facebook's RocksDB](https://github.com/facebook/rocksdb), which, although still being experimental, is known to
be faster and could lead to lower syncing times. If you want to try out RocksDB you can take a look at
our [RocksDB installation guide](rocksdb-installation.md) before proceeding further.
:::

In your terminal, run the following:

```bash
# Make sure we are inside the home directory
cd $HOME

# Clone the Desmos software
git clone https://github.com/desmos-labs/desmos.git && cd desmos

# Checkout the correct tag
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

Configuration files and chain data will be stored inside the `$HOME/.desmos` directory by default. In order to create
this folder and all the necessary data we need to initialize a new full node using the `desmos init` command.

Starting from `v0.15.0`, you are now able to provide a custom seed when initializing your node. This will be
particularly useful because, in the case that you want to reset your node, you will be able to re-generate the same
private node key instead of having to create a new node.

In order to provide a custom seed to your private key, you can do as follows:

1. Get a new random seed by running
   ```shell
   desmos keys add node --dry-run

   # Example
   # desmos keys add node --dry-run
   # - name: node
   #   type: local
   #   address: desmos126cw9j2wy099lttf2qx0qds6k7t4kdea5ualh9
   #   pubkey: desmospub1addwnpepqdpfv4lh0vqjvmu43spz4lq0l92qret9hv6007j4r28z05wuthw2jz3frd4
   #   mnemonic: ""
   #   threshold: 0
   #   pubkeys: []
   #
   #
   # **Important** write this mnemonic phrase in a safe place.
   # It is the only way to recover your account if you ever forget your password.
   #
   # sort curious village display voyage oppose dice idea mutual inquiry keep swim team direct tired pink clinic figure tiny december giant obvious clump chest
   ```
   This will create a new key **without** adding it to your keystore, and output the underlying seed.

2. Run the `init` command using the `--recover` flag.
   ```shell
   desmos init <your_node_moniker> --recover
   ```
   You can choose any `moniker` value you like. It will be saved in the `config.toml` under the `.desmos` working
   directory.

3. Insert the previously outputed mnemonic phrase:
   ```
   > Enter your bip39 mnemonic
   sort curious village display voyage oppose dice idea mutual inquiry keep swim team direct tired pink clinic figure tiny december giant obvious clump chest
   ```

   This will generate the working files in `~/.desmos`

   :::tip Tip
   By default, running `desmos init <your_node_moniker>` without the `--recover` will randomly generate a `priv_validator_key.json`. There is no way to regenerate this key if you lose it.\
   We recommend running this command with the `--recover` so that you can regenerate the same `priv_validator_key.json` from the mnemonic phrase.
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

## 4. Setup seeds, peers and state sync

The next thing you have to do now is telling your node how to connect with other nodes that are already present on the
network. In order to do so, we will use the `seeds` and `persistent_peers` values of the `~/.desmos/config/config.toml`
file.

Seed nodes are a particular type of nodes present on the network. Your fullnode will connect to them, and they will
provide it with a list of other fullnodes that are present on the network. Then, your fullnode will automatically
connect to such nodes. Our team is running three seed nodes, and we advise you to use them by setting the
following `seeds` value:

```toml
seeds = "be3db0fe5ee7f764902dbcc75126a2e082cbf00c@seed-1.morpheus.desmos.network:26656,4659ab47eef540e99c3ee4009ecbe3fbf4e3eaff@seed-2.morpheus.desmos.network:26656,1d9cc23eedb2d812d30d99ed12d5c5f21ff40c23@seed-3.morpheus.desmos.network:26656"
```

Next, you will need to set some persistent peers of your node. Such nodes are going to be a particular type of peer
nodes to which your fullnode will always try to connect. You need to set them as the following value so that your node
can start syncing faster with the rest of the chain:

```toml
persistent_peers = "67dcef828fc2be3c3bcc19c9542d2b228bd7cff9@seed-4.morpheus.desmos.network:26656,fcf8207fb84a7238089bd0cd8db994e0af9016b6@seed-5.morpheus.desmos.network:26656"
```

### Using state sync

Starting from Desmos `v0.15.0`, we've added the support for Tendermint'
s [state sync](https://docs.tendermint.com/master/tendermint-core/state-sync.html). This feature allows new nodes to
sync with the chain extremely fast, by downloading snapshots created by other full nodes.

In order to use this feature, you will have to edit a couple of things inside your `~/.desmos/config/config.toml` file,
under the `statesync` section:

1. Enable state sync by setting `enable = true`

2. Set the RPC addresses from where to get the snapshots using the `rpc_servers` field
   to `seed-4.morpheus.desmos.network:26657,seed-5.morpheus.desmos.network:26657`.
   These are two of our fullnodes that are set up to create periodic snapshots every 600 blocks.

3. Get a trusted chain height, and the associated block hash. To do this, you will have to:
1. Get the current chain height by running:
   ```bash
   curl -s http://seed-4.morpheus.desmos.network:26657/commit | jq "{height: .result.signed_header.header.height}"
   ```
2. Once you have the current chain height, get a height that is a little bit lower (200 blocks) than the current one. To
   do this you can execute:
   ```bash
   curl -s http://seed-4.morpheus.desmos.network:26657/commit?height=<your-height> | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"

   # Example
   # curl -s http://seed-4.morpheus.desmos.network:26657/commit?height=100000 | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
   ```

4. Now that you have a trusted height and block hash, use those values as the `trust_height` and `trust_hash` values.

Here is an example of what the `statesync` section of your `~/.desmos/config/config.toml` file should look like in the
end (the `trust_height` and `trust_hash` should contain your values instead):

```toml
enable = true

rpc_servers = "seed-4.morpheus.desmos.network:26657,seed-5.morpheus.desmos.network:26657"
trust_height = 16962
trust_hash = "E8ED7A890A64986246EEB02D7D8C4A6D497E3B60C0CAFDDE30F2EE385204C314"
trust_period = "336h0m0s"
```

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

## 5. (Optional) Snapshots
Snapshots allow nodes to use state sync and quickly sync up. While it is optional, we extremely recommend turning this feature on to keep the chain stable.

We will assume you have set pruning to `default`. If that is not the case, please adjust accordingly

In order to turn this on, please to inside the `app.toml` file and change the following:

```toml
# snapshot-interval specifies the block interval at which local state sync snapshots are
# taken (0 to disable). Must be a multiple of pruning-keep-every.
snapshot-interval = 1500
```

This will tell the node to keep a snapshot image every 1500 blocks.


## 6. Open the proper ports

Now that everything is in place to start the node, the last thing to do is to open up the proper ports.

Your node uses vary different ports to interact with the rest of the chain. Particularly, it relies on:

- port `26656` to listen for incoming connections from other nodes
- port `26657` to expose the RPC service to clients

A part from those, it also uses:

- port `9090` to expose the [gRPC](https://grpc.io/) service that allows clients to query the chain state
- port `1317` to expose the REST APIs service

While opening any ports are optional, it is beneficial to the whole chain if
you open port `26656`. This would allow new nodes to connect to you as a peer, making their sync faster and the connection more reliable.

For this reason, we will be opening port `26656` using `ufw`. \
By default, `ufw` is not enabled. In order to enable it please run the following:

```bash
# running this should show it is inactive
sudo ufw status

# Turn on ssh if you need it
sudo ufw allow ssh

# Accept connections to port 26656 from any address
sudo ufw allow from any to any port 26656 proto tcp

# enable ufw
sudo ufw enable

# check ufw is running
sudo ufw status
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
  "NodeInfo": {
    "protocol_version": {
      "p2p": "8",
      "block": "11",
      "app": "0"
    },
    "id": "84cc13d6acf22c32c209f4205d2693f70f458dde",
    "listen_addr": "tcp://0.0.0.0:26656",
    "network": "morpheus-13001",
    "version": "",
    "channels": "40202122233038606100",
    "moniker": "fullnode",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://0.0.0.0:26657"
    }
  },
  "SyncInfo": {
    "latest_block_hash": "9BA7801C0935C4E18B4E2F8C6E8A2FF4C598C8E5F71A94113D2F0595524FD4E3",
    "latest_app_hash": "375C9F0E4E63B7ACAD497F8DEDF5E2382F469668DE671B2FF92A5D2B8B50C6D2",
    "latest_block_height": "204393",
    "latest_block_time": "2021-02-03T08:03:06.448549383Z",
    "earliest_block_hash": "839FEB9ED0257B71116CE54618C7E3C15189239CB571066DDBE9E0F1C101DCC8",
    "earliest_app_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
    "earliest_block_height": "1",
    "earliest_block_time": "2021-01-20T07:00:00Z",
    "catching_up": false
  },
  "ValidatorInfo": {
    "Address": "E457913A98EC0F9251BB40008E6680A8EFF93F99",
    "PubKey": {
      "type": "tendermint/PubKeyEd25519",
      "value": "BLT8jjQ+ODKa0ERcrhHfOVFVVrJDq7hxyXx6bLXnCdw="
    },
    "VotingPower": "0"
  }
}
```

If you see that the `catching_up` value is `false` under the `sync_info`, it means that you are fully synced. If it
is `true`, it means your node is still syncing. You can get the `catching_up` value by simply running:

```shell
desmos status 2>&1 | jq "{catching_up: .SyncInfo.catching_up}"

# Example
# $ desmos status 2>&1 | jq "{catching_up: .SyncInfo.catching_up}"
# {
#   "catching_up": false
# }
```

After your node is fully synced, you can consider running your full node as a [validator node](../validators/setup.md).

## 8. (Optional) Configure the background service

To allow your `desmos` instance to run in the background as a service you need to execute the following command

```bash
tee /etc/systemd/system/desmosd.service > /dev/null <<EOF
[Unit]
Description=Desmos Full Node
After=network-online.target

[Service]
User=$USER
ExecStart=$GOBIN/desmos start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
```

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

## 9. (Optional) Cosmovisor
In order to do automatic on-chain upgrades we will be using cosmovisor. Please check out [Using Cosmovisor](cosmovisor.md) for information on how to set this up.
