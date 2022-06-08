---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts 

## Minimum fees
A minimum fee represents the minimum amount of `sdk.Coins` that must be present inside the transaction fees when broadcasting a message of a given type. Setting a minimum fee of `100token` for a message types means that users will need to pay at least `100token` for each message of that type they broadcast. Trying to broadcast a transaction with multiple message of such types will require the user to pay `n * min_fee` fees or above. 

For each message type, there can only be a single minimum fee amount that needs to be paid at any given time. 

A single minimum fee amount can be made of multiple coin amounts, so that if a minimum fee amount is set to `100tokenA,50tokenB` this means that for each message of such type the user will have to pay at least 100 `tokenA` **and** 50 `tokenB` to make sure the transaction is valid. 

If a transaction contains multiple messages of different kinds, each one having a custom minimum fee amount, the overall transaction fee must be greater or equal to the sum of all the minimum fee amounts. So that if a transaction contains one message which minimum fee amount is `100tokenA` and other one which minimum fee is `100tokenB`, the overall transaction fee will have to be `100tokenA + 100tokenB` or greater.

Failing to provide a transaction fee amount that is not enough to satisfy all the minimum fee requirements will lead to an invalid transaction.  