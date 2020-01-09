# Join the public testnet

:::tip Current Testnet  
See the [testnet repo](https://github.com/desmos-labs/morpheus) for information on the latest testnet, including the correct version of Desmos to use and details about the genesis file.  
:::

## Validators
:::warning Requires desmos executables  
To join the public testnet you **must** have [`desmoscli` and `desmosd` installed](../install.md).  
:::

To become a testnet validators, the mainnet instructions apply: 

1. [Create a full node](../validators/node-setup.md).
2. [Become a validator](../validators/validator-setup.md)

The only difference is the SDK version and genesis file. See the [testnet repo](https://github.com/desmos-labs/morpheus) for information on testnets, including the correct version of Desmos to use and details about the genesis file.

## Upgrading Your Node
These instructions are for full nodes that have ran on previous versions of and would like to upgrade to the latest testnet.

### Reset Data
First, remove the outdated files and reset the data.

```bash
rm $HOME/.desmosd/config/addrbook.json $HOME/.desmosd/config/genesis.json
desmosd unsafe-reset-all
```

Your node is now in a pristine state while keeping the original `priv_validator.json` and `config.toml`. If you had any sentry nodes or full nodes setup before, your node will still try to connect to them, but may fail if they haven't also been upgraded.

:::warning  
Make sure that every node has a unique `priv_validator.json`. Do not copy the `priv_validator.json` from an old node to multiple new nodes. Running two nodes with the same `priv_validator.json` will cause you to double sign.  
:::

## Software Upgrade
Now it is time to upgrade the software.

Go into the directory in which you have installed `desmos`. If you have followed the [installation instructions](../install.md) and didn't change the path, it should be `/home/$USER/desmos`: 

```bash
cd <installation-path> 

# E.g
# cd /home/$USER/desmos
``` 

Now, update the `desmoscli` and `desmosd` software:

```bash
git clone https://github.com/desmos-labs/desmos.git .
git fetch --all && git checkout master
make install
```

:::tip Note   
If you have issues at this step, please check that you have the [latest stable version](https://golang.org/dl/) of Go installed.  
:::

Note we use `master` here since it contains the latest stable release. See the [testnet](https://github.com/desmos-labs/morpheus) repo for details on which version is needed for which testnet, and the [Desmos release page](https://github.com/desmos-labs/desmos/releases) for details on each release.

Your full node has been cleanly upgraded!

