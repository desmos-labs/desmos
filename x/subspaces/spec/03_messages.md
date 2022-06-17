<!--
order: 3
-->

# Msg Service

## Msg/CreateSubspace
A subspace can be created with the `MsgCreateSubspace`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L65-L75 
```

## Msg/EditSubspace
A subspace can be edited with the `MsgEditSubspace`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L85-L100 
```

It's expected to fail if:
* the subspace does not exist.
* the sender has no permission to edit subspace inside the specified subspace.

## Msg/DeleteSubspace
A subspace can be deleted using `MsgDeleteSubspace`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L105-115 
```

It's expected to fail if:
* the subspace does not exist.
* the sender has no permission to delete subspace inside the specified subspace.

## Msg/CreateSection
A section can be created using the `MsgCreateSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L122-L141 
```

It's expected to fail if:
* the subspace does not exist
* the parent section does not exist
* the sender has no permission to manage sections inside the specified subspace

## Msg/EditSection
A section can be edited using the `MsgEditSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L156-L179 
```

It's expected to fail if:
* the subspace does not exist.
* the section does not exist.
* the sender has no permission to manage sections inside the specified subspace.

## Msg/MoveSection
A section can be moved to under another section using the `MsgMoveSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L184-L207 
```

It's expected to fail if:
* the subspace does not exist.
* the section does not exist.
* the destination section does not exist.
* the sender has no permission to manage sections inside the specified subspace.
* the section path is invalid not to be able to reach the root section.

## Msg/DeleteSection
A section can be deleted using the `MsgDeleteSection`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L212-L224 
```

It's expected to fail if:
* the subspace does not exist.
* the section does not exist.
* the sender has no permission to manage sections inside the specified subspace.

## Msg/CreateUserGroup
A user group can be created using the `MsgCreateUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L235-L261 
```

It's expected to fail if:
* the subspace does not exist.
* the destination section does not exist.
* the sender has no permissions to create a user group and set permissions inside the specified section.

## Msg/EditUserGroup
A user group can be edited using the `MsgEditUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L271-L293
```

It's expected to fail if:
* the subspace does not exist.
* the user group does not exist.
* the sendor has no permission to manage user groups.

## Msg/MoveUserGroup
A user group can be moved to another section group using the `MsgMoveUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L298-L317
```

It's expected to fail if:
* the subspace does not exist.
* the user group does not exist.
* the destination section does not exist.
* the sender has no permission to manage user groups inside the section.
* the sender has no permissions to manage user groups and set permissions inside the destination section.

## Msg/SetUserGroupPermissions
A user group permissions can be set using the `MsgSetUserGroupPermissions`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L326-L347
```

It's expected to fail if:
* the subspace does not exist.
* the user group does not exist.
* the sender has no permission to set permissions.
* the sender is inside the user group but not the subspace owner.

## Msg/DeleteUserGroup
A user group permissions can be deleted using the `MsgDeleteUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L353-L369
```

It's expected to fail if:
* the subspace does not exist.
* the user group does not exist.
* the sender has no permission to manage sections inside the specified section.

## Msg/AddUserToUserGroup
A user can be added to a user group using the `MsgAddUserToUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L376-L396
```

It's expected to fail if:
* the subspace does not exist.
* the user group does not exist.
* the sender has no permission to set permissions inside the section where user group live.
* the sender already is the member of the user group.

## Msg/RemoveUserFromUserGroup
A user can be removed from a user group using the `MsgRemoveUserFromUserGroup`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L402-L422
```

It's expected to fail if:
* the subspace does not exist.
* the user group does not exist.
* the sender has no permission to set permissions inside the section where user group live.
* the user is not the member of the usr group.

## Msg/SetUserPermissions
A user permissions can be set using the `MsgSetUserPermissions`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/subspaces/v2/msgs.proto#L430-L454
```

It's expected to fail if:
* the subspace does not exist.
* the destination section does not exist.
* the sender has no permission to set permissions inside the destination section.