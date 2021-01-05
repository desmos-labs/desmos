# Installing and running a Desmos fullnode

:::warning This guide is for new fullnodes only  
If you have previously run a fullnode and you wish to update it instead, please follow the [updating guide](../update.md).   
:::

## Requirements
### Understanding pruning
In order to run a full node, different hardware requirements should be met based on the pruning strategy you would like to use.

*Pruning* is the term used to identify the periodic action that can be taken in order to free some disk space on your full node. This is done by removing old blocks data from the disk, freeing up space.

Inside Desmos, there are various types of pruning strategies that can be applied. The main ones are: 

- `default`: the last 100 states are kept in addition to every 500th state; pruning at 10 block intervals

- `nothing`: all historic states will be saved, nothing will be deleted (i.e. archiving node)

- `everything`: all saved states will be deleted, storing only the current state; pruning at 10 block intervals

### Hardware requirements
You can easily understand how using a pruning strategy of `nothing` will use more disk space than `everything`. For this reason, there are different disk space that we recommend based on the pruning strategy you choose:

| Pruning strategy | Minimum disk space | Recommended disk space |
| :--------------: | :----------------: | :--------------------: |
| `everything` | 20 GB | 40 GB | 
| `default` | 80 GB | 120 GB |
| `nothing` | 120 GB | \> 240 GB |

A part from disk space, the following requirements should be met.

| Minimum CPU cores | Recommended CPU cores |
| :---------------: | :-------------------: |
| 2 | 4 |

| Minimum RAM | Recommended RAM |
| :---------------: | :-------------------: |
| 4 GB | 8 GB |


## 1. Setup your environment
In order to run a fullnode, you need to build `desmosd` and `desmoscli` which require `Go`, `git`, `gcc` and `make` installed.

This process depends on your working environment.

:::: tabs

::: tab Linux
The following example is based on **Ubuntu (Debian)** and assumes you are using a terminal environment by default. Please run the equivalent commands if you are running other Linux distributions.

```bash
# Update the system
sudo apt update 
sudo apt upgrade 

# Install git, gcc and make
sudo apt install git build-essential --yes

# Install Go with Snap
sudo snap install go --classic

# Export environment variables
echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export GOBIN="$GOPATH/bin"' >> ~/.profile
echo 'export PATH="$GOBIN:$PATH"' >> ~/.profile
source ~/.profile
```

:::

::: tab MacOS
To install the required build tools, simple [install Xcode from the Mac App Store](https://apps.apple.com/hk/app/xcode/id497799835?l=en&mt=12).

To install `Go` on __MacOS__, the best option is to install with [__Homebrew__](https://brew.sh/). To do so, open the `Terminal` application and run the following command: 

```bash
# Install Homebrew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```


> If you don't know how to open a `Terminal`, you can search it by typing `terminal` in `Spotlight`. 

After __Homebrew__ is installed, run

```bash
# Install Go using Homebrew
brew install go

# Install Git using Homebrew
brew install git

# Export environment variables
echo 'export GOPATH="$HOME/go"' >> ~/.profile
echo 'export GOBIN="$GOPATH/bin"' >> ~/.profile
echo 'export PATH="$GOBIN:$PATH"' >> ~/.profile
source ~/.profile
```

:::

::: tab Windows
The software has not been tested on __Windows__. If you are currently running a __Windows__ PC, the following options should be considered:

1. Switch to a __Mac__ 👨‍💻. 
2. Wipe your hard drive and install a __Linux__ system on your PC.
3. Install a separate __Linux__ system using [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
4. Run a __Linux__ instance on a cloud provider.

Note that is still possible to build and run the software on __Windows__ but it may give you unexpected results and it may require additional setup to be done. If you insist to build and run the software on __Windows__, the best bet would be installing the [Chocolatey](https://chocolatey.org/) package manager.

:::

::::

## 2. Install the software
Once you have setup your environment correctly, you are now ready to install the Desmos software and start your full node. 

In order to do so, you have two options: 

1. [Automatic installation](automatic.md).  
   This requires you to run a few commands so that you can easily have a Desmos full node running in a couple of minutes. This is recommended to most users that do not yet have the skills to go through a manual setup. 
   
2. [Manual installation](manual.md).  
   This is recommended to the those who want to understand each step or want to customize their setup accordingly to their needs. It is not recommended to people running a validator for the first time, although everyone should take a look at it once. 
