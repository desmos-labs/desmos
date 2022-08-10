---
id: messages
title: Messages
sidebar_label: Messages
slug: messages
---

# Msg Service

## Msg/CreateReport
A report can be created using the `MsgCreateReport`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/reports/v1/msgs.proto#L33-L58
```

It's expected to fail if:
* the reporter does not have a profile;
* the subspace does not exist;
* one of the specified reasons ids does not exist inside the subspace;
* the reported does not have the permission to report content within the subspace;
* another report for the same target has already been created by the same user;
* the report target does not exist.

## Msg/DeleteReport
A report can be deleted using the `MsgDeleteReport`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/reports/v1/msgs.proto#L76-L92
```

It's expected to fail if:
* the subspace does not exist;
* the report does not exist;
* the signer does not have the permission to delete a report within the subspace.

## Msg/SupportStandardReason
A standard reason can be supported within a subspace using the `MsgSupportStandardReason`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/reports/v1/msgs.proto#L97-L114
```

It's expected to fail if:
* the subspace does not exist;
* the reason does not exist;
* the signer does not have the permission to manage registered within inside the subspace.

## Msg/AddReason
A reason can be added to a subspace using the `MsgAddReason`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/reports/v1/msgs.proto#L126-L143
```

It's expected to fail if:
* the subspace reason does not exist;
* the signer does not have the permission to manage reasons within the subspace;
* the reason name is either empty or blank.

## Msg/RemoveReason
A previously added reason can be removed using the `MsgRemoveReason`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/reports/v1/msgs.proto#L154-L171
```

It's expected to fail if:
* the subspace reason does not exist;
* the reason does not exist;
* the signer does not have the permission to manage registered reasons within the subspace.