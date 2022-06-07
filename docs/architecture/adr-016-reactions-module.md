# ADR 016: Reactions module

## Changelog

- 2022, June 06th: First draft;

## Status

DRAFT Not Implemented

## Abstract

This ADR contains the specification of the `x/reactions` module which will allow users to react to different posts inside a subspace.

## Context

One of the most commonly used features of any centralized social network are the so-called _reactions_. Initially known as _like_, this feature allows users to react to different posts with a limited set of emojis or a free text limited in length (in order to distinguish it from a comment).

Inside Desmos we MUST give all subspace owners the ability to decide whether reactions should be enabled or not inside their subspaces, and if each reaction should be composed of only emojis or allow even for texts. We SHOULD also allow subspace owners to register custom reactions the same way that Discord does, so that they can provide a customized user experience inside their platforms. 

## Decision

We will implement a new module named `x/reactions` that allows users to react to each post inside a subspace using one or more of the registered reactions if they have the permission to do so. 

The same module will also allow subspace owners to register their own custom reactions that users can use withing that subspace.

### Types 

#### Reaction
In order to allow for a better management of different reactions, each `Reaction` will accept two different types of values within: 

- `RegisteredReactionValue` should be used when the reaction references a registered reaction; 
- `FreeTextValue` should be used in all other cases. 

This will allow to perform more easily the custom checks based on the different reaction params that each subspace can set.

```protobuf
syntax = "proto3";

// Reaction contains the data of a single post reaction
message Reaction {
  // Id of the subspace inside which the reaction has been put 
  uint64 subspace_id = 1;
  
  // Id of the reaction within the subspace
  uint64 id = 2;
  
  // Id of the post to which the reaction is associated
  uint64 post_id = 3;
  
  // Value of the reaction.
  google.proto.Any value = 4;
  
  // Author of the reaction
  string author = 5;
}

// RegisteredReactionValue contains the details of a reaction value that
// references a reaction registered within the subspace
message RegisteredReactionValue {
  // Id of the registered reaction
  uint32 registered_reaction_id = 1; 
}

// FreeTextValue contains the details of a reaction value that
// is made of free text
message FreeTextValue {
  string text = 1;
}
```

### RegisteredReaction

```protobuf
syntax = "proto3";

// RegisteredReaction contains the details of a registered reaction within a subspace
message RegisteredReaction {
  // Id of the subspace for which this reaction has been registered 
  uint64 subspace_id = 1;
  
  // Id of the registered reaction
  uint32 id = 2;
  
  // Unique shorthand code associated to this reaction
  string shorthand_code = 3;
  
  // Value that should be displayed when using this reaction
  string display_value = 4;
}
```

### ReactionParams
In order to allow subspace owners to customize their users experience with reactions, each subspace will be able to set its own reaction params based on its need. 

```protobuf
syntax = "proto3";

// SubspaceReactionsParams contains the params related to a single subspace reactions
message SubspaceReactionsParams {
  // Id of the subspace for which these params are valid
  uint64 subspace_id = 1;
  
  // Params related to RegisteredReactionValue reactions
  RegisteredReactionValueParams registered = 2;
  
  // Params related to FreeTextValue reactions
  FreeTextValueParams free_text = 3;
}

// FreeTextValueParams contains the params for FreeTextValue based reactions
message FreeTextValueParams {
  // Whether FreeTextValue reactions should be enabled
  bool enabled = 1;
  
  // The max length that FreeTextValue reactions should have
  uint64 max_length = 2;
}

// RegisteredReactionValueParams contains the params for RegisteredReactionValue based reactions
message RegisteredReactionValueParams {
  // Whether RegisteredReactionValue reactions should be enabled
  bool enabled = 1;
}
```

### The `Msg` Service
We will allow the following operations:  
- add a post reaction; 
- remove a post reaction; 
- register a new reaction;
- remove a registered reaction;
- set a subspace's reaction params; 

