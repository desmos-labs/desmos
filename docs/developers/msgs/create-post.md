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
    "creation_date": "<RFC3339-formatted date representing the creation date of the post>",
    "medias": "<Media's array that contains all the medias associated with the post",
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
| `optional_data` | Map | Optional arbitrary data that you might want to store |
| `creator` | String | Desmos address of the user that is creating the post |
| `creation_date` | String | Date in RFC3339 format (e.g. `"2006-01-02T15:04:05Z07:00"`) in which the post has been created. Cannot be a future date. |
| `medias` | PostMedias | PostMedia's array contains all the medias related to the post |
| `poll_data` | *PollData | Pointer to Poll Data that contains all the information related to post's poll, if exists |
## Example
### With optional data, medias and poll data
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "0",
    "message": "Desmos is great!",
    "allows_comments": true,
    "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
    "optional_data": {
      "custom_field": "My custom value"
    },
    "creator": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "creation_date": "2020-01-01T12:00:00Z",
    "medias": [
      {
        "uri": "https://example.com/media1",
        "mime_type": "text/plain"
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
      "is_open": true,
      "allows_multiple_answers": true,
      "allows_answer_edits": true
    }
  }
}
``` 

### Without optional data, medias and poll data
```json
{
  "type": "desmos/MsgCreatePost",
  "value": {
    "parent_id": "0",
    "message": "Desmos is great!",
    "allows_comments": true,
    "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
    "creator": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax",
    "creation_date": "2020-01-01T12:00:00Z"
  }
}
```

## Message action
The action associated to this message is the following: 

```
answer_poll
```