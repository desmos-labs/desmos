# Installing Desmos
Desmos is actually composed of two different executables: `desmoscli` and `desmosd`. 

The first one, `desmoscli`, represents the Command Line Interface (CLI) that you can use to interface with the Desmos blockchain.

The second one, `desmosd` (a.k.a. Desmos Daemon), represents instead the executable that allows you take part to a Desmos blockchain either as a full node or a [validator node](validators/overview.md). 

## Installing the Desmos CLI and Daemon
All you have to do to install both the `desmoscli` and `desmosd` executable is following the below procedure. 

### Requirements
The requirements you must satisfy before attempting to install Desmos CLI and Daemon are the following: 

- Having Go 1.13 or later installed.  
   If you dont have it, you can download it [here](https://golang.org/dl/).
   
- Having Git installed.  
  If you need to install it, you can download the installer on the [official website](https://git-scm.com/downloads).
   
### Installation procedure 
To install `desmoscli` and `desmosd` execute the following commands: 

```shell
cd /home/$USER
git clone https://github.com/desmos-labs/desmos && cd desmos
make install
```

Once all the procedure has completed properly, you should see an output similar to the following: 

```shell
$ make install
go install -mod=readonly -tags "netgo ledger" -ldflags '-X "github.com/cosmos/cosmos-sdk/version.Name=Desmos" -X "github.com/cosmos/cosmos-sdk/version.ServerName=desmosd" -X "github.com/cosmos/cosmos-sdk/version.ClientName=desmoscli" -X github.com/cosmos/cosmos-sdk/version.Version=0.1.0-19-g8df3833 -X github.com/cosmos/cosmos-sdk/version.Commit=8df3833b6967b776b9378fc11872c20f563113ae -X "github.com/cosmos/cosmos-sdk/version.BuildTags=netgo ledger"' ./cmd/desmosd
go install -mod=readonly -tags "netgo ledger" -ldflags '-X "github.com/cosmos/cosmos-sdk/version.Name=Desmos" -X "github.com/cosmos/cosmos-sdk/version.ServerName=desmosd" -X "github.com/cosmos/cosmos-sdk/version.ClientName=desmoscli" -X github.com/cosmos/cosmos-sdk/version.Version=0.1.0-19-g8df3833 -X github.com/cosmos/cosmos-sdk/version.Commit=8df3833b6967b776b9378fc11872c20f563113ae -X "github.com/cosmos/cosmos-sdk/version.BuildTags=netgo ledger"' ./cmd/desmoscli
```

### Verify the installation 
To verify you have correctly installed `desmoscli` and `desmosd`, try running: 

```shell
desmoscli version
desmosd version
``` 

If you get an error like `No command found`, please make sure you have appended your `GOBIN` folder path to your system's `PATH` environmental variable value.    

:::tip Congratulations!   
You have successfully installed `desmoscli` and `desmosd`!  
:::