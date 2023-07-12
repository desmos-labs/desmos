---
id: events
title: Events
sidebar_label: Events
slug: events
---

# Events

The tokenfactory module emits the following events:

## Handlers

### MsgCreateDenom

| **Type**     | **Attribute Key** | **Attribute Value**                   | 
|:-------------|:------------------|:--------------------------------------|
| create_denom | subspace_id       | {subspaceID}                          |
| create_denom | creator           | {userAddress}                         |
| create_denom | new_token_denom   | {newTokenDenom}                       |
| message      | module            | subspaces                             |
| message      | action            | desmos.tokenfactory.v1.MsgCreateDenom |
| message      | sender            | {userAddress}                         |

## MsgMint

| **Type** | **Attribute Key** | **Attribute Value**            | 
|:---------|:------------------|:-------------------------------|
| tf_mint  | subspace_id       | {subspaceID}                   |
| tf_mint  | mint_to_address   | {subspaceTreasuryAddress}      |
| tf_mint  | amount            | {tokenAmount}                  |
| message  | module            | subspaces                      |
| message  | action            | desmos.tokenfactory.v1.MsgMint |
| message  | sender            | {userAddress}                  |

## MsgBurn

| **Type** | **Attribute Key** | **Attribute Value**            | 
|:---------|:------------------|:-------------------------------|
| tf_burn  | subspace_id       | {subspaceID}                   |
| tf_burn  | burn_from_address | {subspaceTreasuryAddress}      |
| tf_burn  | amount            | {tokenAmount}                  |
| message  | module            | subspaces                      |
| message  | action            | desmos.tokenfactory.v1.MsgBurn |
| message  | sender            | {userAddress}                  |

## MsgSetDenomMetadata

| **Type**           | **Attribute Key** | **Attribute Value**                        | 
|:-------------------|:------------------|:-------------------------------------------|
| set_denom_metadata | subspace_id       | {subspaceID}                               |
| set_denom_metadata | denom             | {denom}                                    |
| set_denom_metadata | denom_metadata    | {denomMetadata}                            |
| message            | module            | subspaces                                  |
| message            | action            | desmos.tokenfactory.v1.MsgSetDenomMetadata |
| message            | sender            | {userAddress}                              |

## MsgUpdateParams

| **Type** | **Attribute Key** | **Attribute Value**                    | 
|:---------|:------------------|:---------------------------------------|
| message  | module            | subspaces                              |
| message  | action            | desmos.tokenfactory.v1.MsgUpdateParams |
| message  | sender            | {userAddress}                          |
