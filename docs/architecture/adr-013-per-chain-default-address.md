# ADR 013: Per-chain default external address

## Changelog

- May 18th, 2022: Initial draft;

## Status

ACCEPTED Not Implemented

## Abstract

This ADR contains the specification of _per-chain default external addresses_, a new feature that will allow users to specify which external address should be used as the default one for each external chain they have linked to their Desmos profile.

## Context

Currently, Desmos allows users to link external chain accounts to their profile by specifying a chain name and an external address associated to such chain. For each chain, each user can link multiple external addresses to the same profile. Although this allows for greater extensibility and usability, there are some occasions in which it might make it harder to properly use this feature. For example, one application might want to deal with a single external address per chain. Right now this is not easy to do as there are many ways in which that single chain link can be chosen among multiple ones (e.g. the first created, or the last one). 

## Decision

We will allow users to specify a _default address_ for each external chain they have connected to their Desmos profile. Initially, when creating an external chain link, the first link will be set as the default one. Later, each user will be able to set a specific link as the default one to be used when dealing with a specific chain.

### Store

Each default chain link will be stored inside the state using the following key:
```
DefaultChainLinkPrefix | UserAddress | ChainName | -> ExternalAddress 
```

This will allow to easily get a default external address based on the user address and chain name, which is going to be the most used query. It will also allow to easily iterate over all default external addresses for each user. 

### `Msg` Service
In order to allow each user to set their own custom default external addresses, we will need to implement a new message. 

```protobuf
syntax = "proto3";

// Msg defines the profiles Msg service.
service Msg {
  // SetDefaultExternalAddress allows to set a specific external address as the default one for a given chain
  rpc SetDefaultExternalAddress(MsgSetDefaultExternalAddress) returns (MsgSetDefaultExternalAddressResponse);
}

// MsgSetDefaultExternalAddress represents the message used to set a default address for a specific chain
message MsgSetDefaultExternalAddress {
  // Name of the chain for which to set the default address
  string chain_name = 1;
  
  // Address to be set as the default one
  string external_address = 2;
  
  // User signing the message
  string signer = 3;
}

// MsgSetDefaultExternalAddressResponse represents the Msg/SetDefaultExternalAddress response type
message MsgSetDefaultExternalAddressResponse {}
```

### `Query` Service
In order to allow clients to easily query for default chain links, we will implement a new query.

```protobuf
syntax = "proto3";

// Query defines the gRPC querier service
service Query {
  //  DefaultExternalAddresses queries the default addresses associated to the given user and (optionally) chain name
  rpc DefaultExternalAddresses(QueryDefaultExternalAddressesRequest) returns (QueryDefaultExternalAddressesResponse) {
    option (google.api.http).get = "/desmos/profiles/v2/default-addresses";
  }
}

// QueryDefaultExternalAddressesRequest is the request type for Query/DefaultExternalAddresses RPC method
message QueryDefaultExternalAddressesRequest {
  // User for which to query the default addresses
  string user = 1;
  
  // (optional) Chain name to query the default address for
  string chain_name = 2;
  
  // pagination defines an optional pagination for the request.
  cosmos.base.query.v1beta1.PageRequest pagination = 3;
}

// QueryDefaultExternalAddressesResponse is the response type for Query/DefaultExternalAddresses RPC method
message QueryDefaultExternalAddressesResponse {
  // List of default addresses, each one represented by the associated chain link 
  repeated ChainLink chain_links = 1 [ (gogoproto.nullable) = false ];
  
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}
```

## Consequences

### Backwards Compatibility

As the only major change will be to introduce a new set of store keys, this change is completely backward compatible.

### Positive

- Allow to get a user's default external addresses, allowing for more use cases to be implemented

### Negative

- Increase storage space by introducing a new set of keys

### Neutral

{neutral consequences}

## References

- Issue [#853](https://github.com/desmos-labs/desmos/issues/853).