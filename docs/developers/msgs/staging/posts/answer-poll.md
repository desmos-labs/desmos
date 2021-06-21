# `MsgAnswerPoll`
This message allows you to answer to a post's poll. 

## Structure
````json
{
  "@type": "/desmos.posts.v1beta1.MsgAnswerPoll",
  "post_id": "<ID of the post associated with the poll to be answered>",
  "answers": "<Array of answers' IDs matching the ones provided by the poll>",
  "answerer": "<Desmos address that's answering the poll>"
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post associated with the poll to which answer |
| `answers` | Array | Array of the answers' IDs |
| `answerer` | String | Desmos address of the user that is answering the poll |

## Example

```json
{
  "@type": "/desmos.posts.v1beta1.MsgAnswerPoll",
  "post_id": "a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc",
  "answers": [
    "1",
    "2"
  ],
  "answerer": "desmos174vlnmnfj34zckfeueqfl6s6vsq9hrek3emuy2"
}
```

## Message action
The action associated to this message is the following: 
```
answer_poll
```