---
id: state-sync
title: State Sync
sidebar_label: State Sync
slug: state-sync
---

# State sync mainnet
:::caution Mainnet only   
The following seed nodes are to be used when configuring a full node for the **mainnet**. If you are looking for testnet seed nodes, please refer to [this](../05-testnets/03-join-public/state-sync.md) instead.  
:::

In order to use this feature, you will have to edit a couple of things inside your `~/.desmos/config/config.toml` file,
under the `statesync` section:

1. Enable state sync by setting `enable = true`;

2. Set the RPC addresses from where to get the snapshots using the `rpc_servers` field.  
   You can ask inside our [discord](https://discord.desmos.network/) for them.

3. Get a trusted chain height, and the associated block hash. To do this, you will have to:
    - Get the current chain height by running:
        ```bash
        curl -s <rpc-address>/commit  | jq "{height: .result.signed_header.header.height}"
        ```
    - Once you have the current chain height, get a height that is a little bit lower (200 blocks) than the current one.  
      To do this you can execute:
        ```bash
        curl -s <rpc-address>/commit?height=<your-height> | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
  
        # Example
        # curl -s https://rpc-desmos.itastakers.com/commit?height=100000 | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
        ```
4. Now that you have a trusted height and block hash, use those values as the `trust_height` and `trust_hash` values. Also,
   make sure they're the right values for the Desmos version you're starting to synchronize:

      | **State sync height range** | **Desmos version** |
      | :-------------------------: | :----------------: |
      |           `0 - 1149679`     |      `v1.0.1`      |
      |     `1149680 - 1347304`     |      `v2.3.0`      |
      |     `> 1347305`             |      `v2.3.1`      |

Here is an example of what the `statesync` section of your `~/.desmos/config/config.toml` file should look like in the end (the `trust_height` and `trust_hash` should contain your values instead):

```toml
enable = true

rpc_servers = "rpc-desmos.itastakers.com:26657,135.181.60.250:26557"
trust_height = 139142
trust_hash = "F55CA4C56CAC348E453A38D6BEBD70B1CD92F7431214AE167B09EFDA478186BE"
trust_period = "336h0m0s"
```
