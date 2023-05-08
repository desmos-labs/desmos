# ADR 022: Per-subspace token factory

## Changelog
- April 28th, 2023: First draft;

## Status

Draft

## Abstract

This ADR introduces a new feature that allows subspace admins to create, mint and burn new tokens. 

## Context

Desmos is a social network protocol that allows users to create, share, and engage with content 
on a decentralized platform. 
Currently, to post or interact with the content in a subspace you need to pay the gas fee using DSM. 
This have the consequence of not giving any financial incentive on creating a social network on Desmos.

## Decision

We will wrap the Osmosis [Token Factory module](https://docs.osmosis.zone/osmosis-core/modules/tokenfactory/)
with the following modifications:
1. Instead of use the token creator address to compose the coin denom we will use the subspace treasury address
   as following: `factory/{trasury_address}/subdenom`;
2. All the operations of `CreateDenom`, `Mint`, `Burn`, `SetDenomMetadata`, `SetBeforeSendHook` and
   `ForceTransfer` can only be performed by subspace treasury;
3. The `CreateDenom` action will burn the coins instead of send them to the community pool.  

With this module subspace admins will be able to create a coin that can be used to pay for 
subspace related transactions after a governance proposal.

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

  // SetBeforeSendHook allows the subspace treasury to set the address of a CosmWasm contract
  // that will be called before the coins are sent to track or block the send action.
  rpc SetBeforeSendHook(MsgSetBeforeSendHook)
      returns (MsgSetBeforeSendHookResponse);

  // ForceTransfer allows the subspace treasury to force transfer some coins from one
  // address to another.
  rpc ForceTransfer(MsgForceTransfer) returns (MsgForceTransferResponse);
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

// MsgSetBeforeSendHook is the sdk.Msg type for allowing the subspace treasury to
// assign a CosmWasm contract to call with a BeforeSend hook.
message MsgSetBeforeSendHook {
  // Address of who is setting the hook.
  string sender = 1 [ (gogoproto.moretags) = "yaml:\"sender\"" ];

  // Subspace id where the coin has been created.
  uint64 subspace_id = 2 [ (gogoproto.moretags) = "yaml:\"subspace_id\"" ];

  // Coin denom.
  string denom = 3 [ (gogoproto.moretags) = "yaml:\"denom\"" ];
  
  // Address of the CosmWasm address that will be called.
  string cosmwasm_address = 4
  [ (gogoproto.moretags) = "yaml:\"cosmwasm_address\"" ];
}

// MsgSetBeforeSendHookResponse defines the response structure for an executed
// MsgSetBeforeSendHook message.
message MsgSetBeforeSendHookResponse {}

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

// MsgForceTransfer allows the subspace treasury to transfer some coins from one
// address to another one.
message MsgForceTransfer {
  // Address of who is performing the force transfer.
  string sender = 1 [ (gogoproto.moretags) = "yaml:\"sender\"" ];

  // Subspace id where the coin has been created.
  uint64 subspace_id = 2 [ (gogoproto.moretags) = "yaml:\"subspace_id\"" ];
  
  // Amount of coin to transfer.
  cosmos.base.v1beta1.Coin amount = 3 [
    (gogoproto.moretags) = "yaml:\"amount\"",
    (gogoproto.nullable) = false
  ];
  
  // Address from which the coins will be taken.
  string transfer_from_address = 4
  [ (gogoproto.moretags) = "yaml:\"transfer_from_address\"" ];
  
  // Address where will be sent the coins.
  string transfer_to_address = 5
  [ (gogoproto.moretags) = "yaml:\"transfer_to_address\"" ];
}

// MsgBurnResponse defines the Msg/BurnResponse response type.
message MsgForceTransferResponse {}
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