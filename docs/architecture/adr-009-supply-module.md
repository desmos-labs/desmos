# ADR 009: Supply module

## Changelog

- March 21, 2022: Initial draft;
- March 21, 2022: First review.

## Status

DRAFTED

## Abstract

This ADR defines the `x/supply` module which will expose a set of APIs that can be called by data aggregator websites (such as CoinGecko and CoinMarketCap) in order to fetch updated data about a specific supply.

## Context

Currently, inside our [CoinGecko](https://www.coingecko.com/en/coins/desmos) and [CoinMarketCap](https://coinmarketcap.com/currencies/desmos/) some important information about current and total supply are missing or not correctly updated. To solve this, we can implement a series of APIs that read those data from the chain. Data aggregator websites that have the token listed can later use them.

## Decision

The APIs will be exposed in a new module called `x/supply` that will have the purpose to fetch the given information using different cosmos-SDK modules (namely `x/bank`, `x/distribution`, `x/staking`) and apply some conversions on them in order to display them the best way possible for the client.

### Queries

All the following APIs will have a custom param named `multiplier` that allows to set the multiplier to be used when returning the values. A `multiplier` of `0` will identify the whole token amount, while a multiplier of `1.000.000` will make the result display the values in millionth of units.

#### Total Supply

This query will fetch the total supply of a given token `denom`.

#### Current supply

This query will return the circulating supply by subtracting the total vested tokens amount and community pool amount from the total supply of the given token.

## Consequences

### Positive
- All the chain nodes will expose these APIs automatically after their node updates;
- No state-breaking changes introduced;
- Easily extensible.

## Test Cases

- Check API to correctly return the converted Total Supply from millionth to non-millionth representation;
- Check API to correctly return the current supply.

## References

Related [Issue #733](https://github.com/desmos-labs/desmos/issues/773).