```protobuf
syntax = "proto3";

// Msg defines the reactions Msg service.
service Msg {
  // AddReaction allows to add a post reaction
  rpc AddReaction(MsgAddReaction) returns (MsgAddReactionResponse);
  
  // RemoveReaction allows to remove an existing post reaction
  rpc RemoveReaction(MsgRemoveReaction) returns (MsgRemoveReactionResponse);
  
  // AddRegisteredReaction allows to registered a new supported reaction
  rpc AddRegisteredReaction(MsgAddRegisteredReaction) returns (MsgAddRegisteredReactionResponse);
  
  // RemoveRegisteredReaction allows to remove an existing supported reaction
  rpc RemoveRegisteredReaction(MsgRemoveRegisteredReaction) returns (MsgRemoveRegisteredReactionResponse);
  
  // SetReactionsParams allows to set the reactions params
  rpc SetReactionsParams(MsgSetReactionsParams) returns (MsgSetReactionsParamsResponse);
}

// MsgAddReaction represents the message to be used to add a post reaction
message MsgAddReaction {
  // Id of the subspace inside which the post to react to is 
  uint64 subspace_id = 1;
  
  // Id of the post to react to
  uint64 post_id = 2;
  
  // Value of the reaction
  google.proto.Any value = 3;
  
  // User reacting to the post
  string user = 4;
}

// MsgAddReactionResponse represents the Msg/AddReaction response type
message MsgAddReactionResponse {
  // Id of the newly added reaction
  uint64 reaction_id = 1;
}

// MsgRemoveReaction represents the message to be used to remove an 
// existing reaction from a post
message MsgRemoveReaction {
  // Id of the subspace inside which the reaction to remove is
  uint64 subspace_id = 1;
  
  // Id of the reaction to be removed
  uint64 reaction_id = 2;
  
  // User removing the reaction
  string user = 3;
}

// MsgRemoveReactionResponse represents the Msg/RemoveReaction response type
message MsgRemoveReactionResponse {}

// MsgAddRegisteredReaction represents the message to be used to 
// register a new supported reaction
message MsgAddRegisteredReaction {
  // Id of the subspace inside which this reaction should be registered
  uint64 subspace_id = 1;
  
  // Shorthand code of the reaction
  string shorthand_code = 2;
  
  // Display value of the reaction
  string display_value = 3;
  
  // User adding the supported reaction
  string user = 4;
}

// MsgAddRegisteredReactionResponse represents the 
// Msg/AddRegisteredReaction response type
message MsgAddRegisteredReactionResponse {
  // Id of the newly registered reaction
  uint32 registered_reaction_id = 1;
}

// MsgRemoveRegisteredReaction represents the message to be used to
// remove an existing registered reaction
message MsgRemoveRegisteredReaction {
  // Id of the subspace from which to remove the registered reaction
  uint64 subspace_id = 1;
  
  // Id of the registered reaction to be removed
  uint32 registered_reaction_id = 2;
  
  // User removing the registered reaction
  string user = 3;
}

// MsgRemoveRegisteredReactionResponse represents the 
// Msg/RemoveRegisteredReaction response type
message MsgRemoveRegisteredReactionResponse {}

// MsgSetReactionsParams represents the message to be used when setting 
// a subspace reactions params
message MsgSetReactionsParams {
  // Id of the subspace for which to set the params
  uint64 subspace_id = 1;

  // Params related to RegisteredReactionValue reactions
  RegisteredReactionValueParams registered = 2;

  // Params related to FreeTextValue reactions
  FreeTextValueParams free_text = 3;
}

// MsgSetReactionsParamsResponse represents the Msg/SetReactionsParams response type
message MsgSetReactionsParamsResponse {}
```


### The `Query` service
```protobuf
syntax = "proto3";

// Query defines the gRPC querier service.
service Query {
  // Reactions allows to query the reactions present inside a subspace
  rpc Reactions(QueryReactionsRequest) returns QueryReactionsResponse {
    option (google.api.http).get = "/desmos/reactions/v1/{subspace_id}/reactions";
  }

  // RegisteredReactions allows to query the registered reaction of a subspace
  rpc RegisteredReactions(QueryRegisteredReactionsRequest) returns QueryRegisteredReactionsResponse {
    option (google.api.http).get = "/desmos/reactions/v1/{subspace_id}/registered_reactions";
  }
  
  // ReactionsParams allows to query the reaction params of a subspace
  rpc ReactionsParams(QueryReactionsParamsRequest) returns QueryReactionsParamsResponse {
    option (google.api.http).get = "/desmos/reactions/v1/{subspace_id}/params";
  }
}

// QueryReactionsRequest is the request type for the Query/Reactions RPC method
message QueryReactionsRequest {
  // Id of the subspace to query the reactions for
  uint64 subspace_id = 1;
  
  // (optional) Post id to query the reactions for
  uint64 post_id = 2;

  // pagination defines an optional pagination for the request.
  optional cosmos.base.query.v1beta1.PageRequest pagination = 3;
}

// QueryReactionsResponse is the response type for the Query/Reactions RPC method
message QueryReactionsResponse {
  repeated Reaction reactions = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryReactionsParamsRequest is the request type for the 
// Query/ReactionsParams RPC method
message QueryReactionsParamsRequest {
  uint64 subspace_id = 1;
}

// QueryReactionsParamsResponse is the response type for the 
// Query/ReactionsParam RPC method
message QueryReactionsParamsResponse {
  // Params related to RegisteredReactionValue reactions
  RegisteredReactionValueParams registered = 1;

  // Params related to FreeTextValue reactions
  FreeTextValueParams free_text = 2;
}
```

## Consequences

### Backwards Compatibility

The changes described inside this ADR are **not** backward compatible. To solve this, we will rely on the `x/upgrade` module in order to properly add these new features inside a running chain. If necessary, to make sure no extra operation is performed, we should make sure that `fromVm[reportstypes.ModuleName]` is set to `1` before running the migrations, so that the `InitGenesis` method does not get called. At the same time, we should make sure that we set the following for all existing subspaces:
- `NextReactionID` to `1`;
- `NextRegisteredReactionID` to `1`;
- `ReactionParams` to `DefaultReactionParams`. 

### Positive

- Allow users to react to posts

### Negative

- Not known

### Neutral

- Not known

## Further Discussions

## Test Cases

- Make sure the migration works and `NextReactionID`, `NextRegisteredReactionID` and `ReactionParams` are properly set for all existing modules.

## References

- [Issue #890](https://github.com/desmos-labs/desmos/issues/890)
- [Discord reactions](https://support.discord.com/hc/en-us/articles/360041139231-Adding-Emojis-and-Reactions)