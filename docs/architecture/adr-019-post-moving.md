# ADR 019: Post moving

## Changelog
- April 20th, 2023: First draft;

## Status

Proposed

## Abstract

This ADR introduces a new feature that enables post owners to move their own content to a different subspaces.

## Context

Currently, the only way to move a post to another subspace is by following process:
1. Create an identical post as the original post in the target subspace
2. Delete the original post on source subspace

This process is overly complicated and time-consuming because it requires the user to perform two steps and manually recreate the post in the new location. Additionally, it can lead to inconsistencies if the user forgets to include certain details. A more user-friendly process would be beneficial for improving efficiency and reducing the risk of inconsistencies.

## Decision

We will implement a new handler that allows users to move their own posts through a message.

## `Msg` Service

In order to enable users to move their own posts, we will need to implement a new message. 

```proto
service Msg {
    // MovePost allows users to move their own posts to another subspace
    rpc MovePost(MsgMovePost) returns (MsgMovePostResponse);
}

// MsgMovePost move a post to another subspace
message MsgMovePost {
    // Id of the subspace where the post is currently located
    uint64 subspace_id = 1;
    
    // Id of the post to be moved
    uint64 post_id = 2;

    // Id of the target subspace to which the post will be moved
    uint64 target_subspace_id = 3;
    
    // Id of the target section to which the post will be moved
    uint64 target_section_id = 4;

    // Address of the post owner
    string owner = 5;
}

// MsgMovePostResponse defines the Msg/MsgMovePost response type
message MsgMovePostResponse {
    // New id of the post in the target subspace
    uint64 post_id = 1;
}
```

## Consequences

### Backwards Compatibility

As the only major change will be to introduce a new message handler, this change is completely backward compatible.

### Positive

- Enable users to move their own posts easily.

### Negative

(none known)

### Neutral

(none known)

## References
