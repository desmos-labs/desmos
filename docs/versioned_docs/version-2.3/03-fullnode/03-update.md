---
id: update
title: Update
sidebar_position: 3
---

# Updating your Desmos fullnode
These instructions are for fullnodes that are running on previous versions of Desmos and would like to update to the
latest testnet.

## 1. Reset the data
First, remove the outdated files and reset the data.

```bash
rm $HOME/.desmos/config/addrbook.json $HOME/.desmos/config/genesis.json
desmos unsafe-reset-all
```

Your node is now in a pristine state while keeping the original `priv_validator.json` and `config.toml`. If you had any sentry nodes or full nodes setup before, your node will still try to connect to them, but may fail if they haven't also been upgraded.

:::warning  
Make sure that every node has a unique `priv_validator.json`. Do not copy the `priv_validator.json` from an old node to multiple new nodes. Running two nodes with the same `priv_validator.json` will cause you to **double sign**.  
:::

## 2. Upgrade the software
Now it is time to upgrade the software.

Go into the directory in which you have installed `desmos`. If you have followed
the [installation instructions](02-setup.md) and didn't change the path, it should be `/home/$USER/desmos`:

```bash
cd <installation-path> 

# E.g
# cd /home/$USER/desmos
``` 

Now, update the `desmos` software:

```bash
git fetch --all
git checkout tags/$(git describe --tags `git rev-list --tags --max-count=1`)
make install
```

:::tip Select another version  
The above commands checks out the latest release that has been tagged on our repository. If you wish to check out a
specific version instead, use the following commands:

1. List all the tags  
   ```bash
   git tags --list
   ```
   
2. Checkout the tag you want 
   ```bash
   git checkout tags/<tag>
   # Example: git checkout tags/v2.3.1
   ```
   
:::

:::tip Note   
If you have issues at this step, please check that you have the [latest stable version](https://golang.org/dl/) of Go installed.  
:::

Your full node has been cleanly updated!
