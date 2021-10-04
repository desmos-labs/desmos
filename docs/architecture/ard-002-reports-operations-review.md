# ADR 002: Reports save operation

## Changelog

- September 13th, 2021: Initial draft;
- September 14th, 2021: Moved from DRAFT to PROPOSED;
- September 15th, 2021: First review;
- October 4th, 2021: Moved the "delete report" part in a different ADR (ADR-007)

## Status

PROPOSED

## Abstract

The `report` type actually gives only the possibility to `create` a report that can't be edited later.   
This method SHOULD be implemented to improve the `reports` management and standardize the way in which
structures can be manipulated in Desmos.

## Context

The idea behind Desmos `reports` is to give users a way to express dissent toward a post. Later, 
developers will choose how to deal with these reports by implementing their own ToS.
Since Desmos is a protocol that will serve as a common field for many social-enabled applications, 
it would have been extremely complicated to build a one-for-all moderation system. 
Different social apps have different scopes, and want to deal with contents in different ways.   
In order to leave everyone free to manage the contents of their social, the creation of the best
ToS-agnostic system to report contents become crucial.

## Decision

To enhance the actual report system evolving it from what we pointed out before, 
we SHOULD edit the actual one and expand it to match the common
operations we're already using in other types like `Profile`.   
The two main operations that will comprise the new reporting system will be:
 * `Save`: to both create and edit a `report`;
 * `Delete`: to delete a previously made `report`. (Discussed here: [ADR-007]())

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
Currently, the inner logic doesn't allow users to edit a previously made report, making it de-facto immutable.   
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

## Consequences

### Backwards Compatibility

Considering that the `posts` module where the report system live is still in staging mode, 
the backwards compatibility is not relevant as there won't be any issue related to it.

### Positive

* A CRUD complete report system;
* Alignment with the `profiles` module way to handle CRUD ops;
* Relieves developers of the burden of extending the system on their own on the top of dAPPs.

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#575](https://github.com/desmos-labs/desmos/issues/575)