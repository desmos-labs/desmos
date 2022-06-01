# ADR 011: Reports module

## Changelog

- April 12th, 2022: Initial draft

## Status

DRAFT

## Abstract

This ADR contains the specification of the `x/reports` module which will allow users to report other users and/or content posted inside subspaces.

## Context

One of the most important parts of a functional social network is the set of rules that constitute the Terms of Service, and how they are enforced. In most social networks the key for a good user experience is the reporting system: users are able to report misbehaving users or inappropriate content that gets later reviewed and taken care of by moderators. 

Inside Desmos we MUST give users a set of tools that allow them to perform the same actions on all the social networks that they will have access to. At the same time, we SHOULD NOT enforce any high-level Term of Service since each subspace that will be built on Desmos might want to have very different rules from the other ones (i.e. an adult content social network will most likely have a very different set of rules from a kids social network). Instead, we should allow subspace owners and admins to register the various reasons a user/content can be reported for.

## Decision

We will implement a new module named `x/reports` that allows users to report either a misbehaving user or a bad content inside subspaces where they have the permission to do so.

The same module will also allow subspace owners to register their own custom supported reasons for reports. 

### Types
Reports can be of two different types: `UserReport` for reporting a misbehaving user, or `PostReport` for reporting a bad post. The whole system must be thought in order to support additional report types in the future (i.e. `ProposalReport` to report a bad on-chain governance proposal). 

#### Report
```protobuf
// Report contains the data of a generic report
message Report {
  // Id of the report inside the subspace
  required uint64 id = 1;
  
  // Id of the reason this report has been created for
  required uint32 reason_id = 2;
  
  // Message attached to this report
  optional string message = 3;
  
  // Address of the reporter 
  required string reporter = 4;

  // Target of the report
  oneof Target {
    UserTarget user_data = 5;
    PostTarget post_data = 6;
  }
}

// UserTarget contains the data of a report about a user
message UserTarget {
  // Id of the subspace inside which the user has been reported
  required uint64 subspace_id = 1;
  
  // Address of the reported user
  required string user = 2;
}

// PostTarget contains the data of a report about a post
message PostTarget {
  // Id of the subspace inside which the reported post is
  required uint64 subspace_id = 1;
  
  // Id of the reported post 
  required uint64 post_id = 2;
}
```

### Reason
Each subspace owner/admin SHOULD be able to define a custom set of supported reasons for which users/posts can be reported for. Each reason MUST contain a human-readable title that identifies the reason itself. 

```protobuf
// Reason contains the data about a reporting reason
message Reason {
  // Id of the reason inside the subspace 
  required uint32 id = 1;
  
  // Title of the reason
  required string title = 2;
  
  // Extended description of the reason and the cases it applies to
  optional string description = 3;
}
```

### Params 
In order to make it easier to set up initially supported reporting reasons for new subspaces, we are going to provide all the subspaces with a list from which they can easily pick reasons that should be supported inside their own subspace as well.

```protobuf
// Params contains the module parameters
message Params {
  // List of available reasons from which new subspaces can pick their default ones
  repeated Reason reasons = 1;
}
```

### `Msg` Service
We will allow the following operations: 
- crete a new report
- delete an existing report 
- manage supported reporting reasons (add, remove a reason)

> NOTE  
> The ability to edit a report will **not** be allowed in order to avoid ever-changing reports that can make the moderation work a lot more complicated.

