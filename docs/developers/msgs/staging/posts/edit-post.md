# `MsgEditPost`
This message allows you to edit the message of a previously published public post.

## Structure
```json
{
  "@type": "/desmos.posts.v1beta1.MsgEditPost",
  "post_id": "<ID of the post to be edited>",
  "message": "<New post message>",
  "comments_state": "<Indicates if the post allows comments or not>",
  "attachments": "<Attachment's array that contains all the attachments associated with the post",
  "poll": "<Poll contains all useful data of the poll's post>", 
  "editor": "<Desmos address of the user editing the message>"
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post to edit |
| `message` | String | New message that should be set as the post message |
| `comments_state` | Enum | Tells whenever the post will allow other posts to reference to it as parent or not |
| `attachments` | Array | (Optional) Array containing all the attachments related to the post |
| `poll` | Object | (Optional) Object containing all the information related to post's poll, if exists |
| `editor` | String | Desmos address of the user that is editing the post. This must be the same address of the original post creator. |

## Example
### Without attachments and pollData

```json
{
  "@type": "/desmos.posts.v1beta1.MsgEditPost",
  "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc",
  "message": "Desmos you rock!",
  "editor": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax"
}
```

### With attachments and pollData

```json
{
  "@type": "/desmos.posts.v1beta1.MsgEditPost",
  "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc",
  "message": "Desmos you rock!",
  "comments_state": "blocked",
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
  "poll": {
    "question": "Which dog do you prefer?",
    "provided_answers": [
      {
        "answer_id": 0,
        "text": "Beagle"
      },
      {
        "answer_id": 1,
        "text": "Pug"
      },
      {
        "answer_id": 2,
        "text": "German Sheperd"
      }
    ],
    "end_date": "2020-02-10T15:00:00Z",
    "is_open": true,
    "allows_multiple_answers": true,
    "allows_answer_edits": true
  },
  "editor": "desmos1w3fe8zq5jrxd4nz49hllg75sw7m24qyc7tnaax"
}
```

## Message action
The action associated to this message is the following: 

```
edit_post
```