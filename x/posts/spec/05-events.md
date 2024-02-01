---
id: events
title: Events
sidebar_label: Events
slug: events
---

# Events

The posts module emits the following events:

## Handlers

### MsgCreatePost

| **Type**     | **Attribute Key** | **Attribute Value**           | 
|:-------------|:------------------|:------------------------------|
| created_post | subspace_id       | {subspaceID}                  |
| created_post | section_id        | {sectionID}                   |
| created_post | post_id           | {postID}                      |
| created_post | author            | {userAddress}                 |
| created_post | creation_time     | {CreationTime}                |
| message      | module            | posts                         |
| message      | action            | desmos.posts.v3.MsgCreatePost |
| message      | sender            | {userAddress}                 |

### MsgEditPost

| **Type**    | **Attribute Key** | **Attribute Value**         | 
|:------------|:------------------|:----------------------------|
| edited_post | subspace_id       | {subspaceID}                |
| edited_post | post_id           | {postID}                    |
| edited_post | last_edit_time    | {LastEditTime}              |
| message     | module            | posts                       |
| message     | action            | desmos.posts.v3.MsgEditPost |
| message     | sender            | {userAddress}               |

### MsgDeletePost

| **Type**     | **Attribute Key** | **Attribute Value**           | 
|:-------------|:------------------|:------------------------------|
| deleted_post | subspace_id       | {subspaceID}                  |
| deleted_post | post_id           | {postID}                      |
| message      | module            | posts                         |
| message      | action            | desmos.posts.v3.MsgDeletePost |
| message      | sender            | {userAddress}                 |

### MsgAddPostAttachment

| **Type**              | **Attribute Key** | **Attribute Value**                  | 
|:----------------------|:------------------|:-------------------------------------|
| added_post_attachment | subspace_id       | {subspaceID}                         |
| added_post_attachment | post_id           | {postID}                             |
| added_post_attachment | attachment_id     | {attachmentID}                       |
| added_post_attachment | last_edit_time    | {lastEditTime}                       |
| message               | module            | posts                                |
| message               | action            | desmos.posts.v3.MsgAddPostAttachment |
| message               | sender            | {userAddress}                        |

### MsgRemovePostAttachment

| **Type**                | **Attribute Key** | **Attribute Value**                     | 
|:------------------------|:------------------|:----------------------------------------|
| removed_post_attachment | subspace_id       | {subspaceID}                            |
| removed_post_attachment | post_id           | {postID}                                |
| removed_post_attachment | attachment_id     | {attachmentID}                          |
| removed_post_attachment | last_edit_time    | {lastEditTime}                          |
| message                 | module            | posts                                   |
| message                 | action            | desmos.posts.v3.MsgRemovePostAttachment |
| message                 | sender            | {userAddress}                           |    

### MsgAnswerPoll

| **Type**      | **Attribute Key** | **Attribute Value**           | 
|:--------------|:------------------|:------------------------------|
| answered_poll | subspace_id       | {subspaceID}                  |
| answered_poll | post_id           | {postID}                      |
| answered_poll | poll_id           | {pollID}                      |
| message       | module            | posts                         |
| message       | action            | desmos.posts.v3.MsgAnswerPoll |
| message       | sender            | {userAddress}                 |

### MsgUpdateParams

| **Type** | **Attribute Key** | **Attribute Value**             | 
|:---------|:------------------|:--------------------------------|
| message  | module            | posts                           |
| message  | action            | desmos.posts.v3.MsgUpdateParams |
| message  | sender            | {userAddress}                   |

### MsgMovePost

| **Type**   | **Attribute Key** | **Attribute Value**         | 
|:-----------|:------------------|:----------------------------|
| moved_post | subspace_id       | {subspaceID}                |
| moved_post | post_id           | {postID}                    |
| moved_post | new_subspace_id   | {newSubspaceID}             |
| moved_post | new_post_id       | {newPostID}                 |
| message    | module            | posts                       |
| message    | action            | desmos.posts.v3.MsgMovePost |
| message    | sender            | {userAddress}               |

### MsgRequestPostOwnerTransfer

| **Type**                      | **Attribute Key** | **Attribute Value**                         | 
|:------------------------------|:------------------|:--------------------------------------------|
| requested_post_owner_transfer | subspace_id       | {subspaceID}                                |
| requested_post_owner_transfer | post_id           | {postID}                                    |
| requested_post_owner_transfer | receiver          | {receiverAddress}                           |
| requested_post_owner_transfer | sender            | {senderAddress}                             |
| message                       | module            | posts                                       |
| message                       | action            | desmos.posts.v3.MsgRequestPostOwnerTransfer |
| message                       | sender            | {userAddress}                               |

### MsgCancelPostOwnerTransferRequest

| **Type**                     | **Attribute Key** | **Attribute Value**                               | 
|:-----------------------------|:------------------|:--------------------------------------------------|
| canceled_post_owner_transfer | subspace_id       | {subspaceID}                                      |
| canceled_post_owner_transfer | post_id           | {postID}                                          |
| canceled_post_owner_transfer | sender            | {senderAddress}                                   |
| message                      | module            | posts                                             |
| message                      | action            | desmos.posts.v3.MsgCancelPostOwnerTransferRequest |
| message                      | sender            | {userAddress}                                     |

### MsgAcceptPostOwnerTransferRequest

| **Type**                     | **Attribute Key** | **Attribute Value**                               | 
|:-----------------------------|:------------------|:--------------------------------------------------|
| accepted_post_owner_transfer | subspace_id       | {subspaceID}                                      |
| accepted_post_owner_transfer | post_id           | {postID}                                          |
| accepted_post_owner_transfer | new_subspace_id   | {newSubspaceID}                                   |
| accepted_post_owner_transfer | new_post_id       | {newPostID}                                       |
| accepted_post_owner_transfer | receiver          | {receiverAddress}                                 |
| message                      | module            | posts                                             |
| message                      | action            | desmos.posts.v3.MsgAcceptPostOwnerTransferRequest |
| message                      | sender            | {userAddress}                                     |

### MsgRefusePostOwnerTransferRequest

| **Type**                    | **Attribute Key** | **Attribute Value**                               | 
|:----------------------------|:------------------|:--------------------------------------------------|
| refused_post_owner_transfer | subspace_id       | {subspaceID}                                      |
| refused_post_owner_transfer | post_id           | {postID}                                          |
| refused_post_owner_transfer | receiver          | {receiverAddress}                                 |
| message                     | module            | posts                                             |
| message                     | action            | desmos.posts.v3.MsgRefusePostOwnerTransferRequest |
| message                     | sender            | {userAddress}                                     |

## Keeper

| **Type**     | **Attribute Key** | **Attribute Value** | 
|:-------------|:------------------|:--------------------|
| tallied_poll | subspace_id       | {subspaceID}        |
| tallied_poll | post_id           | {postID}            |
| tallied_poll | poll_id           | {pollID}            |