# ADR 003: Subspaces: End-blocker checks on unregistered users

## Changelog

- September 15th, 2021: Initial draft;
- September 16th, 2021: Moved from DRAFT to PROPOSED.

## Status

PROPOSED

## Abstract

When a user unregister himself from a subspace, all the related `relationships` and `blocks` with
users should be deleted. To do this, we SHOULD add some check to the `subspaces` end-blocker
in order to let it deal with these situations.

## Context

The `subspaces` module has the purpose to let dApp developers create a "space" where their dApps can live
on with their own Term of Services. By doing this, all their users and posts will be associated with
a particular subspace and should follow a set of rules and be subject to the changes that happens on the subspace itself. 
Currently, inside a subspace, to obtain the right to post, add a reaction, answer a poll and report a post, users firstly need to register
into it. After doing it, he will become a user of that subspace and be able to perform the set of ops we mentioned above here.  
To ensure a user can perform these operations, Desmos modules needs to check user's status on the subspace. Most of these checks have already been implemented inside the `posts` module and are crucial to make the whole `subspaces`
aspect to work correctly.

## Decision

The implementation idea for this is pretty straightforward, and it's split across the 3 modules:
`subspaces`, `profiles` and `posts`.

### Subspaces
When a user unregister from a subspace, the system creates an `UnregisteredPair` of (subspace + user). Later the pair will be used to check for the user, and delete any eventual relationship or block.

````protobuf
message UnregisteredPair {
   option (gogoproto.goproto_getters) = false;

   // the id of the subspace where user unregisters himself from
   string subspace_id = 1 [
     (gogoproto.moretags) = "yaml:\"subspace_id\"",
     (gogoproto.customname) = "SubspaceID"
   ];

   // the address of the unregistered user
   string user = 2 [ (gogoproto.moretags) = "yaml:\"user\"" ];
 }
````

The system saves the pair inside `subspaces` keeper and deleted it after it complete the operations
regarding users' `relationships` and `blocks`.  
To do so, two methods needs to be implemented:
```go
 func (k Keeper) SaveUnregisteredPair(ctx sdk.Context, subspaceID, user string) {
 	store := ctx.KVStore(k.storeKey)
 	pair := types.NewUnregisteredPair(subspaceID, user)
 	store.Set(types.UnregisteredPairKey(subspaceID, user), k.cdc.MustMarshalBinaryBare(&pair))
 }
```

```go
func (k Keeper) DeleteUnregisteredPair(ctx sdk.Context, subspaceID, user string) {
 	store := ctx.KVStore(k.storeKey)
 	unregisteredKey := types.UnregisteredPairKey(subspaceID, user)
 	if store.Has(unregisteredKey) {
 		store.Delete(unregisteredKey)
 	}
 }
```

It's also necessary to: 
* Check if an `uregisterPair` exists when a user is registering himself into a subspace;
* Save the `unregisterPair` when a user unregister himself from a subspace.

All the pairs needs to be constantly checked and in order to do it we will introduce an iterator that will be implemented
inside the following function:  
```go
 func (k Keeper) IterateUnregisteredPairs(ctx sdk.Context, fn func(index int64, pair types.UnregisteredPair) (stop bool)) {
 	store := ctx.KVStore(k.storeKey)
 	iterator := sdk.KVStorePrefixIterator(store, types.UnregisteredPairPrefix)
 	defer iterator.Close()
 	i := int64(0)
 	for ; iterator.Valid(); iterator.Next() {
 		var pair types.UnregisteredPair
 		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &pair)
 		stop := fn(i, pair)
 		if stop {
 			break
  		}
 		i++
 	}
 }
```

###Profiles

The `profiles` module MUST use the following `subspaces` interface to enable interactions with 
`UnregisteredPair` related methods:

```go
type SubspacesKeeper interface {
 	IterateUnregisteredPairs(ctx sdk.Context, fn func(index int64, pair subspacestypes.UnregisteredPair) (stop bool))
 	DeleteSubspaceUnregisteredPair(ctx sdk.Context, subspaceID, user string)
 }
```

The `profiles` module handles the relationships and blocks of any user, so whenever one of them unregister from a subspace,
the module itself need to act and delete both of them.
To handle this, we will first need to implement some iterators for the two structures and later use them in order
to find out which ones to delete. 

```go
func (k Keeper) IterateSubspaceUserRelationships(ctx sdk.Context, user, subspaceID string, fn func(index int64, relationship types.Relationship) (stop bool)) {
func (k Keeper) IterateSubspaceUserBlocks(ctx sdk.Context, user, subspaceID string, fn func(index int64, block types.UserBlock) (stop bool))
```

These two methods above will be called inside other two methods that will take care of delete any relationship or block
of a user:

```go
func (k Keeper) DeleteSubspaceUserBlocks(ctx sdk.Context, subspaceID, user string) {
	store := ctx.KVStore(k.storeKey)

	k.IterateSubspaceUserBlocks(ctx, user, subspaceID, func(index int64, block types.UserBlock) (stop bool) {
		store.Delete(types.UserBlockStoreKey(block.Blocker, block.Subspace, block.Blocked))
		return false
	})
}
```

```go
func (k Keeper) DeleteSubspaceUserRelationships(ctx sdk.Context, subspaceID, user string) {
	store := ctx.KVStore(k.storeKey)
	
	k.IterateSubspaceUserRelationships(ctx, user, subspaceID, func(index int64, relationship types.Relationship) (stop bool) {
		store.Delete(types.RelationshipsStoreKey(relationship.Creator, relationship.Subspace, relationship.Recipient))
		return false
	})
}
```

After their implementations, we need to implement the `EndBlocker` function to allow unregistered users'
checks being executed at the end of each block:
```go
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
    k.sk.IterateUnregisteredPairs(ctx, func(_ int64, pair subspacestypes.UnregisteredPair) (stop bool) {
        k.DeleteSubspaceUserRelationships(ctx, pair.SubspaceID, pair.User)
        k.DeleteSubspaceUserBlocks(ctx, pair.SubspaceID, pair.User)
        k.sk.DeleteSubspaceUnregisteredPair(ctx, pair.SubspaceID, pair.User)
        return false
    })
}
```

## Consequences

### Backwards Compatibility

`Subspaces` module is still a staging module, so from its side there are no backwards compatibility problems.  
`Profiles` module is currently in mainnet but at the moment there are no BC problems related to this ADR.


### Positive

* Improve the memory management by deleting unnecessary data; 

### Negative

* Extra logic running on each `EndBlock`;

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#488](https://github.com/desmos-labs/desmos/issues/488)
- PR [#556](https://github.com/desmos-labs/desmos/pull/556)