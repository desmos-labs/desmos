# `MsgRegisterReaction`
This message allows you to register a new reaction.  
If you want to know more about the `Reaction` type, you can do so inside the [`Reaction` type documentation page](../../types/reaction.md)

## Structure
```json
{
  "type": "desmos/MsgRegisterReaction",
  "value": {
    "short_code": "<reaction short code>",
    "value": "<url (identifing gif or image) or unicode (identifing emoji)>",
    "subspace": "<Reaction subspace>",
    "creator": "<Desmos address that's registering the reaction>"
  }
}
```

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `short_code` | String | Short code that identifies the reaction (e.g. `":smile:"`)  |
| `value` | String | Value can be a URL or a UNICODE identifing gif, images, emoji   |
| `subspace` | String | Required string that identifies the subspace inside which the reaction will be registered |
| `creator` | String | Desmos address of the user that is registering the reaction |

## Example
```json
{
  "type": "desmos/MsgRegisterReaction",
  "value": {
    "shortcode": ":smile:",
    "value": "U+1F604",
    "subspace": "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
    "creator": "desmos13s7p4jx3rj5pxjzlecxdvua68ex0sg7rug0pt3"
  }
}
```

## Message action
The action associated to this message is the following: 

```
register_reaction
```