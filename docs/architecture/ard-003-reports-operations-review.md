# ADR 007: Reports delete operation

## Changelog

- October 4th, 2021: Proposed ADR

## Status

PROPOSED

## Abstract

The `report` type doesn't have a `delete` method to trash a report after creating it.   
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
 * `Save`: to both create and edit a `report` (discussed in [ADR-002]());
 * `Delete`: to delete a previously made `report`.

### Delete 
The `delete` operation will be the new introduction to the report system. It will add the possibility to remove a previously made report to a post.  
This operation, SHOULD be considered as a requirement and not only an add-on. It's in the interest of the devs that users can revert their actions for any reason:
re-consideration, mistake...etc.   
The following implementation, takes the message proposed inside [#575](https://github.com/desmos-labs/desmos/issues/575)
and add a new field to it. The `reason` field will be used to specify why a user is removing the report.  
This could be useful to understand users point of view and get some considerations out of a hot topic.
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

Considering that the `posts` module where the report system live is still in staging mode, 
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