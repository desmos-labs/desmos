# ADR 018: Improve subspace treasury account

## Changelog
- Jan 4th, 2023: First draft;

## Status

DRAFT

## Abstract

This ADR introduces a new treasury account structure, whish guarantee the account 100% controlled by the subspace.

## Context

Currently, the subspace address is a third-party account assigned by the subspace manager(s). The subspace has no methods to control their treasury. It would cause the issues below:
1. the managers can assign a rich account no controlled by them as their treasury in order to scam their users;
2. if we implement the feature to allow managers to spend money from treasury account in the future, the managers can assign any other account to thieve money.

## Decision

In order to solve the issues above, we will implement a new treasury account structure. The new treasury address will be generated from the subspace with `authtypes.NewModuleAddress` when the subspace is created, and can not be later edited. In addition, we need a new permission and method to allow user to spend from the treasury.

## `Msg` Service

To disallow the later edit the treasury account, we need to remove the `treasury` field in the current `MsgCreateSubspace` and `MsgEditSubspace`. The new ones will be like:

```protobuf
// MsgCreateSubspace represents the message used to create a subspace
message MsgCreateSubspace {
  // Name of the subspace
  string name = 1 [ (gogoproto.moretags) = "yaml:\"name\"" ];

  // (optional) Description of the subspace
  string description = 2 [ (gogoproto.moretags) = "yaml:\"description\"" ];

  // (optional) Owner of this subspace. If not specified, the creator will be
  // the default owner.
  string owner = 3 [ (gogoproto.moretags) = "yaml:\"owner\"" ];

  // Address creating the subspace
  string creator = 4 [ (gogoproto.moretags) = "yaml:\"creator\"" ];
}

// MsgEditSubspace represents the message used to edit a subspace fields
message MsgEditSubspace {
  // Id of the subspace to edit
  uint64 subspace_id = 1 [
    (gogoproto.customname) = "SubspaceID",
    (gogoproto.moretags) = "yaml:\"subspace_id\""
  ];

  // New name of the subspace. If it shouldn't be changed, use [do-not-modify]
  // instead.
  string name = 2 [ (gogoproto.moretags) = "yaml:\"name\"" ];

  // New description of the subspace. If it shouldn't be changed, use
  // [do-not-modify] instead.
  string description = 3 [ (gogoproto.moretags) = "yaml:\"description\"" ];

  // New owner of the subspace. If it shouldn't be changed, use [do-not-modify]
  // instead.
  string owner = 4 [ (gogoproto.moretags) = "yaml:\"owner\"" ];

  // Address of the user editing the subspace
  string signer = 5 [ (gogoproto.moretags) = "yaml:\"signer\"" ];
}

// MsgEditSubspaceResponse defines the Msg/EditSubspace response type
message MsgEditSubspaceResponse {}
```
In addition, we need implement a new operation to allow managers to spend money from treasury account bellow:

```protobuf
service Msg {
    // SpendTreasury allows user who has the spend permission to transfer tokens out from the treasury
    rpc SpendTreasury(MsgSpendTreasury) returns (MsgSpendTreasuryResponse);
}

// MsgSpendTreasury transfers the amount of money from the treasury of the given subspace money to other address
message MsgSpendTreasury {
    // Id of the subspace where the spender transfer money out from the treasury
    uint64 subspace_id = 1;
    // Address of the destination
    string to_address = 2;
    // Address who spends money
    string spender = 3;
    // Amount of money spent by spender
    repeated Coins amount = 4; 
}

// MsgSpendTreasuryResponse defines the Msg/MsgSpendTreasury response type.
message MsgSpendTreasuryResponse{}
```

