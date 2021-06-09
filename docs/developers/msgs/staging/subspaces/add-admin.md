# `MsgAddAdmin`
This message allows you to add a new admin in a subspace.

## Structure
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgAddAdmin",
  "subspace_id": "<ID of the subspace where the admin should be added>",
  "admin": "<Desmos address of the new admin>",
  "owner": "<Desmos address of the subspace owner that will add the admin>"
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :---: | :--------- |
| `subspace_id` | String | ID of the subspace where the admin should be added |
| `admin` | String | Desmos address of the user who will be admin |
| `owner` | String | Desmos address of the owner that will add the new admin |

## Example
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgAddAdmin",
  "subspace_id": "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
  "admin": "desmos1yjk03uwmglpsfmyyfl803jxfr5dg0ltn7jyf24",
  "owner": "desmos19ayrtxc7aj5hv5j74uzfk0cdmackhpqq35hkpl"
}
```

## Message action
The action associated to this message is the following:

````
add_admin
````