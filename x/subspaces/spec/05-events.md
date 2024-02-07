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

| **Type**         | **Attribute Key** | **Attribute Value**                   | 
|:-----------------|:------------------|:--------------------------------------|
| created_subspace | subspace_id       | {subspaceID}                          |
| created_subspace | subspace_name     | {subspaceName}                        |
| created_subspace | subspace_creator  | {subspaceCreator}                     |
| created_subspace | creation_date     | {subspaceCreationTime}                |
| message          | module            | subspaces                             |
| message          | action            | desmos.subspaces.v3.MsgCreateSubspace |
| message          | sender            | {userAddress}                         |

### MsgEditSubspace

| **Type**        | **Attribute Key** | **Attribute Value**                 | 
|:----------------|:------------------|:------------------------------------|
| edited_subspace | subspace_id       | {subspaceID}                        |
| message         | module            | subspaces                           |
| message         | action            | desmos.subspaces.v3.MsgEditSubspace |
| message         | sender            | {userAddress}                       |

### MsgDeleteSubspace

| **Type**         | **Attribute Key** | **Attribute Value**                   | 
|:-----------------|:------------------|:--------------------------------------|
| deleted_subspace | subspace_id       | {subspaceID}                          |
| message          | module            | subspaces                             |
| message          | action            | desmos.subspaces.v3.MsgDeleteSubspace |
| message          | sender            | {userAddress}                         |

### MsgCreateSection

| **Type**        | **Attribute Key** | **Attribute Value**                  |
|:----------------|:------------------|:-------------------------------------|
| created_section | subspace_id       | {subspaceID}                         |
| created_section | section_id        | {sectionID}                          |
| message         | module            | subspaces                            |
| message         | action            | desmos.subspaces.v3.MsgCreateSection |
| message         | sender            | {userAddress}                        |

### MsgEditSection

| **Type**       | **Attribute Key** | **Attribute Value**                | 
|:---------------|:------------------|:-----------------------------------|
| edited_section | subspace_id       | {subspaceID}                       |
| edited_section | section_id        | {sectionID}                        |
| message        | module            | subspaces                          |
| message        | action            | desmos.subspaces.v3.MsgEditSection |
| message        | sender            | {userAddress}                      |

### MsgMoveSection

| **Type**      | **Attribute Key** | **Attribute Value**                | 
|:--------------|:------------------|:-----------------------------------|
| moved_section | subspace_id       | {subspaceID}                       |
| moved_section | section_id        | {sectionID}                        |
| message       | module            | subspaces                          |
| message       | action            | desmos.subspaces.v3.MsgMoveSection |
| message       | sender            | {userAddress}                      |

### MsgDeleteSection

| **Type**        | **Attribute Key** | **Attribute Value**                  | 
|:----------------|:------------------|:-------------------------------------|
| deleted_section | subspace_id       | {subspaceID}                         |
| deleted_section | section_id        | {sectionID}                          |
| message         | module            | subspaces                            |
| message         | action            | desmos.subspaces.v3.MsgDeleteSection |
| message         | sender            | {userAddress}                        |

### MsgCreateUserGroup

| **Type**           | **Attribute Key** | **Attribute Value**                    | 
|:-------------------|:------------------|:---------------------------------------|
| created_user_group | subspace_id       | {subspaceID}                           |    
| created_user_group | user_group_id     | {userGroupID}                          |    
| message            | module            | subspaces                              |
| message            | action            | desmos.subspaces.v3.MsgCreateUserGroup |
| message            | sender            | {userAddress}                          |

### MsgEditUserGroup

| **Type**          | **Attribute Key** | **Attribute Value**                  | 
|:------------------|:------------------|:-------------------------------------|
| edited_user_group | subspace_id       | {subspaceID}                         |
| edited_user_group | user_group_id     | {userGroupID}                        |
| message           | module            | subspaces                            |
| message           | action            | desmos.subspaces.v3.MsgEditUserGroup |
| message           | sender            | {userAddress}                        |

### MsgMoveUserGroup

| **Type**         | **Attribute Key** | **Attribute Value**                  | 
|:-----------------|:------------------|:-------------------------------------|
| moved_user_group | subspace_id       | {subspaceID}                         |
| moved_user_group | user_group_id     | {userGroupID}                        |
| message          | module            | subspaces                            |
| message          | action            | desmos.subspaces.v3.MsgMoveUserGroup |
| message          | sender            | {userAddress}                        |

### MsgSetUserGroupPermissions

| **Type**                   | **Attribute Key** | **Attribute Value**                            | 
|:---------------------------|:------------------|:-----------------------------------------------|
| set_user_group_permissions | subspace_id       | {subspaceID}                                   |
| set_user_group_permissions | user_group_id     | {userGroupID}                                  |
| set_user_group_permissions | permissions       | {permissions}                                  |
| message                    | module            | subspaces                                      |
| message                    | action            | desmos.subspaces.v3.MsgSetUserGroupPermissions |
| message                    | sender            | {userAddress}                                  |

