---
id: messages
title: Messages
sidebar_label: Messages
slug: messages
---

# Msg Service

## Msg/CreateSubspace
A subspace can be created with the `MsgCreateSubspace` specifying the fields showed below:

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L65-L75 
```

The message is expected to fail if one of the following checks inside the function is matched:
```js reference
https://github.com/desmos-labs/desmos/blob/48b7d58e9b09f26639f7a767fbd921dba309350f/x/subspaces/types/models.go#L51-L82
```

## Msg/EditSubspace
A subspace can be edited with the `MsgEditSubspace`:

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L85-L100 
```

The message is expected to fail if:
* The subspace id doesn't exist;
* The updated subspaces is invalid;
* The signer has no permission to edit subspace inside the specified subspace.

## Msg/DeleteSubspace
A subspace can be deleted using `MsgDeleteSubspace`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L105-115 
```

The message is expected to fail if:
* The subspace does not exist;
* The signer has no permission to delete subspace inside the specified subspace.

## Msg/CreateSection
A section can be created using the `MsgCreateSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L122-L141 
```

The message is expected to fail if:
* The subspace does not exist;
* The parent section does not exist;
* The creator has no permission to manage sections inside the specified subspace.

## Msg/EditSection
A section can be edited using the `MsgEditSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L156-L179 
```

The message is expected to fail if:
* The subspace does not exist;
* The section does not exist;
* The editor has no permission to manage sections inside the specified subspace.

## Msg/MoveSection
A section can be moved to under another section using the `MsgMoveSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L184-L207 
```

The message is expected to fail if:
* The subspace does not exist;
* The section does not exist;
* The destination section does not exist;
* The updated section is not valid;
* The signer has no permission to manage sections inside the specified subspace;
* The section path is invalid.

## Msg/DeleteSection
A section can be deleted using the `MsgDeleteSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L212-L224 
```

The message is expected to fail if:
* The subspace does not exist;
* The section does not exist;
* The signer has no permission to manage sections inside the specified subspace.

## Msg/CreateUserGroup
A user group can be created using the `MsgCreateUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L235-L261 
```

The message is expected to fail if:
* The subspace does not exist;
* The destination section does not exist;
* The signer has no permissions to create a user group or set permissions inside the specified section;
* The permissions values are not valid;
* The group is not valid;

## Msg/EditUserGroup
A user group can be edited using the `MsgEditUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L271-L293
```

The message is expected to fail if:
* The subspace does not exist;
* The user group does not exist;
* The signer has no permission to manage user groups;
* The updated group is invalid.

## Msg/MoveUserGroup
A user group can be moved to another section group using the `MsgMoveUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L298-L317
```

The message is expected to fail if:
* The subspace does not exist;
* The destination section does not exist;
* The user group does not exist;
* The signer has no permission to manage user groups inside the section;
* The signer has no permissions to manage user groups or set permissions inside the destination section.

## Msg/SetUserGroupPermissions
A user group permissions can be set using the `MsgSetUserGroupPermissions`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L326-L347
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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L353-L369
```

The message is expected to fail if:
* The subspace does not exist;
* The user group does not exist;
* The signer has no permission to manage sections inside the specified section.

## Msg/AddUserToUserGroup
A user can be added to a user group using the `MsgAddUserToUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L376-L396
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the signer has no permission to set permissions inside the subspace and section where user group is;
* the user already is a member of the user group.

## Msg/RemoveUserFromUserGroup
A user can be removed from a user group using the `MsgRemoveUserFromUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L402-L422
```

It's expected to fail if:
* the subspace does not exist;
* the user group does not exist;
* the sender has no permission to set permissions inside the subspace and section where user group is;
* the user is not the member of the user group.

## Msg/SetUserPermissions
A user permissions can be set using the `MsgSetUserPermissions`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L430-L454
```

It's expected to fail if:
* the subspace does not exist;
* the section does not exist;
* the signer has no permission to set permissions inside the destination section;
* the permissions values are not valid.