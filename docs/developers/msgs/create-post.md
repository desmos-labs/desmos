# `MsgCreatePost`
This message allows you to create a new public post. 

## Structure
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "<ID of the post for which this post should be a comment of>",
    "message": "<Post message>",
    "allows_comments": false,
    "external_reference": "<Arbitrary external reference>",
    "creator": "<Desmos address that's creating the post>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `parent_id` | String | ID of the parent post for which this post should be a comment of (Set to `0` if you do not want to have a parent) |
| `message` | String | Message of the post |
| `allows_comments` | Boolean | Tells whenever the post will allow other posts to reference to it as parent or not | 
| `external_reference` | String | Arbitrary external reference to the post |
| `creator` | String | Desmos address of the user that is creating the post |

## Example
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "0",
    "message": "Desmos is great!",
    "allows_comments": true,
    "external_reference": "",
    "creator": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax"
  }
}
``` 

## Message action
The action associated to this message is the following: 

```
create_post
```