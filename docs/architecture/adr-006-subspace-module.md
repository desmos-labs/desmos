# ADR 006: Subspace module

## Changelog

- December 15th, 2021: Initial draft;
- December 16th, 2021: First review;
- January 06th, 2022: Second review;
- January 13th, 2022: Third review;
- January 14th, 2022: Fourth review;
- January 17th, 2022: Fifth review;
- February 10th, 2022: Sixth review.

## Status

ACCEPTED Implemented

## Abstract
This ADR defines the `x/subspaces` module which allows users to create and manage the representation of different social networks inside which contents will be created.  

## Context
In order to support building independent social networks, it is vital for Desmos to expose a mechanism that allows to replicate as much as possible the current stace of social networks in which each platform has different economic model, set of administrators, Terms of Services and ways of enforcing such terms.  

## Decision
We will create a module named `subspaces` which provided users the ability to create independent spaces inside Desmos, each one having its own administrators, ToS and tokenomics.

### Types 
Subspaces must always have an account that is elected as the _owner_ and should act as the final decision maker and accountable user. Additionally, each subspace can have the following data:

* a _human-readable name_ and a _description_, in order to allow users to easily identify the scope of such subspace;
* a set of _connections_ with external applications, allowing users to verify the validity of such subspace and avoid any possible fakes;
* an _ACL_ (access-control list) defining the permissions that different kind of users have inside the subspace itself (e.g. administrators, moderators, etc.).

#### Subspace

```go
// Subspace contains all the data related to a Desmos subspace
type Subspace struct {
  // Unique identifier of this subspace
  ID uint64

  // Human-readable name of the subspace
  Name string
  
  // Optional description of this subspace
  Description string
  
  // Address of the user that owns the subspace
  Owner  string

  // Address of the subspace creator
  Creator string
  
  // Represents the account that is associated with the subspace and
  // should be used to connect external applications to verify this subspace
  Treasury string

  // the creation time of the subspace
  CreationTime time.Time
}
```

#### Permissions
Since each subspace is thought to represent an independent application, we SHOULD allow different subspaces owners to set different permissions for each user. 

For this reason, we will implement an ACL (*Access Control List*) system that can be customized for each subspace. Each ACL MUST support setting permissions for both individual users and user groups.

##### Permission value
To easily implement a composable system, we will use byte-based permissions: each value will be represented by an integer value, and composed permissions can be obtained by using the byte-wide *or* (`|`) operator. Also, this will allow us to easily check whether a user has a specific permission by using the byte-wide *and* (`&`) operator.

```go
const (
    // PermissionNothing represents the permission to do nothing
    PermissionNothing = Permission(0b000000)
    
    // PermissionWrite identifies users that can create content inside the subspace
    PermissionWrite = Permission(0b000001)
    
    // PermissionModerateContent allows users to moderate contents of other users (e.g. deleting it)
    PermissionModerateContent = Permission(0b000010)
    
    // PermissionChangeInfo allows to change the information of the subspace
    PermissionChangeInfo = Permission(0b000100)
    
    // PermissionManageGroups allows users to manage user groups and members
    PermissionManageGroups = Permission(0b001000)
    
    // PermissionSetPermissions allows users to set other users' permissions (except PermissionSetPermissions).
    // This includes managing user groups and the associated permissions
    PermissionSetPermissions = Permission(0b010000)
    
    // PermissionDeleteSubspace allows users to delete the subspace.
    PermissionDeleteSubspace = Permission(0b100000)
    
    // PermissionEverything allows to do everything.
    // This should usually be reserved only to the owner (which has it by default)
    PermissionEverything = Permission(0b111111)
)

userPermissions := PermissionWrite | PermissionManageGroups | PermissionChangeInfo

canWrite := (userPermissions & PermissionWrite) == PermissionWrite  // True
canModerateContent := (userPermissions & PermissionModerateContent) == PermissionModerateContent // False
```

> **Note**:  
> Only the `Owner` account will be able to grant other users the `PermissionSetPermissions`

##### User groups
In order to properly implement user groups, we are going to define the following structure: 

```golang
type UserGroup struct {
    // SubspaceID represents the ID of the subspace inside which the group exists
    SubspaceID  uint64
    // Unique id that identifies the group inside the subspace
    ID          uint32
    // Human-readable name of the user group
    Name        string
    // Optional description of the group
    Description string
    // Permissions that will be granted to all the users inside this group
    Permissions uint32
}

```

The `ID` of each group will be a sequential value starting from `1` for each subspace.

The user group with ID `0` will be used to identify a default user group that contains all users that are not part of any other group. This will be useful if a subspace owner wants to assign a default permission to all users (e.g. they want all users to be able to post inside that subspace by default making that subspace not require a specific registration). This group will be present by default on all subspaces, and it will only be possible to edit its permissions or details. It won't be possible to delete it nor to remove or add people to it.

To store a group and its members, we will use the following key: 
```
GroupPrefix + SubspaceID + GroupID -> Group
```
This will allow us to easily iterate over all the groups inside a subspace.

