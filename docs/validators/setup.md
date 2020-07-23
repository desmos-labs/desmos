# Become a Desmos validator
[Validators](overview.md) are responsible for committing new blocks to the blockchain through voting. A validator's stake is slashed if they become unavailable or sign blocks at the same height. Please read about [Sentry Node Architecture](validator-faq.md#how-can-validators-protect-themselves-from-denial-of-service-attacks) to protect your node from DDOS attacks and to ensure high-availability.

::: danger Warning
If you want to become a validator for the `mainnet`, you should [research security](security.md).
:::

## 1. Run a fullnode
To become a validator, you must first have `desmosd` and `desmoscli` installed and be able to run a fullnode. You can first [setup your fullnode](../fullnode/overview.md) if you haven't yet.

:::tip Not enough funds? Google Cloud can help you  
Running a validator node should be done on a separate machine, not your local computer. This is due to the fact that validators need to be constantly running to avoid getting slashed (and thus loosing funds). We highly recommend setting up a local machine that can run 24/7, even a Raspberry can do the job. 

If you do not have the possibility of using a local machine, even an hosted server can be perfect. If you wish to get started for free, you can use the [300$ Google Cloud bonus](https://cloud.google.com/free/docs/gcp-free-tier). This should be enough to run a validator for 5-6 months.  
:::  

:::tip Previously gone thought a Primer challenge? Reset your `desmoscli`
If you've previously gone through a Primer challenge, you need to make sure your `demoscli` is properly setup to communicate with your node. In order to do so, please execute the following command: 

```bash
rm $HOME/.desmoscli/config/config.toml
```
:::

## 2. Create your validator
In order to create a validator, you need to have to create a local wallet first. This will be used in order to hold the tokens that you will later delegate to your validator node, allowing him to properly work. In order to create this wallet, please run: 

```shell
desmoscli keys add <key_name>
```  

:::warning Key name  
Please select a key name that you will easily remember and be able to type fast. This name will be used all over the places inside other commands later.   
:::

Once that you have created your local wallet, it's time to get some tokens to be used as the initial validator stake so that it can run properly. If you are setting up a validator inside one of our testnets, please refer to our [testnet repo](https://github.com/desmos-labs/morpheus) to know the faucet address. If you are running a validator on our mainnet, you will need to purchase the tokens.

To run a validator node you need to first get your current validator public key that was created when you ran `desmod init`. Your `desmosvalconspub` (Desmos Validator Consensus Pubkey) can be used to create a new validator by staking tokens. You can find your validator pubkey by running:

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
  --moniker="<Your moniker here>" \
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
When specifying the value of the `moniker` flag, please keep in mind this is going to be the public name associated to your validator. For this reason, it should represent your company name or something else that can easily identify you among all the other validators.  
:::

::: tip
When specifying commission parameters, the `commission-max-change-rate` is used to measure % _point_ change over the `commission-rate`. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point.
:::

::: tip
`Min-self-delegation` is a stritly positive integer that represents the minimum amount of self-delegated staking token your validator must always have. A `min-self-delegation` of 1 means your validator will never have a self-delegation lower than `1udaric`. A valdiator self delegate lower than this number will automatically be unbonded.
:::

You can confirm that you are in the validator set by using a block explorer, e.g. [Big Dipper](https://morpheus.desmos.network).

## 3. Edit the validator description
You can edit your validator's public description. This info is to identify your validator, and will be relied on by delegators to decide which validators to stake to. Make sure to provide input for every flag below. If a flag is not included in the command the field will default to empty (`--moniker` defaults to the machine name) if the field has never been set or remain the same if it has been set in the past.

The <key_name> specifies which validator you are editing. If you choose to not include certain flags, remember that the --from flag must be included to identify the validator to update.

The `--identity` can be used as to verify identity with systems like Keybase or UPort. When using with Keybase `--identity` should be populated with a 16-digit string that is generated with a [keybase.io](https://keybase.io) account. It's a cryptographically secure method of verifying your identity across multiple online networks. The Keybase API allows some block explorers to retrieve your Keybase avatar. This is how you can add a logo to your validator profile.

```bash
desmoscli tx staking edit-validator
  --moniker="choose a moniker" \
  --website="https://desmos.network" \
  --identity=6A0D65E29A4CBC8E \
  --details="To infinity and beyond!" \
  --commission-rate="0.10" \
  --chain-id=<chain_id> \
  --from=<key_name>
```

__Note__: The `commission-rate` value must adhere to the following invariants:

- Must be between 0 and the validator's `commission-max-rate`
- Must not exceed the validator's `commission-max-change-rate` which is maximum
  % point change rate **per day**. In other words, a validator can only change
  its commission once per day and within `commission-max-change-rate` bounds.

### View the validator description
View the validator's information with this command:

```bash
desmoscli query staking validator <account_desmos>
```

```bash
desmoscli tx slashing unjail \
	--from=<key_name> \
	--chain-id=<chain_id>
```

## 4. Confirm your validator is running
Your validator is active if the following command returns anything:

```bash
desmoscli query tendermint-validator-set | grep "$(desmosd tendermint show-validator)"
```

You should now see your validator in one of the Desmos explorers. You are looking for the `bech32` encoded `address` in the `~/.desmosd/config/priv_validator.json` file.

::: warning Note
To be in the validator set, you need to have more total voting power than the 100th validator.
:::
