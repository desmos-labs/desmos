# ADR 003: Subspaces: End-blocker checks on unregistered users

## Changelog

- September 15th, 2021: Initial draft;

## Status

DRAFT

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

All the pairs needs to be constantly checked, to do so we will introduce an iterator that will be implemented
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

## Consequences

> This section describes the resulting context, after applying the decision. All consequences should be listed here, not just the "positive" ones. A particular decision may have positive, negative, and neutral consequences, but all of them affect the team and project in the future.

### Backwards Compatibility

> All ADRs that introduce backwards incompatibilities must include a section describing these incompatibilities and their severity. The ADR must explain how the author proposes to deal with these incompatibilities. ADR submissions without a sufficient backwards compatibility treatise may be rejected outright.

### Positive

{positive consequences}

### Negative

{negative consequences}

### Neutral

{neutral consequences}

## Further Discussions

While an ADR is in the DRAFT or PROPOSED stage, this section should contain a summary of issues to be solved in future iterations (usually referencing comments from a pull-request discussion).
Later, this section can optionally list ideas or improvements the author or reviewers found during the analysis of this ADR.

## Test Cases [optional]

Test cases for an implementation are mandatory for ADRs that are affecting consensus changes. Other ADRs can choose to include links to test cases if applicable.

## References

- {reference link}