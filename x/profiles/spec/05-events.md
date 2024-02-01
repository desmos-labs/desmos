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

| **Type**      | **Attribute Key**     | **Attribute Value**               | 
|:--------------|:----------------------|:----------------------------------|
| saved_profile | profile_dtag          | {profileDTag}                     |
| saved_profile | profile_creator       | {userAddress}                     |
| saved_profile | profile_creation_time | {profileCreationTime}             |
| message       | module                | profiles                          |
| message       | action                | desmos.profiles.v3.MsgSaveProfile |
| message       | sender                | {userAddress}                     |

## MsgDeleteProfile

| **Type**        | **Attribute Key** | **Attribute Value**                 | 
|:----------------|:------------------|:------------------------------------|
| deleted_profile | profile_creator   | {userAddress}                       |
| message         | module            | profiles                            | 
| message         | action            | desmos.profiles.v3.MsgDeleteProfile |
| message         | sender            | {userAddress}                       |

## MsgRequestDTagTransfer

| **Type**                      | **Attribute Key** | **Attribute Value**                       | 
|:------------------------------|:------------------|:------------------------------------------|
| created_dtag_transfer_request | dtag_to_trade     | {dTagToTrade}                             | 
| created_dtag_transfer_request | request_sender    | {requestSenderAddress}                    | 
| created_dtag_transfer_request | request_receiver  | {requestReceiverAddress}                  |
| message                       | module            | profiles                                  | 
| message                       | action            | desmos.profiles.v3.MsgRequestDTagTransfer |
| message                       | sender            | {requestSenderAddress}                    |

## MsgCancelDTagTransferRequest

| **Type**                       | **Attribute Key** | **Attribute Value**                             | 
|:-------------------------------|:------------------|:------------------------------------------------|
| canceled_dtag_transfer_request | request_sender    | {requestSenderAddress}                          | 
| canceled_dtag_transfer_request | request_receiver  | {requestReceiverAddress}                        |
| message                        | module            | profiles                                        | 
| message                        | action            | desmos.profiles.v3.MsgCancelDTagTransferRequest |
| message                        | sender            | {userAddress}                                   |

## MsgAcceptDTagTransferRequest

| **Type**                       | **Attribute Key** | **Attribute Value**                             | 
|:-------------------------------|:------------------|:------------------------------------------------|
| accepted_dtag_transfer_request | dtag_to_trade     | {dTagToTrade}                                   |
| accepted_dtag_transfer_request | new_dtag          | {newDTag}                                       |
| accepted_dtag_transfer_request | request_sender    | {requestSenderAddress}                          | 
| accepted_dtag_transfer_request | request_receiver  | {requestReceiverAddress}                        |
| message                        | module            | profiles                                        | 
| message                        | action            | desmos.profiles.v3.MsgAcceptDTagTransferRequest |
| message                        | sender            | {userAddress}                                   |

## MsgRefuseDTagTransferRequest

| **Type**                      | **Attribute Key** | **Attribute Value**                             | 
|:------------------------------|:------------------|:------------------------------------------------|
| refused_dtag_transfer_request | request_sender    | {requestSenderAddress}                          | 
| refused_dtag_transfer_request | request_receiver  | {requestReceiverAddress}                        |
| message                       | module            | profiles                                        | 
| message                       | action            | desmos.profiles.v3.MsgRefuseDTagTransferRequest |
| message                       | sender            | {userAddress}                                   |

## MsgLinkChainAccount

| **Type**           | **Attribute Key**            | **Attribute Value**                    | 
|:-------------------|:-----------------------------|:---------------------------------------|
| created_chain_link | chain_link_account_target    | {targetAddress}                        |
| created_chain_link | chain_link_source_chain_name | {chainName}                            | 
| created_chain_link | chain_link_account_owner     | {userAddress}                          |
| created_chain_link | chain_link_creation_time     | {creationTime}                         |
| message            | module                       | profiles                               | 
| message            | action                       | desmos.profiles.v3.MsgLinkChainAccount |
| message            | sender                       | {userAddress}                          |

