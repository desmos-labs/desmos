# `MsgEditProfile`
This message allows you to edit the a previously created profile.

## Structure
````json
{
  "type": "desmos/MsgEditProfile",
  "value": {
    "previous_moniker": "<Profile previous moniker>",
    "new_moniker": "<Profile new moniker>",
    "name": "<Profile name>",
    "surname": "<Profile surname>",
    "bio": "<Profile biography>",
    "pictures": {
      "profile": "<Profile account's picture>",
      "cover": "<Profile cover picture>"
    },
    "creator": "<Desmos address that's creating the profile>"
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `previous_moniker` | String | Moniker of the user's profile |
| `new_moniker` | String | (Optional) New moniker of the user's profile |
| `name` | String | (Optional) Name of the user |
| `surname` | String | (Optional) Surname of the user |
| `bio` | String | (Optional) Biography of the user |
| `pictures` | Object | (Optional) Object containing all the information related to the profile's pictures |
| `creator` | String | Desmos address of the user that is creating the profile |

## Example
````json
{
  "type": "desmos/MsgEditProfile",
  "value": {
    "previous_moniker": "Eva01",
    "new_moniker": "Eva00",
    "name": "Rei",
    "surname": "Ayanami",
    "bio": "evaPilot",
    "pictures": {
      "profile": "eva00Pic",
      "cover": "nervCover"
    },
    "creator": "desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t"
  }
}
````

## Message action
The action associated to this message is the following:

```
edit_profile
```