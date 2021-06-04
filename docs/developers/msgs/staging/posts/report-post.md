# `MsgReportPost`
This message allows you to report a post. If you want to know more about the `Report` type, you can do so inside
the [`Report` type documentation page](../../../types/staging/posts/report.md).

## Structure
```json
{
  "@type": "/desmos.reports.v1beta1.MsgReportPost",
  "post_id": "<ID of the post to report>",
  "report": {
    "type": "<Report's type>",
    "message": "<Report's message>",
    "user": "<Desmos address that's creating the post>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `post_id` | String | ID of the post to report |
| `type`    | String | Type of the report |
| `message` | String | Message of the report |
| `user`    | String | Desmos address of the user that is reporting the post. |

The `type` field will only accept the following values:
```json
"nudity",  
"violence",  
"intimidation",  
"suicide or self-harm",  
"fake news",  
"spam",  
"unauthorized sale",  
"hatred incitement",  
"promotion of drug use",  
"non-consensual intimate images",  
"pornography",  
"children abuse",  
"animals abuse",  
"bullying",  
"scam" 
```

## Example
```json
{
  "@type": "/desmos.reports.v1beta1.MsgReportPost",
  "post_id": "301921ac3c8e623d8f35aef1886fea20849e49f08ec8ddfdd9b96feaf0c4fd15",
  "report": {
    "type": "scam",
    "message": "it's a trap",
    "user": "desmos1jnntz0xrql68mhjjsp82nlj9jrhgzc9t2ydtd5"
  }
}
```

## Message action
The action associated to this message is the following: 

```
report_post
```
