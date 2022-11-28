---
id: cosmovisor
title: Use Cosmovisor
sidebar_label: Use Cosmovisor
slug: cosmovisor
---

# Cosmovisor 
The Cosmos team provides a tool named _Cosmovisor_ that allows your node to perform some automatic operations when needed. This is particularly useful when dealing with on-chain upgrades, because Cosmovisor can help you by taking care of downloading the updated binary and restarting the node for you.  

If you want to learn how to setup Cosmovisor inside your full or validator node, please follow the guide below. 

## Setup
### 1. Downloading Cosmovisor
To install the latest version of `cosmovisor`, run the following command:
```
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest
```

To install a previous version, you can specify the version. 

```
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0
```

You can also install from source by pulling the cosmos-sdk repository and switching to the correct version and building as follows:
```
git clone git@github.com:cosmos/cosmos-sdk
cd cosmos-sdk
git checkout cosmovisor/vx.x.x
cd cosmovisor
make
```

This will build cosmovisor in your current directory. Afterwards you may want to put it into your machine's PATH like as follows:
```
cp cosmovisor ~/go/bin/cosmovisor
```
To check your cosmovisor version, you can use this command and make sure it matches the version you've installed:
```
strings $(which cosmovisor) | egrep -e "mod\s+github.com/cosmos/cosmos-sdk/cosmovisor"
```

*Note: If you are using go `v1.15` or earlier, you will need to use `go get`, and you may want to run the command outside a project directory.*

### 2. Setting up environmental variables
Cosmovisor relies on the following environmental variables to work properly:

* `DAEMON_HOME` is the location where upgrade binaries should be kept (e.g. `$HOME/.desmos`).
* `DAEMON_NAME` is the name of the binary itself (eg. `desmos`).
* `DAEMON_ALLOW_DOWNLOAD_BINARIES` (*optional*, default = `false`) if set to `true` will enable auto-downloading of new binaries
  (for security reasons, this is intended for full nodes rather than validators).
* `DAEMON_RESTART_AFTER_UPGRADE` (*optional*, default = `true`) if set to `true` it will restart the sub-process with the same
  command line arguments and flags (but new binary) after a successful upgrade. By default, `cosmovisor` dies
  afterwards and allows the supervisor to restart it if needed. Note that this will not auto-restart the child
  if there was an error.
* `DAEMON_POLL_INTERVAL` (*optional*, default = `300ms`) is the interval length for polling the upgrade plan file. The value can either be a number (in milliseconds) or a duration (e.g. `1s`).
* `UNSAFE_SKIP_BACKUP` (*optional*, default = `false`), if set to `true`, upgrades directly without performing a backup. Otherwise (`false`) backs up the data before trying the upgrade. The default value of `false` is useful and recommended in case of failures and when a backup needed to rollback. We recommend using the default backup option `UNSAFE_SKIP_BACKUP=false`.
  
To properly set those variables, we suggest you to edit the `~/.profile` file so that they are loaded when you log into your machine. To edit this file you can simply run 

```shell
sudo nano ~/.profile
```

Once you're in, we suggest you to set the following values: 

```
export DAEMON_HOME=$HOME/.desmos
export DAEMON_NAME=desmos
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_RESTART_AFTER_UPGRADE=true
export UNSAFE_SKIP_BACKUP=false
```

**IMPORTANT**: If you don't have much free disk space, please set `UNSAFE_SKIP_BACKUP=true` to avoid your node failing the upgrade due to insufficient disk space when creating the backup.

Once you're done, please reload the `~/.profile` file by running 

```shell
source ~/.profile
```

You can verify the values set by running 

```
echo $DAEMON_NAME
```

If this outputs `desmos` you are ready to go.

### 3. Copying Desmos files in the proper folders
In order to work properly, Cosmovisor needs the `desmos` binary to be placed in the `~/.desmos/cosmovisor/genesis/bin` folder. To do this you can simply run the following command: 

```shell
mkdir -p ~/.desmos/cosmovisor/genesis/bin/
cp $(which desmos) ~/.desmos/cosmovisor/genesis/bin/
```

To verify that you have setup everything correctly, you can run the following command: 

```shell
cosmovisor version
```

This should output the Desmos version.

### 4. Restarting your node
Finally, if you've setup everything correctly you can now restart your node. To do this you can simply stop the running `desmos start` and re-start your Desmos node using the following command: 

```
cosmovisor start
```

#### Updating the service file
If you are running your node using a service, you need to update your service file to use `cosmovisor` instead of `desmos`. To do this you can simply run the following command:

```shell
sudo tee /etc/systemd/system/desmosd.service > /dev/null <<EOF  
[Unit]
Description=Desmos Full Node
After=network-online.target

[Service]
User=$USER
ExecStart=$(which cosmovisor) start
Restart=always
RestartSec=3
LimitNOFILE=4096

Environment="DAEMON_HOME=$HOME/.desmos"
Environment="DAEMON_NAME=desmos"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="UNSAFE_SKIP_BACKUP=false"

[Install]
WantedBy=multi-user.target
EOF
```

**IMPORTANT**: If you don't have much free disk space, please set `UNSAFE_SKIP_BACKUP=true` to avoid your node failing the upgrade due to insufficient disk space when creating the backup.

Once you have edited your system file, you need to reload it using the following command:

```shell
sudo systemctl daemon-reload
```

Finally, you can restart is as follows: 

```shell
sudo systemctl restart desmosd
```
