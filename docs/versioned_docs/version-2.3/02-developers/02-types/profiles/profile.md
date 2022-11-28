---
id: profile
title: Profile
sidebar_label: Profile
slug: profile
---

# Profile
Inside Desmos, profile are the way that users could register their own identity.

Since this is a decentralized profile, every data inside it can be omitted except for the `DTag` which could be used to identify a Desmos wallet once the profile is created. 

Profile can be enriched with some of your personal data that could be verified later on chain by providing
some additional information.

## Contained data
A profile contains different fields. Most of them are optional and can be edited.

### `DTag` (`string`)
The `DTag` is a `string` of min `6` and max `30` characters that uniquely identifies the user.
In order to be valid it needs to match the following RegEx:

```
^[A-Za-z0-9_]+$
``` 

### `Nickname` (`string`)
The `Nickname` represents the name of the user. 
It can be either a combination of first, second and last name, or a completely invented name. 
Although we always suggest setting one, this field is completely optional.
If specified, it has to be at least of `2` characters long.

### `Bio` (`string`)
The `Bio` contains the biography of the user. It can be at most `1000` characters long.

### `Pictures` (`object`)
The [Pictures](profile-pictures.md) contains the pictures of the profile.

### `Account` (`object`)
The `Account` represents the base Cosmos account associated with the profile and contains information
such as the user's Bech32 address, public key etc...
