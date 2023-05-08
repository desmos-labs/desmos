# ADR 022: Per-subspace token factory

## Changelog
- April 28th, 2023: First draft;

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

```protobuf
// Msg defines the tokenfactory module's gRPC message service.
service Msg {
  // CreateDenom allows the subspace treasury to create a new coin that will have
  // denom: factory/{subspace_treasury_address}/{subdenom}.
  rpc CreateDenom(MsgCreateDenom) returns (MsgCreateDenomResponse);

  // Mint allows the subspace treasury to mint an amount of coins of a previously
  // created denom.
  rpc Mint(MsgMint) returns (MsgMintResponse);

  // Burn allows the subspace treasury to burn an amount of coins of a previously
  // created denom.
  rpc Burn(MsgBurn) returns (MsgBurnResponse);

  // SetDenomMetadata allows the subspace treasury to change the metadata of a
  // previously created denom.
  rpc SetDenomMetadata(MsgSetDenomMetadata)
      returns (MsgSetDenomMetadataResponse);
}

// MsgCreateDenom defines the message structure for the CreateDenom gRPC service
// method. It allows the subspace treasury to create a new denom. It requires a subspace
// id and a sub denomination. The (subspace_id, sub_denomination) tuple
// must be unique and cannot be re-used.
//
// The resulting denom created is defined as
// <factory/{subspace_treasury_address}/{subdenom}>. 
message MsgCreateDenom {
  // Address of who is creating the denom.
  string sender = 1 [ (gogoproto.moretags) = "yaml:\"sender\"" ];
  
  // Subspace id where the coin can be used.
  uint64 subspace_id = 2 [ (gogoproto.moretags) = "yaml:\"subspace_id\"" ];

  // subdenom can be up to 44 "alphanumeric" characters long.
  string subdenom = 3 [ (gogoproto.moretags) = "yaml:\"subdenom\"" ];
}

// MsgCreateDenomResponse is the return value of MsgCreateDenom
// It returns the full string of the newly created denom.
message MsgCreateDenomResponse {
  // Denom of the newly created coin.
  string new_token_denom = 1
  [ (gogoproto.moretags) = "yaml:\"new_token_denom\"" ];
}

// MsgMint is the sdk.Msg type for allowing the subspace treasury to mint
// more of a token.
message MsgMint {
  // Address of who is performing the mint action.
  string sender = 1 [ (gogoproto.moretags) = "yaml:\"sender\"" ];
  
  // Subspace id where the coin has been created.
  uint64 subspace_id = 2 [ (gogoproto.moretags) = "yaml:\"subspace_id\"" ];
  
  // Amount of coins to mint.
  cosmos.base.v1beta1.Coin amount = 3 [
    (gogoproto.moretags) = "yaml:\"amount\"",
    (gogoproto.nullable) = false
  ];
  
  // Optional address to which the coins will be sent, if this is empty
  // the coins will be sent to the sender. 
  string mint_to_address = 4
  [ (gogoproto.moretags) = "yaml:\"mint_to_address\"" ];
}

// MsgMintResponse defines the Msg/MintResponse response type.
message MsgMintResponse {}

// MsgBurn is the sdk.Msg type for allowing the subspace treasury to burn
// a token.
message MsgBurn {
  // Address of who is performing the burn action.
  string sender = 1 [ (gogoproto.moretags) = "yaml:\"sender\"" ];

  // Subspace id where the coin has been created.
  uint64 subspace_id = 2 [ (gogoproto.moretags) = "yaml:\"subspace_id\"" ];

  // Amount of coins to burn.
  cosmos.base.v1beta1.Coin amount = 3 [
    (gogoproto.moretags) = "yaml:\"amount\"",
    (gogoproto.nullable) = false
  ];

  // Optional address from which the tokens will be burned, if this is empty
  // the coins will be burned from the sender. 
  string burn_from_address = 4
  [ (gogoproto.moretags) = "yaml:\"burn_from_address\"" ];
}

// MsgBurnResponse defines the Msg/BurnResponse response type.
message MsgBurnResponse {}

// MsgSetDenomMetadata is the sdk.Msg type for allowing the subspace treasury to set
// the denom's bank metadata.
message MsgSetDenomMetadata {
  // Address of who is setting the coin's metadata.
  string sender = 1 [ (gogoproto.moretags) = "yaml:\"sender\"" ];

  // Subspace id where the coin has been created.
  uint64 subspace_id = 2 [ (gogoproto.moretags) = "yaml:\"subspace_id\"" ];

  // Coin metadata to be set.
  cosmos.bank.v1beta1.Metadata metadata = 3 [
    (gogoproto.moretags) = "yaml:\"metadata\"",
    (gogoproto.nullable) = false
  ];
}

// MsgSetDenomMetadataResponse defines the response structure for an executed
// MsgSetDenomMetadata message.
message MsgSetDenomMetadataResponse {}
```

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