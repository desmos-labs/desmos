# Desmos

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
