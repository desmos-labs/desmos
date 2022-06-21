---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts 

## Profile
A Desmos profile is a wrap around a common Cosmos on-chain account that adds some fields to it.
It is built in order to represent an on-chain social identity that will be unique across the multiple social networks built on top of Desmos. 

Each profile is supposed to identify a single **on-chain identity**, but one human user can have multiple profiles connected each one to a different on-chain address. This is particularly useful if you want to have one private profile, one company profile and so on.

### DTag
The most important information about a profile it's the `DTag`, which represents the **unique human-readable name** associated to such profile. 

To understand it better, the best reference for this would be Twitter's handles: every user has a unique handle that other users can use to tag them. For example, Jack Dorsey's Twitter handle is `jack` ([@jack](https://twitter.com/jack) on Twitter).

If Jack registered a profile inside Desmos, he would probably like to have the same DTag as his Twitter handle: `jack` (which would be referenced as `@jack` inside Desmos too). 

When creating your profile, make sure you select a DTag that other people can easily remember if you want them to later find you inside Desmos.

Note that you **cannot choose whatever DTag you want**. Instead, you should reference the on-chain parameters to make sure that you respect the rules that have been set. Such rules might contain a set of characters allowed and/or a min/max length the DTag must adhere to. 

### Nickname 
A profile nickname represents what on other social networks is often called _username_. This is a non-unique name the user decides should be used to identify them. 

Such nickname has pretty much no restriction: it can contain any character and has only to be at least 2 characters long. 

### Bio
A profile biography allows the user to shortly describe their profile and their personality. It allows up to 1.000 characters and can contain anything the user might want to display publicly to everyone.

### Pictures
A profile owner can specify the profile picture and the cover picture that should be used when displaying the full details of their identity.

## DTag Transfer Request
A DTag transfer request is the method that a user A has in order to ask another user B if they are willing to give them their DTag. This process is particularly useful when trading DTags. Once a request has been made, it can be either accepted/rejected by the recipient, or it can be canceled by the sender themselves if they changed their mind.

## Chain Link
A chain link represents a link to an external chain account that has been created by the user to connect their Desmos profile to such account. These links can be created either offline or using the IBC protocol and the provided packet data types.

### Creating a chain link
#### 1. Create the ownership proofs
When creating a chain link, you need to provide two different proofs to make sure the link is valid:

1. The proof that you own the external chain account;
2. The proof that you own the Desmos profile to which you want to link.

In order to create a proof, the following steps are needed:

1. Get the address of the profile that should be connected to the external address;
2. Sign the hex-encoded external address using your private key (`sign(utf8.decode(address))`);
3. Assemble the signature, hex-encoded Desmos address and public key into a `Proof` object.

Here is an example of how to create a proof (in JavaScript):

```js
const desmosAddress = "cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm";
const hexEncodedDesmosAddress = hex.encode(utf8.decode(desmosAddress));
const signature = hex.encode(externalWallet.sign(utf8.decode(desmosAddress)));

const proof = {
    "pub_key": {
        "@type": "/cosmos.crypto.secp256k1.PubKey",
        "key": base64.encode(externalWallet.pubKeyBytes)
    },
    "signature": signature,
    "plain_text": hexEncodedDesmosAddress
}
```

Following an example of a proof JSON encoded:

```json
{
    "pub_key": {
        "@type": "/cosmos.crypto.secp256k1.PubKey",
        "key": "A58DXR/lXKVkIjLofXgST/OHi+pkOQbVIiOjnTy7Zoqo"
    },
    "signature": "ecc6175e730917fb289d3a9f4e49a5630a44b42d972f481342f540e09def2ec5169780d85c4e060d52cc3ffb3d677745a4d56cd385760735bc6db0f1816713be",
    "plain_text": "636f736d6f73313575633839766e7a756675356b7568687378646b6c747433387a66783876637967677a77666d"
}
```

Note that the `pub_key` field must be encoded using Protobuf and must be compatible with the public key types that are currently supported by Cosmos. You can see a list of such key types [here](https://github.com/cosmos/cosmos-sdk/tree/master/proto/cosmos/crypto).

#### 2. Create the link
Once you have created the two ownership proofs, you are now ready to create the link. This can be done in two ways:

1. [Using IBC](#using-ibc);
2. [Using the CLI](#using-the-cli).

##### Using IBC
This is the way that you want to use when integrating the Desmos connection from your chain.  
To implement the IBC capability of connecting an external account to a Desmos profile, the `x/profiles` module supports the following packet data type.

###### LinkChainAccountPacketData
`LinkChainAccountPacketData` defines the object that should be sent inside a `MsgSendPacket` when wanting to link an external chain to a Desmos profile using IBC.

```js reference
https://github.com/desmos-labs/desmos/blob/v3.0.0/x/profiles/types/models_packets.pb.go#L28-L43
```

Note that the `SourceAddress` field must be one of the currently supported types:

- `Base58Address` if the external address is represented by the Base58 encoded public key of the account;
- `Bech32Address` if the external address is Bech32 encoded;
- `HexAddress` if the external address is hex encoded.

##### Using the CLI
You can easily create a chain link using the CLI by running two commands:

1.`desmos create-chain-link-json`
This will start an interactive prompt session allowing you to generate the proper JSON file containing all the linkage information.

2. `desmos tx profiles link-chain [/path/to/link_file.json]`
   This will effectively link your Desmos profile to the external chain address. The required argument is the (absolute) path to the file generated using the `create-chain-link-json` command.


## Application Link
An application link (abbr. _app link_) represents a link to an external (and possibly centralized) application. Such links are one of the easiest way for Desmos profile owners to verify they are trustworthy: if you have a lot of followers on Twitter, connecting your Desmos profile to your Twitter account will make it easier for users to make sure no other people is trying to impersonate you.

### Creating an application link
Application links validity is checked using a multi-step verification process described inside the [_"Themis"_ repository](https://github.com/desmos-labs/themis).

### States of an application link
During its lifetime, an application link can have different states. Here it is a comprehensive lists of all the values the `State` field can assume: 

- `APPLICATION_LINK_STATE_INITIALIZED_UNSPECIFIED` if the link has just been created, and the verification process is pending to be started;
- `APPLICATION_LINK_STATE_VERIFICATION_STARTED` if the verification process has started;
- `APPLICATION_LINK_STATE_VERIFICATION_ERROR` if the verification process ended with an error;
- `APPLICATION_LINK_STATE_VERIFICATION_SUCCESS` if the verification process ended with success;
- `APPLICATION_LINK_STATE_TIMED_OUT` if the verification process expired due to a timeout.

Based on whether the application has been verified properly (`APPLICATION_LINK_STATE_VERIFICATION_SUCCESS`) or an error happened (`APPLICATION_LINK_STATE_VERIFICATION_ERROR`), the `Result` field will either be of type `Result_Success` or `Result_Failed`. Note that the `Result` field will remain empty in all other cases (including the `APPLICATION_LINK_STATE_TIMED_OUT`).
