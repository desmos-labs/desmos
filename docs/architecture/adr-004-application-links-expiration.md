# ADR 004: Expiration of application links

## Changelog

- September 20th, 2021: Initial draft;
- September 21th, 2021: Moved from DRAFT to PROPOSED;
- December  22th, 2021: First review.

## Status

PROPOSED

## Abstract

Currently when a user links their centralized application with the Desmos profile, the created link contains a timestamp of when it has been created.   
Since centralized applications username can be switched and sold, we SHOULD implement an "expiration date" system on links. This means that after a certain amount of time passes, the link will be automatically marked as invalid, and the user has to connect it again in order to keep it valid.

## Context

Desmos `profiles` module give the possibility to link the desmos profile with both centralized application and 
other blockchains accounts. By doing this, a user can be verified as the owner of those accounts and prove to the system
that they're not impersonating anyone else. This verification however remains valid only if the user
never trade or sell his centralized-app username to someone else. If they do the link to such username MUST be invalidated. 
Unfortunately for us, it's not possible to understand when this happens since it's off-chain. 
To prevent this situation, an "expiration time" SHOULD be added to the centralized `ApplicationLink` object.

## Decision

First we expand the `Applicationlink` structure by including an `ExpirationHeight` field:  
```protobuf
message ApplicationLink {
  [...]
  // ExpirationBlockHeight represents the block height in which the link will be expired
   google.protobuf.Timestamp expiration_time = 7 [ (gogoproto.moretags) = "yaml:\"expiration_time\"" ];
}
```

These additions should be reflected later on the `keeper` code itself:
1) Save the new `expiration_time` field with a key composed by `ExpirationPrefix + BlockTime + ClientID`:  
   The pair will be formed `ExpirationTimeApplicationLinkKey -> ClientID`
```go
     // ExpiringApplicationLinkPrefix returns the store prefix used to identify the expiration time of an application link with the given timestamp
     func ExpiringApplicationLinkPrefix(timestamp int64) []byte {
         return append(ExpiringAppLinkPrefix, []byte(strconv.FormatInt(timestamp, 10))...)
     }

     // ExpirationTimeApplicationLinkKey returns the key used to store the data about the expiration time
     // of the application link associated with the given clientID
     func ExpirationTimeApplicationLinkKey(timestamp int64, clientID string) []byte {
         return append(ExpiringApplicationLinkPrefix(timestamp), []byte(clientID)...)
     }
```

2) Add a new function that will iterate over the expiring links references: 
```go
// IterateExpiringApplicationLinks iterates through all the expiring application links references at the given block height
 // The key will be skipped and deleted if the application link has been deleted
func (k Keeper) IterateExpiringApplicationLinks(ctx sdk.Context, expirationTimestamp int64, fn func(index int64, link types.ApplicationLink) (stop bool)) {
    store := ctx.KVStore(k.storeKey)

    iterator := sdk.KVStorePrefixIterator(store, types.ExpiringApplicationLinkPrefix(expirationTimestamp))
    defer iterator.Close()

    i := int64(0)
    for ; iterator.Valid(); iterator.Next() {
    // Skip if application link has been deleted already
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
3) Add a new function that will take care to delete all the expired application links:
```go
// DeleteExpiredApplicationLinks deletes all the expired application links in the given context
func (k Keeper) DeleteExpiredApplicationLinks(ctx sdk.Context) {
    blockTimestamp := ctx.BlockTime().UnixNano()
    k.IterateExpiringApplicationLinks(ctx, blockTimestamp, func(_ int64, link types.ApplicationLink) (stop bool) {
    store := ctx.KVStore(k.storeKey)
    store.Delete(types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username))
    store.Delete(types.ExpirationTimeApplicationLinkKey(blockTimestamp, link.OracleRequest.ClientID))
    return false
})
}
```

The last step need to update the `BeginBlock` function in order to handle the checks and perform
the according expiring actions on the links:

```go
// BeginBlock returns the begin blocker for the profiles module.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
    am.keeper.DeleteExpiredApplicationLinks(ctx)
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