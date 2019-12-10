# `MsgLikePost`
This message allows you to like a post that is already existing on the chain. 

## Structure
```json
{
  "type": "desmos/MsgLikePost",
  "value": {
    "liker": "<Desmos address of the user liking the post>",
    "post_id": "<Id of the post to like>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `liker` | String | Desmos address of the user liking the post | 
| `post_id` | String | ID of the post to be liked | 

## Example
```json
{
  "type": "desmos/MsgLikePost",
  "value": {
    "liker": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "post_id": "12"
  }
}
```

## Message action
The action associated to this message is the following: 

```
like_post
```