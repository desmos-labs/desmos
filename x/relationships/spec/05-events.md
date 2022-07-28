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

| Type                | Attribute Key | Attribute Value                                | 
|:--------------------|:--------------|:-----------------------------------------------|
| create_relationship | creator       | {userAddress}                                  |
| create_relationship | counterparty  | {counterpartyAddress}                          |
| create_relationship | subspace      | {subspaceID}                                   |
| message             | module        | relationships                                  |
| message             | action        | /desmos.relationships.v1.MsgCreateRelationship |
| message             | sender        | {userAddress}                                  |

### MsgDeleteRelationship

| Type                | Attribute Key | Attribute Value                                | 
|:--------------------|:--------------|:-----------------------------------------------|
| delete_relationship | creator       | {userAddress}                                  |
| delete_relationship | counterparty  | {counterpartyAddress}                          |
| delete_relationship | subspace      | {subspaceID}                                   |
| message             | module        | relationships                                  |
| message             | action        | /desmos.relationships.v1.MsgDeleteRelationship |
| message             | sender        | {userAddress}                                  |

### MsgBlockUser

| Type       | Attribute Key | Attribute Value                       | 
|:-----------|:--------------|:--------------------------------------|
| block_user | blocker       | {blockerAddress}                      |
| block_user | blocked       | {blockedAddress}                      |
| block_user | subspace      | {subspaceID}                          |
| message    | module        | relationships                         |
| message    | action        | /desmos.relationships.v1.MsgBlockUser |
| message    | sender        | {userAddress}                         |

### MsgUnblockUser

| Type         | Attribute Key | Attribute Value                         | 
|:-------------|:--------------|:----------------------------------------|
| unblock_user | blocker       | {blockerAddress}                        |
| unblock_user | blocked       | {blockedAddress}                        |
| unblock_user | subspace      | {subspaceID}                            |
| message      | module        | relationships                           |
| message      | action        | /desmos.relationships.v1.MsgUnblockUser |
| message      | sender        | {userAddress}                           |