# Setting up Tendermint KMS + Softsign

::: danger Warning
KMS and Ledger Tendermint app are currently work in progress. Details may vary. Use with care under your own risk.
:::

## Config file

### Create the folders and file
```bash
mkdir -p kms/home
cd kms/home
nano tmkms.toml
```

### Add your configuration
```
# Example Tendermint KMS configuration file

## Chain Configuration

[[chain]]
id = "<chain-id>"
key_format = { type = "bech32", account_key_prefix = "desmospub", consensus_key_prefix = "desmosvalconspub" }
state_file = "/root/kms/home/state/cosmoshub-3-consensus.json"

## Signing Provider Configuration

### Software-based Signer Configuration

[[providers.softsign]]
chain_ids = ["<chain-id>"]
key_type = "consensus"
path = "/root/kms/home/secrets/validator-consensus.key"

## Validator Configuration

[[validator]]
chain_id = "morpheus-apollo-1"
addr = "tcp://127.0.0.1:26658"
secret_key = "/root/kms/home/secrets/kms-identity.key"
protocol_version = "v0.34"
reconnect = true

```

### Get the Identity key
 
```bash
desmos query staking validator <valoperaddress>
```

You will get a similar response:
```
commission:
  commission_rates:
    max_change_rate: "0.050000000000000000"
    max_rate: "0.500000000000000000"
    rate: "0.050000000000000000"
  update_time: "2021-06-08T13:56:42.931427534Z"
consensus_pubkey:
  '@type': /cosmos.crypto.ed25519.PubKey
  key: eUhoKzRsVUhPMDlvUWdjWmo1RmNtODFRqTT0=
delegator_shares: "9999999999.000080008000800080"
description:
  details: ""
  identity: 12FA04A22E47GN17
  moniker: testman
  security_contact: ""
  website: ""
jailed: false
min_self_delegation: "1"
operator_address: desmosvaloper1...
status: BOND_STATUS_BONDED
tokens: "9999999999"
unbonding_height: "617524"
unbonding_time: "2021-05-13T05:05:09.783549624Z"

```

Copy the `consensus_pubkey` `key` value and put it inside the `kms-identity.key` file

### Import the private validator key

```bash
cd ~/.desmos/config

tmkms softsign import priv_validator_key.json "/root/kms/home/secret/morpheus-apollo-1.consensus.key"
```

### Chain configuration

Now you need to enable KMS access by editing .desmos/config/config.toml.   
In this file, modify `priv_validator_laddr` to create a listening address/port or a unix socket in desmos.

```
...
# TCP or UNIX socket address for Tendermint to listen on for
# connections from an external PrivValidator process
priv_validator_laddr = "tcp://127.0.0.1:26658"
...
```

### Start the tmkms

```bash
tmkms start -c ~/.tmkms/tmkms.toml
```


If you've setup everything properly you should see a log like this:
```bash
2021-06-09T14:23:51.525184Z  INFO tmkms::commands::start: tmkms 0.10.1 starting up...
2021-06-09T14:23:51.525380Z  INFO tmkms::keyring: [keyring:softsign] added consensus Ed25519 key: desmosvalconspub1zcjduepqepu8acj4qua576zzquvcly2un0xkzhwh0ehvgmx8gxgl34zhkceskthfp6
2021-06-09T14:23:51.526030Z  INFO tmkms::connection::tcp: KMS node ID: 7489C3177FB578314099657A7B56CC241144CCC8
2021-06-09T14:23:51.527591Z  INFO tmkms::session: [morpheus-apollo-1@tcp://127.0.0.1:26659] connected to validator successfully
2021-06-09T14:23:51.527625Z  WARN tmkms::session: [morpheus-apollo-1@tcp://127.0.0.1:26659]: unverified validator peer ID! (6CA3A1674B1AE6774D1BA6E100D559C5BF80F82B)
2021-06-09T14:35:20.416891Z  INFO tmkms::session: [morpheus-apollo-1@tcp://127.0.0.1:26659] signed PreCommit:78AD7099DE at h/r/s 609456/0/2 (0 ms)
2021-06-09T14:35:25.792444Z  INFO tmkms::session: [morpheus-apollo-1@tcp://127.0.0.1:26659] signed PreVote:2DA528546B at h/r/s 609457/0/1 (0 ms)
2021-06-09T14:35:26.126467Z  INFO tmkms::session: [morpheus-apollo-1@tcp://127.0.0.1:26659] signed PreCommit:2DA528546B at h/r/s 609457/0/2 (0 ms)
2021-06-09T14:35:31.529730Z  INFO tmkms::session: [morpheus-apollo-1@tcp://127.0.0.1:26659] signed PreVote:AD37ACB851 at h/r/s 609458/0/1 (0 ms)
2021-06-09T14:35:31.793969Z  INFO tmkms::session: [morpheus-apollo-1@tcp://127.0.0.1:26659] signed PreCommit:AD37ACB851 at h/r/s 609458/0/2 (0 ms)
```