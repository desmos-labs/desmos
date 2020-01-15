# Validators Overview

## Introduction
[Desmos](../README.md) is based on [Tendermint](https://github.com/tendermint/tendermint/tree/master/docs/introduction), which relies on a set of validators that are responsible for committing new blocks in the blockchain. These validators participate in the consensus protocol by broadcasting votes which contain cryptographic signatures signed by each validator's private key.

Validator candidates can bond their own Desmos tokens and have Desmos tokens "delegated", or staked, to them by token holders. Desmos will have 150 validators, but over time this will increase to 300 validators depends on the network performance and needs. The validators are determined by who has the most stake delegated to them — the top 150 validator candidates with the most stake will become Desmos validators.

If validators double sign, are frequently offline or do not participate in governance, their staked Desmos tokens (including the tokens of users that delegated to them) can be slashed. The penalty depends on the severity of the violation.

## Hardware
There currently exists no appropriate cloud solution for validator key management. This may change in 2020 when cloud SGX becomes more widely available. For this reason, validators must set up a physical operation secured with restricted access. A good starting place, for example, would be co-locating in secure data centers.

Validators should expect to equip their datacenter location with redundant power, connectivity, and storage backups. Expect to have several redundant networking boxes for fiber, firewall and switching and then small servers with redundant hard drive and failover. Hardware can be on the low end of datacenter gear to start out with.

We anticipate that network requirements will be low initially. The current testnet requires minimal resources. Then bandwidth, CPU and memory requirements will rise as the network grows. Large hard drives are recommended for storing years of blockchain history.

## Seek legal advice
Seek legal advice if you intend to run a Validator.

## Community
Discuss the finer details of being a validator on our community chat:

* [Validator Chat](https://discord.gg/J6VsHDT)
