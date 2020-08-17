# `MsgEditPost`
This message allows you to edit the message of a previously published public post.

## Structure
```json
{
  "type": "desmos/MsgEditPost",
  "value": {
    "post_id": "<ID of the post to be edited>",
    "message": "<New post message>",
    "editor": "<Desmos address of the user editing the message>",
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post to edit |
| `message` | String | New message that should be set as the post message |
| `editor` | String | Desmos address of the user that is editing the post. This must be the same address of the original post creator. |

## Example
### With optional data
```json
{
  "type": "desmos/MsgEditPost",
  "value": {
    "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc",
    "message": "Desmos you rock!",
    "editor": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
  }
}
```

## Message action
The action associated to this message is the following: 

```
edit_post
```