# ADR 006: Subspace module

## Changelog

- December 15th, 2021: Initial draft

## Status

DRAFT

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

#### ACL
In order to easily implement an ACL, we will use a simple set of keys made as follows: 
```
ACLPrefix + User Address -> ACL Value
```

The `ACL Value` will be a simple binary value allowing us to perform bitwise operations to combine the following different permissions: 

```go
const (
  // Identifies users that can create content inside the subspace
  PermissionWrite           = 0b000001
  
  // Allows users to moderate contents of other users (e.g. deleting it) 
  PermissionModerateContent = 0b000010
  
  // Allows to add a link for this subspace
  PermissionAddLink         = 0b000100
  
  // Allows to change the information of the subspace
  PermissionChangeInfo      = 0b001000
  
  // Allows to set other users' permissions (except PermissionSetPermissions)
  PermissionSetPermissions  = 0b010000
)
```

> **Note**:  
> Only the `Owner` account will be able to grant other users the `PermissionSetPermissions`

Using this kind of permissions will allow us to easily set permissions and check whether a user has a permission or not: 
```go
userPermissions := PermissionWrite | PermissionAddLink | PermissionChangeInfo

canWrite := (userPermissions & PermissionWrite) == PermissionWrite  // True
canModerateContent := (userPermissions & PermissionModerateContent) == PermissionModerateContent // False
```

### `Msg` Service
We will allow the following operations to be performed:

* delete contents that do not respect the ToS;
* ban users that do not respect the ToS.

```protobuf
// Msg defines subspaces Msg service.
service Msg {

  // CreateSubspace allows to create a subspace
  rpc CreateSubspace(MsgCreateSubspace) returns (MsgCreateSubspaceResponse);

  // EditSubspace allows to edit a subspace
  rpc EditSubspace(MsgEditSubspace) returns (MsgEditSubspaceResponse);
  
  // SetUserPermissions allows to set another user's permissions
  rpc SetUserPermissions(MsgSetUserPermissions) returns (MsgSetUserPermissionsResponse);
}

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