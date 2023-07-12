---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client

## CLI 

A user can interact with the `x/tokenfactory` module using the following CLI commands.

### Query

The `query` command allows users to query the state of the module.

```bash
desmos query tokenfactory --help
```

#### params 
The `params` query command allows users to query for the module's parameters.

```bash
desmos query tokenfactory params [flags]
```

Example: 
```bash
desmos query tokenfactory params
```

Example output: 
```yaml
params:
  denom_creation_fee:
  - amount: "10000000000"
    denom: udaric
```

#### subspace-denoms
The `subspace-denoms` query command allows users to query for the denoms that have been created from within a given subspace.

```bash
desmos query tokenfactory subspace-denoms [subspace-id] [flags]
```

Example: 
```bash
desmos query tokenfactory subspace-denoms 1
```

Example output: 
```yaml
denoms:
- factory/desmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat3xsfpa/test
```

## gRPC
Users can query the `tokenfactory` module gRPC endpoints.

### Params
The `Params` gRPC endpoint allows users to query the module's parameters.

```
desmos.tokenfactory.v1.Query/Params
```

Example: 
```bash
grpcurl -plaintext localhost:9090 desmos.tokenfactory.v1.Query/Params
```

Example output:
```json
{
  "params": {
    "denomCreationFee": [
      {
        "denom": "udaric",
        "amount": "10000000000"
      }
    ]
  }
}
```

### SubspaceDenoms
The `SubspaceDenoms` gRPC endpoint allows users to query the denoms that have been created from within a given subspace.

```
desmos.tokenfactory.v1.Query/SubspaceDenoms
```

Example: 
```bash
grpcurl -plaintext -d '{"subspace_id":1}' localhost:9090 desmos.tokenfactory.v1.Query/SubspaceDenoms
```

Example output: 
```json
{
  "denoms": [
    "factory/desmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat3xsfpa/test"
  ]
}
```

## REST
A user can query the `tokenfactory` module REST endpoints.

### Params
The `Params` REST endpoint allows users to query the module's parameters.

```
/desmos/tokenfactory/v1/params
```

### SubspaceDenoms
The `SubspaceDenoms` REST endpoint allows users to query the denoms that have been created from within a given subspace.

```
/desmos/tokenfactory/v1/subspaces/{subspace_id}/denoms
```