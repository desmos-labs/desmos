---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---

# Validators Overview

## Introduction
[Desmos](../01-intro.md) is based on [Tendermint](https://github.com/tendermint/tendermint/tree/master/docs/introduction), which relies on a set of validators that are responsible for committing new blocks in the blockchain. These validators participate in the consensus protocol by broadcasting votes which contain cryptographic signatures signed by each validator's private key.

Validator candidates can bond their own Desmos tokens and have Desmos tokens "delegated", or staked, to them by token holders. Desmos will have 150 validators, but over time this will increase to 300 validators depends on the network performance and needs. The validators are determined by who has the most stake delegated to them — the top 150 validator candidates with the most stake will become Desmos validators.

If validators double sign, are frequently offline or do not participate in governance, their staked Desmos tokens (including the tokens of users that delegated to them) can be slashed. The penalty depends on the severity of the violation.

## Seek legal advice
Seek legal advice if you intend to run a Validator.

## Community
Discuss the finer details of being a validator on our community chat:

* [Validators Chat](https://discord.gg/J6VsHDT)
