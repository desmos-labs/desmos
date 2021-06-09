# `MsgUnbanUser`
This message allows you to unban a user in a subspace.

## Structure
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgUnregisterUser",
  "subspace_id": "<ID of the subspace from where the user will be unbanned>",
  "user": "<Desmos address of the user that will be unbanned>",
  "admin": "<Desmos address of the admin that will unban the user>"
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :---: | :--------- |
| `subspace_id` | String | ID of the subspace from where the user will be unbanned |
| `user` | String | Desmos address of the user that will be unbanned |
| `admin` | String | Desmos address of the admin that will unban the user |

## Example
```json
{
  "@type": "/desmos.subspaces.v1beta1.MsgBanUser",
  "subspace_id": "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
  "user": "desmos1tqzrfy9ujrk883e2wezsumyvq64gcm65vhdyr7",
  "admin": "desmos14dz9drkw0dyagnht5fnj6s63cwpxxkw8zsx7x9"
}
```

## Message action
The action associated to this message is the following:

````
unban_user
````



