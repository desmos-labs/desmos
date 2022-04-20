<!--
order: 4
-->

# Events

The profiles module emits the following events:

## Handlers

### MsgSaveProfile

| **Type**      | **Attribute Key**     | **Attribute Value**               | 
|:--------------|:----------------------|:----------------------------------|
| profile_saved | profile_dtag          | {profileDTag}                     |
| profile_saved | profile_creator       | {userAddress}                     |
| profile_saved | profile_creation_time | {profileCreationTime}             |
| message       | module                | profiles                          |
| message       | action                | desmos.profiles.v2.MsgSaveProfile |
| message       | sender                | {userAddress}                     |

## MsgDeleteProfile

| **Type**        | **Attribute Key** | **Attribute Value**                 | 
|:----------------|:------------------|:------------------------------------|
| profile_deleted | profile_creator   | {userAddress }                      |
| message         | module            | profiles                            | 
| message         | action            | desmos.profiles.v2.MsgDeleteProfile |
| message         | sender            | {userAddress}                       |

## MsgRequestDTagTransfer

| **Type**              | **Attribute Key** | **Attribute Value**                       | 
|:----------------------|:------------------|:------------------------------------------|
| dtag_transfer_request | dtag_to_trade     | {dTagToTrade}                             | 
| dtag_transfer_request | request_sender    | {requestSenderAddress}                    | 
| dtag_transfer_request | request_receiver  | {requestReceiverAddress}                  |
| message               | module            | profiles                                  | 
| message               | action            | desmos.profiles.v2.MsgRequestDTagTransfer |
| message               | sender            | {requestSenderAddress}                    |

## MsgCancelDTagTransferRequest

| **Type**             | **Attribute Key** | **Attribute Value**                             | 
|:---------------------|:------------------|:------------------------------------------------|
| dtag_transfer_cancel | request_sender    | {requestSenderAddress}                          | 
| dtag_transfer_cancel | request_receiver  | {requestReceiverAddress}                        |
| message              | module            | profiles                                        | 
| message              | action            | desmos.profiles.v2.MsgCancelDTagTransferRequest |
| message              | sender            | {userAddress}                                   |

## MsgAcceptDTagTransferRequest

| **Type**             | **Attribute Key** | **Attribute Value**                             | 
|:---------------------|:------------------|:------------------------------------------------|
| dtag_transfer_accept | dtag_to_trade     | {dTagToTrade}                                   |
| dtag_transfer_accept | new_dtag          | {newDTag}                                       |
| dtag_transfer_accept | request_sender    | {requestSenderAddress}                          | 
| dtag_transfer_accept | request_receiver  | {requestReceiverAddress}                        |
| message              | module            | profiles                                        | 
| message              | action            | desmos.profiles.v2.MsgAcceptDTagTransferRequest |
| message              | sender            | {userAddress}                                   |

## MsgRefuseDTagTransferRequest

| **Type**             | **Attribute Key** | **Attribute Value**                             | 
|:---------------------|:------------------|:------------------------------------------------|
| dtag_transfer_refuse | request_sender    | {requestSenderAddress}                          | 
| dtag_transfer_accept | request_receiver  | {requestReceiverAddress}                        |
| message              | module            | profiles                                        | 
| message              | action            | desmos.profiles.v2.MsgRefuseDTagTransferRequest |
| message              | sender            | {userAddress}                                   |

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

| **Type**                 | **Attribute Key**              | **Attribute Value**                   | 
|:-------------------------|:-------------------------------|:--------------------------------------|
| application_link_created | user                           | {userAddress}                         |
| application_link_created | application_name               | {applicationName}                     | 
| application_link_created | application_username           | {applicationUsername}                 |
| application_link_created | application_link_creation_time | {creationTime}                        |
| message                  | module                         | profiles                              | 
| message                  | action                         | desmos.profiles.v2.MsgLinkApplication |
| message                  | sender                         | {userAddress}                         |

## MsgUnlinkApplication

| **Type**                 | **Attribute Key**    | **Attribute Value**                     | 
|:-------------------------|:---------------------|:----------------------------------------|
| application_link_deleted | user                 | {userAddress}                           |
| application_link_deleted | application_name     | {applicationName}                       | 
| application_link_deleted | application_username | {applicationUsername}                   |
| message                  | module               | profiles                                | 
| message                  | action               | desmos.profiles.v2.MsgUnlinkApplication |
| message                  | sender               | {userAddress}                           |

## Keeper

### Application Link Saved
| **Type**               | **Attribute Key**    | **Attribute Value**                     | 
|:-----------------------|:---------------------|:----------------------------------------|
| application_link_saved | user                 | {userAddress}                           |
| application_link_saved | application_name     | {applicationName}                       | 
| application_link_saved | application_username | {applicationUsername}                   |