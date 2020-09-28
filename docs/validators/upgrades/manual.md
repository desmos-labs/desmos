# Migrate your validator node to a new network

:::warning Upgrade-only guide
The following guide is intended for all validators that are currently operating on a Desmos chain version and would like to upgrade to a new version.

If you wish to run a new validator node instead, please reference the [setup instructions](../setup.md).  
:::  

## 1. Stop the currently running node. 
First of all we need to stop the currently running validator node. To do so you can go inside the console where you have run `desmosd start` and type `Ctrl + C`. This will halt your fullnode gracefully. 

If you have also setup a [background service](../../fullnode/setup/manual.md#optional-configure-the-service), please stop that too by executing the following command: 

```bash
systemctl stop desmosd
``` 

## 2. Export the current state
Once the fullnode has been properly stopped, you can export the current chain state. To do so execute the following command: 

```bash
desmosd export --for-zero-height > old-state.json
```

This will allow you to write the current chain state to the `old-state.json` file inside the current directory. 

:::warning Beware of state changes  
During some updates it might happen that you need to perform some justified state changes before updating the fullnode and migrating to the new version.  

Before performing such changes make sure the people that tell you to do so are allowed to do it. Any required change will however be also written inside the [testnets repo](https://github.com/desmos-labs/morpheus) so make sure to always check that before performing any modification.  
:::

## 3. Update the underlying fullnode
Once you have stopped your validator, it is now time to update the Desmos binaries that allow your fullnode to run properly. To do so please reference the [fullnode updating guide](../../fullnode/update.md). 

:::warning Do not start the fullnode yet  
After updating the fullnode software **do not** start it yet. If you have mistakenly started it, please shut it down before continuing.  
:::

## 4. Migrate to a new network   
After updating the fullnode, it is now time to migrate the old chain state to a new genesis state. 
First of all, do a backup of your current genesis file: 

```bash
cp ~/.desmosd/config/genesis.json ~/.desmosd/config/genesis.json.bak
```

Then, you can migrate the old state and replace it with the new one: 
```bash
desmosd migrate <version> old-state.json \
  --chain-id <new-chain-id> \
  --genesis-time <new-genesis-time> \
  > ~/.desmosd/config/genesis.json
```

Please note that the `version`, `new-chain-id` and the `new-genesis-time` will be communicated to you in advance and will also be available inside the proper folder [on the testnets repo](https://github.com/desmos-labs/morpheus). 

Once you have migrated the genesis file, you need to reset the status of your node.

## 5. Reset your node
:::danger Make sure you understand what you are doing  
Please be cautious when you reset your node. Unintended mistakes may lead to missing validator key or double sign.  
:::

Resetting the node will have the followings happen.
1. **Reset validating state**
2. **Remove chain data**
3. **Remove address book**

To reset the node, you need to execute the following command:

```bash
desmosd unsafe-reset-all
```

Please make sure your validator key located at `~/.desmosd/config/priv_validator_key.json` is intact.

## 6. Start the fullnode again
After you have properly migrated the genesis state, you can start again the fullnode and the validator by running 

```bash
desmosd start
``` 

:::warning Peers not connecting  
After updating the node to a new chain, it is normal that it cannot find any peer in the first blocks. However, if your node keeps having problem after a couple of hours, please contact us on the dedicated channels to help you with the troubleshooting.  
::: 
