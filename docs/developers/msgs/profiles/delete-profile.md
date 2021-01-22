# `MsgDeleteProfile`
This message allows you to delete a previously created profile.

## Structure
````json
{
  "@type": "/desmos.profiles.v1beta1.MsgDeleteProfile",
  "creator": "<Address of the profile owner>"
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `creator` | String | Desmos address of the user that is deleting the profile |

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgDeleteProfile",
  "creator": "desmos1qchdngxk8zkl4c4mheqdlpgcegkdrtucmwllpx"
}
````

## Message action
The action associated to this message is the following:

```
delete_profile
```