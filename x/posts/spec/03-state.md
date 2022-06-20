---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Next Post ID
The Next Post ID is stored tied to the subspace where it lives. It allows to query the ID to be used next for the newest
post created:

* `0x00 | Subspace ID | ->  bytes(NextPostID)`

## Post
A Post is stored tied to the subspace in which it was created. This allows to easily query:
- All the posts of a given Subspace;
- A specific post of a given Subspace.

* `0x01 | Subspace ID | Post ID | -> ProtocolBuffer(Post)` 

## Post section
The section reference is stored to enable the possibility of querying posts for a particular subspace's section:

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