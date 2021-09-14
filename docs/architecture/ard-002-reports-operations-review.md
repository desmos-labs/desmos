# ADR 002: Reports save and delete operations

## Changelog

- September 13th, 2021: Initial draft;
- September 14th, 2021: Moved from DRAFT to PROPOSED

## Status

PROPOSED

## Abstract

Inside Desmos, most of the common types you deal with has two methods: `save` (to create & edit) and `delete`.
The `report` type can only be created, without the possibility to edit or delete it later.
These methods SHOULD be implemented to improve the handling of `reports` and standardize the way in which
structures can be manipulated in Desmos.

## Context

The idea behind Desmos `reports` is to give users a way to express dissent toward a post leaving 
developers the freedom to implement their own ToS around them.
Since Desmos is a protocol that will serve as a common field for many social-enabled applications, 
it would have been extremely complicated to build a one-for-all reporting system. 
Different social apps have different scopes, and want to deal with contents in different ways.   
In order to leave everyone free to manage the contents of their social, the creation of the best
ToS-agnostic system to reports contents is crucial.

## Decision

In order to build a complete report system, we SHOULD edit the actual one and expand it to match the common
operations we're already using in other types like `Profile`.
The two main operations that will comprise the new reporting system will be:
 * `Save`: to both create and edit a `report`;
 * `Delete`: to delete a previously made `report`.

### Save
The `save` operation COULD be a revised version of what's already in staging mode.
The revision will introduce the improvements proposed inside the ADR-001 (if it will be implemented) and will be exposed 
to the CLI under the name of `MsgReportPost` or `MsgSaveReportPost`.   
For simplicity, I will now use the first one, as it's the current one:
```protobuf
message MsgReportPost {
  option (gogoproto.equal) = false;
  option (gogoproto.goproto_getters) = false;

  // Post id of the post to report
  string post_id = 1 [
    (gogoproto.customname) = "PostID",
    (gogoproto.jsontag) = "post_id",
    (gogoproto.moretags) = "yaml:\"post_id\""
  ];

  // The reasons of the report
  repeated string reasons = 2 [
    (gogoproto.jsontag) = "reasons",
    (gogoproto.moretags) = "yaml:\"reasons\""
  ];

  // User message
  string message = 3 [ 
    (gogoproto.moretags) = "yaml:\"message\"" 
  ];

  // The reporting user
  string user = 4 [ 
    (gogoproto.moretags) = "yaml:\"user\"" 
  ];
}
```
Actually, the intern logic doesn't allow users to edit a previously made report, making the report de-facto immutable. 
What we MAY do instead, is making this editable by removing the actual check over its uniqueness.  
Here an example of it:  
```go
// If the same report has already been inserted, it will be updated.
 func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) error {
 	store := ctx.KVStore(k.storeKey)
 	key := types.ReportStoreKey(report.PostID, report.User)
 	
 	if err := k.CheckReportValidity(ctx, report); err != nil {
 		return err
 	}
	store.Set(key, types.MustMarshalReport(k.cdc, report))
	k.Logger(ctx).Info("reported post", "post-id", report.PostID, "from", report.User)
 	return nil
 }
```

OLD version:
```go
// If the same report has already been inserted, it will be updated.
 func (k Keeper) SaveReport(ctx sdk.Context, report types.Report) error {
 	store := ctx.KVStore(k.storeKey)
 	key := types.ReportStoreKey(report.PostID, report.User)

 	// Check the existance of the report
 	if store.Has(key) {
 		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s already reported post with id %s",
 			report.User, report.PostID)
 	}

 	if err := k.CheckReportValidity(ctx, report); err != nil {
 		return err
 	}
	store.Set(key, types.MustMarshalReport(k.cdc, report))
	k.Logger(ctx).Info("reported post", "post-id", report.PostID, "from", report.User)
 	return nil
 }
```
### Delete 
The `delete` operation will be the new introduction to the report system.   
It will add the possibility to remove a previously made report to a post.  
This operation, SHOULD be considered as a requirement and not only an add-on.  
It's in the interest of the devs that users can revert their actions for any reason:
re-consideration, mistake...etc.   
The following implementation, takes the message proposed inside [#575](https://github.com/desmos-labs/desmos/issues/575)
and add a new field to it. The `reason` field will be used to specify why a user is removing the report.  
This could be useful in the future for possible social-analysis cases and studies on users beliefs around a specific theme.
Here the representation of the CLI message to delete the report:
```protobuf
 message MsgDeleteReport {
  option (gogoproto.equal) = false;
  option (gogoproto.goproto_getters) = false;

  string post_id = 1 [
    (gogoproto.customname) = "PostID",
    (gogoproto.jsontag) = "post_id",
    (gogoproto.moretags) = "yaml:\"post_id\""
  ];
  
  string reason = 2 [
    (gogoproto.jsontag) = "reason",
    (gogoproto.moretags) = "yaml:\"reason\""
  ];

  string user = 3 [
    (gogoproto.moretags) = "yaml:\"user\"" 
  ];
}
```

The inner logic will be handled by `DeleteReport` method of the `Keeper`:
```go
 func (k Keeper) DeleteReport(ctx sdk.Context, postID, user string) error {
 	store := ctx.KVStore(k.storeKey)
 	key := types.ReportStoreKey(postID, user)

 	if !store.Has(key) {
 		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
 			"The report of the post with id %s doesn't exist and cannot be removed",
 			postID)
 	}

 	store.Delete(key)
 	k.Logger(ctx).Info("deleted post report", "post-id", postID, "from", user)
 	return nil
 }
```

## Consequences

### Backwards Compatibility

Considering that the `posts` module where the report system live is still in staging, 
the backwards compatibility is not relevant as there won't be any issue related to it.

### Positive

* A CRUD complete report system;
* Alignment with the `profiles` module way to handle stuff;
* Relieves developers of the burden of extending the system on their own on the top of dAPPs.

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#575](https://github.com/desmos-labs/desmos/issues/575)