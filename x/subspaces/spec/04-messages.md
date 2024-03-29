---
id: messages
title: Messages
sidebar_label: Messages
slug: messages
---

# Msg Service

## Msg/CreateSubspace
A subspace can be created with the `MsgCreateSubspace`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L96-L119
```

It's expected to fail if:
* the provided name is either empty or blank;
* the specified treasury address (if any) is invalid;
* the specified owner address (if any) is invalid.

## Msg/EditSubspace
A subspace can be edited with the `MsgEditSubspace`:

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L130-L161
```

It's expected to fail if:
* the subspace does not exist;
* the updated subspace is invalid;
* the signer has no permission to edit the subspace.

## Msg/DeleteSubspace
A subspace can be deleted using `MsgDeleteSubspace`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L166-L182
```

It's expected to fail if:
* the subspace does not exist;
* the signer has no permission to delete the subspace.

## Msg/CreateSection
A section can be created using the `MsgCreateSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L189-L218
```

The message is expected to fail if:
* the subspace does not exist;
* the parent section (if specified) does not exist;
* the creator has no permission to manage sections within the subspace;
* the provided section name is either empty or blank.

## Msg/EditSection
A section can be edited using the `MsgEditSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L229-L258
```

It's expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the editor has no permission to manage sections within the subspace.

## Msg/MoveSection
A section can be moved to under another section using the `MsgMoveSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L263-L292
```

It's expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the destination section does not exist;
* the signer has no permission to manage sections within the subspace;
* the new section path is invalid (this means that is not possible to reach the moved section starting from the root section, or that a circular path is detected).

## Msg/DeleteSection
A section can be deleted using the `MsgDeleteSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L297-L319
```

It's is expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the signer has no permission to manage sections within the subspace.

## Msg/CreateUserGroup
A user group can be created using the `MsgCreateUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L326-L364
```

The message is expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the signer has no permissions to create a user group or set permissions within the section;
* the permissions values are not valid;
* the provided user group name is either blank or empty.

## Msg/EditUserGroup
A user group can be edited using the `MsgEditUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L374-L402
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the signer has no permission to manage user groups within the subspace;
* the updated group is invalid.

## Msg/MoveUserGroup
A user group can be moved to another section group using the `MsgMoveUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L407-L436
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the destination section does not exist;
* the signer has no permission to manage user groups inside the current group's section;
* the signer has no permissions to manage user groups or set permissions inside the destination section.

## Msg/SetUserGroupPermissions
A user group permissions can be set using the `MsgSetUserGroupPermissions`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L441-L468
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the signer has no permission to set permissions within the group's section;
* the permissions values are not valid;
* the signer is inside the user group and it is not the subspace owner.

## Msg/DeleteUserGroup
A user group permissions can be deleted using the `MsgDeleteUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L474-L496
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the signer has no permission to manage sections inside the group's section.

## Msg/AddUserToUserGroup
A user can be added to a user group using the `MsgAddUserToUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L503-L532
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the signer has no permission to set permissions inside the subspace and section where user group is;
* the user already is a member of the user group.

## Msg/RemoveUserFromUserGroup
A user can be removed from a user group using the `MsgRemoveUserFromUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L538-L567
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the sender has no permission to set permissions inside the subspace and section where user group is;
* the user is not the member of the user group.

## Msg/SetUserPermissions
A user permissions can be set using the `MsgSetUserPermissions`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L575-L608
```

It's expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the signer has no permission to set permissions inside the destination section;
* the permissions values are not valid.

## Msg/GrantAllowance
A subspace admin can grant a user or a user group a fee allowance within the subspace using a `MsgGrantAllowance`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L616-L645
```

It's expected to fail if:
* the subspace does not exist;
* the signer has no permission to grant allowances within the subspace;
* the grantee is not a valid user or user group;
* the grantee is already allowed to spend fees within the subspace;
* the allowance is not valid.

## Msg/RevokeAllowance
A subspace admin can revoke a user or a user group a fee allowance within the subspace using a `MsgRevokeAllowance`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L653-L676
```

It's expected to fail if:
* the subspace does not exist;
* the signer has no permission to revoke allowances within the subspace;
* the grantee is not a valid user or user group;
* the grantee does not have an allowance within the subspace.

## Msg/GrantTreasuryAuthorization
A subspace admin can grant a user the authorization to perform an action on behalf of the subspace treasury using a `MsgGrantTreasuryAuthorization`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L684-L714
```

It's expected to fail if:
* the subspace does not exist;
* the signer has no permission to grant treasury authorizations within the subspace;
* the grantee is not a valid user;
* the grantee is already authorized to perform the given action;
* the authorization is not valid.

## Msg/RevokeTreasuryAuthorization
A subspace admin can revoke a user the authorization to perform an action on behalf of the subspace treasury using a `MsgRevokeTreasuryAuthorization`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L720-L746
```

It's expected to fail if:
* the subspace does not exist;
* the signer has no permission to revoke treasury authorizations within the subspace;
* the grantee is not a valid user;
* the grantee does not already have an authorization to perform the given action.

## Msg/UpdateSubspaceFeeTokens
The `MsgUpdateSubspaceFeeTokens` represents the message to be used to update the accepted fee tokens inside a given subspace, using an on-chain governance proposal.

This is done through a governance proposal in order to make sure that validators are aware that they will start receiving fees in a new token when validating the messages related to the given subspace.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L754-L782
```

It's expected to fail if:
* the subspace does not exist;
* the additional fee tokens are invalid;
* the authority is not the address of the governance module.