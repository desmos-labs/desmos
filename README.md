# Desmos

[![CircleCI](https://circleci.com/gh/desmos-labs/desmos/tree/master.svg?style=shield)](https://circleci.com/gh/desmos-labs/desmos/tree/master)
[![codecov](https://codecov.io/gh/desmos-labs/desmos/branch/develop/graph/badge.svg)](https://codecov.io/gh/desmos-labs/desmos/branch/develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/desmos-labs/desmos)](https://goreportcard.com/report/github.com/desmos-labs/desmos)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/desmos-labs/desmos.svg)
![GitHub](https://img.shields.io/github/license/desmos-labs/desmos.svg)

Desmos social chain

## Installation
``` sh
make install
```

## Start a testnet

### Single node, manual
Run the [desmos-faucet](https://github.com/desmos-labs/desmos-faucet).

``` sh
desmosd init {moniker} --chain-id {chain id}
desmoscli keys add {your name}
desmosd add-genesis-account {your name} 100000000000desmos
desmoscli keys add {faucet}
desmosd add-genesis-account {faucet} 10000000000000phanero
desmosd gentx --amount 1000000000desmos --name {your name}
desmosd collect-gentxs
```

``` sh
desmosd start
```

### Multi node, automatic
To start a 4 node testnet run:

```
make localnet-start
```

This command creates a 4-node network using the `desmoslabs/desmosdnode` image.
The ports for each node are found in this table:

| Node ID | P2P Port | RPC Port |
| --------|-------|------|
| `desmosdnode0` | `26656` | `26657` |
| `desmosdnode1` | `26659` | `26660` |
| `desmosdnode2` | `26661` | `26662` |
| `desmosdnode3` | `26663` | `26664` |

To update the binary, just rebuild it and restart the nodes:

```
make build-linux localnet-start
```

#### Configuration

The `make localnet-start` creates files for a 4-node testnet in `./build` by
calling the `desmosd testnet` command. This outputs a handful of files in the
`./build` directory:

```bash
$ tree -L 2 build/
build/
├── desmoscli
├── desmosd
├── gentxs
│   ├── node0.json
│   ├── node1.json
│   ├── node2.json
│   └── node3.json
├── node0
│   ├── desmoscli
│   │   ├── key_seed.json
│   │   └── keys
│   └── desmosd
│       ├── ${LOG:-desmosd.log}
│       ├── config
│       └── data
├── node1
│   ├── desmoscli
│   │   └── key_seed.json
│   └── desmosd
│       ├── ${LOG:-desmosd.log}
│       ├── config
│       └── data
├── node2
│   ├── desmoscli
│   │   └── key_seed.json
│   └── desmosd
│       ├── ${LOG:-desmosd.log}
│       ├── config
│       └── data
└── node3
    ├── desmoscli
    │   └── key_seed.json
    └── desmosd
        ├── ${LOG:-desmosd.log}
        ├── config
        └── data
```

Each `./build/nodeN` directory is mounted to the `/desmosd` directory in each container.

#### Logging
Logs are saved under each `./build/nodeN/desmosd/desmos.log`. You can also watch logs
directly via Docker, for example:

```
docker logs -f desmosdnode0
```

#### Keys & Accounts
To interact with `desmoscli` and start querying state or creating txs, you use the
`desmoscli` directory of any given node as your `home`, for example:

```shell
desmoscli keys list --home ./build/node0/desmoscli
```

Now that accounts exists, you may create new accounts and send those accounts
funds!

**Note**: Each node's seed is located at `./build/nodeN/desmoscli/key_seed.json` and can be restored to the CLI using the `desmoscli keys add --restore` command
