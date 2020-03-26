# Common Problems

## Problem #1: My validator has `voting_power: 0`
Your validator has become jailed. Validators get jailed, i.e. get removed from the active validator set, if they do not vote on `500` of the last `10000` blocks, or if they double sign. 

If you got jailed for downtime, you can get your voting power back to your validator. First, if `desmosd` is not running, start it up again:

```bash
desmosd start
```

Wait for your full node to catch up to the latest block. Then, you can [unjail your validator](#unjail-validator)

Lastly, check your validator again to see if your voting power is back.

```bash
desmoscli status
```

You may notice that your voting power is less than it used to be. That's because you got slashed for downtime!

## Problem #2: My `desmosd` crashes because of `too many open files`
The default number of files Linux can open (per-process) is `1024`. `desmosd` is known to open more than `1024` files. This causes the process to crash. A quick fix is to run `ulimit -n 4096` (increase the number of open files allowed) and then restart the process with `desmosd start`. If you are using `systemd` or another process manager to launch `desmosd` this may require some configuration at that level. A sample `systemd` file to fix this issue is below:

```toml
# /etc/systemd/system/desmosd.service
[Unit]
Description=Desmos Full Node
After=network.target

[Service]
Type=simple
User=ubuntu # This is the user that running the software in the background. Change it to your username if needed.
WorkingDirectory=/home/ubuntu # This is the home directory of the user that running the software in the background. Change it to your username if needed.
ExecStart=/home/ubuntu/go/bin/desmosd start # The path should point to the correct location of the software you have installed.
Restart=on-failure
RestartSec=3
LimitNOFILE=4096 # To compensate the "Too many open files" issue.

[Install]
WantedBy=multi-user.target
```

## Problem #3: My validator is inactive/unbonding
When creating a validator you have the minimum self delegation amount using the `--min-self-delegation` flag. What this means is that if your validator has less than that specific value of tokens seldelegated, it will automatically enter the unbonding state and then be marked as inactive. 

To solve this, what you can do is getting more tokens delegated to it by following these steps: 

1. Get your address: 
   ```bash
   desmoscli keys show <your_key> --address
   ```
   
2. Require more tokens using the [faucet](https://faucet.desmos.network). 

3. Make sure the tokens have been sent properly: 
   ```bash
   desmoscli query account $(desmoscli keys show <your_key> --address) --chain-id <chain_id>
   ```
   
4. Delegate the tokens to your validator: 
   ```bash
   desmoscli tx staking delegate \
     $(desmoscli keys show <your_key> --bech=val --address) \
     <amount> \
     --chain-id <chain_id> \
     --from <your_key> --yes
   
   # Example
   # desmoscli tx staking delegate \
   #  $(desmoscli keys show validator --bech=val --address) \
   #  10000000udaric \
   #  --chain-id morpheus-1001 \
   #  --from validator --yes
   ```

## Problem #4: My validator is jailed
If your validator is jailed it probably means that it have been inactive for a log period of time missing a consistent number of blocks. We suggest you checking the Desmos daemon status to make sure it hasn't been interrupted by some error.

If your service is running properly, you can also try and reset your `desmoscli` configuration by running the following command: 

```bash
rm $HOME/.desmoscli/config/config.toml
``` 

After doing so, remember to restart your validator service to apply the changes.

Once you have fixed the problems, you can unjail your validator by executing the following command: 

```bash
desmoscli tx slashing unjail --chain-id <chain_id> --from <your_key>

# Example
# desmoscli tx slashing unjail --chain-id morpheus-1001 --from validator
```

This will perform an unjail transaction that will set your validator as active again from the next block. 

If the problem still persists, please make sure you have [enough tokens delegated to your validator](#problem-3-my-validator-is-inactiveunbonding).
