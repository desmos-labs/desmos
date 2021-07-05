# `MsgRequestDTagTransfer`
This message allows you to request a transfer to your profile for a `DTag` owned by another user.

## Structure
````json
{
  "@type": "/desmos.profiles.v1beta1.MsgRequestDTagTransfer",
  "receiver": "<Desmos address of the DTag owner>",
  "sender": "<Desmos address that's making the DTag's request>"
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `receiver`  | String | Desmos address of the user that is the owner of the requested `DTag` |
| `sender`| String | Desmos address of the user that request the `DTag` |

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgRequestDTagTransfer",
  "receiver": "desmos1k99c8htyk32srx78efzg7sxm965prtz0j9qrc7",
  "sender": "desmos1nhgk008jvrxwa9tufr9tcr6zfrhe2uz0v90r2a"
}
````

## Message action
The action associated to this message is the following:

```
request_dtag_transfer
```
