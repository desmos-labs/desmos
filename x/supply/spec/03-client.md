---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client 

## CLI

A user can query the `supply` module using the CLI. 

### Query 

The `query` commands allow users to query the `supply` state. 

```
desmos query supply --help
```

##### About the divider exponent
Both the `total` and the `circulating` queries allow to specify an optional `divider exponent`.
If provided, such exponent will be used to divide the resulting amount by `10^(divider exponent)`.  

Example:
```
supply = 1.000.000
divider_exponent = 3
result = 1.000.000 / 10^3 = 1.000
```

#### total
The `total` command allows users to query the total supply of a token given a denomination and an optional divider exponent. 
If a divider exponent is provided, the resulting supply amount will be divided by `10^(divider_exponent)`.  

```bash
desmos query supply total [denom] [[divider_exponent]] [flags]
```

Example: 
```bash
desmos query supply total udsm 2
```

Example Output: 
```yaml
total_supply: "100003895600953035670"
```

#### circulating
The `circulating` command allows users to query the circulating supply of a token given a denomination. 

```bash
desmos query supply circulating [denom] [[divider_exponent]] [flags]
```

Example:
```bash
desmos query supply circulating udsm 2
```

Example Output:
```yaml
circulating_supply: "100003882303991703831"
```

## gRPC 
A user can query the `supply` module gRPC endpoints. 

### Total
The `Total` endpoint allows users to query for the total supply of a token given a denomination. 

```bash
desmos.supply.v1.Query/Total
```

Example:
```bash
grpcurl -plaintext \
  -d '{"denom": "stake", "divider_exponent": "2"}' localhost:9090 desmos.supply.v1.Query/Total
```

Example Output:
```json
{
  "totalSupply": "1000040727987145688"
}
```

### Circulating
The `Circulating` endpoint allows users to query for the circulating supply of a token given a denomination.

```bash
desmos.supply.v1.Query/Circulating
```

Example:
```bash
grpcurl -plaintext \
  -d '{"denom": "stake", "divider_exponent": "2"}' localhost:9090 desmos.supply.v1.Query/Circulating
```

Example Output:
```json
{
  "circulatingSupply": "1000040236507203206"
}
```

## REST 
A user can query the `supply` module using REST endpoints. 

### Total 
The `/total` endpoint allows users to query for the total supply of a token given a denomination. 

```bash
/supply/total/{denom}
```

Example: 
```bash
curl localhost:1317/supply/total/stake?divider-exponent=2
```

Example Output:
```json
1000040727987145688
```

### Circulating
The `/circulating` endpoint allows users to query for the circulating supply of a token given a denomination.

```bash
/supply/circulating/{denom}
```

Example:
```bash
curl localhost:1317/supply/circulating/stake?divider-exponent=2
````

Example Output:
```json
1000040236507203206
```