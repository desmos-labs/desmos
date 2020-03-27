# `MsgRemovePostReaction`
With this message you can remove a reaction from a post you have previously reacted to.

## Structure
```json
{
  "type": "desmos/MsgRemovePostReaction",
  "value": {
    "reaction": "<Value of the reaction to be removed>",
    "liker": "<Desmos address of the user unliking the post>",
    "post_id": "<Id of the post to unlike>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `value` | String | Value of the reaction to be removed | 
| `liker` | String | Desmos address of the user that had added the reaction | 
| `post_id` | String | ID of the post from which to remove the reaction |

## Example
```json
{
  "type": "desmos/MsgRemovePostReaction",
  "value": {
    "value": "like",
    "liker": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "post_id": "12"
  }
}
```

## Message action
The action associated to this message is the following: 

```
remove_post_reaction
```