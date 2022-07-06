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
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reports/v1/msgs.proto#L34-L58
```

The message is expected to fail if any of the following situations occurs:
* The reporter does not have a profile;
* The subspace does not exist;
* One or more of the specified reasons IDs do not exist;
* The reported does not have the permission to report content inside the subspace;
* The report has already been made;
* The report validation fails;
* The report target type is not valid.

## Msg/DeleteReport
A report can be deleted using the `MsgDeleteReport`, specifying the information you see below:

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reports/v1/msgs.proto#L77-L92
```

The message is expected to fail if any of the following situations occurs:

* The subspace does not exist;
* The report does not exist;
* The signer does not have the permission to delete a report inside the subspace;
* The signer is not the reporter;

## Msg/SupportStandardReason
The `MsgSupportStandardReason` can be used if you want to use the set of standard reasons specified in the module 
params inside your dApp.

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reports/v1/msgs.proto#L99-L114
```

It's expected to fail if:
* the subspace does not exist;
* the reason does not exist;
* the signer does not have the permission to manage registered within inside the subspace.

## Msg/AddReason
A reason can be added to a subspace using the `MsgAddReason`.

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reports/v1/msgs.proto#L128-L143
```

It's expected to fail if:
* the subspace reason does not exist;
* the signer does not have the permission to manage reasons within the subspace;
* the reason name is either empty or blank.

## Msg/RemoveReason
A previously added reason can be removed using the `MsgRemoveReason`.

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reports/v1/msgs.proto#L156-L171
```

It's expected to fail if:
* the subspace reason does not exist;
* the reason does not exist;
* the signer does not have the permission to manage registered reasons within the subspace.