### MsgDeleteUserGroup

| **Type**           | **Attribute Key** | **Attribute Value**                    | 
|:-------------------|:------------------|:---------------------------------------|
| deleted_user_group | subspace_id       | {subspaceID}                           |
| deleted_user_group | user_group_id     | {userGroupID}                          |
| message            | module            | subspaces                              |
| message            | action            | desmos.subspaces.v3.MsgDeleteUserGroup |
| message            | sender            | {userAddress}                          |

### MsgAddUserToUserGroup

| **Type**           | **Attribute Key** | **Attribute Value**                       | 
|:-------------------|:------------------|:------------------------------------------|
| added_group_member | subspace_id       | {subspaceID}                              |
| added_group_member | user_group_id     | {userGroupID}                             |
| added_group_member | user              | {userAddress}                             |
| message            | module            | subspaces                                 |
| message            | action            | desmos.subspaces.v3.MsgAddUserToUserGroup |
| message            | sender            | {userAddress}                             |

### MsgRemoveUserFromUserGroup

| **Type**             | **Attribute Key** | **Attribute Value**                            | 
|:---------------------|:------------------|:-----------------------------------------------|
| removed_group_member | subspace_id       | {subspaceID}                                   |
| removed_group_member | user_group_id     | {userGroupID}                                  |
| removed_group_member | user              | {userAddress}                                  |
| message              | module            | subspaces                                      |
| message              | action            | desmos.subspaces.v3.MsgRemoveUserFromUserGroup |
| message              | sender            | {userAddress}                                  |

### MsgSetUserPermissions

| **Type**             | **Attribute Key** | **Attribute Value**                 | 
|:---------------------|:------------------|:------------------------------------|
| set_user_permissions | subspace_id       | {subspaceID}                        |
| set_user_permissions | section_id        | {sectionID}                         |
| set_user_permissions | permissions       | {permissions}                       |
| set_user_permissions | user              | {userAddress}                       |
| message              | module            | subspaces                           |
| message              | action            | desmos.subspaces.v3.MsgEditSubspace |
| message              | sender            | {userAddress}                       |

## MsgGrantTreasuryAuthorization

| **Type**                       | **Attribute Key** | **Attribute Value**                               | 
|:-------------------------------|:------------------|:--------------------------------------------------|
| granted_treasury_authorization | subspace_id       | {subspaceID}                                      |
| granted_treasury_authorization | granter           | {userAddress}                                     |
| granted_treasury_authorization | grantee           | {userAddress}                                     |
| message                        | module            | subspaces                                         |
| message                        | action            | desmos.subspaces.v3.MsgGrantTreasuryAuthorization |
| message                        | sender            | {userAddress}                                     |

## MsgRevokeTreasuryAuthorization

| **Type**                       | **Attribute Key** | **Attribute Value**                                | 
|:-------------------------------|:------------------|:---------------------------------------------------|
| revoked_treasury_authorization | subspace_id       | {subspaceID}                                       |
| revoked_treasury_authorization | granter           | {userAddress}                                      |
| revoked_treasury_authorization | grantee           | {userAddress}                                      |
| message                        | module            | subspaces                                          |
| message                        | action            | desmos.subspaces.v3.MsgRevokeTreasuryAuthorization |
| message                        | sender            | {userAddress}                                      |

## MsgGrantAllowance

| **Type**          | **Attribute Key** | **Attribute Value**                               | 
|:------------------|:------------------|:--------------------------------------------------|
| granted_allowance | subspace_id       | {subspaceID}                                      |
| granted_allowance | granter           | {userAddress}                                     |
| granted_allowance | user_grantee      | {userAddress}                                     |
| granted_allowance | group_grantee     | {groupID}                                         |
| message           | module            | subspaces                                         |
| message           | action            | desmos.subspaces.v3.MsgGrantTreasuryAuthorization |
| message           | sender            | {userAddress}                                     |

## MsgRevokeAllowance

| **Type**          | **Attribute Key** | **Attribute Value**                               | 
|:------------------|:------------------|:--------------------------------------------------|
| revoked_allowance | subspace_id       | {subspaceID}                                      |
| revoked_allowance | granter           | {userAddress}                                     |
| revoked_allowance | user_grantee      | {userAddress}                                     |
| revoked_allowance | group_grantee     | {groupID}                                         |
| message           | module            | subspaces                                         |
| message           | action            | desmos.subspaces.v3.MsgGrantTreasuryAuthorization |
| message           | sender            | {userAddress}                                     |

## MsgUpdateSubspaceFeeTokens

| **Type**                    | **Attribute Key** | **Attribute Value**                               | 
|:----------------------------|:------------------|:--------------------------------------------------|
| updated_subspace_fee_tokens | subspace_id       | {subspaceID}                                      |
| updated_subspace_fee_tokens | user              | {authorityAddress}                                |
| message                     | module            | subspaces                                         |
| message                     | action            | desmos.subspaces.v3.MsgGrantTreasuryAuthorization |
| message                     | sender            | {userAddress}                                     |