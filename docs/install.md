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

```bash
cd /home/$USER
git clone https://github.com/desmos-labs/desmos && cd desmos
make install
```

### Verify the installation 
To verify you have correctly installed `desmoscli` and `desmosd`, try running: 

```bash
desmoscli version
desmosd version
``` 

If you get an error like `No command found`, please make sure you have appended your `GOBIN` folder path to your system's `PATH` environmental variable value.    

:::tip Congratulations!   
You have successfully installed `desmoscli` and `desmosd`!  
:::
