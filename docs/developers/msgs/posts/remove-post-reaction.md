# `MsgRemovePostReaction`
With this message you can remove a reaction from a post you have previously reacted to.

## Structure
```json
{
  "@type": "/desmos.posts.v1beta1.MsgRemovePostReaction",
  "post_id": "<Id of the post to unlike>",
  "user": "<Desmos address of the user who is removing the reaction>",
  "value": "<Shortcode of the reaction or Emoji to be removed>"
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post from which to remove the reaction |
| `user` | String | Desmos address of the user who is removing the reaction | 
| `value` | String | Shortcode of the reaction or Emoji to add | 

## Example with emoji

```json
{
  "@type": "/desmos.posts.v1beta1.MsgRemovePostReaction",
  "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc",
  "user": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
  "value": "👍"
}
```

## Example with shortcode

```json
{
  "@type": "/desmos.posts.v1beta1.MsgRemovePostReaction",
  "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc",
  "user": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
  "value": ":+1:"
}
```

## Message action
The action associated to this message is the following: 

```
remove_post_reaction
```
