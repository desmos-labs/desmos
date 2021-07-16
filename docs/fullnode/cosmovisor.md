# Cosmovisor 
The Cosmos team provides a tool named _Cosmovisor_ that allows your node to perform some automatic operations when needed. This is particularly useful when dealing with on-chain upgrades, because Cosmovisor can help you by taking care of downloading the updated binary and restarting the node for you.  

If you want to learn how to setup Cosmovisor inside your full or validator node, please follow the guide below. 

## Setup
### 1. Downloading Cosmovisor
The first thing you have to do is downloading the `cosmovisor` binary file. To do this you can execute the following command: 

```shell
go get github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor
```

This will download all the dependencies and build `cosmovisor` for your machine. Once that's done, you can execute the following command to make sure that `cosmovisor` is installed: 

```shell
cosmovisor version
```

It should print `DAEMON_NAME is not set`. If that's the case, you have installed `cosmovisor` successfully.

### 2. Setting up environmental variables
Cosmovisor relies on the following environmental variables to work properly:

* `DAEMON_HOME` is the location where upgrade binaries should be kept (e.g. `$HOME/.desmos`).
* `DAEMON_NAME` is the name of the binary itself (eg. `desmos`).
* `DAEMON_ALLOW_DOWNLOAD_BINARIES` (*optional*) if set to `true` will enable auto-downloading of new binaries
  (for security reasons, this is intended for full nodes rather than validators).
* `DAEMON_RESTART_AFTER_UPGRADE` (*optional*) if set to `true` it will restart the sub-process with the same
  command line arguments and flags (but new binary) after a successful upgrade. By default, `cosmovisor` dies
  afterwards and allows the supervisor to restart it if needed. Note that this will not auto-restart the child
  if there was an error.
  
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
```

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

[Install]
WantedBy=multi-user.target
EOF
```

Once you have edited your system file, you need to reload it using the following command:

```shell
sudo systemctl daemon-reload
```

Finally, you can restart is as follows: 

```shell
sudo systemctl restart desmosd
```
