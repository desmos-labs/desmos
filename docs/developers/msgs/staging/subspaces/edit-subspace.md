# `MSgEditSubspace`
This message allows you to edit an existent subspace.

## Structure
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgEditSubspace",
  "id": "<ID of the subspace to edit>",
  "owner": "<Desmos address of the new owner of the subspace>",
  "name": "<New subspace name>",
  "description": "<Description of the subspace>",
  "logo": "<URI of the picture that identifies the subspace>",
  "subspace_type": "<Indicates if users can post in it freely or not>",
  "editor": "<Desmos address of the subspace editor>",
}
```

## Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `subspace_id` | String | ID of the subspace to edit |
| `owner` |  String | Desmos address of the new owner of the subspace |
| `name` | String | New name of the subspace |
| `description` | String | Description of the subspace |
| `logo` | String | URI of the picture that identifies the subspace |
| `subspace_type` | Enum | Tells if users can post in it without being registered |
| `editor` |  String | Desmos address of the subspace editor |

The `subspace_type` field will only accept the following values:
```json
"open",
"close"
```

## Example
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgEditSubspace",
  "id": "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
  "owner": "desmos1tqzrfy9ujrk883e2wezsumyvq64gcm65vhdyr7",
  "name": "mooncake",
  "description": "a good cake with secret messages in it",
  "logo": "https://mooncake-logo-png.com",
  "subspace_type": "<Indicates if users can post in it freely or not>",
  "editor": "desmos14dz9drkw0dyagnht5fnj6s63cwpxxkw8zsx7x9",
}
```

# Message action
The action associated to this message is the following:

```
edit_subspace
```