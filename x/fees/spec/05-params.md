---
id: params
title: Parameters
sidebar_label: Parameters
slug: params
---

# Parameters

The fees module contains the following parameters:

| Key     | Type     | Example                                                                                                         |
|---------|----------|-----------------------------------------------------------------------------------------------------------------|
| MinFees | []MinFee | `[{ "message_type": "/desmos.profiles.v2.SaveProfile", "amount": [ { "amount": "100", "denom": "tokenA" } ] }]` |

## MinFees 
The `MinFees` is an array of `MinFee` object, each one made of two different fields: 

* `MessageType` (string), representing the type url of the message which the min fees refer to 
* `Amount` (`[]Coin`), representing the amount of min fees associated with the message

Inside the `MinFees`, there can only be a single entry for each `MessageType`.