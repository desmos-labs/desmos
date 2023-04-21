# ADR 020: Post ownership transfer

## Changelog
- April 21th, 2023: First draft;

## Status

Proposed

## Abstract

This ADR introduces a new feature that enables users to transfer post ownership to another person.

## Context

Desmos is a social network protocol that allows users to create, share, and engage with content on a decentralized platform. As of now, Desmos does not provide a feature that allows users to transfer the ownership of their posts to other users. This has caused inconvenience for users who wish to transfer ownership of their posts, as they have to create new posts and lose the engagement history and feedback of the original post. Therefore, the introduction of the new feature that enables users to transfer post ownership to another person aims to address this issue and provide users with more control over their content. The proposed feature is expected to improve user experience and enhance the functionality of the Desmos protocol.

## Decision

We will add an __owner__ field to the post structure and implement a new handler that allows users to transfer post ownership to another person. The upcoming changes are as follows:

### Type

```proto
message Post {
  uint64 subspace_id = 1;
  uint32 section_id = 2;
  
  ...skip

  google.protobuf.Timestamp last_edited_date = 13;
  
  // Owner of the post
  string owner = 14;
}
```

### `Msg` Service

```proto
service Msg {
    // ChangePostOwner allows users to transfer the ownership of a post to another person
    rpc ChangePostOwner(MsgChangePostOwner) returns (MsgChangePostOwnerResponse);
}

// MsgChangePostOwner move a post to another subspace
message MsgChangePostOwner {
    // Id of the subspace
    uint64 subspace_id = 1;
    
    // Id of the post
    uint64 post_id = 2;

    // Address of the new post owner
    string new_owner = 3;
    
    // Address of the current post owner
    string owner = 5;
}
// MsgChangePostOwnerResponse defines the Msg/MsgChangePostOwner response type
message MsgChangePostOwnerResponse {}
```

## Consequences

### Backwards Compatibility

The solution outlined above is **not** backwards compatible and will require a migration script to update all existing posts to the new version. This script will handle the following tasks:
- migrate all posts to have a new __owner__ field.

### Positive

- Enable users to transfer the ownership of the post

### Negative

(none known)

### Neutral

(none known)

## References