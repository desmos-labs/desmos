# `MsgSaveProfile`
This message allows you to save a new profile or edit an existent one.

## Structure
````json
{
  "@type": "/desmos.profiles.v1beta1.MsgSaveProfile",
  "dtag": "<Profile dtag>",
  "nickname": "<Profile nickname>",
  "bio": "<Profile biography>",  
  "profile_picture": "<URI of the profile account's picture>",
  "cover_picture": "<URI of the profile cover picture>",
  "creator": "<Desmos address that's creating the profile>"
}
````

### Attributes
| Attribute | Type | Description |
| :-------: | :----: | :-------- |
| `dtag` | String | DTag of the user |
| `nickname` | String (Optional) | Nickname of the user | 
| `bio` | String | (Optional) Biography of the user |
| `profile_picture` | String | (Optional) URL to the user profile picture |
| `cover_picture` | String | (Optional) URL to the user cover picture |
| `creator` | String | Desmos address of the user that is editing the profile |

If you are editing an existing profile and do not want to edit some fields, you can use the special
value `[do-not-modify]`.

## Example

````json
{
  "@type": "/desmos.profiles.v1beta1.MsgSaveProfile",
  "dtag": "Eva00",
  "nickname": "Rei Ayanami",
  "bio": "The real pilot",
  "profile_picture": "[do-not-modify]",
  "cover_picture": "https://shorturl.at/cgpyF",
  "creator": "desmos12a2y7fflz6g4e5gn0mh0n9dkrzllj0q5vx7c6t"
}
````

## Message action
The action associated to this message is the following:

```
save_profile
```

