# `MsgCancelDTagTransferRequest`
This message allows you to cancel a `DTag` transfer request made by yourself.

## Structure
````json
{
  "@type": "/desmos.profiles.v1beta1.MsgCancelDTagTransferRequest",
  "sender": "<Desmos address that sent the DTag's request>",
  "receiver": "<Desmos address of the DTag owner>"
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `sender`| String | Desmos address of the user that request the `DTag` |
| `receiver`  | String | Desmos address of the user that is the owner of the requested `DTag` |

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgCancelDTagTransferRequest",
  "sender": "desmos1nhgk008jvrxwa9tufr9tcr6zfrhe2uz0v90r2a",
  "receiver": "desmos1k99c8htyk32srx78efzg7sxm965prtz0j9qrc7"
}
````

## Message action
The action associated to this message is the following:

```
cancel_dtag_transfer_request
```
