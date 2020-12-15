# `MsgRefuseDTagTransferRequest`
This message allows you to refuse a `DTag` transfer request made by a user.

## Structure
````json
{
  "type": "desmos/MsgRefuseDTagTransferRequest",
  "value": {
    "sender": "<Desmos address that sent the DTag's request>",
    "receiver": "<Desmos address of the DTag owner>"
  }
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
  "type": "desmos/MsgRefuseDTagTransferRequest",
  "value": {
    "sender": "desmos1nhgk008jvrxwa9tufr9tcr6zfrhe2uz0v90r2a",
    "receiver": "desmos1k99c8htyk32srx78efzg7sxm965prtz0j9qrc7"
  }
}
````

## Message action
The action associated to this message is the following:

```
refuse_dtag_request
```
