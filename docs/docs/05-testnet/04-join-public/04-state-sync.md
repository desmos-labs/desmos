---
id: state-sync
title: State Sync
sidebar_label: State Sync
slug: state-sync
---

# State sync testnet
:::caution Testnet only   
The following seed nodes are to be used when configuring a full node for the **testnet**. If you are looking for mainnet seed nodes, please refer to [this](../../06-mainnet/03-state-sync.md) instead.
:::

In order to use this feature, you will have to edit a couple of things inside your `~/.desmos/config/config.toml` file,
under the `statesync` section:

1. Enable state sync by setting `enable=true`;

2. Set the RPC addresses from where to get the snapshots using the `rpc_servers` field.
    You can find the state sync nodes at the [testnet repo](https://github.com/desmos-labs/morpheus) inside
    the folder with the name of the testnet that you want to join.
   
3. Get a trusted chain height, and the associated block hash. To do this, you will have to:
    - Get the current chain height by running:
       ```bash
       curl -s curl -s <rpc-address>/commit | jq "{height: .result.signed_header.header.height}"
       ```
    - Once you have the current chain height, get a height that is a little lower (200 blocks) than the current one.  
      To do this you can execute:
       ```bash
       curl -s curl -s <rpc-address>/commit?height=<your-height> | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
 
       # Example
       # curl -s curl -s <rpc-address>/commit?height=100000 | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
       ```
      
4. Now that you have a trusted height and block hash, use those values as the `trust_height` and `trust_hash` values. 
   Also, make sure they're the right values for the Desmos version you're starting to synchronize.
   You can check them looking inside the associated testnet folder [here](https://github.com/desmos-labs/morpheus).

Here is an example of what the `statesync` section of your `~/.desmos/config/config.toml` file should look like in the end (the `trust_height` and `trust_hash` should contain your values instead):

```toml
enable = true

rpc_servers = "seed-4.morpheus.desmos.network:26657,seed-5.morpheus.desmos.network:26657"
trust_height = 16962
trust_hash = "E8ED7A890A64986246EEB02D7D8C4A6D497E3B60C0CAFDDE30F2EE385204C314"
trust_period = "336h0m0s"
```

5. Add peers to `~/.desmos/config/config.toml` file:

 ```toml
persistent_peers = "67dcef828fc2be3c3bcc19c9542d2b228bd7cff9@seed-4.morpheus.desmos.network:26656,fcf8207fb84a7238089bd0cd8db994e0af9016b6@seed-5.morpheus.desmos.network:26656"
 ```
