# `MsgAnswerPoll`
This message allows you to answer to a post's poll. 

## Structure
````json
{
  "type": "desmos/MsgAnswerPoll",
  "value": {
    "post_id": "<ID of the post associated with the poll to be answered>",
    "answers": "<Array of answers' IDs matching the ones provided by the poll>",
    "answerer": "<Desmos address that's answering the poll>" 
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post associated with the poll to which answer |
| `answers` | []Uint | Array of the answers' IDs |
| `answerer` | String | Desmos address of the user that is answering the poll |

## Example
```json
{
  "type": "desmos/MsgAnswerPoll",
  "value": {
    "post_id": "1",
    "answers": [
      "1",
      "2"
    ],
    "answerer": "desmos174vlnmnfj34zckfeueqfl6s6vsq9hrek3emuy2"
  }
}
```

## Message action
The action associated to this message is the following: 