## MsgUnlinkChainAccount

| **Type**           | **Attribute Key**            | **Attribute Value**                      | 
|:-------------------|:-----------------------------|:-----------------------------------------|
| deleted_chain_link | chain_link_account_target    | {targetAddress}                          |
| deleted_chain_link | chain_link_source_chain_name | {chainName}                              | 
| deleted_chain_link | chain_link_account_owner     | {userAddress}                            |
| message            | module                       | profiles                                 | 
| message            | action                       | desmos.profiles.v3.MsgUnlinkChainAccount |
| message            | sender                       | {userAddress}                            |

## MsgSetDefaultExternalAddress

| **Type**                     | **Attribute Key**           | **Attribute Value**                     | 
|:-----------------------------|:----------------------------|:----------------------------------------|
| set_default_external_address | chain_link_chain_name       | {chainName}                             | 
| set_default_external_address | chain_link_external_address | {externalAddress}                       |
| set_default_external_address | chain_link_owner            | {chainLinkOwner}                        |
| message                      | module                      | profiles                                | 
| message                      | action                      | desmos.profiles.v3.MsgSetDefaultAddress |
| message                      | sender                      | {userAddress}                           |

## MsgLinkApplication

| **Type**                 | **Attribute Key**              | **Attribute Value**                   | 
|:-------------------------|:-------------------------------|:--------------------------------------|
| created_application_link | user                           | {userAddress}                         |
| created_application_link | application_name               | {applicationName}                     | 
| created_application_link | application_username           | {applicationUsername}                 |
| created_application_link | application_link_creation_time | {creationTime}                        |
| message                  | module                         | profiles                              | 
| message                  | action                         | desmos.profiles.v3.MsgLinkApplication |
| message                  | sender                         | {userAddress}                         |

## MsgUnlinkApplication

| **Type**                 | **Attribute Key**    | **Attribute Value**                     | 
|:-------------------------|:---------------------|:----------------------------------------|
| deleted_application_link | user                 | {userAddress}                           |
| deleted_application_link | application_name     | {applicationName}                       | 
| deleted_application_link | application_username | {applicationUsername}                   |
| message                  | module               | profiles                                | 
| message                  | action               | desmos.profiles.v3.MsgUnlinkApplication |
| message                  | sender               | {userAddress}                           |

## Keeper

### Chain Link Saved

| **Type**         | **Attribute Key**           | **Attribute Value** | 
|:-----------------|:----------------------------|:--------------------|
| saved_chain_link | chain_link_owner            | {userAddress}       |
| saved_chain_link | chain_link_chain_name       | {chainName}         | 
| saved_chain_link | chain_link_external_address | {externalAddress}   |

### Application Link Saved

| **Type**               | **Attribute Key**    | **Attribute Value**   | 
|:-----------------------|:---------------------|:----------------------|
| saved_application_link | user                 | {userAddress}         |
| saved_application_link | application_name     | {applicationName}     | 
| saved_application_link | application_username | {applicationUsername} |

## IBC

### Received link chain account IBC packet

| **Type**                           | **Attribute Key**           | **Attribute Value** | 
|:-----------------------------------|:----------------------------|:--------------------|
| received_link_chain_account_packet | module                      | profiles            |
| received_link_chain_account_packet | chain_link_owner            | {userAddress}       |
| received_link_chain_account_packet | chain_link_chain_name       | {chainName}         | 
| received_link_chain_account_packet | chain_link_external_address | {externalAddress}   |
| received_link_chain_account_packet | success                     | true                |

### Received oracle response IBC packet

| **Type**                        | **Attribute Key** | **Attribute Value** | 
|:--------------------------------|:------------------|:--------------------|
| received_oracle_response_packet | module            | profiles            |
| received_oracle_response_packet | client_id         | {clientID}          |
| received_oracle_response_packet | request_id        | {requestID}         | 
| received_oracle_response_packet | resolve_status    | {resolveStatus}     |
| received_oracle_response_packet | success           | true                |
