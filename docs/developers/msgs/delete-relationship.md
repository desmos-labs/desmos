# `MsgDeleteRelationship`
This message allows you to delete a mono directional relationship with a specified counterparty.

## Structure
````json
{
  "type": "desmos/MsgDeleteRelationship",
  "value": {
    "sender": "<Desmos address that's deleting the relationship>",
    "counterparty": "<Desmos address that's with which sender want to cut-off the relationship>"
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `sender`  | String | Desmos address of the user that is deleting the relationship |
| `counterparty`| String | Desmos address of the relationship's counterparty |

## Example
````json
{
  "type": "desmos/MsgDeleteRelationship",
  "value": {
    "sender": "desmos1e209r8nc8qdkmqujahwrq4xrlxhk3fs9k7yzmw",
    "counterparty": "desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud"
  }
} 
````

## Message action
The action associated to this message is the following: 

```
delete_relationship
```
