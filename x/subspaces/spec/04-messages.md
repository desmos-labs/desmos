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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L65-L82 
```

It's expected to fail if:
* the provided name is either empty or blank;
* the specified treasury address (if any) is invalid;
* the specified owner address (if any) is invalid.

## Msg/EditSubspace
A subspace can be edited with the `MsgEditSubspace`:

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L93-L119 
```

It's expected to fail if:
* the subspace does not exist;
* the updated subspace is invalid;
* the signer has no permission to edit the subspace.

## Msg/DeleteSubspace
A subspace can be deleted using `MsgDeleteSubspace`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L124-L134 
```

It's expected to fail if:
* the subspace does not exist;
* the signer has no permission to delete the subspace.

## Msg/CreateSection
A section can be created using the `MsgCreateSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L141-L164 
```

The message is expected to fail if:
* the subspace does not exist;
* the parent section (if specified) does not exist;
* the creator has no permission to manage sections within the subspace;
* the provided section name is either empty or blank.

## Msg/EditSection
A section can be edited using the `MsgEditSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L175-L198 
```

It's expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the editor has no permission to manage sections within the subspace.

## Msg/MoveSection
A section can be moved to under another section using the `MsgMoveSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L203-L226 
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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L231-L247 
```

It's is expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the signer has no permission to manage sections within the subspace.

## Msg/CreateUserGroup
A user group can be created using the `MsgCreateUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L254-L284 
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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L294-L316
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the signer has no permission to manage user groups within the subspace;
* the updated group is invalid.

## Msg/MoveUserGroup
A user group can be moved to another section group using the `MsgMoveUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L321-L344
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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L349-L370
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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L376-L392
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the signer has no permission to manage sections inside the group's section.

## Msg/AddUserToUserGroup
A user can be added to a user group using the `MsgAddUserToUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L399-L419
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the signer has no permission to set permissions inside the subspace and section where user group is;
* the user already is a member of the user group.

## Msg/RemoveUserFromUserGroup
A user can be removed from a user group using the `MsgRemoveUserFromUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L425-L445
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the sender has no permission to set permissions inside the subspace and section where user group is;
* the user is not the member of the user group.

## Msg/SetUserPermissions
A user permissions can be set using the `MsgSetUserPermissions`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v3/msgs.proto#L453-L477
```

It's expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the signer has no permission to set permissions inside the destination section;
* the permissions values are not valid.
