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

| **Type**    | **Attribute Key** | **Attribute Value**           | 
|:------------|:------------------|:------------------------------|
| create_post | subspace_id       | {subspaceID}                  |
| create_post | section_id        | {sectionID}                   |
| create_post | post_id           | {postID}                      |
| create_post | author            | {userAddress}                 |
| create_post | creation_time     | {CreationTime}                |
| message     | module            | posts                         |
| message     | action            | desmos.posts.v2.MsgCreatePost |
| message     | sender            | {userAddress}                 |

### MsgEditPost

| **Type**  | **Attribute Key** | **Attribute Value**         | 
|:----------|:------------------|:----------------------------|
| edit_post | subspace_id       | {subspaceID}                |
| edit_post | post_id           | {postID}                    |
| edit_post | last_edit_time    | {LastEditTime}              |
| message   | module            | posts                       |
| message   | action            | desmos.posts.v2.MsgEditPost |
| message   | sender            | {userAddress}               |

### MsgDeletePost

| **Type**    | **Attribute Key** | **Attribute Value**           | 
|:------------|:------------------|:------------------------------|
| delete_post | subspace_id       | {subspaceID}                  |
| delete_post | post_id           | {postID}                      |
| message     | module            | posts                         |
| message     | action            | desmos.posts.v2.MsgDeletePost |
| message     | sender            | {userAddress}                 |

### MsgAddPostAttachment

| **Type**            | **Attribute Key** | **Attribute Value**                  | 
|:--------------------|:------------------|:-------------------------------------|
| add_post_attachment | subspace_id       | {subspaceID}                         |
| add_post_attachment | post_id           | {postID}                             |
| add_post_attachment | attachment_id     | {attachmentID}                       |
| add_post_attachment | last_edit_time    | {lastEditTime}                       |
| message             | module            | posts                                |
| message             | action            | desmos.posts.v2.MsgAddPostAttachment |
| message             | sender            | {userAddress}                        |

### MsgRemovePostAttachment

| **Type**               | **Attribute Key** | **Attribute Value**                     | 
|:-----------------------|:------------------|:----------------------------------------|
| remove_post_attachment | subspace_id       | {subspaceID}                            |
| remove_post_attachment | post_id           | {postID}                                |
| remove_post_attachment | attachment_id     | {attachmentID}                          |
| remove_post_attachment | last_edit_time    | {lastEditTime}                          |
| message                | module            | posts                                   |
| message                | action            | desmos.posts.v2.MsgRemovePostAttachment |
| message                | sender            | {userAddress}                           |    


### MsgAnswerPoll

| **Type**    | **Attribute Key** | **Attribute Value**           | 
|:------------|:------------------|:------------------------------|
| answer_poll | subspace_id       | {subspaceID}                  |
| answer_poll | post_id           | {postID}                      |
| answer_poll | poll_id           | {pollID}                      |
| message     | module            | posts                         |
| message     | action            | desmos.posts.v2.MsgAnswerPoll |
| message     | sender            | {userAddress}                 |

## Keeper

| **Type**   | **Attribute Key** | **Attribute Value** | 
|:-----------|:------------------|:--------------------|
| tally_poll | subspace_id       | {subspaceID}        |
| tally_poll | post_id           | {postID}            |
| tally_poll | poll_id           | {pollID}            |