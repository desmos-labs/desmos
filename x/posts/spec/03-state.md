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
The next attachment ID is stored tied to the subspace ID and the post ID where the next attachment will be added:

* `0x10 | Subspace ID | Post ID | -> bytes(NextAttachmentID)`

## Attachment
The attachment are stored in a way that allows to easily query all the posts' attachments:

* `0x11 | SubspaceID | PostID | Attachment ID | -> ProtocolBuffer(Attachment)`

## User Answer
The user answers are stored to allow an easy way to query a poll's users answers:

* `0x20 | Subspace ID | Post ID | Poll ID | -> ProtocolBuffer(UserAnswer)`

## Active poll queue 
The active poll queue allows to append an active poll to a queue of active polls.
Later on this key is used to check on each of them and set them as inactive at the correct poll's end time:

* `0x21 | End Time | Subspace ID | Post ID | poll ID | -> bytes(PollID)`