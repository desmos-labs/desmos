# ADR 018: Improve subspace treasury account

## Changelog
- Jan 4th, 2023: First draft;
- Jan 9th, 2023: First review;

## Status

PROPOSED

## Abstract

This ADR introduces a new treasury account structure, which guarantees the account is fully controlled by the subspace admins.

## Context

Currently, each subspace's treasury is a third-party accounts assigned by the subspace's manager(s). The subspace itself has no control over its treasury, which can lead to the following issues:
1. subspace managers can assign a wealthy account not controlled by them as their treasury in order to scam users;
2. if we implement a feature that allows managers to spend money from the treasury in the future, they could assign any other account in order to steal funds.

## Decision

To address the issues mentioned above, we propose implementing a new treasury account structure. The new treasury address will be generated using `authtypes.NewModuleAddress` from its subspace id when the subspace is created, and will not be able to be edited thereafter. This will ensure that the subspace's treasury is fully controlled by the subspace itself, rather than being managed by external third-party accounts.

Additionally, we will introduce a new permission and method that will grant users to perform operations on the treasury.

## `Msg` Service

To prevent the treasury account from being edited later, we will remove the `treasury` field from the current `MsgCreateSubspace` and `MsgEditSubspace` messages.

Additionally, we will implement new methods to grant or revoke treasury authorization for users to perform operations on the treasury. These methods will be as follows:

```protobuf
service Msg {
    // GrantTreasuryAuthorization allows managers who have the permission to grant a treasury authorization to a user
    rpc GrantTreasuryAuthorization(MsgGrantTreasuryAuthorization) returns (MsgGrantTreasuryAuthorizationResponse);

    // RevokeTreasuryAuthorization allows managers who have the permission to revoke an existing treasury authorization
    rpc RevokeTreasuryAuthorization(MsgRevokeTreasuryAuthorization) returns
    (MsgRevokeTreasuryAuthorizationResponse);
}

// MsgGrantTreasuryAuthorization grants an authorization of the treasury to a user
message MsgGrantTreasuryAuthorization {
    // Id of the subspace where the granter grants a treasury authorization
    uint64 subspace_id = 1;
    // Address of the user granting a treasury authorization
    string granter = 2;
    // Address of the user who is being granted a treasury authorization
    string grantee = 3;
    // Grant represents the authorization to execute the provided methods
    cosmos.authz.v1beta1.Grant grant = 4 [(gogoproto.nullable) = false];
}

// MsgGrantTreasuryAuthorizationResponse defines the Msg/MsgGrantTreasuryAuthorization response type
message MsgGrantTreasuryAuthorizationResponse{}

// MsgRevokeTreasuryAuthorization revokes an authorization of the treasury from a user
message MsgRevokeTreasuryAuthorization {
    // Id of the subspace where the granter revokes a treasury authorization
    uint64 subspace_id = 1;
    // Address of the user revoking a treasury authorization
    string granter = 2;
    // Address of the user who is being revoked a treasury authorization
    string grantee = 3;
}

// MsgRevokeTreasuryAuthorizationResponse defines the Msg/MsgRevokeTreasuryAuthorization response type
message MsgRevokeTreasuryAuthorizationResponse{}
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