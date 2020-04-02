# Profile
Inside Desmos, profile are the way that users could register their own identity.

Since this is a decentralized profile, every data inside it can be omitted except for the `moniker` which could be used
to identify a desmos `bech32addr` once the profile is created. 

Profile can be enriched with some of your personal data. 
These data that could be verified later on chain by providing some additional information.

## Contained data
A profile contains different fields. Most of them could be omitted and each of them could be edited.

### `Name`
The `Name` identifies the name of the profile owner. It's omittable. Cannot be longer than `500` characters

### `Surname`
The `Surname` identifies the surname of the profile owner. It's omittable. Cannot be longer than `500` characters

### `Moniker`
The `Moniker` identifies the `moniker` of the profile owner. It can be made by at most `30` characters.

### `Bio`
The `Bio` represents the biography of the user. It can be at most `1000` characters long.

### `Pictures`
The [`Pictures`](docs/types/profile-pictures.md) contains the pictures of the account. This field is omittable.

### `VerifiedServices`
The `VerifiedServices` identifies a list of all trusted services ([`ServiceLink`](docs/types/profile-service-link.md)) linked to the profile.
They can be used to further improve the verification level of the profile. They can be omitted.

### `ChainLinks`
The `ChainLinks` identifies a list of [`ChainLink`](docs/types/profile-chain-link.md) that have been used to prove an external 
chain account is owned by this user of the profile.


### `Creator`
The `Creator` field is used to specify the Bech32 address of the creator of the profile. 
In order for a creator address to be valid, it must begin with the `desmos` Bech32 human-readable part. 

