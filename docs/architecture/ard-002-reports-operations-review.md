# ADR 002: Reports save operation

## Changelog

- September 13th, 2021: Initial draft;
- September 14th, 2021: Moved from DRAFT to PROPOSED;
- September 15th, 2021: First review;
- October 4th, 2021: Moved the "delete report" part in a different ADR (ADR-007)
- October 7th, 2021: Further corrections

## Status

PROPOSED

## Abstract

Currently, we give users only the possibility to create a report that can't be edited later.
This missing option SHOULD be implemented to improve the reports management and standardize 
the way in which structures can be manipulated in Desmos.

## Context

The idea behind Desmos reports is to give users a way to express dissent toward a post. Later, 
developers will choose how to deal with these reports by implementing their own ToS.

## Decision

The current report system only allows a user to create a report and never edit it later.  
To allow this possibility, we should use the concept of __saving__ rather than __creating/editing__.

### Save
The `save` operation COULD be a revised version of what's already in staging mode.
The revision will introduce the improvements proposed and will be exposed 
to the CLI under the name of `MsgSaveReport`.
For simplicity, I will now use the first one, as it's the current one:
```protobuf
message MsgSaveReport {
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

There are no known backwards compatibility issues with this implementation. 

### Positive

* A CRUD partially-complete report system;
* Alignment with the `profiles` module way to handle CRUD operations;

### Negative

(none known)

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#575](https://github.com/desmos-labs/desmos/issues/575)