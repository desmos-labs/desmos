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
    "subspace": "<Subspace of a post>",
    "optional_data": {},
    "creator": "<Desmos address that's creating the post>",
    "creation_date": "<RFC3339-formatted date representing the creation date of the post>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `parent_id` | String | ID of the parent post for which this post should be a comment of (Set to `0` if you do not want to have a parent) |
| `message` | String | Message of the post |
| `allows_comments` | Boolean | Tells whenever the post will allow other posts to reference to it as parent or not | 
| `susbspace` | String | Required string that identifies the posting app |
| `optional_data` | Map | Optional arbitrary data that you might want to store |
| `creator` | String | Desmos address of the user that is creating the post |
| `creation_date` | String | Date in RFC3339 format (e.g. `"2006-01-02T15:04:05Z07:00"`) in which the post has been created. Cannot be a future date. |

## Example
### With optional data
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "0",
    "message": "Desmos is great!",
    "allows_comments": true,
    "subspace": "desmos",
    "optional_data": {
      "custom_field": "My custom value"
    },
    "creator": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "creation_date": "2020-01-01T12:00:00Z"
  }
}
``` 

### Without optional data
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "0",
    "message": "Desmos is great!",
    "allows_comments": true,
    "subspace": "desmos",
    "creator": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "creation_date": "2020-01-01T12:00:00Z"
  }
}
```

## Message action
The action associated to this message is the following: 

```
create_post
```