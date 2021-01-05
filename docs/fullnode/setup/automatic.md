# Automatic full node setup
Following you will find how to download and execute the script that allows you to run a Desmos full node in minutes.  

:::warning Requirements  
Before proceeding, make sure you have read the overview page and your system satisfied the [needed requirements](overview.md#requirements).  

Also, make sure you setup your environment properly as described inside the [_"Setup your environment"_ section](overview.md#1-setup-your-environment).   
:::

## 1. Download the script
You can get the script by executing 

```shell
wget -O install-desmos-fullnode https://raw.githubusercontent.com/desmos-labs/desmos/master/contrib/validators/automatic-fullnode-installer.sh 
```

Once you downloaded it properly, you need to change its permissions making it executable: 

```shell
chmod +x ./install-desmos-fullnode
```

## 2. Execute the script
Once you got the script, you are now ready to use it. 

### Parameters
In order to work, it needs the following parameters: 

1. The `moniker` of the validator you are creating.  
   This is just a string that will allow you to identify the validator you are running locally. It can be anything you want. 
   
### Running the script
Once you are ready to run the script, just execute: 

```shell
./install-desmos-fullnode <PARAMETERS>
```

E.g: 

```
./install-desmos-fullnode my-validator
```

## How it works
The script will perform the following operations.

1. **Environment setup**   
   It will create all the necessary environmental variables. 
   
2. **Cosmovisor setup**  
   It will download Cosmovisor and set it up so that your node is able to automatically update based on on-chain upgrades.

3. **Desmos setup**  
   It will download and install Desmos properly so that your node is able to start syncing and also update itself based on all the on-chain upgrades that have been done until now.
   
4. **Service setup**  
   It will setup a system service to make sure Desmos runs properly in the background. 
   
5. **Service start and log output**  
   Finally, it will start the system service and output the logs from it. You will see it syncing the blocks properly and catching up with the rest of the chain. 
