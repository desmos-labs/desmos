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

| **Type**       | **Attribute Key** | **Attribute Value**               | 
|:---------------|:------------------|:----------------------------------|
| created_report | subspace_id       | {subspaceID}                      |
| created_report | report_id         | {reportID}                        |
| created_report | reporter          | {userAddress}                     |
| created_report | creation_time     | {creationTime}                    |
| message        | module            | reports                           |
| message        | action            | desmos.reports.v1.MsgCreateReport |
| message        | reporter          | {userAddress}                     |

Other events attributes depending on the type of the report target are attached.

#### `Post Target`

| **Type**      | **Attribute Key** | **Attribute Value** | 
|:--------------|:------------------|:--------------------|
| reported_post | subspace_id       | {subspaceID}        |
| reported_post | post_id           | {reportID}          |
| reported_post | reporter          | {userAddress}       |

#### `User Target`

| **Type**      | **Attribute Key** | **Attribute Value** | 
|:--------------|:------------------|:--------------------|
| reported_user | subspace_id       | {subspaceID}        |
| reported_user | user              | {userAddress}       |
| reported_user | reporter          | {userAddress}       |

### MsgDeleteReport

| **Type**       | **Attribute Key** | **Attribute Value**               | 
|:---------------|:------------------|:----------------------------------|
| deleted_report | subspace_id       | {subspaceID}                      |
| deleted_report | report_id         | {reportID}                        |
| message        | module            | reports                           |
| message        | action            | desmos.reports.v1.MsgDeleteReport |
| message        | signer            | {userAddress}                     |

### MsgSupportStandardReason

| **Type**                  | **Attribute Key**  | **Attribute Value**                        | 
|:--------------------------|:-------------------|:-------------------------------------------|
| supported_standard_reason | subspace_id        | {subspaceID}                               |
| supported_standard_reason | standard_reason_id | {reasonID}                                 |
| supported_standard_reason | reason_id          | {reasonID}                                 |
| message                   | module             | reports                                    |
| message                   | action             | desmos.reports.v1.MsgSupportStandardReason |
| message                   | signer             | {userAddress}                              |

### MsgAddReason

| **Type**               | **Attribute Key** | **Attribute Value**            | 
|:-----------------------|:------------------|:-------------------------------|
| added_reporting_reason | subspace_id       | {subspaceID}                   |
| added_reporting_reason | reason_id         | {reasonID}                     |
| message                | module            | reports                        |
| message                | action            | desmos.reports.v1.MsgAddReason |
| message                | signer            | {userAddress}                  |

### MsgRemoveReason

| **Type**                 | **Attribute Key** | **Attribute Value**               | 
|:-------------------------|:------------------|:----------------------------------|
| removed_reporting_reason | subspace_id       | {subspaceID}                      |
| removed_reporting_reason | reason_id         | {reasonID}                        |
| message                  | module            | reports                           |
| message                  | action            | desmos.reports.v1.MsgRemoveReason |