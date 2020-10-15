# `MsgCreatePost`
This message allows you to create a new public post. If you want to know more about the `Post` type, you can do so inside the [`Post` type documentation page](../../../types/posts/post.md).

## Structure
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "<ID of the post for which this post should be a comment of>",
    "message": "<Post message>",
    "allows_comments": false,
    "subspace": "<Subspace of a post>",
    "optional_data": [],
    "creator": "<Desmos address that's creating the post>",
    "attachments": "<Attachment's array that contains all the attachments associated with the post",
    "poll_data": "<Poll data contains all useful data of the poll's post>"
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
| `optional_data` | Array | Optional arbitrary data that you might want to store |
| `creator` | String | Desmos address of the user that is creating the post |
| `attachments` | Array | (Optional) Array containing all the attachments related to the post |
| `poll_data` | Object | (Optional) Object containing all the information related to post's poll, if exists |

## Example
### With optional data, attachments and poll data
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "",
    "message": "Desmos is great!",
    "allows_comments": true,
    "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
    "optional_data": [
      {
        "key": "My custom key",
        "value": "My custom value"
      }
    ],
    "creator": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "attachments": [
      {
        "uri": "https://example.com/media1",
        "mime_type": "text/plain",
        "tags": [
            "desmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
            "desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru"
        ]   
      },
      {
        "uri": "https://example.com/media2",
        "mime_type": "application/json"
      }
    ],
    "poll_data": {
      "question": "Which dog do you prefer?",
      "provided_answers": [
        {
          "id": 0,
          "text": "Beagle"
        },
        {
          "id": 1,
          "text": "Pug"
        },
        {
          "id": 2,
          "text": "German Sheperd"
        }
      ],
      "end_date": "2020-02-10T15:00:00Z",
      "allows_multiple_answers": true,
      "allows_answer_edits": true
    }
  }
}
``` 

### Without optional data, attachments and poll data
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "",
    "message": "Desmos is great!",
    "allows_comments": true,
    "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
    "creator": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
  }
}
```

## Message action
The action associated to this message is the following: 

```
create_post
```
