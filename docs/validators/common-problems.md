# Common Problems

## Problem #1: My validator has `voting_power: 0`
Your validator has become jailed. Validators get jailed, i.e. get removed from the active validator set, if they do not vote on `500` of the last `10000` blocks, or if they double sign. 

If you got jailed for downtime, you can get your voting power back to your validator. First, if `desmosd` is not running, start it up again:

```bash
desmosd start
```

Wait for your full node to catch up to the latest block. Then, you can [unjail your validator](#unjail-validator)

Lastly, check your validator again to see if your voting power is back.

```bash
desmoscli status
```

You may notice that your voting power is less than it used to be. That's because you got slashed for downtime!

## Problem #2: My `desmosd` crashes because of `too many open files`
The default number of files Linux can open (per-process) is `1024`. `desmosd` is known to open more than `1024` files. This causes the process to crash. A quick fix is to run `ulimit -n 4096` (increase the number of open files allowed) and then restart the process with `desmosd start`. If you are using `systemd` or another process manager to launch `desmosd` this may require some configuration at that level. A sample `systemd` file to fix this issue is below:

```{12}
# /etc/systemd/system/desmosd.service
[Unit]
Description=Desmos Full Node
After=network.target

[Service]
Type=simple
User=ubuntu # This is the user that is running the software in the background. Change it to your username if needed.
WorkingDirectory=/home/ubuntu # This is the home directory of the user that running the software in the background. Change it to your username if needed.
ExecStart=/home/ubuntu/go/bin/desmosd start # The path should point to the correct location of the software you have installed.
Restart=on-failure
RestartSec=3
LimitNOFILE=4096 # To compensate the "Too many open files" issue.

[Install]
WantedBy=multi-user.target
```

## Problem #3: My validator is inactive/unbonding
When creating a validator you have the minimum self delegation amount using the `--min-self-delegation` flag. What this means is that if your validator has less than that specific value of tokens seldelegated, it will automatically enter the unbonding state and then be marked as inactive. 

To solve this, what you can do is getting more tokens delegated to it by following these steps: 

1. Get your address: 
   ```bash
   desmoscli keys show <your_key> --address
   ```
   
2. Require more tokens using the [faucet](https://faucet.desmos.network). 

3. Make sure the tokens have been sent properly: 
   ```bash
   desmoscli query account $(desmoscli keys show <your_key> --address) --chain-id <chain_id>
   ```
   
4. Delegate the tokens to your validator: 
   ```bash
   desmoscli tx staking delegate \
     $(desmoscli keys show <your_key> --bech=val --address) \
     <amount> \
     --chain-id <chain_id> \
     --from <your_key> --yes
   
   # Example
   # desmoscli tx staking delegate \
   #  $(desmoscli keys show validator --bech=val --address) \
   #  10000000udaric \
   #  --chain-id morpheus-1001 \
   #  --from validator --yes
   ```

## Problem #4: My validator is jailed
If your validator is jailed it probably means that it have been inactive for a log period of time missing a consistent number of blocks. We suggest you checking the Desmos daemon status to make sure it hasn't been interrupted by some error.

If your service is running properly, you can also try and reset your `desmoscli` configuration by running the following command: 

```bash
rm $HOME/.desmoscli/config/config.toml
``` 

After doing so, remember to restart your validator service to apply the changes.

Once you have fixed the problems, you can unjail your validator by executing the following command: 

```bash
desmoscli tx slashing unjail --chain-id <chain_id> --from <your_key>

# Example
# desmoscli tx slashing unjail --chain-id morpheus-1001 --from validator
```

This will perform an unjail transaction that will set your validator as active again from the next block. 

If the problem still persists, please make sure you have [enough tokens delegated to your validator](#problem-3-my-validator-is-inactiveunbonding).

## Problem #5: The persistent peers do not work properly
Sometimes, it might happen that your node cannot connect to the persistent peers we have provided inside the [testnet repository](https://github.com/desmos-labs/morpheus). This happens because all nodes have a limit of inbound connections that they can accept. Once that limit is exceed, the nodes will not accept any more connections. 

In order to solve this problem, there are two alternative way: 

1. Use a seed node instead of a persistent peer;
2. **OR**, use different persistent peers.

### Using a seed node
Seed nodes are a particular type of nodes that provide every validator with a set of peers to connect with, based on the current network status. What will happen when you use seed nodes is the following: 

1. Your node will connect to a seed node. 
2. The seed node will provide your node with a list of peers. 
3. Your node will disconnect from the seed node and connect to the peers. 
4. Your node will start syncing with the chain. 

In order to use this particular type of nodes, all you have to do is:

1. Open the file located as
   ```
   ~/.desmosd/config/config.toml
   ```

2. Find the line starting with 
   ```
   seeds = ""
   ```

3. Replace that line with the following: 
   ```
   seeds = "cd4612957461881d5f62367c589aaa0fdf933bd8@seed-1.morpheus.desmos.network:26656,fc4714d15629e3b016847c45d5648230a30a50f1@seed-2.morpheus.desmos.network:26656"
   ```
   
4. Save the file and exit the editor. 
5. Restart your node.

### Changing your persistent peers
Instead of using a seed node, you can also keep relying on persistent peers. In this case, you will need to find new ones to connect your node to. To do this, you can query the current peers of any chain node using the following RPC endpoint: 

```
/net_info
```

For example, you can use the public RPC endpoint [here](http://rpc.morpheus.desmos.network:26657/net_info). 

From that page, you can see all the peers that a node is connect to. To do this, you can reference the `peers` field, which contains a list of objects made as follows: 

```json{4,5,17}
{
  "node_info": {
    "protocol_version": {},
    "id": "d45d4e0a6a6c393d58cfa1c5fed6286164fbfceb",
    "listen_addr": "tcp://0.0.0.0:26656",
    "network": "morpheus-10000",
    "version": "0.33.7",
    "channels": "4020212223303800",
    "moniker": "Maria",
    "other": {
      "tx_index": "on",
      "rpc_address": "tcp://127.0.0.1:26657"
    }
  },
  "is_outbound": false,
  "connection_status": {},
  "remote_ip": "35.193.251.165"
}
```

In order to get new peer addresses, all you have to do is to combine the `id`, `remote_ip` and `listen_addr` field values as follows: 

```
id@remote_ip:listen_addr(port)
```

In the above case, this peer's address would be: 

```
d45d4e0a6a6c393d58cfa1c5fed6286164fbfceb@35.193.251.165:26656
``` 

You can do this with as many peers as you want. Once you have a list of peers, you can use those inside the `persisten_peers` field of your `~/.desmosd/config/config.toml` file.

## Problem #6: The `desmoscli keys list` command does not work
Starting with v0.38, the Cosmos SDK uses os-native keyring to store all the private keys. Unfortunately, in some cases this does not work well by default. For example, it might return some errors when used in GUI-less machines.

In order to solve this problem, you have two options: 

1. Store the private keys inside a file on your machine
2. **OR**, use a password manager.


We highly suggest you to use a password manager. However, if you want to use a file-based approach you can execute the following commands:

```
mkdir -p ~/.desmoscli
desmoscli config keyring-backend file
```

Once you have executed those commands, you will be required to re-add your key by using your mnemonic phrase. To do so, run:

```
desmoscli keys add <your_key_name> --recover
```

This will require you to input the mnemonic phrase and then the keyring password. Once done, you should be able to execute all the `desmoscli keys`-related commands properly.
