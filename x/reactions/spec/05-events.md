---
id: events
title: Events
sidebar_label: Events
slug: events
---

# Events

The reactions module emits the following events:

## Handlers

### MsgAddReaction

| **Type**     | **Attribute Key** | **Attribute Value**                | 
|:-------------|:------------------|:-----------------------------------|
| add_reaction | subspace_id       | {subspaceID}                       |
| add_reaction | post_id           | {postID}                           |
| add_reaction | reaction_id       | {reactionID}                       |
| add_reaction | user              | {userAddress}                      |
| message      | module            | reactions                          |
| message      | action            | desmos.reactions.v1.MsgAddReaction |
| message      | sender            | {userAddress}                      |

### MsgRemoveReaction

| **Type**        | **Attribute Key** | **Attribute Value**                   | 
|:----------------|:------------------|:--------------------------------------|
| remove_reaction | subspace_id       | {subspaceID}                          |
| remove_reaction | post_id           | {postID}                              |
| remove_reaction | reaction_id       | {reaction_id}                         |
| message         | module            | reactions                             |
| message         | action            | desmos.reactions.v1.MsgRemoveReaction |
| message         | sender            | {userAddress}                         |

### MsgAddRegisteredReaction

| **Type**                | **Attribute Key**      | **Attribute Value**                          | 
|:------------------------|:-----------------------|:---------------------------------------------|
| add_registered_reaction | subspace_id            | {subspaceID}                                 |
| add_registered_reaction | registered_reaction_id | {registeredReactionID}                       |
| message                 | module                 | reactions                                    |
| message                 | action                 | desmos.reactions.v1.MsgAddRegisteredReaction |
| message                 | sender                 | {userAddress}                                |

### MsgEditRegisteredReaction

| **Type**                 | **Attribute Key**      | **Attribute Value**                           | 
|:-------------------------|:-----------------------|:----------------------------------------------|
| edit_registered_reaction | subspace_id            | {subspaceID}                                  |
| edit_registered_reaction | registered_reaction_id | {registeredReactionID}                        |
| message                  | module                 | reactions                                     |
| message                  | action                 | desmos.reactions.v1.MsgEditRegisteredReaction |
| message                  | sender                 | {userAddress}                                 |

### MsgRemoveRegisteredReaction

| **Type**                   | **Attribute Key**      | **Attribute Value**                             | 
|:---------------------------|:-----------------------|:------------------------------------------------|
| remove_registered_reaction | subspace_id            | {subspaceID}                                    |
| remove_registered_reaction | registered_reaction_id | {registeredReactionID}                          |
| message                    | module                 | reactions                                       |
| message                    | action                 | desmos.reactions.v1.MsgRemoveRegisteredReaction |
| message                    | sender                 | {userAddress}                                   |

### MsgSetReactionsParams

| **Type**             | **Attribute Key** | **Attribute Value**                       | 
|:---------------------|:------------------|:------------------------------------------|
| set_reactions_params | subspace_id       | {subspaceID}                              |
| message              | module            | reactions                                 |
| message              | action            | desmos.reactions.v1.MsgSetReactionsParams |
| message              | sender            | {userAddress}                             | 
