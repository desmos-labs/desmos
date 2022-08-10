---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client

## CLI

A user can query and interact with the `reports` module using the CLI.

### Query

The `query` commands allow users to query the `reports` state.

```
desmos query reports --help
```

#### report
The `report` query command allows users to get the report with the given id inside the given subspace id.

```bash
 desmos query reports report [subspace-id] [report-id] [flags]
```

Example:
```bash
 desmos query reports report 1 1
```

Example output:
```yaml
report:
  creation_date: "2022-07-01T10:11:09.229623Z"
  id: "1"
  message: This is a test report
  reasons_ids:
  - 1
  reporter: desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
  subspace_id: "1"
  target:
    '@type': /desmos.reports.v1.PostTarget
    post_id: "1"
```

#### reports
The `reports` query command allows users to get all the reports inside the subspace with the given id.

```bash
desmos query reports reports [subspace-id] [flags]
```

Example:
```bash
desmos query reports reports 1 --page=1 --limit=100
```

Example output:
```yaml
pagination:
  next_key: null
  total: "0"
reports:
- creation_date: "2022-07-01T10:11:09.229623Z"
  id: "1"
  message: This is a test report
  reasons_ids:
  - 1
  reporter: desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
  subspace_id: "1"
  target:
    '@type': /desmos.reports.v1.PostTarget
    post_id: "1"
```

### reason
The `reason` query command allows users to get the reason with the given id inside the subspace with the given id.

```bash
desmos query reports reason [subspace-id] [reason-id] [flags]
```

Example:
```bash
desmos query reports reason 1 1
```

Example output:
```yaml
reason:
  description: Spam content or user
  id: 1
  subspace_id: "1"
  title: Spam
```

### reasons
The `reasons` query command allows users to get all the reasons inside the subspace with the given id.

```bash
desmos query reports reasons [subspace-id] [flags]
```

Example:
```bash
desmos query reports reasons 1 --page=1 --limit=100
```

Example output:
```yaml
pagination:
  next_key: null
  total: "0"
reasons:
- description: Spam content or user
  id: 1
  subspace_id: "1"
  title: Spam
```

### params
The `params` query command allows users to get the currently set parameters of the module.

```bash
desmos query reports params [flags]
```

Example:
```bash
desmos query reports params
```

Example output:
```bash
standard_reasons:
  - id: 1
    title: "Spam"
    description: "Spam user or content"
```

## gRPC
A user can query the `reports` module gRPC endpoints.

### Report
The `Report` endpoint allows users to query a report given its ID and the ID of the subspace where its made.

```bash
desmos.reports.v1.Query/Report
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1, "report_id":1}' localhost:9090 desmos.reports.v1.Query/Report
```

Example output:
```json
{
  "report": {
    "subspaceId": "1",
    "id": "1",
    "reasonsIds": [
      1
    ],
    "message": "This is a test report",
    "reporter": "desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
    "target": {"@type":"/desmos.reports.v1.PostTarget","postId":"1"},
    "creationDate": "2022-07-01T10:11:09.229623Z"
  }
}
```

### Reports
The `Reports` endpoint allows users to query all the reports of the subspace with the given ID.

```bash
desmos.reports.v1.Query/Reports
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1}' localhost:9090 desmos.reports.v1.Query/Reports
```

Example output:
```json
{
  "reports": [
    {
      "subspaceId": "1",
      "id": "1",
      "reasonsIds": [
        1
      ],
      "message": "This is a test report",
      "reporter": "desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
      "target": {"@type":"/desmos.reports.v1.PostTarget","postId":"1"},
      "creationDate": "2022-07-01T10:11:09.229623Z"
    }
  ],
  "pagination": {
    "total": "1"
  }
}

```

### Reason
The `Reason` endpoint allows users to query the reason given its ID and the ID of the subspace where its made.

```bash
desmos.reports.v1.Query/Reason
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1, "reason_id":1}' localhost:9090 desmos.reports.v1.Query/Reason 
```

Example output:
```json
{
  "reason": {
    "subspaceId": "1",
    "id": 1,
    "title": "Spam",
    "description": "Spam content or user"
  }
}
```

### Reasons
The `Reasons` endpoint allows users to query all the reasons of the subspace with the given ID.

```bash
desmos.reports.v1.Query/Reasons
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":1}' localhost:9090 desmos.reports.v1.Query/Reasons 
```

Example output:
```json
{
  "reasons": [
    {
      "subspaceId": "1",
      "id": 1,
      "title": "Spam",
      "description": "Spam content or user"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### Params
The `Params` endpoint allows users to query the module's parameters.

```bash
desmos.reports.v1.Query/Params
```

Example:
```bash
grpcurl -plaintext localhost:9090 desmos.reports.v1.Query/Params 
```

Example output:
```json
{
  "params": {
    "standard_reasons": [
      {
        "id": "1",
        "title": "Spam",
        "description": "Spam user or content"
      }
    ]
  }
}
```

## REST
A user can query the `reports` module using REST endpoints.

### Report
The `Report` endpoint allows users to query a report given its ID and the ID of the subspace where its made.

```
/desmos/reports/v1/subspaces/{subspace_id}/reports
```

### Reports
The `Reports` endpoint allows users to query all the reports of the subspace with the given ID.

```
/desmos/reports/v1/subspaces/{subspace_id}/reports
```

### Reason
The `Reason` endpoint allows users to query the reason given its ID and the ID of the subspace where its made.

```
/desmos/reports/v1/subspaces/{subspace_id}/reasons/{reason_id}
```

### Reasons
The `Reasons` endpoint allows users to query all the reasons of the subspace with the given ID.

```
/desmos/reports/v1/subspaces/{subspace_id}/reasons
```

### Params
The `Params` endpoint allows users to query the module's parameters.

```
/desmos/reports/v1/params
```