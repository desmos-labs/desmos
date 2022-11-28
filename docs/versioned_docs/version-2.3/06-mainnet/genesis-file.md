---
id: genesis-file
title: Genesis File
sidebar_label: Genesis File
slug: genesis-file
---

# Genesis file
:::caution Mainnet only   
The following seed nodes are to be used when configuring a full node for the **mainnet**. If you are looking for testnet seed nodes, please refer to [this](../05-testnets/03-join-public/genesis-file.md) instead.  
:::

To connect to the `desmos-mainnet`, you will need the corresponding genesis file.

## 1. Download
Visit the [mainnet repo](https://github.com/desmos-labs/mainnet) and
download the correct genesis file by running the following command.

```bash
# Download the existing genesis file for the mainnet
# Replace <chain-id> with the id of the testnet you would like to join
curl https://raw.githubusercontent.com/desmos-labs/mainnet/main/genesis.json > ~/.desmos/config/genesis.json
```

## 2. Check
After the download, ensure it's the correct one by checking that it has the same hashsum below:

```bash
jq -S -c -M '' /root/.desmos/config/genesis.json | shasum -a 256
619c9462ccd9045522300c5ce9e7f4662cac096eed02ef0535cca2a6826074c4  -
```
