---
id: events
title: Events
sidebar_label: Events
slug: events
---

# Events

The relationships module emits the following events:

## Handlers

### MsgCreateRelationship

| Type                 | Attribute Key | Attribute Value                                | 
|:---------------------|:--------------|:-----------------------------------------------|
| created_relationship | creator       | {userAddress}                                  |
| created_relationship | counterparty  | {counterpartyAddress}                          |
| created_relationship | subspace_id   | {subspaceID}                                   |
| message              | module        | relationships                                  |
| message              | action        | /desmos.relationships.v1.MsgCreateRelationship |
| message              | sender        | {userAddress}                                  |

### MsgDeleteRelationship

| Type                 | Attribute Key | Attribute Value                                | 
|:---------------------|:--------------|:-----------------------------------------------|
| deleted_relationship | creator       | {userAddress}                                  |
| deleted_relationship | counterparty  | {counterpartyAddress}                          |
| deleted_relationship | subspace_id   | {subspaceID}                                   |
| message              | module        | relationships                                  |
| message              | action        | /desmos.relationships.v1.MsgDeleteRelationship |
| message              | sender        | {userAddress}                                  |

### MsgBlockUser

| Type         | Attribute Key | Attribute Value                       | 
|:-------------|:--------------|:--------------------------------------|
| blocked_user | blocker       | {blockerAddress}                      |
| blocked_user | blocked       | {blockedAddress}                      |
| blocked_user | subspace_id   | {subspaceID}                          |
| message      | module        | relationships                         |
| message      | action        | /desmos.relationships.v1.MsgBlockUser |
| message      | sender        | {userAddress}                         |

### MsgUnblockUser

| Type           | Attribute Key | Attribute Value                         | 
|:---------------|:--------------|:----------------------------------------|
| unblocked_user | blocker       | {blockerAddress}                        |
| unblocked_user | blocked       | {blockedAddress}                        |
| unblocked_user | subspace_id   | {subspaceID}                            |
| message        | module        | relationships                           |
| message        | action        | /desmos.relationships.v1.MsgUnblockUser |
| message        | sender        | {userAddress}                           |