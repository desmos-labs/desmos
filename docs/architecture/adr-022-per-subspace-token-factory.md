# ADR 022: Per-subspace token factory

## Changelog
- April 28th, 2023: First draft;
- May 8th, 2023: First review;

## Status

Draft

## Abstract

This ADR introduces a new feature that allows subspace admins to create, mint and burn new tokens. 

## Context

Desmos is a social network protocol that allows users to create, share, and engage with content on a decentralized platform. It also provides the ability to create subspaces, which represent applications built on top of Desmos.

Currently, to post or interact with content within a subspace on Desmos, users are required to pay the gas fee using DSM (Desmos native token). This means that there is no direct financial incentive for creating a social network on Desmos itself.

By allowing subspace admins to mint custom tokens, we are going to enable the implementation of a custom tokenomic system within the subspace. This means that when users create or interact with content stored within a particular subspace, they may earn or spend these custom tokens. This will provide a new level of flexibility and incentive for users to actively participate within the subspaces.

## Decision

We will wrap the CosmWasm [Token Factory module](https://github.com/CosmWasm/token-factory/blob/main/x/tokenfactory)
with the following modifications:
1. Instead of using the token creator address to compose the coin denom, we will use the subspace treasury address: `factory/{trasury_address}/subdenom`;
2. Only the subspace treasury will be able to perform the admin-related operations.
3. The `CreateDenom` action will burn `dsm` instead of send them to the community pool.  


### `Msg` Service
The messages used from this module will be the same as [CosmWasm](https://github.com/CosmWasm/token-factory/blob/main/proto/osmosis/tokenfactory/v1beta1/tx.proto)
with the addition of a `subspace_id` field and without the `MsgChangeAdmin` since in our case the admin will be the 
subspace treasury and it will not be allowed to change it.

## Consequences

### Backwards Compatibility

The solution outlined above is fully backward compatible since we are just adding a new module.

### Positive

- Enable other projects to have economic incentives on building on Desmos.

### Negative

(none known)

### Neutral

(none known)

## References