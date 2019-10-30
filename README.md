# Desmos

[![CircleCI](https://circleci.com/gh/desmos-labs/desmos/tree/master.svg?style=shield)](https://circleci.com/gh/desmos-labs/desmos/tree/master)
[![codecov](https://codecov.io/gh/desmos-labs/desmos/branch/master/graph/badge.svg)](https://codecov.io/gh/desmos-labs/desmos)
[![Go Report Card](https://goreportcard.com/badge/github.com/desmos-labs/desmos)](https://goreportcard.com/report/github.com/desmos-labs/desmos)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/desmos-labs/desmos.svg)
![GitHub](https://img.shields.io/github/license/desmos-labs/desmos.svg)

Desmos social chain

## Installation

``` sh
make install
```

## Start a testnet

Run the [desmos-faucet](https://github.com/kwunyeung/desmos-faucet).

``` sh
desmosd init {moniker} --chain-id {chain id}
desmoscli keys add {your name}
desmosd add-genesis-account {your name} 100000000000desmos
desmoscli keys add {faucet}
desmosd add-genesis-account {faucet} 10000000000000phanero
desmosd gentx --amount 1000000000desmos --name {your name}
desmosd collect-gentxs
```

Edit `$HOME/.desmosd/config/genesis.json` and update the `bond_denom` from `stake` to `desmos`.

``` sh
desmosd start
```