```protobuf
service Msg {
  // CreateReport allows to create a new report
  rpc CreateReport(MsgCreateReport) returns (MsgCreateReportResponse);
  
  // DeleteReport allows to delete an existing report
  rpc DeleteReport(MsgDeleteReport) returns (MsgDeleteReportResponse);
  
  // SupportReason allows to support one of the reasons present inside the module params 
  rpc SupportReason(MsgSupportReason) returns (MsgSupportReasonResponse);
  
  // AddReason allows to add a new supported reporting reason
  rpc AddReason(MsgAddReason) returns (MsgAddReasonResponse);
  
  // RemoveReason allows to remove a supported reporting reason
  rpc RemoveReason(MsgRemoveReason) returns (MsgRemoveReasonResponse);
}

// MsgCreateReport represents the message to be used to create a report
message MsgCreateReport {
  // Id of the reason this report has been created for
  required uint32 reason_id = 1;

  // Message attached to this report
  optional string message = 2;

  // Address of the reporter 
  required string reporter = 3;

  // Target of the report
  oneof Target {
    UserTarget user_data = 4;
    PostTarget post_data = 5;
  }
}

// MsgCreateReportResponse represents the Msg/CreateReport response type
message MsgCreateReportResponse {
  // Id of the newly created report
  required uint64 report_id = 1;
}

// MsgDeleteReport represents the message to be used when deleting a report
message MsgDeleteReport {
  // Id of the subspace that contains the report to be deleted 
  required uint64 subspace_id = 1;
  
  // Id of the report to be deleted 
  required uint64 report_id = 2; 
  
  // Address of the user deleting the report 
  required string signer = 3;
}

// MsgDeleteReportResponse represents the Msg/DeleteReport response type
message MsgDeleteReportResponse {}

// MsgSupportReason represents the message to be used when wanting to support a reason from the module params
message MsgSupportReason {
  // Id of the subspace for which to support the reason
  required uint64 subspace_id = 1;
  
  // Id of the reason that should be supported
  required uint32 reason_id = 2;
  
  // Address of the user signing the message
  required string signer = 3;
}

// MsgSupportReasonResponse represents the Msg/SupportReason response type
message MsgSupportReasonResponse {}

// MsgAddReason represents the message to be used when adding a new supported reason
message MsgAddReason {
  // Id of the subspace for which to add the reason 
  required uint64 subspace_id = 1;
  
  // Title of the reason
  required string title = 2;

  // Extended description of the reason and the cases it applies to
  optional string description = 3;
  
  // Address of the user adding the supported reason
  required string signer = 4;
}

// MsgAddReasonResponse represents the Msg/AddReason response type
message MsgAddReasonResponse {
  // Id of the newly supported reason
  required uint32 reason_id = 1;
}

// MsgRemoveReason represents the message to be used when removing an exiting reporting reason
message MsgRemoveReason {
  // Id of the subspace from which to remove the reason 
  required uint64 subspace_id = 1;

  // Id of the reason to be deleted
  required uint32 reason_id = 2;

  // Address of the user adding the supported reason
  required string signer = 3;
}

// MsgRemoveReasonResponse represents the Msg/RemoveReason response type
message MsgRemoveReasonResponse {
  // Id of the newly supported reason
  required uint32 reason_id = 1;
}

```

### `Query` Service
```protobuf
service Query {
  // Reports allows to query the reports for a specific target
  rpc Reports(QueryReportsRequest) returns (QueryReportsResponse) {
    option (google.api.http).get = "/desmos/reports/v1/{subspace_id}/reports";
  }
  
  // Reasons allows to query the supported reporting reasons for a subspace
  rpc Reasons(QueryReasonsRequest) returns (QueryReasonsResponse) {
    option (google.api.http).get = "/desmos/reports/v1/{subspace_id}/reasons";
  }
  
  // Params allows to query the module parameters
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse) {
    option (google.api.http).get = "/desmos/reports/v1/params";
  }
}

// QueryReportsResponse is the request type for Query/Reports RPC method
message QueryReportsRequest {
  // Target to query the reports for
  oneof Target {
    UserTarget user_data = 1;
    PostTarget post_data = 2;
  }

  // pagination defines an optional pagination for the request.
  optional cosmos.base.query.v1beta1.PageRequest pagination = 3;
} 

// QueryReportsResponse is the response type for Query/Reports RPC method
message QueryReportsResponse {
  repeated Report reports = 1;
  required cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryReasonsRequest is the request type for Query/Reasons RPC method
message QueryReasonsRequest {
  // Id of the subspace to query the supported reporting reasons for
  required uint64 subspace_id = 1;
  
  // pagination defines an optional pagination for the request.
  optional cosmos.base.query.v1beta1.PageRequest pagination = 3;
}

// QueryReasonsResponse is the response type for Query/Reasons RPC method
message QueryReasonsResponse {
  repeated Reason reasons = 1;
  required cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryParamsRequest is the request type for Query/Params RPC method
message QueryParamsRequest {}

// QueryParamsResponse is the response type for Query/Params RPC method
message QueryParamsResponse {
  required Params params = 1;
}
```

## Consequences

### Backwards Compatibility

The changes described inside this ADR are **not** backward compatible. To solve this, we will rely on the `x/upgrade` module in order to properly add these new features inside a running chain. If necessary, to make sure no extra operation is performed, we should make sure that `fromVm[reportstypes.ModuleName]` is set to `1` before running the migrations, so that the `InitGenesis` method does not get called.

### Positive

- Allow users to report misbehaving users or bad content
- Allow subspace owners and admins to make sure ToS are respected more easily

### Negative

- Not known

### Neutral

- Require the `x/subspaces` to implement the following new permissions: 
  - `PermissionCreateReport` to allow creating new reports inside a subspace;
  - `PermissionDeleteReport` to allow deleting existing reports inside a subspace;
  - `PermissionManageReasons` to allow managing supported reporting reasons inside a subspace.

## Further Discussions

## References
