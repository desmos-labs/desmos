# ADR 017: Subspace fee grant

## Changelog

- Nov 14th, 2022: First draft;
- Nov 21th, 2022: First review.

## Status

PROPOSED

## Abstract

This ADR introduces a new subspace-based fee grant method, which allows subspace owners/admins to pay the fees of subspace-related transactions on behalf of the users.

## Context

Currently, one of the major problems of current Web3 services is complicated to use, it requires users to have tokens paying transaction fees so that they can use the service. For instance, in order to create a post on Desmos-based social network, users have to understand:
- what `transaction` is and why they need some DSM to broadcast it;
- get some DSM either via an on-ramp or by swapping existing funds.

The `x/feegrant` gives the possibility to pay fees for the users, meaning that users can use the service without understanding how Web3 works behind. However, the `x/feegrant` allowance is not subspace-specified so subspace fees providers might unexpected pay the fees of the transactions to any other subspaces.

## Decision

We will implement a subspace-specified fee grant process based on `x/feegrant` that allows subspace fees providers to pay fees for the users inside the specified subspace. This system will allow subspace owners/admins to issue a fee grant towards either a users group or a single user, so that they can later leverage this grant to execute subspace-related transactions without having to worry about paying fees.

### DeductFeeDecorator

Currently, `x/auth` provides a `DeductFeeDecorator` based on `x/feegrant` to execute the action deducting the fees of a transaction from the signer/feepayer. We will build a new subspace-specified `DeductFeeDecorator` to replace the current one.

The new subspace-specified `DeductFeeDecorator` will work as follows:
1. check if all the messages in the transaction are messages related to the same subspace;
2. if the transaction contains subspace messages from the same subspace and a fee grant exists, run the `x/subspaces` `DeductFeeDecorator`; otherwise run the `x/auth` `DeductFeeDecorator`,

### Types

#### Allowance

The `x/feegrant` module already provides an interface (`FeeAllowanceI`) to represent a generic allowance, along with a set if useful implementatations like `BasicAllowance`, `AllowedMsgAllowance` and `PeriodicAllowance`. Since we are not going to introduce new kinds of allowances, nor we are going to edit how an allowance is represented, we will reuse these types.

#### User Grant

The `x/feegrant` module provides a `Grant` object used to store the `granter`, `grantee` and what kind of `allowance` is granted to a specific user. Since this is enough information for us as well, we will use the same object to store user-related fee grants.

#### GroupGrant

#### Group Grant

In order to represent a fee grant granted to a user group, we will implement a `GroupGrant` type. This will contain the `granter`, `group_id` and what kind of `allowance` is granted to the group:

```protobuf
message GroupGrant {
    // granter is the address of the user granting an allowance of their funds.
  string granter = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];

  // the id of the group being granted an allowance of another user's funds.
  uint32 group_id = 2;

  // allowance can be any of basic, periodic, allowed fee allowance.
  google.protobuf.Any allowance = 3 [(cosmos_proto.accepts_interface) = "FeeAllowanceI"];
}
```

### Store

#### Group grant

In order to simplify granters to manage the group allowance, the granted group allowance will be stored in the keys having the structure like:
```
SubspaceGroupAllowancePrefix | SubspaceID | GranterAddress | GroupID | -> Protobuf(GroupGrant)
```

This structure allows granters to easily manage the group allowance inside a subspace by iterating over all allowances for the granters, which will be the most used query. In the other hand, grantees must know who the granter is when using the application, they can find their grant with O(N) time complexity, N is the number of granted groups by the granter.

#### User grant

Each subspace fee granted allowance will be stored in the keys having the structure as follows:
```
SubspaceUserAllowancePrefix | SubspaceID | GranterAddress | GranteeAddress |-> Protobuf(Grant)
```

This structure allows granters to easily manage their grants inside a subspace by iterating over all grants for the granters, which will be the most used query. In the other hand, grantees must know who the granter is when using the application, they can directly find their grant with O(1) time complexity.


### `Msg` Service

In order to allow subspace fees providers to grant an allowance for the users, we will have the following operations:

- grant an allowance to a user
- revoke an allowance to a user
- grant an allowance to a group
- revoke an allowance to a group

