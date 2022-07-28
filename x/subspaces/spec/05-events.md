---
id: events
title: Events
sidebar_label: Events
slug: events
---

# Events

The subspaces module emits the following events:

## Handlers

### MsgCreateSubspace

| **Type**        | **Attribute Key** | **Attribute Value**                   | 
|:----------------|:------------------|:--------------------------------------|
| create_subspace | subspace_id       | {subspaceID}                          |
| create_subspace | subspace_name     | {subspaceName}                        |
| create_subspace | subspace_creator  | {subspaceCreator}                     |
| create_subspace | creation_date     | {subspaceCreationTime}                |
| message         | module            | subspaces                             |
| message         | action            | desmos.subspaces.v2.MsgCreateSubspace |
| message         | sender            | {userAddress}                         |

### MsgEditSubspace

| **Type**      | **Attribute Key** | **Attribute Value**                 | 
|:--------------|:------------------|:------------------------------------|
| edit_subspace | subspace_id       | {subspaceID}                        |
| message       | module            | subspaces                           |
| message       | action            | desmos.subspaces.v2.MsgEditSubspace |
| message       | sender            | {userAddress}                       |


### MsgDeleteSubspace

| **Type**        | **Attribute Key** | **Attribute Value**                   | 
|:----------------|:------------------|:--------------------------------------|
| delete_subspace | subspace_id       | {subspaceID}                          |
| message         | module            | subspaces                             |
| message         | action            | desmos.subspaces.v2.MsgDeleteSubspace |
| message         | sender            | {userAddress}                         |

### MsgCreateSection

| **Type**       | **Attribute Key** | **Attribute Value**                  |
|:---------------|:------------------|:-------------------------------------|
| create_section | subspace_id       | {subspaceID}                         |
| create_section | section_id        | {sectionID}                          |
| message        | module            | subspaces                            |
| message        | action            | desmos.subspaces.v2.MsgCreateSection |
| message        | sender            | {userAddress}                        |

### MsgEditSection

| **Type**     | **Attribute Key** | **Attribute Value**                | 
|:-------------|:------------------|:-----------------------------------|
| edit_section | subspace_id       | {subspaceID}                       |
| edit_section | section_id        | {sectionID}                        |
| message      | module            | subspaces                          |
| message      | action            | desmos.subspaces.v2.MsgEditSection |
| message      | sender            | {userAddress}                      |

### MsgMoveSection

| **Type**     | **Attribute Key** | **Attribute Value**                | 
|:-------------|:------------------|:-----------------------------------|
| move_section | subspace_id       | {subspaceID}                       |
| move_section | section_id        | {sectionID}                        |
| message      | module            | subspaces                          |
| message      | action            | desmos.subspaces.v2.MsgMoveSection |
| message      | sender            | {userAddress}                      |

### MsgDeleteSection

| **Type**       | **Attribute Key** | **Attribute Value**                  | 
|:---------------|:------------------|:-------------------------------------|
| delete_section | subspace_id       | {subspaceID}                         |
| delete_section | section_id        | {sectionID}                          |
| message        | module            | subspaces                            |
| message        | action            | desmos.subspaces.v2.MsgDeleteSection |
| message        | sender            | {userAddress}                        |

### MsgCreateUserGroup

| **Type**          | **Attribute Key** | **Attribute Value**                    | 
|:------------------|:------------------|:---------------------------------------|
| create_user_group | subspace_id       | {subspaceID}                           |    
| create_user_group | user_group_id     | {userGroupID}                          |    
| message           | module            | subspaces                              |
| message           | action            | desmos.subspaces.v2.MsgCreateUserGroup |
| message           | sender            | {userAddress}                          |

### MsgEditUserGroup

| **Type**        | **Attribute Key** | **Attribute Value**                  | 
|:----------------|:------------------|:-------------------------------------|
| edit_user_group | subspace_id       | {subspaceID}                         |
| edit_user_group | user_group_id     | {userGroupID}                        |
| message         | module            | subspaces                            |
| message         | action            | desmos.subspaces.v2.MsgEditUserGroup |
| message         | sender            | {userAddress}                        |

### MsgMoveUserGroup

| **Type**        | **Attribute Key** | **Attribute Value**                  | 
|:----------------|:------------------|:-------------------------------------|
| move_user_group | subspace_id       | {subspaceID}                         |
| move_user_group | user_group_id     | {userGroupID}                        |
| message         | module            | subspaces                            |
| message         | action            | desmos.subspaces.v2.MsgMoveUserGroup |
| message         | sender            | {userAddress}                        |

### MsgSetUserGroupPermissions

| **Type**                   | **Attribute Key** | **Attribute Value**                            | 
|:---------------------------|:------------------|:-----------------------------------------------|
| set_user_group_permissions | subspace_id       | {subspaceID}                                   |
| set_user_group_permissions | user_group_id     | {userGroupID}                                  |
| message                    | module            | subspaces                                      |
| message                    | action            | desmos.subspaces.v2.MsgSetUserGroupPermissions |
| message                    | sender            | {userAddress}                                  |

### MsgDeleteUserGroup

| **Type**          | **Attribute Key** | **Attribute Value**                    | 
|:------------------|:------------------|:---------------------------------------|
| delete_user_group | subspace_id       | {subspaceID}                           |
| delete_user_group | user_group_id     | {userGroupID}                          |
| message           | module            | subspaces                              |
| message           | action            | desmos.subspaces.v2.MsgDeleteUserGroup |
| message           | sender            | {userAddress}                          |

### MsgAddUserToUserGroup

| **Type**         | **Attribute Key** | **Attribute Value**                       | 
|:-----------------|:------------------|:------------------------------------------|
| add_group_member | subspace_id       | {subspaceID}                              |
| add_group_member | user_group_id     | {userGroupID}                             |
| add_group_member | user              | {userAddress}                             |
| message          | module            | subspaces                                 |
| message          | action            | desmos.subspaces.v2.MsgAddUserToUserGroup |
| message          | sender            | {userAddress}                             |

### MsgRemoveUserFromUserGroup

| **Type**            | **Attribute Key** | **Attribute Value**                            | 
|:--------------------|:------------------|:-----------------------------------------------|
| remove_group_member | subspace_id       | {subspaceID}                                   |
| remove_group_member | user_group_id     | {userGroupID}                                  |
| remove_group_member | user              | {userAddress}                                  |
| message             | module            | subspaces                                      |
| message             | action            | desmos.subspaces.v2.MsgRemoveUserFromUserGroup |
| message             | sender            | {userAddress}                                  |

### MsgSetUserPermissions

| **Type**             | **Attribute Key** | **Attribute Value**                 | 
|:---------------------|:------------------|:------------------------------------|
| set_user_permissions | subspace_id       | {subspaceID}                        |
| set_user_permissions | user              | {userAddress}                       |
| message              | module            | subspaces                           |
| message              | action            | desmos.subspaces.v2.MsgEditSubspace |
| message              | sender            | {userAddress}                       |