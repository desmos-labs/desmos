# `MsgRemovePostReport`
This message allows you to remove a report you have previously created from a post.

## Structure
```json
{
    "@type": "/desmos.posts.v1beta1.MsgRemovePostReport",
    "post_id": "<ID of the post from which to remove the report>",
    "user": "<Desmos address of the user who is removing the report>",
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post from which to remove the report |
| `user` | String | Desmos address of the user who is removing the report | 

## Example
```json
{
    "@type": "/desmos.posts.v1beta1.MsgRemovePostReport",
    "post_id": "301921ac3c8e623d8f35aef1886fea20849e49f08ec8ddfdd9b96feaf0c4fd15",
    "user": "desmos1jnntz0xrql68mhjjsp82nlj9jrhgzc9t2ydtd5",
}
```

## Message action
The action associated to this message is the following: 

```
remove_post_report
```
