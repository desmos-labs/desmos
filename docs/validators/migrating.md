# Migrating a validator
In some cases, you might want to move your running validator from one server to another. For example, this might be the case if you find a cheaper alternative or if your host does not meet the requirements. 

In this case, what you need to do is following a three steps procedure: 

1. Stop the current running validator and backup the data.
2. Setup the new server and transfer the backed up files to it. 
3. Start the new node.

## 1. Backup the data
In order to properly transfer your validator to another server, you first have to stop the running node. To do this, you can execute the following command: 

```
systemctl stop desmosd
```

Once you have done so, you need to back up the following data:

1. The mnemonic phrase associated with your key.  
   You should have written it down on a piece of paper when you first created the node.
   
2. The validator private key.  
   This is located inside the `~/.desmosd/config/priv_validator_key.json` file.
   
3. The validator consensus state.  
   This is located inside the `~/.desmosd/data/priv_validator_state.json` file.

:::tip Do not delete your old node yet   
We highly suggest you to delete your old node instance once that the new node is running properly. This will allow you to recover any information if you forgot to do so.  
:::

## 2. Setup your new node
### Setting up the full node
Once you are ready to get your new node running, the first thing to do is setup your node as a full node. To do this, you can read the guide [here](../fullnode/setup/overview.md). 

Once that is done, stop your new node from running by using 

```
systemctl stop desmosd
```

Now, transfer the following backed up files from the old node into the new one:

- `~/.desmosd/config/priv_validator_key.json`
- `~/.desmosd/data/priv_validator_state.json` 

### Restoring your key
To restored your validator key, all you have to do is execute the following command:

```
desmoscli keys add <your_key_name> --recover
```

Now, insert your mnemonic phrase and then your keyring password. 

## 3. Start the new node
Once everything is in place, you can start your new node. To do this, run the following commands: 

```
desmosd unsafe-reset-all
sudo systemctl start desmosd
```

:::warning Make sure the old node is not running  
Before starting the new node, make sure you stop the old one. If you start the new node, and the old one is still running, you will run into a double signature, and your node will be slashed of 5% of its staked amount. 

To prevent this, stop the old server from running if possible.  
::: 

When you are sure your node caught up to the chain properly, you can now delete the old server instance and all the files contained inside it. 

:::tip Wait before deleting the old node  
We suggest you to wait before deleting the old node. Instead of waiting only to see the blocks syncing, make sure your new node is actually signing blocks as your validator. You will see this by looking at your validator uptime on our [explorer](https://morpheus.desmos.network/validators). If everything is working properly, the uptime should slowly increase.  
:::
