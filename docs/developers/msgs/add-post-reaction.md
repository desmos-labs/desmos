# `MsgAddPostReaction`
This message allows you to add a reaction to a post that is already existing on the chain. 

## Structure
```json
{
  "type": "desmos/MsgAddPostReaction",
  "value": {
    "value": "<Value of the reaction>",
    "user": "<Desmos address of the user liking the post>",
    "post_id": "<Id of the post to like>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `value` | String | Value of the reaction | 
| `user` | String | Desmos address of the user adding the reaction to the post | 
| `post_id` | String | ID of the post to which add the reaction | 

## Example
```json
{
  "type": "desmos/MsgAddPostReaction",
  "value": {
    "value": "like",
    "liker": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc"
  }
}
```

## Message action
The action associated to this message is the following: 

```
add_post_reaction
```