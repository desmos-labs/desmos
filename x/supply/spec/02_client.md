<!--
order: 2
-->

# Client 

## CLI

A user can query the `supply` module using the CLI. 

### Query 

The `query` commands allow users to query the `supply` state. 

```
desmos query supply --help
```

#### total
The `total` command allows users to query the total supply of a token given a denomination. 

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
A user can query the `profiles` module gRPC endpoints. 

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

```
/supply/total/{denom}
```

### Circulating
The `/circulating` endpoint allows users to query for the circulating supply of a token given a denomination.

```
/supply/circulating/{denom}
```