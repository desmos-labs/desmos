# ADR 004: Expiration of application links

## Changelog

- September 20th, 2021: Initial draft;
- September 21th, 2021: Moved from DRAFT to PROPOSED

## Status

PROPOSED

## Abstract

Currently when a user links their centralized application with the Desmos profile, the created link contains a timestamp of when it has been created.   
Since centralized applications username can be switched and sold, we SHOULD implement an "expiration date" system on links. This means that after a certain amount of time passes, the link will be automatically marked as invalid, and the user has to connect it again in order to keep it valid.

## Context

Desmos `profiles` module give the possibility to link the desmos profile with both centralized application and 
other blockchains addresses. By doing this, a user can be verified as the owner of those accounts and prove to the system
that he's not impersonating anyone else. This verification however remains valid only if the user
never trade/sell his centralized-app username to someone else. If it does, and he already linked his Desmos profile
with some application, the link should be invalidated. Unfortunately for us, it's not possible to understand
when this happens since it's off-chain. To prevent this situation, an "expiration time" SHOULD be added 
to the centralized `app_links` object.

## Decision

First we expand the `Application link` structure by including a `expiration_height` field:  
```protobuf
message ApplicationLink {
  [...]
  // ExpirationBlockHeight represents the block height in which the link will be expired
  uint64 expiration_block_height = 7 [ (gogoproto.moretags) = "yaml:\"expiration_block_height\"" ];
}
```

Then we add a new `ApplicationLinkState` `EXPIRED` necessary to address the no-longer valid links.

These additions should be reflected later on the `keeper` code itself:
1) Save the new `expiration_block_height` field with a key composed by `ExpirationPrefix + BlockHeight + ClientID`:  
   The pair will be formed `ExpiringApplicationLinkKey -> ClientID`  
2) Add a new function that will handle the checks on all links by using an iterator over them: 
```go
// IterateExpiringApplicationLinks iterates through all the expiring application links at the given block height
 // The key will be skipped and deleted if the application link has been deleted
 func (k Keeper) IterateExpiringApplicationLinks(ctx sdk.Context, blockHeight uint64, fn func(index int64, link types.ApplicationLink) (stop bool)) {
 	store := ctx.KVStore(k.storeKey)
 	iterator := sdk.KVStorePrefixIterator(store, types.ExpiringApplicationLinkPrefix(blockHeight))
 	defer iterator.Close()

 	i := int64(0)
 	for ; iterator.Valid(); iterator.Next() {
 		// Skip if application link has been deleted
 		clientID := iterator.Value()
 		if !store.Has(clientID) {
 			store.Delete(iterator.Key())
 			continue
 		}
 		applicationKey := store.Get(clientID)
 		link := types.MustUnmarshalApplicationLink(k.cdc, store.Get(applicationKey))
 		stop := fn(i, link)
 		if stop {
 			break
 		}
 		i++
 	}
 }
```

```go
// UpdateExpiringApplicationLinks updates the states of all the expiring application links to be expired
 func (k Keeper) UpdateExpiringApplicationLinks(ctx sdk.Context) {
 	k.IterateExpiringApplicationLinks(ctx, uint64(ctx.BlockHeight()), func(_ int64, link types.ApplicationLink) (stop bool) {
 		store := ctx.KVStore(k.storeKey)
 		link.State = types.AppLinkStateVerificationExpired
 		userApplicationLinkKey := types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username)
 		store.Set(userApplicationLinkKey, types.MustMarshalApplicationLink(k.cdc, link))

 		store.Delete(types.ExpiringApplicationLinkKey(uint64(ctx.BlockHeight()), link.OracleRequest.ClientID))
 		return false
 	})
 }
```

The last step need to update the `EndBlock` function in order to handle the checks and perform
the according expiring actions on the links:

```go
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
    k.UpdateExpiringApplicationLinks(ctx)
}
```

## Consequences

### Backwards Compatibility

This update will affect the `ApplicationLink` object by adding the new `ExpirationBlockHeight` 
field on it breaking the compatibility with the previous versions of the software. To allow
a smooth update and overcome these compatibility issues, we need to set up a proper migration
from the previous versions to the one that will include the additions contained in this ADR.

```go
func Migrate(actualHeight uint64, genState legacy.GenesisState) *GenesisState {
    return &GenesisState{
        DTagTransferRequests: genState.DTagTransferRequests,
        Relationships:        genState.Relationships,
        Blocks:               genState.Blocks,
        Params:               genState.Params,
        IBCPortID:            genState.IBCPortID,
        ChainLinks:           genState.ChainLinks,
        ApplicationLinks:     migrateApplicationLinks(actualHeight, genState.ApplicationLinks),
    }
}

func migrateApplicationLinks(actualHeight uint64, legacyAppLinks []legacy.ApplicationLink) (appLinks []ApplicationLink) {
    appLinks = make([]ApplicationLink, len(legacyAppLinks))
    
    for index, link := range legacyAppLinks {
        appLinks[index] = ApplicationLink{
            User:                  link.User,
            Data:                  link.Data,
            State:                 link.State,
            OracleRequest:         link.OracleRequest,
            Result:                link.Result,
            CreationTime:          link.CreationTime,
            ExpirationBlockHeight: actualHeight,
        }
    }   
    return appLinks
}
```

### Positive

* Considerably reduce the possibility of impersonation of entities and users of centralized apps;

### Negative

* Some more operations to perform on the end Blocker side

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#516](https://github.com/desmos-labs/desmos/issues/516)
- PR [#562](https://github.com/desmos-labs/desmos/pull/562)