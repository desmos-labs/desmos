---
id: ante-handlers
title: Ante Handlers
sidebar_label: Ante Handlers
slug: ante-handlers
---

# Ante Handlers

The `fees` module presently has no transaction handlers of its own, but does expose the special `AnteHandler`, used for performing a validity check on a transaction, such that it could be thrown out of the mempool.
The `AnteHandler` can be seen as a set of decorators that check transactions within the current context, per [ADR 010](https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-alpha1/docs/architecture/adr-010-modular-antehandler.md).

Note that the `AnteHandler` is called on both `CheckTx` and `DeliverTx`, as Tendermint proposers presently have the ability to include in their proposed block transactions which fail `CheckTx`.

## Decorators

The fees module provides the following `AnteDecorator`:

* `MinFeeDecorator`: Checks if the `tx` fee is greater or equal to the minimum fee amount computed looking the messages present inside the transaction itself.