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

| **Type**       | **Attribute Key** | **Attribute Value**                | 
|:---------------|:------------------|:-----------------------------------|
| added_reaction | subspace_id       | {subspaceID}                       |
| added_reaction | post_id           | {postID}                           |
| added_reaction | reaction_id       | {reactionID}                       |
| added_reaction | user              | {userAddress}                      |
| message        | module            | reactions                          |
| message        | action            | desmos.reactions.v1.MsgAddReaction |
| message        | sender            | {userAddress}                      |

### MsgRemoveReaction

| **Type**         | **Attribute Key** | **Attribute Value**                   | 
|:-----------------|:------------------|:--------------------------------------|
| removed_reaction | subspace_id       | {subspaceID}                          |
| removed_reaction | post_id           | {postID}                              |
| removed_reaction | reaction_id       | {reaction_id}                         |
| message          | module            | reactions                             |
| message          | action            | desmos.reactions.v1.MsgRemoveReaction |
| message          | sender            | {userAddress}                         |

### MsgAddRegisteredReaction

| **Type**                  | **Attribute Key**      | **Attribute Value**                          | 
|:--------------------------|:-----------------------|:---------------------------------------------|
| added_registered_reaction | subspace_id            | {subspaceID}                                 |
| added_registered_reaction | registered_reaction_id | {registeredReactionID}                       |
| message                   | module                 | reactions                                    |
| message                   | action                 | desmos.reactions.v1.MsgAddRegisteredReaction |
| message                   | sender                 | {userAddress}                                |

### MsgEditRegisteredReaction

| **Type**                   | **Attribute Key**      | **Attribute Value**                           | 
|:---------------------------|:-----------------------|:----------------------------------------------|
| edited_registered_reaction | subspace_id            | {subspaceID}                                  |
| edited_registered_reaction | registered_reaction_id | {registeredReactionID}                        |
| message                    | module                 | reactions                                     |
| message                    | action                 | desmos.reactions.v1.MsgEditRegisteredReaction |
| message                    | sender                 | {userAddress}                                 |

### MsgRemoveRegisteredReaction

| **Type**                    | **Attribute Key**      | **Attribute Value**                             | 
|:----------------------------|:-----------------------|:------------------------------------------------|
| removed_registered_reaction | subspace_id            | {subspaceID}                                    |
| removed_registered_reaction | registered_reaction_id | {registeredReactionID}                          |
| message                     | module                 | reactions                                       |
| message                     | action                 | desmos.reactions.v1.MsgRemoveRegisteredReaction |
| message                     | sender                 | {userAddress}                                   |

### MsgSetReactionsParams

| **Type**             | **Attribute Key** | **Attribute Value**                       | 
|:---------------------|:------------------|:------------------------------------------|
| set_reactions_params | subspace_id       | {subspaceID}                              |
| message              | module            | reactions                                 |
| message              | action            | desmos.reactions.v1.MsgSetReactionsParams |
| message              | sender            | {userAddress}                             | 
