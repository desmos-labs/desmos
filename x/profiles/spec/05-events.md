---
id: events
title: Events
sidebar_label: Events
slug: events
---

# Events

The profiles module emits the following events:

## Handlers

### MsgSaveProfile

| **Type**     | **Attribute Key**     | **Attribute Value**               | 
|:-------------|:----------------------|:----------------------------------|
| save_profile | profile_dtag          | {profileDTag}                     |
| save_profile | profile_creator       | {userAddress}                     |
| save_profile | profile_creation_time | {profileCreationTime}             |
| message      | module                | profiles                          |
| message      | action                | desmos.profiles.v2.MsgSaveProfile |
| message      | sender                | {userAddress}                     |

## MsgDeleteProfile

| **Type**       | **Attribute Key** | **Attribute Value**                 | 
|:---------------|:------------------|:------------------------------------|
| delete_profile | profile_creator   | {userAddress}                       |
| message        | module            | profiles                            | 
| message        | action            | desmos.profiles.v2.MsgDeleteProfile |
| message        | sender            | {userAddress}                       |

## MsgRequestDTagTransfer

| **Type**                     | **Attribute Key** | **Attribute Value**                       | 
|:-----------------------------|:------------------|:------------------------------------------|
| create_dtag_transfer_request | dtag_to_trade     | {dTagToTrade}                             | 
| create_dtag_transfer_request | request_sender    | {requestSenderAddress}                    | 
| create_dtag_transfer_request | request_receiver  | {requestReceiverAddress}                  |
| message                      | module            | profiles                                  | 
| message                      | action            | desmos.profiles.v2.MsgRequestDTagTransfer |
| message                      | sender            | {requestSenderAddress}                    |

## MsgCancelDTagTransferRequest

| **Type**                     | **Attribute Key** | **Attribute Value**                             | 
|:-----------------------------|:------------------|:------------------------------------------------|
| cancel_dtag_transfer_request | request_sender    | {requestSenderAddress}                          | 
| cancel_dtag_transfer_request | request_receiver  | {requestReceiverAddress}                        |
| message                      | module            | profiles                                        | 
| message                      | action            | desmos.profiles.v2.MsgCancelDTagTransferRequest |
| message                      | sender            | {userAddress}                                   |

## MsgAcceptDTagTransferRequest

| **Type**                     | **Attribute Key** | **Attribute Value**                             | 
|:-----------------------------|:------------------|:------------------------------------------------|
| accept_dtag_transfer_request | dtag_to_trade     | {dTagToTrade}                                   |
| accept_dtag_transfer_request | new_dtag          | {newDTag}                                       |
| accept_dtag_transfer_request | request_sender    | {requestSenderAddress}                          | 
| accept_dtag_transfer_request | request_receiver  | {requestReceiverAddress}                        |
| message                      | module            | profiles                                        | 
| message                      | action            | desmos.profiles.v2.MsgAcceptDTagTransferRequest |
| message                      | sender            | {userAddress}                                   |

## MsgRefuseDTagTransferRequest

| **Type**                     | **Attribute Key** | **Attribute Value**                             | 
|:-----------------------------|:------------------|:------------------------------------------------|
| refuse_dtag_transfer_request | request_sender    | {requestSenderAddress}                          | 
| refuse_dtag_transfer_request | request_receiver  | {requestReceiverAddress}                        |
| message                      | module            | profiles                                        | 
| message                      | action            | desmos.profiles.v2.MsgRefuseDTagTransferRequest |
| message                      | sender            | {userAddress}                                   |

## MsgLinkChainAccount

| **Type**           | **Attribute Key**            | **Attribute Value**                    | 
|:-------------------|:-----------------------------|:---------------------------------------|
| link_chain_account | chain_link_account_target    | {targetAddress}                        |
| link_chain_account | chain_link_source_chain_name | {chainName}                            | 
| link_chain_account | chain_link_account_owner     | {userAddress}                          |
| link_chain_account | chain_link_creation_time     | {creationTime}                         |
| message            | module                       | profiles                               | 
| message            | action                       | desmos.profiles.v2.MsgLinkChainAccount |
| message            | sender                       | {userAddress}                          |

## MsgUnlinkChainAccount

| **Type**             | **Attribute Key**            | **Attribute Value**                      | 
|:---------------------|:-----------------------------|:-----------------------------------------|
| unlink_chain_account | chain_link_account_target    | {targetAddress}                          |
| unlink_chain_account | chain_link_source_chain_name | {chainName}                              | 
| unlink_chain_account | chain_link_account_owner     | {userAddress}                            |
| message              | module                       | profiles                                 | 
| message              | action                       | desmos.profiles.v2.MsgUnlinkChainAccount |
| message              | sender                       | {userAddress}                            |

## MsgLinkApplication

| **Type**         | **Attribute Key**              | **Attribute Value**                   | 
|:-----------------|:-------------------------------|:--------------------------------------|
| link_application | user                           | {userAddress}                         |
| link_application | application_name               | {applicationName}                     | 
| link_application | application_username           | {applicationUsername}                 |
| link_application | application_link_creation_time | {creationTime}                        |
| message          | module                         | profiles                              | 
| message          | action                         | desmos.profiles.v2.MsgLinkApplication |
| message          | sender                         | {userAddress}                         |

## MsgUnlinkApplication

| **Type**           | **Attribute Key**    | **Attribute Value**                     | 
|:-------------------|:---------------------|:----------------------------------------|
| unlink_application | user                 | {userAddress}                           |
| unlink_application | application_name     | {applicationName}                       | 
| unlink_application | application_username | {applicationUsername}                   |
| message            | module               | profiles                                | 
| message            | action               | desmos.profiles.v2.MsgUnlinkApplication |
| message            | sender               | {userAddress}                           |

## Keeper

### Application Link Saved
| **Type**               | **Attribute Key**    | **Attribute Value**                     | 
|:-----------------------|:---------------------|:----------------------------------------|
| application_link_saved | user                 | {userAddress}                           |
| application_link_saved | application_name     | {applicationName}                       | 
| application_link_saved | application_username | {applicationUsername}                   |