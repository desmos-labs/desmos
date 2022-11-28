---
id: link-application
title: Link application
sidebar_label: Link application
slug: link-application
---

# `MsgLinkApplication`
This message allows you to start the process that will verify
an [application link](../../02-types/profiles/application-link.md#contained-data) and add it to your Desmos profile.

## Structure

```json
{
  "@type": "/desmos.profiles.v1beta1.MsgLinkApplication",
  "link_data": {
    "application": "<Name of the application to link>",
    "username": "<Username of the application account to link>"
  },
  "call_data": "<Hex encoded call data for the data source>",
  "source_channel": "<IBC channel to be used>",
  "source_port": "<IBC port to be used>",
  "sender": "<Desmos address of the profile to which link the application>"
}
```

### Attributes

| Attribute |                                Type                                | Description | Required |
| :-------: |:------------------------------------------------------------------:| :-------- | :------- |
| `link_data`  | [Data](../../02-types/profiles/application-link.md#contained-data) | Data of the link to be verified | yes |
| `call_data`|                               String                               | Hex encoded data that will be sent to the data source to verify the link | yes |
| `source_channel` |                               String                               | ID of the IBC channel to be used in order to send the packet | yes |
| `source_port` |                               String                               | ID of the IBC port to be used in order to send the packet | yes |
| `sender` |                               String                               | Desmos address of the profile to which the link will be associated | yes |

#### Note
You can also specify an optional timeout after which the request will be marked as invalid. This can be done using the
appropriate fields:

- `height` (`int64`), or
- `timeout_timestamp` (nanoseconds).

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgLinkApplication",
  "link_data": {
    "application": "github",
    "username": "RiccardoM"
  },
  "call_data": "7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
  "source_channel": "channel-0",
  "source_port": "profiles",
  "sender": "desmos1qchdngxk8zkl4c4mheqdlpgcegkdrtucmwllpx"
} 
````

## Message action
The action associated to this message is the following:

```
link_application
```
