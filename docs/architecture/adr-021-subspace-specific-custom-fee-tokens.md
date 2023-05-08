# ADR 021: Subspace specific custom fee tokens

## Changelog

- April 24th, 2023: First draft;
- May 8th, 2023: First review;

## Status

Proposed

## Abstract

This ADR enables subspace owners to let content creators pay for transaction fees using different token denominations when creating contents within their subspace.

## Context

Desmos uses subspaces to represent virtual applications built on top of our protocol. Currently, subspaces are mainly used to store content such as posts, reports, and reactions. While these features are sufficient to develop new applications on top of Desmos, they may not be enough to convince existing applications to migrate to our platform.

For instance, let's consider an existing application that uses a centralized database in order to store users discussions. Although such application might easily use Desmos as its backend, they most likely are not going to do so because it offers no advantage to them. Migrating to Desmos would simply transfer their existing users to our chain without any benefits for them.

## Decision

To address the issue mentioned above, we propose implementing a mechanism that allows users to pay content-related fees using additional token denominations within the specified subspace. 

Ideally, the flow that will lead one subspace to accept an additional fee token is going to be the following: 

1. The subspace owner creates an on-chain governance proposal asking validators if they are fine in receiving fees paid in the new token denomination when validating transactions related to that subspace.
2. Validators will agree or reject the proposal through on-chain voting. 
3. If the proposal is accepted, the new token denom will be added to the list of additional fee token denoms that can be used to pay for fees within that subspace.
4. If the proposal is rejected, the new token denom will not be added to the list of additional fee token denoms.

### Types

To represent the list of supported fee token denoms and their default minimum prices within a subspace, we will add a new field `allowed_fee_tokens` to the existing Subspace structure.

```proto
message Subspace {

  ...skip

  // the creation time of the subspace
  google.protobuf.Timestamp creation_time = 7;
  
  // List of fee token denoms with default minimum gas prices allowed inside the subspace
  repeated Coin allowed_fee_tokens = 8;
}
```

### Msgs

We will implement the operation that allows subspace admins to updates subspace allowed fee tokens list by governance;

```proto
service Msg {
  // UpdateSubspaceFeeTokens allows subspace admins to update the allowed tokens to be fee tokens inside the subspace by governance
  rpc UpdateSubspaceFeeTokens(MsgUpdateSubspaceFeeTokens) returns (MsgUpdateSubspaceFeeTokensResponse);
}

// MsgUpdateSubspaceFeeTokens represents the message to be used to update a subspace fee tokens by governance
message MsgUpdateSubspaceFeeTokens {

  // Id of the subspace where the list of allowed fee tokens will be updated
  uint64 subspace_id = 1;
    
  // List of the allowed tokens to be fee token inside the subspace along with their default minimum prices
  repeated Coin allowed_fee_tokens = 2;
    
  // authority is the address that controls the module (defaults to x/gov unless overwritten).
  string authority = 3;
}

// MsgUpdateSubspaceFeeTokensResponse represents the Msg/UpdateSubspaceFeeTokens response type
message MsgUpdateSubspaceFeeTokensResponse {}
```

### Custom TxFeeChecker

To make it easier for validators to manage the minimum prices of fee tokens allowed within a subspace, we will develop a new `TxFeeChecker` based on the existing one, which will function as follows:
- Combine the list of minimum gas prices in the validator's local configuration with the list of allowed fee tokens and their minimum prices within the subspace.
- Follow the same process as the existing `TxFeeChecker`.

### Update `x/fees`

Since `x/fees` module may conflict with the custom fee tokens feature, as demonstrated in the following scenario:
- Governance decides to change the minimum fees of MsgCreatePost to `10dsm` using `x/fees`.
- Governance accepts the custom token `minttoken` as a fee within subspace 1 using `MsgUpdateSubspaceFeeTokens`.

Due to this conflict, `minttoken` will not be accepted as a fee within the subspace 1. Therefore, we decided to update the `CheckFees` logic of `x/fees` to ensure compatibility with this feature.

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

