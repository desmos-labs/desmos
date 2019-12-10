# `MsgUnlikePost`
With this message you can remove the like from a previously liked post. 

## Structure
```json
{
  "type": "desmos/MsgUnlikePost",
  "value": {
    "liker": "<Desmos address of the user unliking the post>",
    "post_id": "<Id of the post to unlike>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `liker` | String | Desmos address of the user unliking the post | 
| `post_id` | String | ID of the post to be unliked |

## Example
```json
{
  "type": "desmos/MsgUnlikePost",
  "value": {
    "liker": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "post_id": "12"
  }
}
```

## Message action
The action associated to this message is the following: 

```
unlike_post
```