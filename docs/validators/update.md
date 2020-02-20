# Migrate your validator node to a new version

:::warning Upgrade-only guide
The following guide is intended for all validators that are currently operating on a Desmos chain version and would like to upgrade to a new version.

If you wish to run a new validator node instead, please reference the [setup instructions](setup.md).  
:::  

## 1. Stop the currently running node. 
First of all we need to stop the currently running validator node. To do so you can go inside the console where you have run `desmosd start` and type `Ctrl + C`. This will halt your fullnode gracefully. 

If you have also setup a [background service](../fullnode/installation.md#optional-configure-the-service), please stop that too by executing the following command: 

```bash
systemctly stop desmosd
``` 

## 2. Export the current state
Once the fullnode has been properly stopped, you can export the current chain state. To do so execute the following command: 

```bash
desmosd export > old-state.json
```

This will allow you to write the current chain state to the `old-state.json` file inside the current directory. 

:::warning Beware of state changes  
During some updates it might happen that you need to perform some justified state changes before updating the fullnode and migrating to the new version.  

Before performing such changes make sure the people that tell you to do so are allowed to do it. Any required change will however be also written inside the [testnets repo](https://github.com/desmos-labs/morpheus) so make sure to always check that before performing any modification.  
:::

## 3. Update the underlying fullnode
Once you have stopped your validator, it is now time to update the Desmos binaries that allow your fullnode to run properly. To do so please reference the [fullnode updating guide](../fullnode/update.md). 

:::warning Do not start the fullnode yet  
After updating the fullnode software **do not** start it yet. If you have mistakenly started it, please shut it down before continuing.  
:::

## 4. Migrate to a new version   
After updating the fullnode, it is now time to migrate the old chain state to a new genesis state. To do so, you need to execute the following command: 

```bash
mv ~/.desmosd/config/
desmosd migrate <version> ./old-state.json \
  --chain-id <new-chain-id> \
  --genesis-time <new-genesis-time> \
  > ~/.desmosd/config
```

Please note that the `version`, `new-chain-id` and the `new-genesis-time` will be communicated to you in advance and will also be available inside the proper folder [on the testnets repo](https://github.com/desmos-labs/morpheus). 

Once you have migrated the 

## 5. Start the fullnode again
After you have properly migrated the genesis state, you can start again the fullnode and the validator by running 

```bash
desmosd start
``` 

:::warning Peers not connecting  
After updating the node to a new chain, it is normal that it cannot find any peer in the first blocks. However, if your node keeps having problem after a couple of hours, please contact us on the dedicated channels to help you with the troubleshooting.  
::: 
