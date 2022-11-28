---
id: block-user
title: Block user
sidebar_label: Block user
slug: block-user
---

# `MsgBlockUser`
This message allows you to block a specific user associated with a given address.

## Structure
```json
{
  "@type": "/desmos.profiles.v1beta1.MsgBlockUser",
  "blocker": "<Desmos address that's blocking someone>",
  "blocked": "<Desmos address that's been blocked>",
  "reason": "<Reason of the block>",
  "subspace": "<Subspace of the block>"
}   
```

### Attributes
| Attribute | Type | Description | Required |
| :-------: | :----: | :-------- | :------- |
| `blocker`  | String | Desmos address of the user that is blocking someone else | yes |
| `blocked`| String | Desmos address of the blocked user | yes |
| `reason` | String | Reason for the block | no |
| `subspace` | String | SHA256 string that identifies the app for which the block should apply | yes |

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgBlockUser",
  "blocker": "desmos1j83hlf5yn5839wgpege3z669r8j3lh2ggmtf5u",
  "blocked": "desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru",
  "reason": "reason",
  "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
}   
````

## Message action
The action associated to this message is the following: 

```
block_user
```
