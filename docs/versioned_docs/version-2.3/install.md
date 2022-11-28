# Installing Desmos

Desmos is represented by the executable named `desmos`.

It contains the Command Line Interface (CLI) that you can use to interface with the Desmos blockchain, as well as the
daemon that allows you take part to a Desmos blockchain either as a full node or
a [validator node](04-validators/01-overview.md).

## Requirements

The requirements you must satisfy before attempting to install Desmos are the following:

- Having Go 1.15 or later installed.  
  If you dont have it, you can download it [here](https://golang.org/dl/).

- Having Git installed.  
  If you need to install it, you can download the installer on the [official website](https://git-scm.com/downloads).

## Installation procedure

To install `desmos` execute the following commands:

```bash
cd /home/$USER
git clone https://github.com/desmos-labs/desmos && cd desmos
make install
```

### Verify the installation

To verify you have correctly `desmos`, try running:

```bash
desmos version
``` 

If you get an error like `No command found`, please make sure you have appended your `GOBIN` folder path to your
system's `PATH` environmental variable value.

:::tip Congratulations!   
You have successfully installed `desmos`!  
:::
