# ADR 009: Supply module

## Changelog

- March 21, 2022: Initial draft.

## Status

DRAFTED

## Abstract

This ADR defines the `x/supply` module which will expose a set of APIs that will be called by data aggregator websites (such as Coingecko and CoinMarketCap)
in order to fetch updated data about $DSM supply.

## Context

Currently, inside our [CoinGecko](https://www.coingecko.com/en/coins/desmos) and [CoinMarketCap](https://coinmarketcap.com/currencies/desmos/) 
some important information (current and total supply) are missing or not correctly updated. To solve this, we can implement a
series of APIs that fetch those data and can be later used by all the data aggregator websites that has the $DSM token listed.

## Decision

The APIs will be exposed in a new module called `x/supply` that will have the purpose to fetch the given coin info from different 
cosmos-SDK modules (`x/bank`, `x/distribution`, `x/staking` ) and apply some conversions on them in order to avoid displaying them
in millionth units.

### Queries

#### Total Supply
This query will:
1. Fetch the total supply of the given token `denom` from the `bank` module
2. Convert it in order to display its non-millionth value

#### Current supply
This query will:
1. Fetch the total supply of the given token `denom`
2. Fetch the total vested tokens amount
3. Fetch the community pool amount
4. Calculate the circulating supply by subtracting data of 2. and 3. from Total supply  
## Consequences

### Positive
- All the chain nodes will expose these APIs automatically after their node updates;
- No state-breaking changes introduced;
- Easily extensible.

## Test Cases

- Check API to correctly return the converted Total Supply from millionth to non-millionth representation;
- Check API to correctly return the current supply.


## References

Actual alpha branch of `x/supply`:  
https://github.com/desmos-labs/desmos/tree/leonardo/coingecko-APIs/x/supply.