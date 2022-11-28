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

2. Set the RPC addresses from where to get the snapshots using the `rpc_servers` field to
   `seed-4.morpheus.desmos.network:26657,seed-5.morpheus.desmos.network:26657`.   
   These are two of our fullnodes that are set up to create periodic snapshots every 600 blocks;
   
3. Get a trusted chain height, and the associated block hash. To do this, you will have to:
    - Get the current chain height by running:
       ```bash
       curl -s http://seed-4.morpheus.desmos.network:26657/commit | jq "{height: .result.signed_header.header.height}"
       ```
    - Once you have the current chain height, get a height that is a little lower (200 blocks) than the current one.  
      To do this you can execute:
       ```bash
       curl -s http://seed-4.morpheus.desmos.network:26657/commit?height=<your-height> | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
 
       # Example
       # curl -s http://seed-4.morpheus.desmos.network:26657/commit?height=100000 | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
       ```
      
   1. Now that you have a trusted height and block hash, use those values as the `trust_height` and `trust_hash` values. Also,
      make sure they're the right values for the Desmos version you're starting to synchronize:

      | **State sync height range** |     **Desmos version**      |
      |:-----------------------------|:---------------------------|
      | `0 - 1235764`               |          `v0.17.0`          |
      | `1235765 - 1415529`         |          `v0.17.4`          |
      | `1415530 - 2121235`         |          `v0.17.6`          |
      | `2121236 - 2226899`         |          `v1.0.4`           |
      | `2226900 - 2589024`         |          `v2.0.0`           |
      | `2589025 - 2643234`         |          `v2.1.0`           |
      | `2643235 - 2756259`         |          `v2.2.0`           |
      | `2756260 - 3130831`         |          `v2.3.0`           |
      | `2756260 - 3130831`         |          `v2.3.0`           |
      | `3130831 - 5842610`         |          `v2.3.1`           |
      | `5842610 - 6233130`         |          `v3.2.0`           |
      | `6233130 - 6339185`         |          `v4.0.1`           |
      | `> 6339185`                 |          `v4.1.0`           |

Here is an example of what the `statesync` section of your `~/.desmos/config/config.toml` file should look like in the end (the `trust_height` and `trust_hash` should contain your values instead):

```toml
enable = true

rpc_servers = "seed-4.morpheus.desmos.network:26657,seed-5.morpheus.desmos.network:26657"
trust_height = 16962
trust_hash = "E8ED7A890A64986246EEB02D7D8C4A6D497E3B60C0CAFDDE30F2EE385204C314"
trust_period = "336h0m0s"
```
