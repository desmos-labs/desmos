# `MsgSaveProfile`
This message allows you to save a new profile or edit an existent one.

## Structure
````json
{
  "type": "desmos/MsgSaveProfile",
  "value": {
    "dtag": "<Profile dtag>",
    "moniker": "<Profile moniker>",
    "bio": "<Profile biography>",
    "pictures": {
      "profile": "<URI of the profile account's picture>",
      "cover": "<URI of the profile cover picture>"
    },
    "creator": "<Desmos address that's creating the profile>"
  }
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `dtag` | String | DTag of the user |
| `moniker` | String (Optional) | Moniker of the user | 
| `bio` | String | (Optional) Biography of the user |
| `pictures` | Object | (Optional) Object containing all the information related to the profile's pictures |
| `creator` | String | Desmos address of the user that is editing the profile |

If you are editing an existing profile you should fill all the existent fields otherwise they will be set as nil.

## Example
````json
{
  "type": "desmos/MsgSaveProfile",
  "value": {
    "dtag": "Eva00",
    "moniker": "Rei Ayanami",
    "bio": "The real pilot",
    "pictures": {
      "profile": "https://shorturl.at/adnX3",
      "cover": "https://shorturl.at/cgpyF"
    },
    "creator": "desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t"
  }
}
````

## Message action
The action associated to this message is the following:

```
save_profile
```

