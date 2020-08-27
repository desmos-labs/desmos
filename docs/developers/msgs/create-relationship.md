# `MsgCreateMonoDirectionalRelationship`
This message allows you to create a mono directional relationship with a specified user.
Mono directional relationships are like the follow of today's social networks.

## Structure
```json
{
  "type": "desmos/MsgCreateMonoDirectionalRelationship",
  "value": {
    "sender": "<Desmos address that's creating the relationship>",
    "receiver": "<Desmos address that's receiving the relationship>"
  }
}      
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `sender`  | String | Desmos address of the user that is creating the relationship |
| `receiver`| String | Desmos address of the relationship's recipient |

## Example
````json
{
  "type": "desmos/MsgCreateMonoDirectionalRelationship",
  "value": {
    "sender": "desmos1e209r8nc8qdkmqujahwrq4xrlxhk3fs9k7yzmw",
    "receiver": "desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud"
  }
}    
````

## Message action
The action associated to this message is the following: 

```
create_relationship
```
