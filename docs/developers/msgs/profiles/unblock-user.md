# `MsgUnblockUser`
This message allows you to unblock a previously blocked user.

## Structure
```json
{
  "@type": "/desmos.profiles.v1beta1.MsgUnblockUser",
  "blocker": "<Desmos address of the user that is unblocking another user>",
  "blocked": "<Desmos address of the unblocked user>",
  "subspace": "<Subspace of the block>"
}   
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `blocker`  | String | Desmos address of the user that is blocking someone else |
| `blocked`| String | Desmos address of the unblocked user |
| `subspace` | String | String that identifies the app for which the block was valid |

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgUnblockUser",
  "blocker": "desmos1j83hlf5yn5839wgpege3z669r8j3lh2ggmtf5u",
  "blocked": "desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru",
  "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
}   
````

## Message action
The action associated to this message is the following: 

```
unblock_user
```
