---
id: refuse-dtag-transfer-request
title: Refuse DTag transfer request
sidebar_label: Refuse DTag transfer request
slug: refuse-dtag-transfer-request
---

# `MsgRefuseDTagTransferRequest`
This message allows you to refuse a `DTag` transfer request made by a user.

## Structure
````json
{
  "@type": "/desmos.profiles.v1beta1.MsgRefuseDTagTransferRequest",
  "sender": "<Desmos address that sent the DTag's request>",
  "receiver": "<Desmos address of the DTag owner>"
}
````

### Attributes
| Attribute | Type | Description | Required |
| :-------: | :----: | :-------- | :------- |
| `sender`| String | Desmos address of the user that request the `DTag` | yes |
| `receiver`  | String | Desmos address of the user that is the owner of the requested `DTag` | yes |

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgRefuseDTagTransferRequest",
  "sender": "desmos1nhgk008jvrxwa9tufr9tcr6zfrhe2uz0v90r2a",
  "receiver": "desmos1k99c8htyk32srx78efzg7sxm965prtz0j9qrc7"
}
````

## Message action
The action associated to this message is the following:

```
refuse_dtag_transfer_request
```
