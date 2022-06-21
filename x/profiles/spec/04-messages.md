---
id: messages
title: Messages
sidebar_label: Messages
slug: messages
---

# Msg Service

## Msg/SaveProfile
A profile can be created or edited with the `MsgSaveProfile`, which allows to specify a DTag, a nickname, a bio and the profile and cover pictures. 

If a profile already exists, and you want to edit only a subset of the fields, you can use `[do-not-modify]` to specify the fields which values should not be changed (i.e. setting the `DTag` to `[do-not-modify]` will preserve the current value of the DTag).

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_profile.proto#L12-L32
```

It's expected to fail if a profile with the same DTag exists.

## Msg/DeleteProfile
A profile can be deleted using the `MsgDeleteProfile`. This will remove all the profile information from the store, but **will not** delete the on-chain account associated to such profile. 

Beware that using this message you will lose the ownership of your DTag and you will delete everything that is related to your profile (i.e. incoming DTag transfer requests). 

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_profile.proto#L39-L45
```

It's expected to fail if the signer does not have a profile. 

## Msg/RequestDTagTransfer
A DTag transfer request can be created using the `MsgRequestDTagTransfer`. 

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_dtag_requests.proto#L12-L25
```

It's expected to fail if the recipient of the request has no profile.

## Msg/CancelDTagTransferRequest
An outgoing DTag transfer request can be canceled using the `MsgCancelDTagTransferRequest`. 

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_dtag_requests.proto#L33-L44
```

It's expected to fail if the request does not exist.

## Msg/AcceptDTagTransferRequest
An incoming DTag transfer request can be accepted using the `MsgAcceptDTagTransferRequest`. When accepting a DTag transfer request, the user accepting it **must** specify a new DTag that they want after their old one gets transferred to the request sender.

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_dtag_requests.proto#L52-L70
```

It's expected to fail if:
* the request does not exist.
* the specified new DTag is already used by someone else.

## Msg/RefuseDTagTransferRequest
An incoming DTag transfer request can be refused using `MsgRefuseDTagTransferRequest`

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_dtag_requests.proto#L78-L89
```

It's expected to fail if the request does not exist.

## Msg/LinkChainAccount
A new chain link can be created using the `MsgLinkChainAccount`

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_chain_links.proto#L11-L35
```

It's expected to fail if:
* the signer does not have a profile.
* the chain link signature is not valid.

## Msg/UnlinkChainAccount
An existing chain link can be deleted using the `MsgUnlinkChainAccount`

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_chain_links.proto#L42-L54
```

It's expected to fail if the chain link does not exist.

## Msg/LinkApplication
A new application link can be created using the `MsgLinkApplication`

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_app_links.proto#L11-L48
```

It's expected to fail if the signer does not have a profile.

## Msg/UnlinkApplication
An existing application link can be deleted using the `MsgUnlinkApplication`

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/proto/desmos/profiles/v2/msgs_app_links.proto#L56-L71
```

It's expected to fail if the application link does not exist.