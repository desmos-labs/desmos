# `MsgDeleteProfile`
This message allows you to delete the a previously created profile.

## Structure
````json
{
  "type": "desmos/MsgDeleteProfile",
  "value": {
    "moniker": "mazinga",
    "creator": "desmos1qchdngxk8zkl4c4mheqdlpgcegkdrtucmwllpx"
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `moniker` | String | Moniker of the user's profile |
| `creator` | String | Desmos address of the user that is creating the profile |

## Example
````json
{
  "type": "desmos/MsgDeleteProfile",
  "value": {
    "moniker": "mazinga",
    "creator": "desmos1qchdngxk8zkl4c4mheqdlpgcegkdrtucmwllpx"
  }
}
````

## Message action
The action associated to this message is the following:

```
delete_profile
```