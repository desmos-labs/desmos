# Upgrade manager setup

The following guide allows you to setup your Desmos node along with
the [`cosmovisor`](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor) binary to ensure that each time a new
on-chain upgrade is required, it is handled properly.

## 1. Setup `desmos`

If you haven't already, please setup the `desmos` binary and your validator node. You can do so by following these
guides:

1. [Setup a full node](../../../fullnode/setup/overview.md)
2. [Setup a validator node](../../setup.md)

## 2. Setup `cosmovisor`

### Installation

In order to setup properly `cosmovisor` you are required to have Go 1.15+ installed on your machine. To make sure you
have it, you can execute the following command, checking that the output matches the one provided:

```shell
$ go version
go version go1.15.1 <os/build>
```

Also, make sure you have `git` installed as well: 

```shell
$ git version
git version 2.25.1
```

Now, we can start the installation process:

1. Create a folder for the `cosmovisor` source files:
    ```shell
   mkdir ~/cosmovisor
   ```

2. Clone the `cosmovisor` sources inside that folder:
   ```shell
   cd ~/cosmovisor 
   git clone https://github.com/cosmos/cosmos-sdk.git . 
   ``` 

3. Build the `cosmovisor` binary and install it inside `$GOBIN` to make it accessible everywhere:
   ```shell
   cd ~/cosmovisor/cosmovisor
   make cosmovisor
   mv cosmovisor $GOBIN/
   ```

Now, you should be able to execute `cosmovisor` from everywhere. If everything has been done correctly, this should be
the output:

```shell
$ cosmovisor
DAEMON_NAME is not set
```

### Environment setup
Now, we need to properly setup the environmental variables for `cosmosd`. To do so, run: 

```shell
echo 'export DAEMON_NAME=desmos' >> ~/.profile
echo 'export DAEMON_HOME=$HOME/.desmos' >> ~/.profile
echo 'export DAEMON_RESTART_AFTER_UPGRADE=on' >> ~/.profile
echo 'export DAEMON_ALLOW_DOWNLOAD_BINARIES=on  >> ~/.profile
source ~/.profile
``` 

This will load the environmental variables into the `~/.profile` file and then refresh the current terminal instance. 

If you want to know more about the set variables, here is a brief description: 

| Variable | Description |
|:-------- | :---------- |
|`DAEMON_HOME` | Location where upgrade binaries should be kept |
|`DAEMON_NAME` | Name of the binary itself |
| `DAEMON_RESTART_AFTER_UPGRADE` | (optional) if set to on it will restart a sub-process with the same args (but new binary) after a successful upgrade. By default, the manager dies afterwards and allows the supervisor to restart it if needed. Note that this will not auto-restart the child if there was an error. |
|`DAEMON_ALLOW_DOWNLOAD_BINARIES` | (optional) if set to `on` will enable auto-downloading of new binaries |


### Folders setup

In order to work properly, the upgrade manager inside `cosmovisor` needs to have a specific folder structure in place:
 
```
- genesis
  - bin
    - desmos
- upgrades
  - <name>
    - bin
      - desmos
```
 
In order to create it, please execute the following commands: 

```shell
mkdir -p ~/.desmos/upgrade_manager/genesis/bin
mkdir -p ~/.desmos/upgrade_manager/upgrades
```

Now we need to copy the current `desmos` binary into the `genesis/bin` folder to ensure that `cosmovisor` can start the
chain properly:

```shell
cp $GOBIN/desmos ~/.desmos/upgrade_manager/genesis/bin
```
