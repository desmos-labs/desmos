# Profile
Inside Desmos, profile are the way that users could register their own identity.

Since this is a decentralized profile, every data inside it can be omitted except for the `dTag` which could be used
to identify a desmos `bech32addr` once the profile is created. 

Profile can be enriched with some of your personal data. 
These data that could be verified later on chain by providing some additional information.

## Contained data
A profile contains different fields. Most of them could be omitted and each of them could be edited.

### `DTag`
The `DTag` is a string of min `3` and max `30` characters that uniquely identifies the user.
In order to be valid it needs to match the following RegEx:

```
^[A-Za-z0-9_]+$
``` 

### `Nickname`
The `Nickname` represents the name of the user. It can be either a combination of first, second and last name, or a completely invented name. Although we always suggest setting one, this field is completely optional. 

### `Bio`
The `Bio` represents the biography of the user. It can be at most `1000` characters long.

### `Pictures`
The [`Pictures`](profile-pictures.md) contains the pictures of the account. This field is omittable.

### `Creator`
The `Creator` field is used to specify the Bech32 address of the creator of the profile. 
In order for a creator address to be valid, it must begin with the `desmos` Bech32 human-readable part. 