At the same time, to store each member of a group we will use the following store keys:
```
GroupMemberPrefix + SubspaceID + GroupID + <User address> -> 0x01
```
In this case, the `0x01` value is just a placeholder that allows the key to exist. This key will allow us to easily iterate over all the members of a subspace, as well as to check whether a user is part of a group or not. 

### `Msg` Service
We will allow the following operations to be performed.

**Subspace administration**
* Create a subspace
* Edit a subspace
* Set a group's permissions
* Delete a subspace

**Content management**
* Delete contents that do not respect the ToS

**Groups management**
* Create a new group
* Edit a group
* Set group permissions
* Delete a group

**Users management**
* Add a user to a group
* Remove a user from a group
* Set user permissions


```protobuf
// Msg defines subspaces Msg service.
service Msg {

  // CreateSubspace allows to create a subspace
  rpc CreateSubspace(MsgCreateSubspace) returns (MsgCreateSubspaceResponse);

  // EditSubspace allows to edit a subspace
  rpc EditSubspace(MsgEditSubspace) returns (MsgEditSubspaceResponse);
  
  // CreateUserGroup allows to create a new user group
  rpc CreateUserGroup(MsgCreateUserGroup) returns (MsgCreateUserGroupResponse);

  // EditUserGroup allows to edit a user group
  rpc EditUserGroup(MsgEditUserGroup) returns (MsgEditUserGroupResponse);

  // SetUserGroupPermissions allows to set the permissions for a specific group
  rpc SetUserGroupPermissions(MsgSetUserGroupPermissions)
      returns (MsgSetUserGroupPermissionsResponse);

  // DeleteUserGroup allows to delete an existing user group
  rpc DeleteUserGroup(MsgDeleteUserGroup) returns (MsgDeleteUserGroupResponse);
  
  // AddUserToUserGroup allows to add a specific user to a specific user group
  rpc AddUserToUserGroup(MsgAddUserToUserGroup) returns (MsgAddUserToUserGroupResponse);

  // RemoveUserFromUserGroup allows to remove a specific user from a specific user group
  rpc RemoveUserFromUserGroup(MsgRemoveUserFromUserGroup) returns (MsgRemoveUserFromUserGroupResponse); 
  
  // SetPermissions allows to set the permissions of a user or user group
  rpc SetUserPermissions(MsgSetUserPermissions) returns (MsgSetUserPermissionsResponse);
}

message MsgCreateSubspace {
  string name = 1;
  string description = 2;
  string treasury = 3;
  string owner = 4;
  string creator = 5;
}

message MsgCreateSubspaceResponse {
  uint64 subspace_id = 1;
}

message MsgEditSubspace {
  uint64 subspace_id = 1;
  string name = 2;
  string description = 3;
  string treasury = 4;
  string owner = 5;
  string signer = 6;
}

message MsgEditSubspaceResponse {}

message MsgDeleteSubspace {
  uint64 subspace_id = 1;
  string signer = 2;
}

message MsgDeleteSubspaceResponse {}

message MsgCreateUserGroup {
  uint64 subspace_id = 1;
  string name = 2;
  string description = 3;
  bytes default_permissions = 4;
  string creator = 5;
}

message MsgCreateUserGroupResponse {
  uint32 group_id = 1;
}

message MsgEditUserGroup {
  uint64 subspace_id = 1;
  uint32 group_id = 2;
  string name = 3;
  string description = 4;
  string signer = 5;
}

message MsgEditUserGroupResponse {}

message MsgSetUserGroupPermissions {
  uint64 subspace_id = 1;
  uint32 group_id = 2;
  uint32 permissions = 3;
  string signer = 4;
}

message MsgSetUserGroupPermissionsResponse {}

message MsgDeleteUserGroup {
  uint64 subspace_id = 1;
  uint32 group_id = 2;
  string signer = 3;
}

message MsgDeleteUserGroupResponse {}

message MsgAddUserToUserGroup { 
  uint64 subspace_id = 1;
  uint32 group_id = 2;
  string user = 3;
  string signer = 4;
}

message MsgAddUserToUserGroupResponse {}

message MsgRemoveUserFromUserGroup {
  uint64 subspace_id = 1;
  uint32 group_id = 2;
  string user = 3; 
  string signer = 4;
}

message MsgRemoveUserFromUserGroupResponse {}

message MsgSetUserPermissions {
  uint64 subspace_id = 1;
  string user = 2;
  bytes permissions = 3;
  string signer = 4;
}

message MsgSetUserPermissionsResponse {}
```

## Consequences
### Positive

* Users will be able to create their own subspace representing a social network inside Desmos
* The ACL implementation proposed is generic enough to allow future permissions to be implemented without much work to be done
* The link-based verification system proposed is generic enough to allow any creator to verify their subspace relying on the already existing `x/profiles` module without the need of new code 

### Negative

### Neutral

## References
- Extend the concept of subspaces: https://github.com/desmos-labs/desmos/discussions/375