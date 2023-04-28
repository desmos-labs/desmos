# ADR 022: Per-subspace token factory

## Changelog
- April 28th, 2023: First draft;

## Status

Draft

## Abstract

This ADR introduces a new feature that enables the subspace admins to create, mint and burn a subspace related
coin. 

## Context

Desmos is a social network protocol that allows users to create, share, and engage with content 
on a decentralized platform. 
Currently, to post or interact with the content in a subspace you need to pay the gas fee using DSM. 
This have the consequence of not giving any financial incentive on creating a social network on Desmos.

## Decision

We will wrap the Osmosis [Token Factory module](https://docs.osmosis.zone/osmosis-core/modules/tokenfactory/)
so that a subspace admin can create, mint and burn a coin related to a subspace.

### Type

```proto
message Post {
  uint64 subspace_id = 1;
  uint32 section_id = 2;
  
  ...skip

  google.protobuf.Timestamp last_edited_date = 13;
  
  // Owner of the post
  string owner = 14;
}
```

### `Msg` Service

```protobuf
service Msg {
    // CreateDenom allows a subspace admin to create a new coin that will have denom:
    // factor/{subspace_tresury_address}/{subdenom}.
    rpc CreateDenom(MsgCreateDenom) returns (MsgCreateDenomResponse);
    
    // Mint allows a subspace admin to mint an amount of coins of previously created denom.
    rpc Mint(MsgMint) returns (MsgMintResponse);

    // Burn allows a subspace admin to burn an amount of coins of previously created denom.
    rpc Burn(MsgBurn) returns (MsgBurnResponse);

    // SetDenomMetadata allows a subspace admin to change the metadata of previously created denom.
    rpc SetDenomMetadata(MsgSetDenomMetadata) returns (MsgSetDenomMetadataResponse);
}

// MsgChangePostOwner move a post to another subspace.
message MsgCreateDenom {
    // Address of the coin creator.
    string creator = 1;
  
    // Id of the subspace.
    uint64 subspace_id = 2;

    // Subdenom of the coin that will be created.
    string subdenom = 3;
}
// MsgCreateDenomResponse defines the Msg/MsgCreateDenom response type.
message MsgCreateDenomResponse {}

// MsgMint allow a subspace admin to mint an amount coins.
message MsgMint {
  // Address of the user that is performing the mint.
  string sender = 1;

  // Id of the subspace.
  uint64 subspace_id = 2;
  
  // Amount of coins to be minted.
  cosmos.base.v1beta1.Coin amount = 3;
}
// MsgMintResponse defines the Msg/MsgMint response type
message MsgMintResponse {}

// MsgBurn allow a subspace admin to mint an amount coins.
message MsgBurn {
  // Address of the user that is performing the burn.
  string sender = 1;

  // Id of the subspace.
  uint64 subspace_id = 2;
  
  // Amount of coins to be burned.
  cosmos.base.v1beta1.Coin amount = 3;
}
// MsgBurnResponse defines the Msg/MsgBurn response type
message MsgBurnResponse {}

// MsgBurn allow a subspace admin to mint an amount coins.
message MsgSetDenomMetadata {
  // Address of the user that is changing the coin metadata.
  string sender = 1;

  // Id of the subspace.
  uint64 subspace_id = 2;
  
  // Amount of coins to be burned.
  cosmos.bank.v1beta1.Metadata metadata = 2;
}
// MsgSetDenomMetadataResponse defines the Msg/MsgSetDenomMetadata response type
message MsgSetDenomMetadataResponse {}
```

## Consequences

### Backwards Compatibility

The solution outlined above is fully backward compatible since we are just adding a new module.

### Positive

- Enable other projects to have economic incentives on building on Desmos 

### Negative

(none known)

### Neutral

(none known)

## References