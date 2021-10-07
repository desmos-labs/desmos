# ADR 001: Reports type field review

## Changelog

- September 9th, 2021: Initial draft
- September 14th, 2021: First proposal
- October 4th, 2021: First review
- October 7th, 2021: Second review

## Status

PROPOSED

## Abstract

Report's `type` field SHOULD be expanded and enhanced to make possible adding 
more reasons and elaborate more why a post has been reported. Instead of a single string field, 
we SHOULD implement an array of strings, renaming the field from `type` to `reasons`.  
In addition to this, each subspace will need to introduce its own default set of `reasons` that will be 
used in all the reports that will be made inside it. Not specifying any default reason will make impossible to 
make valid reports.

## Context

Centralized social networks give users the possibility to select multiple reasons when they're about to report a post.
By doing this, the report become itself more explicit and complete and will later help the moderator/community to judge and choose the best
way to handle it. 
Currently inside Desmos we don't handle the reporting system like described. Instead, we rely on our `Report` object 
which contains the `type` field. The `type` field takes care of displaying the reason behind the report, but at the moment,
it can only store 1 reason at time. Despite this was handy and quick during a first implementation attempt, it
may not be enough for future decentralized social networks and social enabled apps.  

Users need an exhaustive way to report contents on decentralized platforms, so having a repeated `type` field that allows
to do this should be the better option to improve the whole report system. This ADR proposes such changes inside the current
`Report` type.

## Decision

We will rename the `Report` `type` field to be `reasons` instead.  
As briefly described before, the `reasons` will contain an array of strings that specify the reasons why the post has been reported.

```protobuf
// Report is the struct of a post's reports
message Report {
  option (gogoproto.equal) = true;
  option (gogoproto.goproto_stringer) = true;
  
  // ID of the post for which the report has been created
  string post_id = 1 [
    (gogoproto.customname) = "PostID",
    (gogoproto.jsontag) = "post_id",
    (gogoproto.moretags) = "yaml:\"post_id\""
   ];
  
   // Identifies the reasons of the report
   repeated string reasons = 2 [ 
     (gogoproto.jsontag) = "reasons", 
     (gogoproto.moretags) = "yaml:\"reasons\"" 
   ];

   // User message
   string message = 3 [
    (gogoproto.jsontag) = "message",
    (gogoproto.moretags) = "yaml:\"message\""
  ];
   
  // Identifies the reporting user
   string user = 4 [ 
     (gogoproto.jsontag) = "user", 
     (gogoproto.moretags) = "yaml:\"user\"" 
   ];
 }
```

The values provided inside with the `reasons` field will need to be chosen from a set of default ones that developers specify
when creating a subspace. If a subspace has an empty set of valid `reasons` that users can choose from, each report 
created within that subspace will be considered invalid, de-facto disabling the support for reports there.

In order to let users add the report reasons to their subspace, we need to introduce a new
type of message called `MsgStoreReportReason` which will have the following representation:
```protobuf
message MsgStoreReportReason {
  option (gogoproto.equal) = false;
  option (gogoproto.goproto_getters) = false;
  
  string reason = 1 [
    (gogoproto.moretags) = "yaml:\"reason\"",
  ];
  
  string subspace = 2 [
    (gogoproto.moretags) = "yaml:\"subspace\""
  ];
  
  string admin = 3 [
    (gogoproto.moretags) = "yaml:\"admin\""
  ];
}
```

The system will make sure that the `reasons` inside each report emitted are correct with the following method:
```go
func (k Keeper) CheckReportValidity(ctx sdk.Context, report types.Report) error
```

## Consequences

### Backwards Compatibility

Considering the fact that the `Report` type lives inside the `posts` module that has yet to be released there are
no backwards compatibility issues.

### Positive

* A complete and more explicit way to express the reasons of a report;

### Negative

* Raise the overall complexity of the report process due to the fact that the system
  will call `CheckReportValidity` to ensure the correctness of the report.

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#507](https://github.com/desmos-labs/desmos/issues/507)