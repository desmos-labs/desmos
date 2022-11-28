---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts 

## Total Supply 
The total supply of a token is defined as _the overall number of tokens having a given denomination that currently exists inside a chain_. 

The total supply of a token is fetched directly from the `bank` module, which properly tracks such amount each time a new token is created (due to inflation or other) or burned. 

## Circulating Supply
The circulating supply of a token is defined as _the number of tokens having a given denomination that can be transferred freely from one user to another_.

Based on this definition, the circulating supply of a token is computed using the following formula: 
```
circulating_supply = total_supply - community_pool - sum(vested_amount)
```

This is due to the fact that the following amounts are considered as non circulating: 
 
* the amount of a token inside the _community pool_, since such tokens can be transferred to a user only after a `CommunitySpendPropsal` passes. As soon as some tokens are transferred from the community pool to a user, they become immediately part of the circulating supply;
* the amount of vested tokens, since those are subject to a lock period. As soon as such period ends, they become immediately part of the circulating supply.  