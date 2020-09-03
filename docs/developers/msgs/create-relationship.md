# `MsgCreateRelationship`
This message allows you to create a relationship between the signer and a specified user.

## Structure
```json
{
  "type": "desmos/MsgCreateRelationship",
  "value": {
    "sender": "<Desmos address that's creating the relationship>",
    "receiver": "<Desmos address that's receiving the relationship>",
    "subspace": "<Subspace of a relationship>"
  }
}      
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `sender`  | String | Desmos address of the user that is creating the relationship |
| `receiver`| String | Desmos address of the relationship's recipient |
| `subspace`| String | Required string that identifies the app where the user make the relationship |

## Example
````json
{
  "type": "desmos/MsgCreateRelationship",
  "value": {
    "sender": "desmos1e209r8nc8qdkmqujahwrq4xrlxhk3fs9k7yzmw",
    "receiver": "desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud",
    "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
  }
}    
````

## Message action
The action associated to this message is the following: 

```
create_relationship
```
