# ADR 018: Improve subspace treasury account

## Changelog
- Jan 4th, 2023: First draft;

## Status

PROPOSED

## Abstract

This ADR introduces a new treasury account structure, which guarantees the account fully controlled by the subspace.

## Context

Currently, each subspace's treasury is a third-party accounts assigned by the subspace's manager(s). The subspace itself has no control over its treasury, which can lead to the following issues:
1. subspace managers can assign a wealthy account not controlled by them as their treasury in order to scam users.;
2. if we implement a feature that allows managers to spend money from the treasury in the future, they could assign any other account in order to steal funds.

## Decision

To address the issues mentioned above, we propose implementing a new treasury account structure. The new treasury address will be generated using `authtypes.NewModuleAddress` from its subspace id when the subspace is created, and will not be able to be edited thereafter. This will ensure that the subspace's treasury is fully controlled by the subspace itself, rather than being managed by external third-party accounts.

Additionally, we will introduce a new permission and method to allow authorized users to spend funds from the treasury. 

## `Msg` Service

To prevent the treasury account from being edited later, we will remove the `treasury` field from the current `MsgCreateSubspace` and `MsgEditSubspace` messages.

Additionally, we will implement a new operation to allow authorized users to spend funds from the treasury account as follows:

```protobuf
service Msg {
    // SpendTreasury allows users who have the permission to transfer tokens out of the treasury
    rpc SpendTreasury(MsgSpendTreasury) returns (MsgSpendTreasuryResponse);
}

// MsgSpendTreasury transfers funds from the treasury of the given subspace to another address
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

## Consequences

### Backwards Compatibility

The solution outlined above is **not** backwards compatible and will require a migration script to update all existing subspaces to the new version. This script will handle the following tasks:
- migrate all subspaces to have a new treasury address generated from the its subspace's ID.

### Positive

- Ensure that the treasury account is fully controlled by the subspace

### Negative

(none known)

### Neutral

(none known)

## References
- [Issue #1057 discussion](https://github.com/desmos-labs/desmos/pull/1057#discussion_r1059423029)