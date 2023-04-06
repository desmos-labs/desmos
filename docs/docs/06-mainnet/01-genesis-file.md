---
id: genesis-file
title: Genesis File
sidebar_label: Genesis File
slug: genesis-file
---

# Genesis file
:::caution Mainnet only   
If you are looking for testnet genesis file, please refer to
[this](../05-testnet/04-join-public/02-genesis-file.md) instead.  
:::

To connect to the `desmos-mainnet`, you will need to download the correct genesis file.

## 1. Download
Visit the [mainnet repo](https://github.com/desmos-labs/mainnet) and
download the correct genesis file by running the following command.

```bash
# Download the existing genesis file for the mainnet
# Replace <chain-id> with the id of the testnet you would like to join
curl https://raw.githubusercontent.com/desmos-labs/mainnet/main/genesis.json > ~/.desmos/config/genesis.json
```

## 2. Check
After the download, ensure it's the correct one by checking that it has the correct hashsum like below:

```bash
jq -S -c -M '' ~/.desmos/config/genesis.json | shasum -a 256
```

The expected output should be:

`619c9462ccd9045522300c5ce9e7f4662cac096eed02ef0535cca2a6826074c4  -`
