---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Next Post ID
The next post id is stored tied to the subspace to which it refers:

* `0x00 | Subspace ID | ->  bytes(NextPostID)`

## Post
A post is stored using the subspace id and its id as the key. This allows to easily query:
- all the posts of a given subspace;
- a specific post of a given subspace.

* `0x01 | Subspace ID | Post ID | -> ProtocolBuffer(Post)` 

## Post section
The section in which a post is placed is stored to enable the possibility of querying posts for a particular subspace's section:

* `0x02 | Subspace ID | Section ID | Post ID | -> 0x01`

## Next Attachment ID
The next attachment id is stored tied to the subspace id and the post id to which it refers:

* `0x10 | Subspace ID | Post ID | -> bytes(NextAttachmentID)`

## Attachment
A post attachment is stored using the subspace id, post id and its id as the key. This allows to easily query all the posts' attachments:

* `0x11 | SubspaceID | PostID | Attachment ID | -> ProtocolBuffer(Attachment)`

## User Answer
A user answer to a poll is stored using the subspace id, post id and poll id as the key. This allows to easily query all the answers of a specific poll:

* `0x20 | Subspace ID | Post ID | Poll ID | -> ProtocolBuffer(UserAnswer)`

## Active poll queue 
Active polls are stored using the voting end time, subspace id, post id and poll id as the key. This allows to determine, at each block height, which polls should have their results tallied:

* `0x21 | End Time | Subspace ID | Post ID | Poll ID | -> bytes(PollID)`