---
id: genesis-file
title: Genesis File
sidebar_label: Genesis File
slug: genesis-file
---

# Genesis file
:::caution Testnet only   
To configure a full node for the **testnet** you need to use the following seed nodes. If you are looking for mainnet seed nodes, please refer to [this](../../06-mainnet/01-genesis-file.md) instead.
:::

To connect to the `morpheus` testnet, you will need the corresponding genesis file of each testnet. Visit the [testnet repo](https://github.com/desmos-labs/morpheus) and download the correct genesis file by running the following command.

```bash
# Download the existing genesis file for the testnet
# Replace <chain-id> with the id of the testnet you would like to join
curl https://raw.githubusercontent.com/desmos-labs/morpheus/master/<chain-id>/genesis.json > $HOME/.desmos/config/genesis.json
```
