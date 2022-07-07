---
id: events
title: Events
sidebar_label: Events
slug: events
---

# Events

The reports module emits the following events:

## Handlers

### MsgCreateReport

| **Type**      | **Attribute Key** | **Attribute Value**               | 
|:--------------|:------------------|:----------------------------------|
| create_report | subspace_id       | {subspaceID}                      |
| create_report | report_id         | {reportID}                        |
| create_report | reporter          | {userAddress}                     |
| create_report | creation_time     | {creationTime}                    |
| message       | module            | reports                           |
| message       | action            | desmos.reports.v1.MsgCreateReport |
| message       | reporter          | {userAddress}                     |

Other events attributes depending on the type of the report target are attached.

#### `Post Target`

| **Type**    | **Attribute Key** | **Attribute Value** | 
|:------------|:------------------|:--------------------|
| report_post | subspace_id       | {subspaceID}        |
| report_post | post_id           | {reportID}          |
| report_post | reporter          | {userAddress}       |

#### `User Target`

| **Type**    | **Attribute Key** | **Attribute Value** | 
|:------------|:------------------|:--------------------|
| report_user | subspace_id       | {subspaceID}        |
| report_user | user              | {userAddress}       |
| report_user | reporter          | {userAddress}       |

### MsgDeleteReport

| **Type**      | **Attribute Key** | **Attribute Value**               | 
|:--------------|:------------------|:----------------------------------|
| delete_report | subspace_id       | {subspaceID}                      |
| delete_report | report_id         | {reportID}                        |
| message       | module            | reports                           |
| message       | action            | desmos.reports.v1.MsgDeleteReport |
| message       | signer            | {userAddress}                     |

### MsgSupportStandardReason

| **Type**                | **Attribute Key**  | **Attribute Value**                        | 
|:------------------------|:-------------------|:-------------------------------------------|
| support_standard_reason | subspace_id        | {subspaceID}                               |
| support_standard_reason | standard_reason_id | {reasonID}                                 |
| support_standard_reason | reason_id          | {reasonID}                                 |
| message                 | module             | reports                                    |
| message                 | action             | desmos.reports.v1.MsgSupportStandardReason |
| message                 | signer             | {userAddress}                              |

### MsgAddReason

| **Type**   | **Attribute Key** | **Attribute Value**            | 
|:-----------|:------------------|:-------------------------------|
| add_reason | subspace_id       | {subspaceID}                   |
| add_reason | reason_id         | {reasonID}                     |
| message    | module            | reports                        |
| message    | action            | desmos.reports.v1.MsgAddReason |
| message    | signer            | {userAddress}                  |


### MsgRemoveReason

| **Type**      | **Attribute Key** | **Attribute Value**               | 
|:--------------|:------------------|:----------------------------------|
| remove_reason | subspace_id       | {subspaceID}                      |
| remove_reason | reason_id         | {reasonID}                        |
| message       | module            | reports                           |
| message       | action            | desmos.reports.v1.MsgRemoveReason |