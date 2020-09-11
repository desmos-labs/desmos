# `MsgAcceptDTagTransfer`
This message allows you to accept a `dTag`'s transfer request.

## Structure
````json
{
  "type": "desmos/MsgAcceptDTagTransfer",
  "value": {
    "new_d_tag": "<the new dTag for the current owner's profile>",
    "current_owner": "<Desmos address of the dTag owner>",
    "receiving_user": "<Desmos address that's making the dTag's request>"
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `new_d_tag` | String | The new `dTag` for the current owner profile that will replace the traded one |
| `current_owner`  | String | Desmos address of the user that is the owner of the requested `dTag` |
| `receiving_user`| String | Desmos address of the user that request the `dTag` |

## Example
````json
{
  "type": "desmos/MsgAcceptDTagTransfer",
  "value": {
    "new_d_tag": "newDTag",
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
