# ADR 014: Improve the permission system

## Changelog

- May 19th, 2022: First draft;

## Status

ACCEPTED Implemented

## Abstract

This ADR introduces a new way of registering and managing permissions within the `x/subspaces` module all the other modules that depend on it.

## Context

Currently, Desmos implements a permission system that is very similar to the one used within Linux-based systems: each permission is represented by an `uint` value treated reading its bits individually. The combination of a permission is thus obtained using the `|` operator, and to check if a specific permission is set it's sufficient to check the value of the associated bit. Although this works, it has some limitations to it:
- clients need to know the value of each permission, and how to properly combine them to form the final permission value they want to set;
- permission values that are given from outside (e.g. a client setting a permission for a user) need to be sanitized in order to grant forward-compatibility;
- all permissions need to be put inside the `x/subspaces` module, so that a proper validity mask can be computed and used to sanitize permissions values.

## Decision

To allow for an easier way of managing permissions, we will change the entire permission system from being based on a bitwise system to be based on a more human-readable system. Instead of representing each permission value as a `uint` and then reading individual bits, we will represent each permission using a `string` value that can easily tell what the permission is about. Here are some examples:

```go
package types

type Permission string

const (
	PermissionManageSubspace Permission = "PERMISSION_MANAGE_SUBSPACE"
	PermissionEverything     Permission = "PERMISSION_EVERYTHING"
)
```

In order to properly support any kind of permission from different modules, we will implement the following method within the `x/subspaces` module:

```go
var (
	// registeredPermissions represents the list of permissions that are registered and should be considered valid
	registeredPermissions []Permission
)

// RegisterPermission allows to register the permission with the given name and returns its value
func RegisterPermission(permissionName string) Permission {
	permission := Permission(strings.ToUpper(strings.ReplaceAll(permissionName, " ", "_")))
	for _, registeredPermission := range registeredPermissions {
		if registeredPermission == permission {
			panic(fmt.Errorf("permission %s has already been registered", permission))
		}
	}

	registeredPermissions = append(registeredPermissions, permission)
	return permission
}
```

This will allow all modules to register whatever permission they want, and then use them freely while performing checks: 

```go
// Define the permissions for the posts module
var (
	PermissionCreatePost = subspacestypes.RegisterPermission("create post")
	PermissionEditPost   = subspacestypes.RegisterPermission("edit post") 
)

// Check such permissions
keeper.SubspacesKeeper.HasPermissions(ctx, subspaceID, user, PermissionCreatePost, PermisisonEditPost)
```

When setting a user or a group's permissions, we will then require the users to provide them as a `string` array, and then we simply check whether the given permissions are registered inside the supported ones or not. If they are, we store the entire array on the store.  

### Types

#### `UserGroup`
We will change the `UserGroup` type to support this new permission system by replacing the `permissions` field from being a `uint32` to be a `repeated string`: 

```protobuf
syntax = "proto3";

// UserGroup represents a group of users
message UserGroup {
  // ID of the subspace inside which this group exists
  uint64 subspace_id = 1;

  // Unique id that identifies the group
  uint32 id = 2;

  // Human-readable name of the user group
  string name = 3;

  // Optional description of this group
  string description = 4;

  // Permissions that will be granted to all the users part of this group
  repeated string permissions = 5;
}
```

#### `UserPermission`

Instead of storing user permissions as follows: 

```
UserPermissionPrefix | SubspaceID | User | -> uint32 
```

We will define a new `UserPermission` object: 

```protobuf
syntax = "proto3";

// UserPermission contains the details of a single user permissions withing a subspace
message UserPermission {
  // Address of the user
  bytes user = 1;
  
  // Permissions set to the user
  repeated string permissions = 2;
}
```

then the permissions of a user will be stored as follows: 

```
UserPermissionPrefix | SubspaceID | User | -> ProtcolBuffer(UserPermission)
```

### `Msg` Service
We will change the `Msg` service of the `x/subspaces` module to properly allow setting permissions: 

```protobuf
syntax = "proto3";

// MsgSetUserGroupPermissions represents the message used to set the permissions of a user group
message MsgSetUserGroupPermissions {
  uint64 subspace_id = 1;
  uint32 group_id = 2;
  repeated string permissions = 3;
  string signer = 4;
}

// MsgSetUserPermissions represents the message used to set the permissions of a specific user
message MsgSetUserPermissions {
  uint64 subspace_id = 1;
  string user = 2;
  repeated string permissions = 3;
  string signer = 4;
}
```

#### `Query` Service

We will change the `Query` service of the `x/subspaces` module to properly returns new permission values:

```protobuf
syntax = "proto3";

// QueryUserPermissionsRequest is the response type for the Query/UserPermissions method
message QueryUserPermissionsResponse {
  repeated string permissions = 1;
  repeated PermissionDetail details = 2;
}
```

## Consequences

### Backwards Compatibility

The above detailed solution is **not** backward compatible. For this reason, we will need to write a migration code that reads all the currently set permissions and performs the following operations: 
- split the permission value into individual permissions;
- map each individual permission to the new `Permission` type; 
- replace the current store values with the new `Permission` values. 

This will need to be performed for all user groups as well as all individual user permissions. 

### Positive

- Allows to define permission in all modules
- No need to sanitize the values received from clients
- Easier permission setting from clients (values are now human-readable)

### Negative

- Storing `string` instead of `uint` takes up more storage space  

### Neutral


## Further Discussions

In the future, we might even allow developers to register custom permissions within their own subspaces, and then request those permissions during the execution of different messages. This could be done by allowing them to specify a `MsgTypeURL -> []Permission` entry for each kind of message type, so that when a message with `MsgTypeURL` gets executed, the user needs to have the specified permission in order to successfully perform the request.

## Test Cases

- Migrating from old permissions to new permissions work properly and no data is lost

## References

- Issue [#800](https://github.com/desmos-labs/desmos/issues/800)
- Issue [#855](https://github.com/desmos-labs/desmos/issues/855)