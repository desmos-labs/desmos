# Run a Validator on Desmos

## What is a Validator?

[Validators](overview.md) are responsible for committing new blocks to the blockchain through voting. A validator's stake is slashed if they become unavailable or sign blocks at the same height. Please read about [Sentry Node Architecture](validator-faq.md#how-can-validators-protect-themselves-from-denial-of-service-attacks) to protect your node from DDOS attacks and to ensure high-availability.

::: danger Warning
If you want to become a validator for the `mainnet`, you should [research security](security.md).
:::

## Running a Fullnode
To become a validator, you must first have `desmosd` and `desmoscli` installed and be able to run a fullnode. You can first [setup your fullnode](/fullnode/overview.md) if you haven't yet.

:::tip Not enough funds? Google Cloud can help you  
Running a validator node should be done on a separate machine, not your local computer. This is due to the fact that validators need to be constantly running to avoid getting slashed (and thus loosing funds). We highly recommend setting up a local machine that can run 24/7, even a Raspberry can do the job. 

If you do not have the possibility of using a local machine, even an hosted server can be perfect. If you wish to get started for free, you can use the [300$ Google Cloud bonus](https://cloud.google.com/free/docs/gcp-free-tier). This should be enough to run a validator for 5-6 months.  
:::  

## Create Your Validator

Your `desmosvalconspub` (Desmos Validator Consensus Pubkey) can be used to create a new validator by staking tokens. You can find your validator pubkey by running:

```bash
desmosd tendermint show-validator
```

To create your validator, just use the following command:

::: warning 
Don't use more staking token than you have! 

On Morpheus testnet, we are using `udaric` as the staking token and it will be the example below. 

We are going to use `udesmos` as the staking token on Mainnet.
:::

```bash
desmoscli tx staking create-validator \
  --amount=1000000udaric \
  --pubkey=$(desmosd tendermint show-validator) \
  --moniker="choose a moniker" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1" \
  --gas="auto" \
  --gas-adjustment="1.2" \
  --gas-prices="0.025udaric" \
  --from=<key_name>
```

::: tip
When specifying commission parameters, the `commission-max-change-rate` is used to measure % _point_ change over the `commission-rate`. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.
:::

::: tip
`Min-self-delegation` is a stritly positive integer that represents the minimum amount of self-delegated staking token your validator must always have. A `min-self-delegation` of 1 means your validator will never have a self-delegation lower than `1udaric`. A valdiator self delegate lower than this number will automatically be unbonded.
:::

You can confirm that you are in the validator set by using a block explorer, e.g. [Big Dipper](https://morpheus.desmos.network).

## Participate in Genesis as a Validator
If you want to participate in genesis as a validator, you need to justify that
you have some stake at genesis, create one (or multiple) transactions to bond this stake to your validator address, and include this transaction in the genesis file.

Your `desmosvalconspub` can be used to create a new validator by staking tokens. You can find your validator pubkey by running:

```bash
desmosd tendermint show-validator
```

Next, craft your `desmosd gentx` command. 

::: tip
A `gentx` is a JSON file carrying a self-delegation. All genesis transactions are collected by a `genesis coordinator` and validated against an initial `genesis.json`.
:::

::: warning Note
Don't use more `desmos` than you have! 
:::

```bash
desmosd gentx \
  --amount <amount_of_delegation_desmos> \
  --commission-rate <commission_rate> \
  --commission-max-rate <commission_max_rate> \
  --commission-max-change-rate <commission_max_change_rate> \
  --pubkey <consensus_pubkey> \
  --name <key_name>
```

::: tip
When specifying commission parameters, the `commission-max-change-rate` is used to measure % _point_ change over the `commission-rate`. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.
:::

These `gentx` will be used to form the final genesis file. 

## Edit Validator Description
You can edit your validator's public description. This info is to identify your validator, and will be relied on by delegators to decide which validators to stake to. Make sure to provide input for every flag below. If a flag is not included in the command the field will default to empty (`--moniker` defaults to the machine name) if the field has never been set or remain the same if it has been set in the past.

The <key_name> specifies which validator you are editing. If you choose to not include certain flags, remember that the --from flag must be included to identify the validator to update.

The `--identity` can be used as to verify identity with systems like Keybase or UPort. When using with Keybase `--identity` should be populated with a 16-digit string that is generated with a [keybase.io](https://keybase.io) account. It's a cryptographically secure method of verifying your identity across multiple online networks. The Keybase API allows some block explorers to retrieve your Keybase avatar. This is how you can add a logo to your validator profile.

```bash
desmoscli tx staking edit-validator
  --moniker="choose a moniker" \
  --website="https://desmos.network" \
  --identity=6A0D65E29A4CBC8E \
  --details="To infinity and beyond!" \
  --chain-id=<chain_id> \
  --from=<key_name> \
  --commission-rate="0.10"
```

__Note__: The `commission-rate` value must adhere to the following invariants:

- Must be between 0 and the validator's `commission-max-rate`
- Must not exceed the validator's `commission-max-change-rate` which is maximum
  % point change rate **per day**. In other words, a validator can only change
  its commission once per day and within `commission-max-change-rate` bounds.

## View Validator Description
View the validator's information with this command:

```bash
desmoscli query staking validator <account_desmos>
```

```bash
desmoscli tx slashing unjail \
	--from=<key_name> \
	--chain-id=<chain_id>
```

## Confirm Your Validator is Running
Your validator is active if the following command returns anything:

```bash
desmoscli query tendermint-validator-set | grep "$(desmosd tendermint show-validator)"
```

You should now see your validator in one of the Desmos explorers. You are looking for the `bech32` encoded `address` in the `~/.desmosd/config/priv_validator.json` file.

::: warning Note
To be in the validator set, you need to have more total voting power than the 100th validator.
:::

## Halting Your Validator
When attempting to perform routine maintenance or planning for an upcoming coordinated
upgrade, it can be useful to have your validator systematically and gracefully halt.
You can achieve this by either setting the `halt-height` to the height at which
you want your node to shutdown or by passing the `--halt-height` flag to `desmosd`.
The node will shutdown with a zero exit code at that given height after committing
the block.

## Common Problems
### Problem #1: My validator has `voting_power: 0`
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

### Problem #2: My `desmosd` crashes because of `too many open files`
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
