# `MsgDeleteRelationship`
This message allows you to delete an existing relationship with a specified counterparty.

## Structure
````json
{
  "type": "desmos/MsgDeleteRelationship",
  "value": {
    "user": "<Desmos address of the user deleting the relationship>",
    "counterparty": "<Desmos address with which the sender want to end the relationship>",
    "subspace": "<Subspace of the relationship>"
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `user`  | String | Desmos address of the user that is deleting the relationship |
| `counterparty`| String | Desmos address of the relationship's counterparty |
| `subspace`| String | Identifies the app where the relationship should be valid |

## Example
````json
{
  "type": "desmos/MsgDeleteRelationship",
  "value": {
    "user": "desmos1e209r8nc8qdkmqujahwrq4xrlxhk3fs9k7yzmw",
    "counterparty": "desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud",
    "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
  }
} 
````

## Message action
The action associated to this message is the following: 

```
delete_relationship
```
