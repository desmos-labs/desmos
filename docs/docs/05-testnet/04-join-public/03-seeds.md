---
id: seeds
title: Seeds
sidebar_label: Seeds
slug: seeds
---

# Seed nodes
:::caution Testnet only   
The following seed nodes are to be used when configuring a full node for 
the **testnet**. 
If you are looking for mainnet seed nodes, please refer to 
[this](../../06-mainnet/02-seeds.md) instead.
:::

To find the corresponding seeds nodes of each testnet visit the 
[testnet repo](https://github.com/desmos-labs/morpheus). The seed nodes
are inside the folder having the name of the testnet of interest. 

Add the seed nodes to the `~/.desmos/config/config.toml` file each one 
separated by a comma like following:
```toml
seeds = "seed1,seed2,..."
```
