# ADR 019: Subspace post migration

## Changelog
- April 20th, 2023: First draft;
- April 24th, 2023: First review;
- May 2th, 2023: Second review;

## Status

Proposed

## Abstract

This ADR introduces a new feature that enables post owners to move their own content between subspaces.

## Context

Currently, when a user creates a post, they are assigned the author role and become the only user who can edit or manage the post. While this approach works, it may not meet all use cases. For instance, we don't allow users to transfer ownership of a post to another user, which could be useful in certain scenarios. For example, an employee might want to draft a post and later transfer ownership to their company.

## Decision

In order to move the post properly, we will implement an operation that involves the following functions:

1. delete all reactions, reports and references that are incompatible with the new subspace of the post
2. reset the moved post `PostReferences` to be __empty__
3. reset the moved post `Conversation` ID to be __0__
4. set the proper `PostID` within target subspace

### `Msg` Service

We will implement a new handler that allows users to move their own posts through a message.

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

### Hooks

Currently, reactions and report reasons are not compatible across different subspaces. This means that the IDs used to identify reactions and reasons may differ between subspaces. To address this limitation, we propose implementing a hook to notify `x/reactions`, `x/reports` and other references modules, which will be used to remove references to the posts. The hook will be as follows:

```go
func (k Keeper) AfterPostMoved(ctx sdk.Context, subspaceID uint64, postID uint64) {
    // Remove all reactions, reports and other references to the post
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
