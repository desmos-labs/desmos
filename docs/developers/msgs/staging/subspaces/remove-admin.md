# `MsgRemoveAdmin`
This message allows you to remove an admin from a subspace.

## Structure
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgRemoveAdmin",
  "subspace_id": "<ID of the subspace from where the admin will be removed>",
  "admin": "<Desmos address of the admin that will be removed>",
  "owner": "<Desmos address of the subspace owner that will remove the admin>"
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :---: | :--------- |
| `subspace_id` | String | ID of the subspace from where the user will be removed |
| `admin` | String | Desmos address of the admin that will be removed |
| `owner` | String | Desmos address of the admin that will perform the removal |

## Example
```json
 {
  "@type": "/desmos.subspaces.v1beta1.MsgRemoveAdmin",
  "subspace_id": "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
  "admin": "desmos1tqzrfy9ujrk883e2wezsumyvq64gcm65vhdyr7",
  "owner": "desmos14dz9drkw0dyagnht5fnj6s63cwpxxkw8zsx7x9"
}
```

## Message action
The action associated to this message is the following:

````
remove_admin
````