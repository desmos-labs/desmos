# `MSgCreateSubspace`
This message allows you to create a subspace. If you want to know more about the `Subspace` type you can check the related
docs [here](../../../types/staging/subspaces/subspace.md)

## Structure
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgCreateSubspace",
  "id": "<ID of the subspace that will be created>",
  "name": "<Human readable name of the subspace>",
  "subspace_type": "<Indicates if users can post in it freely or not>",
  "creator": "<Desmos address of the subspace creator>"
}
```

## Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `subspace_id` | String | ID of the subspace to create |
| `name` | String | Human readable name of the subspace to create |
| `subspace_type` | Enum | Tells if users can post in it without being registered |
| `creator` |  String | Desmos address of the subspace creator |

The `subspace_type` field will only accept the following values:
```json
"open",
"close"
```

## Example
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgCreateSubspace",
  "id": "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
  "name": "mooncake",
  "subspace_type": "open",
  "creator": "desmos14dz9drkw0dyagnht5fnj6s63cwpxxkw8zsx7x9"
}
```

# Message action
The action associated to this message is the following:

```
create_subspace
```