#Subspace
Inside Desmos, subspaces are the way with which anyone can distinguish a dApp with their users, posts and admins.  
A subspace can be seen as a group that belongs to a user who can decide how to handle users inside it.  
Inside subspaces users can be:
 - **Registered or unregistered**:  
   This will allow them to perform some kind of operations like to post inside the subspace 
   or to put reactions, etc...
 - **Added or Removed as admins**:   
   This will allow them to moderate the subspace alongside the owner.
 - **Banned or Unbanned**:   
   This can happen when a user breaks the rules of a subspace. 
   
## Contained data

## `SubspaceID`
The `SubspaceID` uniquely identifies each subspace. It can be specified from the creator and has to be a valid
SHA-256 hash string.

## `Name`
The `Name` is the human-readable name of the subspace. It must be non-empty nor blank.

## `Description`
The `Description` contains a brief summary of what the subspace is about.

## `Logo`
The `Logo` is the URI of the image representing the subspace.

## `Creator`
The `Creator` identifies the user that has created the subspace.  
It's a string representation of a Bech32 address and, in order to be valid,   
it must begin with the `desmos` Bech32 human-readable part.

## `Owner`
The `Owner` identifies the owner of the subspace.   
It can be equivalent to the `Creator` field.  
It's a string representation of a Bech32 address and, in order to be valid,   
it must begin with the `desmos` Bech32 human-readable part.

## `Type`
The `Type` field tells if users are free to post inside the subspace without being registered in it.
It's an enum value, and it's set to `close` by default.

## `Admins`
The `Admins` field contains all the subspace's admins.   
It has to be an array of valid Bech32 addresses.

## `RegisteredUsers`
The `RegisteredUsers` field contains all the subspace's users.   
It has to be an array of valid Bech32 addresses.

## `BannedUsers`
The `BannedUsers` field contains all the subspace's banned users.   
It has to be an array of valid Bech32 addresses.