```protobuf
service Msg {
    // GrantUserAllowance allows the granter to grant a fee allowance to the grantee.
    rpc GrantUserAllowance(MsgGrantUserAllowance) returns(MsgGrantUserAllowanceResponse);

    // RevokeUserAllowance allows a granter to revoke any existing allowance that has to been granted to the grantee.
    rpc RevokeUserAllowance(MsgRevokeUserAllowance) returns(MsgRevokeUserAllowanceResponse);

    // GrantGroupAllowance allows the granter to grant a fee allowance to the group.
    rpc GrantGroupAllowance(MsgGrantGroupAllowance) returns(MsgGrantGroupAllowanceResponse);

    // RevokeGroupAllowance allows a granter to revoke any existing allowance that has to been granted to the group.
    rpc RevokeGroupAllowance(MsgRevokeGroupAllowance) returns(MsgRevokeGroupAllowanceResponse);
}

// MsgGrantUserAllowance adds permissions for the grantee to spend up allowance of fees from the granter inside the given subspace.
message MsgGrantUserAllowance {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // the address of the user granting an allowance of their funds.
    string granter = 2;

    // the address of the user being granted an allowance of another user's funds.
    string grantee = 3;

    // allowance can be any of fee allowances which implements FeeAllowanceI.
    google.protobuf.Any allowance = 4 [(cosmos_proto.accepts_interface) = "FeeAllowanceI"];
}

// MsgGrantUserAllowanceResponse defines the Msg/GrantAllowanceResponse response type.
message MsgGrantUserAllowanceResponse {}

// MsgRevokeUserAllowance removes any existing allowance from granter to the grantee inside the subspace.
message MsgRevokeUserAllowance {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // the address of the user granting an allowance of their funds.
    string granter = 2;

    // the address of the user being granted an allowance of another user's funds.
    string grantee = 3;
}

// MsgRevokeUserAllowanceResponse defines the Msg/RevokeAllowanceResponse response type.
message MsgRevokeUserAllowanceResponse {}

// MsgGrantGroupAllowance adds permissions for the group to spend up allowance of fees from the granter inside the given subspace.
message MsgGrantGroupAllowance {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // the id of the group being granted an allowance of another user's funds.  
    uint32 group_id = 2;

    // the address of the user granting an allowance of their funds.
    string granter = 3;

    // allowance can be any of fee allowances which implements FeeAllowanceI.
    google.protobuf.Any allowance = 4 [(cosmos_proto.accepts_interface) = "FeeAllowanceI"];
}

// MsgGrantGroupAllowanceResponse defines the Msg/GrantAllowanceResponse response type.
message MsgGrantGroupAllowanceResponse {}

// MsgRevokeGroupAllowance removes any existing allowance from granter to the group inside the subspace.
message MsgRevokeGroupAllowance {
    // the id of the subspace where the granter grants the allowance to the group.
    uint64 subspace_id = 1;

    // the id of the group being granted an allowance of another user's funds.
    uint32 group_id = 2;

    // the address of the user granting an allowance of their funds.
    string granter = 3;
}

// MsgRevokeGroupAllowanceResponse defines the Msg/RevokeAllowanceResponse response type.
message MsgRevokeUserAllowanceResponse {}
```

### `Query` Service

In order to allow clients to easily query for allowances we will implement a new queries:

```protobuf
service Query {
    // UserAllowances returns all the grants for users.
    rpc UserAllowances(QueryUserAllowancesRequest) returns (QueryAllowancesResponse) {
        option (google.api.http).get = "/desmos/subspaces/v3/subspaces/{subspace_id}/granter/{granter}/users/allowances";
    }
    // GroupAllowances returns all the grants for groups.
    rpc GroupAllowances(QueryGroupAllowancesRequest) returns(QueryGroupAllowancesResponse) {
        option (google.api.http).get = "/desmos/subspaces/v3/subspaces/{subspace_id}/granter/{granter}/groups/allowances";
    }
}

// QueryUserAllowancesRequest is the request type for the Query/UserAllowances RPC method.
message QueryUserAllowancesRequest {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // the address of the user granting an allowance of their funds.
    string granter = 2;

    // (Optional) the address of the user being granted an allowance of another user's funds.
    string grantee = 3;

    // pagination defines an pagination for the request.
    cosmos.base.query.v1beta1.PageRequest pagination = 4;
}

// QueryUserAllowancesResponse is the response type for the Query/UserAllowances RPC method.
message QueryUserAllowancesResponse {
    // allowances are allowance's granted for grantee by granter.
    repeated cosmos.feegrant.v1beta1.Grant allowances = 1;

    // pagination defines an pagination for the response.
    cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryGroupAllowancesRequest is the request type for the Query/GroupAllowances RPC method.
message QueryGroupAllowancesRequest {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // the address of the user granting an allowance of their funds.
    string granter = 2;

    // (Optional) the id of the group being granted an allowance of another user's funds.
    uint32 group_id = 3;

    // pagination defines an pagination for the request.
    cosmos.base.query.v1beta1.PageRequest pagination = 4;
}

// QueryGroupAllowancesResponse is the response type for the Query/GroupAllowances RPC method.
message QueryGroupAllowancesResponse {
    // allowances are allowance's granted for grantee by granter.
    repeated cosmos.subspace.v3.GroupGrant allowances = 1;

    // pagination defines an pagination for the response.
    cosmos.base.query.v1beta1.PageResponse pagination = 2;
}
```

## Consequences

### Backwards Compatibility

The proposed solution introduces a new set of store keys, it is completely backward compatible.

### Positive

- Allow subspace fee providers to grant fees allowances to their users.

### Negative

- Storing extra subspace grant info takes up more storage space;
- Having more check in the deducting fees phase slow down the transaction operation;
- Removing expired allowances at the block begin phase increases the block production time.  

### Neutral

- Not known

## References
https://docs.cosmos.network/v0.46/modules/feegrant/01_concepts.html