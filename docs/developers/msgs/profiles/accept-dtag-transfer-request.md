# `MsgAcceptDTagTransfer`
This message allows you to accept a `DTag`'s transfer request.

## Structure
````json
{
  "type": "desmos/MsgAcceptDTagTransfer",
  "value": {
    "new_dtag": "<The new DTag for the current owner's profile>",
    "current_owner": "<Desmos address of the DTag owner>",
    "receiving_user": "<Desmos address that's making the DTag's request>"
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `new_dtag` | String | The new `DTag` for the current owner profile that will replace the traded one |
| `current_owner`  | String | Desmos address of the user that is the owner of the requested `DTag` |
| `receiving_user`| String | Desmos address of the user that request the `DTag` |

## Example
````json
{
  "type": "desmos/MsgAcceptDTagTransfer",
  "value": {
    "new_dtag": "newDTag",
    "current_owner": "desmos1k99c8htyk32srx78efzg7sxm965prtz0j9qrc7",
    "receiving_user": "desmos1nhgk008jvrxwa9tufr9tcr6zfrhe2uz0v90r2a"
  }
}
````

## Message action
The action associated to this message is the following:

```
accept_dtag_request
```
