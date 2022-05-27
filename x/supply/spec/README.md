<!--
order: 0
title: Supply Overview
parent:
  title: "supply"
-->

# `supply`

## Abstract 

This document specifies the supply module of the Cosmos SDK. 

The supply module exposes some query endpoints that can be used by price aggregator services such as [CoinGecko](https://coingecko.com) and [CoinMarketCap](https://coinmarketcap.com) to easily get the total and circulating supply of a token.  

## Concepts 
1. **[Concepts](01_concepts.md)**
   - [Total Supply](01_concepts.md#total-supply)
   - [Circulating Supply](01_concepts.md#circulating-supply)
2. **[Client](02_client.md)**
   - [CLI](02_client.md#cli)
   - [gRPC](02_client.md#grpc)
   - [REST](02_client.md#rest)