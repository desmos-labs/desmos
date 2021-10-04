# ADR 001: Reports type field review

## Changelog

- September 9th, 2021: Initial draft
- September 14th, 2021: First proposal
- October 4th, 2021: First review

## Status

PROPOSED

## Abstract

Report's `type` field SHOULD be expanded and enhanced to make possible adding 
more reasons and elaborate more why a post has been reported. To standardise the different `reasons`
inside the network and across all the social networks there will be a group of default ones. These can be
extended by governance proposal later on.

## Context

Centralized social networks give users the possibility to select multiple reasons when they're about to report a post.
By doing this, the report become itself more explicit and complete and will later help the moderator/community to judge and choose the best
way to handle it. 
Actually inside Desmos we don't handle the reporting system like described. Instead, we rely on our `Report` object 
which contains the `type` field. The `type` field takes care to display the reason behind the report, but at the moment,
it can only store 1 reason at time. Despite this can be handy and quick during a first implementation attempt, it
may be not enough for future decentralized social networks and social enabled apps.  

Users need an exhaustive way to report contents on decentralized platforms, so having a repeated `type` field that allows
to do this should be the better option to improve the whole report system. This ADR proposes such changes inside the current
`Report` type.

## Decision

We will replace the `Report` `type` field with a more explicit `reasons` type.  
As briefly described before,
the `reasons` will be an array of strings that specify the reasons why the post has been reported.

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

The reasons need to be chosen from a set of default ones that developers specify.  
These default values will be stored inside each subspace so any dApp can have
its own set of reports reasons. Despite this, it will not be mandatory to set default values because  
dApps developers can choose to use the reports system or not.

In order to let devs add the default reports reasons to their dApp subspace, we need to introduce a new
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
* Already familiar for end-users.

### Negative

* Raise the overall complexity of the report process due to the fact that the system
  will call `CheckReportValidity` to ensure the correctness of the report.

### Neutral

(none known)

## Further Discussions

## Test Cases [optional]

## References

- Issue [#507](https://github.com/desmos-labs/desmos/issues/507)