# Subspace
Since Desmos is a protocol to create multiple social-enabled applications, a `subspace` represents a single dApp built on top of Desmos. 

Each subspace is independent of the others and has its own users and administrators. When you create a post, you need to specify the subspace inside which it should be visualized. This allows different dApps to coexist on the same protocol.

## Contained data
Following you can find the attributes that distinguish one subspace from the other.

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
The `Type` field tells if users are free to post inside the subspace without being registered in it. The accepted values of this field are:

- `SUBSPACE_TYPE_OPEN` if the subspace should be open to everyone and does not require an admin registering a user before they can post inside such subspace
- `SUBSPACE_TYPE_CLOSED` if the subspace should be closed and users that want to post inside it should be first registered by a subspace admin.

## Other data
Aside from the attributes of each subspace, each subspace has its own set of the following:

- **Admins**  
  They represent the administrators of the subspace and are allowed to edit the subpsace attributes as well as to (un)register and (un)ban users.
  
- **Registered users**  
   These are all the users that are allowed to post inside the subspace. Note that this is valid only for subspaces that have their type set to `Closed`. All subspaces that are set to be `Open` allow any users to post inside them, so the list of registered users will be empty.

- **Banned users**  
   These are all the users that are not permitted to perform any action inside a subspace. 