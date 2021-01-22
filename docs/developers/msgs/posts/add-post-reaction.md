# `MsgAddPostReaction`
This message allows you to add a reaction to a post that is already existing on the chain. 

## Structure
```json
{
  "@type": "/desmos.posts.v1beta1.MsgAddPostReaction",
  "post_id": "<Id of the post to like>",
  "reaction": "<Shortcode of the reaction or Emoji to add>",
  "user": "<Desmos address of the user liking the post>"
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post to which add the reaction | 
| `reaction` | String | Shortcode of the reaction or Emoji to add | 
| `user` | String | Desmos address of the user adding the reaction to the post | 

## Example with emoji

```json
{
  "@type": "/desmos.posts.v1beta1.MsgAddPostReaction",
  "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc",
  "reaction": "👍",
  "user": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax"
}
```

## Example with shortcode

```json
{
  "@type": "/desmos.posts.v1beta1.MsgAddPostReaction",
  "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc",
  "reaction": ":+1:",
  "user": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax"
}
```

## Message action
The action associated to this message is the following: 

```
add_post_reaction
```
