---
id: accept-dtag-transfer-request
title: Accept DTag transfer request
sidebar_label: Accept DTag transfer request
slug: accept-dtag-transfer-request
---

# `MsgAcceptDTagTransferRequest`
This message allows you to accept a `DTag` transfer request.

## Structure
````json
{
  "@type": "/desmos.profiles.v1beta1.MsgAcceptDTagTransferRequest",
  "new_dtag": "<The new DTag for the current owner's profile>",
  "receiver": "<Desmos address of the DTag owner>",
  "sender": "<Desmos address that's making the DTag's request>"
}
````

### Attributes
| Attribute | Type | Description | Required |
| :-------: | :----: | :-------- | :------- |
| `new_dtag` | String | The new `DTag` for the current owner profile that will replace the traded one | yes |
| `receiver`  | String | Desmos address of the user that is the owner of the requested `DTag` | yes |
| `sender`| String | Desmos address of the user that request the `DTag` | yes |

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgAcceptDTagTransferRequest",
  "new_dtag": "newDTag",
  "receiver": "desmos1k99c8htyk32srx78efzg7sxm965prtz0j9qrc7",
  "sender": "desmos1nhgk008jvrxwa9tufr9tcr6zfrhe2uz0v90r2a"
}
````

## Message action
The action associated to this message is the following:

```
accept_dtag_transfer_request
```
