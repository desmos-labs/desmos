---
id: reset-data
title: Reset Data
sidebar_label: Reset Data
slug: reset-data
---

# Reset the data
In case something goes wrong and your node can't be recovered, you can reset it and sync it again.

Reset the data:
```bash
rm $HOME/.desmos/config/addrbook.json $HOME/.desmos/config/genesis.json
desmos unsafe-reset-all
```
    
Your node is now in a pristine state while keeping the original `priv_validator.json` and `config.toml`. If you had any sentry nodes or full nodes setup before, your node will still try to connect to them, but may fail if they haven't also been upgraded.

:::warning  
Make sure that every node has a unique `priv_validator.json`. Do not copy the `priv_validator.json` from an old node to multiple new nodes. Running two nodes with the same `priv_validator.json` will cause you to **double sign**.  
:::

After the reset, you can sync back your node with state-sync, check how depending you are doing this on:
- [Testnet](../05-testnet/03-join-public/04-state-sync.md)
- [Mainnet](../06-mainnet/03-state-sync.md)