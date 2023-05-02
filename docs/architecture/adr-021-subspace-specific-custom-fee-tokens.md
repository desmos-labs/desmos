# ADR 021: Subspace specific custom fee tokens

## Changelog

- April 24th, 2023: First draft;

## Status

Proposed

## Abstract

This ADR introduces a mechanism that allows users to pay fees within a designated subspace using their existing service tokens.

## Context

Desmos uses subspaces to represent virtual applications built on top of our protocol. Currently, subspaces are mainly used to store content such as posts, reports, and reactions. While this is sufficient for developing DApps on Desmos, it may not be enough to convince existing applications to migrate to our platform since Desmos is lack of a mechanism that enables users to pay fees with existing service tokens.

For instance, consider Keplr's on-chain governance discussions feature, which is currently implemented using a custom backend that stores data in a centralized way. Although Keplr could easily use Desmos as their backend, they are not willing to do so because it offers no advantage to them. Migrating to Desmos would simply transfer their existing users to our chain without any benefits for them.

## Decision

To address the issue mentioned above, we propose implementing a mechanism that allows users to pay fees using the existing service tokens within the specified subspace. The mechanism must have the following features:
- Defining a minimum gas prices for non-DSM tokens that validators are willing to accept via governance proposal;
- Allowing the subspace manager(s) to update the list of fee tokens allowed in the subspace;
- Checking the on-chain minimum gas price requirement for non-DSM fees using a custom `TxChecker`.

The upcoming changes are as follows:

### Types

We will define a new `Params` type to show the `x/subspaces` parameters.

```proto
// Params contains the module parameters
message Params {
  // List of the minimum gas prices to be accepted by validators
  repeated Coin min_gas_prices = 2;
}
```

In addition, we will add a __allowed_fee_tokens__ field to current `Subspace` structure.

```proto
message Subspace {

  ...skip

  // the creation time of the subspace
  google.protobuf.Timestamp creation_time = 7;
  
  // List of fee token denoms allowed inside the subspace
  repeated string allowed_fee_tokens = 8;
}
```

### Msgs

We will allow the the following operations:
1. update subspace allowed fee tokens list by managers;
2. update `x/subspaces` parameters by governance.

```proto
service Msg {
  // UpdateSubspaceFeeTokens allows subspace managers to update the allowed tokens to be fee tokens inside the subspace
  rpc UpdateSubspaceFeeTokens(MsgUpdateSubspaceFeeTokens) returns (MsgUpdateSubspaceFeeTokensResponse);
    
  // UpdateParams defines a (governance) operation for updating the module
  rpc UpdateParams(MsgUpdateParams) returns (MsgUpdateParamsResponse);
}

// MsgUpdateSubspaceFeeTokens represents the message to be used to update a subspace fee tokens
message MsgUpdateSubspaceFeeTokens {

  // Id of the subspace where the list of allowed fee tokens will be updated
  uint64 subspace_id = 1;
    
  // List of the allowed tokens to be fee token inside the subspace
  repeated string allowed_fee_tokens = 2;
    
  // Address of the sender/signer
  string signer = 3;
}

// MsgUpdateSubspaceFeeTokensResponse represents the Msg/UpdateSubspaceFeeTokens response type
message MsgUpdateSubspaceFeeTokensResponse {}


// MsgUpdateParams is the Msg/UpdateParams request type.
message MsgUpdateParams {
  string authority = 1;
  Params params = 2;
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
message MsgUpdateParamsResponse {}
```

### Query

We will also provide a new query endpoint that enables users to check `x/subspaces` parameters.

```proto
service Query {
  // Params allows to query the module parameters
  rpc Params(QueryParamsRequest) returns (QueryParamsResponses) {
    option (google.api.http).get = "/desmos/subspaces/v3/params";
  };
}

// QueryParamsRequest is the request type for Query/Params RPC method
message QueryParamsRequest {}

// QueryParamsResponse is the response type for Query/Params RPC method
message QueryParamsResponse {
  Params params = 1;
}
```

### Custom TxFeeChecker

We will build a new `TxFeeChecker` based on the existing one that acts as follows:
1. check if the provided fee tokens are included in the list of subspace allowed fee tokens;
2. perform the following operation:
  1. if the provide fee tokens are all included subspace allowed fee tokens list, run custom `x/subspaces` TxFeeChecker;
  2. otherwise, run the default fee checker `CheckTxFeeWithValidatorMinGasPrices`.

## Consequences

### Backwards Compatibility

The solution outlined above is **not** backwards compatible and will require a migration script to update all existing subspaces to the new version. This script will handle the following tasks:
- migrate all subspaces to have a default allowed fee tokens list.

### Positive

- The subspace manager can enable users to pay fees using custom tokens inside the subspace.

### Negative

- Performing additional checks during the transaction check phase can slow down transaction processing.

### Neutral

(none known)

