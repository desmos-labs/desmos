---
id: seeds
title: Seeds
sidebar_label: Seeds
slug: seeds
---

# Seed nodes
:::caution Mainnet only   
The following seed nodes are to be used when configuring a full node for the **mainnet**. 
If you are looking for testnet seed nodes, please refer to [this](../05-testnet/04-join-public/03-seeds.md) instead.  
:::

Visit the [mainnet repo](https://github.com/desmos-labs/mainnet#seed-nodes) to get the seed nodes.

Add the seed nodes to the `~/.desmos/config/config.toml` file each one
separated by a comma like following:
```toml
seeds = "seed1,seed2,..."
```
