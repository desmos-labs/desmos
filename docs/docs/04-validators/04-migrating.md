---
id: migrating
title: Migrating
sidebar_label: Migrating
slug: migrating
---

# Migrating a validator
In some cases, you might want to move your running validator from one server to another. For example, this might be the case if you find a cheaper alternative or if your host does not meet the requirements.

In this case, what you need to do is following a two-step procedure:

<!-- 1. Stop the current running validator and backup the data.
2. Setup the new server and transfer the backed up files to it.
3. Start the new node. -->

1. Set and fully sync up a fullnode
2. Migrate data from old node to new node

## 1. Setup new fullnode
To avoid as much missed blocks as possible it is recommended that you setup a new server and fully sync up a full node first. To do this,
you can read the guide [here](../03-fullnode/02-setup.md).

Proceed to step 2 only if your new node has caught up.

```
desmos status 2>&1 | jq "{catching_up: .SyncInfo.catching_up}"
```

## 2. Migrate Data
### Stopping both the validator and the newly synced full node
In order to avoid as much side affects as possible we will be stopping both nodes.

```
systemctl stop desmos
```

### Backup the following data in the validator node

In order to properly migrate our validator node to another server we will need to backup the following data:

1. The validator private key.
   This is located inside the `~/.desmos/config/priv_validator_key.json` file.

2. The validator consensus state.
   This is located inside the `~/.desmos/data/priv_validator_state.json` file.
3. If you keep your keys on the node make sure you have the secret recovery phrase (mnemonic phrase) associated with your key(s).

<!-- :::warning Do not move them in to the new fullnode just yet
Back them up somewhere save but don't
::: -->


<!-- In order to properly transfer your validator to another server, you first have to stop the running node. To do this, you can execute the following command:

```
systemctl stop desmos
``` -->

<!-- Once you have done so, you need to back up the following data: -->


:::tip Do not delete your old node yet
We highly suggest you to delete your old node instance once that the new node is running properly. This will allow you to recover any information if you forgot to do so.
:::

### Migrating data to the new fullnode
With both nodes stopped we will be copying the backed up data in to the newly synced fullnode.

:::warning Warning
At this point, both nodes should not be running. This is to prevent any possible side effects such as double signing
:::

Transfer the following backed up files from the old node to the new node:

- `~/.desmos/config/priv_validator_key.json`
- `~/.desmos/data/priv_validator_state.json`


### Startup the new validator node
Once you have moved your `priv_validator_key.json` and `priv_validator_state` to the newly synced fullnode, it will be recognized as the same validator node.

:::warning Warning
The `priv_validator_key` should only be online from a single instance. A good practice would be to remove or rename the `priv_validator_key.json` in the old node so that even if you accidentally start running both nodes, the validator key would only be online for one of them.
:::

With the old validator node stopped, start up the new node:

```
sudo systemctl start desmos
```

:::warning Wait before deleting the old node
We suggest you to wait before deleting the old node. Instead of waiting only to see the blocks syncing, make sure your new node is actually signing blocks as your validator. You will see this by looking at your validator uptime on our [explorer](https://morpheus.desmos.network/validators). If everything is working properly, the uptime should slowly increase.
:::

### (Optional) Recover your key
If you originally had your key in the previous server you can easily add it back using the secret recovery phrase (mnemonic phrase) you had backed up

```
desmos keys add <key_name> --recover
```
