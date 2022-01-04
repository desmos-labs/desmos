# ADR 004: Expiration of application links

## Changelog

- September 20th, 2021: Initial draft;
- September 21th, 2021: Moved from DRAFT to PROPOSED;
- December  22th, 2021: First review;
- January   04th, 2022: Second review.

## Status

PROPOSED

## Abstract

Currently when a user links their centralized application with the Desmos profile, the created link contains a timestamp of when it has been created.   
Since centralized applications' username can be switched and sold, we SHOULD implement an "expiration date" system on links. 
This means that after a certain amount of time passes, the link will be automatically marked as invalid, and the user has to connect it again in order to keep it valid.

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
  // ExpirationTime represents the time in which the link will expire
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
// IterateExpiringApplicationLinks iterates through all the expiring application links references at the given expirationTimestamp
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
        k.deleteApplicationLinkStoreKeys(ctx, link)
        return false
    })
}
```
 * The above function uses `deleteApplicationLinkStoreKeys` which is a private method that's used also from
 the `DeleteApplicationLink` method in order to avoid code repetition and pursue the DRY principle:

 ```go
   // deleteApplicationLinkStoreKeys deletes all the store keys related to an application link with the given
   func (k Keeper) deleteApplicationLinkStoreKeys(ctx sdk.Context, link types.ApplicationLink) {
	    store := ctx.KVStore(k.storeKey)
	    store.Delete(types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username))
	    store.Delete(types.ApplicationLinkClientIDKey(link.OracleRequest.ClientID))
	    store.Delete(types.ExpirationTimeApplicationLinkKey(link.ExpirationTime.UnixNano(), link.OracleRequest.ClientID))
   }
```

* The `DeleteApplicationLink` function is edited accordingly as follows:
```go
// DeleteApplicationLink removes the application link associated to the given user,
// for the given application and username
func (k Keeper) DeleteApplicationLink(ctx sdk.Context, user string, application, username string) error {
    [...]
    // Delete the application link data and keys
    k.deleteApplicationLinkStoreKeys(ctx, link)

    return nil
}
```


4) Update the `BeginBlock` function in order to handle the checks and perform
the according expiring actions on the links:

```go
// BeginBlock returns the begin blocker for the profiles module.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
    am.keeper.DeleteExpiredApplicationLinks(ctx)
}
```

We also need to add a new `AppLinksParams` parameter to the `x/profiles` params, which will contain
an `expiration_time` timestamp that will be added to all the `AppLinks` before they're stored in the chain. 
The implementation for the new parameter is the following:
1) Create the new `AppLinkParams`:
```protobuf
message AppLinksParams {
  option (gogoproto.goproto_getters) = false;

  google.protobuf.Timestamp expiration_time = 1 [
    (gogoproto.stdtime) = true,
    (gogoproto.moretags) = "yaml:\"expiration_time\"",
    (gogoproto.nullable) = false
  ];
}
```

2) Add the `AppLinkParams` to the `Params`
```protobuf
// Params contains the parameters for the profiles module
message Params {
  [...]
  AppLinksParams appLinks = 5 [
    (gogoproto.nullable) = false,
    (gogoproto.moretags) = "yaml:\"applinks\""
  ];
}
```

3) Edit the `StartProfileConnection` in order to insert add the `expirationTime` properly to the `AppLink` at the moment of its creation:
```go
func (k Keeper) StartProfileConnection(
	ctx sdk.Context,
	applicationData types.Data,
	dataSourceCallData string,
	sender sdk.AccAddress,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {
	[...]
    // Calculate the expiration time for the link
    expirationTimeParam := k.GetParams(ctx).AppLinks.ExpirationTime
    creationTime := ctx.BlockTime()
    expirationTime := creationTime.Add(time.Duration(expirationTimeParam.UnixNano()))

    // Store the connection
    err = k.SaveApplicationLink(ctx, types.NewApplicationLink(
            sender.String(),
            applicationData,
            types.ApplicationLinkStateInitialized,
            types.NewOracleRequest(
                0,
                oraclePrams.ScriptID,
                types.NewOracleRequestCallData(applicationData.Application, dataSourceCallData),
                clientID,
            ),
            nil,
            creationTime,
            expirationTime,
        ))
	[...]
}
```

## Consequences

### Backwards Compatibility

This update will affect the `ApplicationLink` object by adding the new `ExpirationBlockHeight` 
field on it breaking the compatibility with the previous versions of the software. To allow
a smooth update and overcome these compatibility issues, we need to set up a proper migration
from the previous versions to the one that will include the additions contained in this ADR.

```go
func (m Migrator) Migrate(ctx sdk.Context) error {
    params := m.keeper.GetParams(ctx)
    blockTime := ctx.BlockTime()
    expirationTimeParam := params.AppLinks.ExpirationTime

    m.keeper.IterateApplicationLinks(ctx, func(index int64, link types.ApplicationLink) (stop bool) {
        expirationTime := link.CreationTime.Add(time.Duration(expirationTimeParam.UnixNano()))
        // if the existent app link is expired, delete it
        if expirationTime.Before(blockTime) {
            m.keeper.deleteApplicationLinkStoreKeys(ctx, link)
            return false
        }

        link.ExpirationTime = expirationTime
        // can this error be unchecked? it checks if the link is associated to a profile
        _ = m.keeper.SaveApplicationLink(ctx, link)
        return false
    })

    return nil
}
```

### Positive

* Considerably reduce the possibility of impersonation of entities and users of centralized apps;

### Negative

* Some more operations to perform in the overall handling of the app links

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#516](https://github.com/desmos-labs/desmos/issues/516)
- PR [#562](https://github.com/desmos-labs/desmos/pull/562)