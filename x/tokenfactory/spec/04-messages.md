---
id: messages
title: Messages
sidebar_label: Messages
slug: messages
---

# Msg Service

## Msg/CreateDenom
A token denomination can be created with a `MsgCreateDenom`. In order to prevent spam, a fee is required to be paid from the subspace treasury in order to create a new denomination. Such fee's amount is set inside the module's params and can be updated by the chain governance if the community decides to do so.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/tokenfactory/v1/msgs.proto#L55-L78
```

It's expected to fail if:
* the subspace does not exist;
* the provided denomination is invalid;
* the sender does not have the permission to mint tokens inside the subspace;
* the subspace treasury does not have enough balance to pay the creation fee.

## Msg/Mint
The admin of a token denomination can mint more tokens using a `MsgMint`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/tokenfactory/v1/msgs.proto#L88-L116
```

It's expected to fail if:
* the subspace does not exist;
* the provided amount does not match the existing denomination;
* the sender does not have the permissions to mint more tokens of the provided denomination.

## Msg/Burn
The admin of a token denomination can burn tokens using a `MsgBurn`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/tokenfactory/v1/msgs.proto#L123-L151
```

It's expected to fail if:
* the subspace does not exist;
* the sender does not have the permissions to burn tokens of the provided denomination.

## Msg/SetDenomMetadata
The admin of a token denomination can set the metadata of the denomination using a `MsgSetDenomMetadata`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/tokenfactory/v1/msgs.proto#L158-L179
```

It's expected to fail if:
* the subspace does not exist;
* the sender does not have the permissions to set the metadata of the provided denomination;
* the provided metadata is invalid.


## Msg/UpdateParams
The `MsgUpdateParams` can be used to update the overall module's params.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/tokenfactory/v1/msgs.proto#L191-L213
```

It's expected to fail if:
* the provided authority is not the address of the `x/gov` module;
* the provided params are invalid.