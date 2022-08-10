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

### Signature
To be properly verified, a chain link must contain a signature that proves the user owns both the Desmos address and the external account that they are trying to link together. To do this, we require the user to cryptographically sign their Desmos address using the private key associated to the external account, and then publish the signature and the public key that should be used to verify it inside the chain link's proof. 

Currently, we support two types of signatures: 

- `CosmosMultiSignature` should be used when 

#### SingleSignature
This signature type should be used when the external wallet is made of a single key. It must specify the type of value that has been signed, as well as the signature bytes. 

Single signature value types can be different to support multiple use cases: 
- `SIGNATURE_VALUE_TYPE_RAW` should be used when you have direct access to the external account, and you can sign the Desmos address directly without the need of any wrapping structure;
- `SIGNATURE_VALUE_TYPE_COSMOS_DIRECT` should be used when you need to wrap the Desmos address into the memo field of a Protobuf-encoded transaction. This might be useful when wanting to support the creation of chain links through external wallets (i.e. Keplr) that only support the signing of transactions;
- `SIGNATURE_VALUE_TYPE_COSMOS_AMINO` should be used when you need to wrap the Desmos address into the memo field of an Amino-encoded transaction. This might be useful when wanting to support the creation of chain links through an external wallet (i.e. Ledger) that only allows signing Amino-encoded transactions;
- `SIGNATURE_VALUE_TYPE_EVM_PERSONAL_SIGN` should be used when you need to wrap the Desmos address within an [EVM `personal_sign` signature](https://github.com/ethereum/go-ethereum/pull/2940). This might be useful when wanting to support the creation of chain links though an external wallet (i.e. MetaMask) that allows this method.

#### CosmosMultiSignature
This signature type should be used when the external account is a multi-sig account (this might be the case of validators wanting to connect their Desmos profile to an external Cosmos-based account).

The `bit_array` field should be populated with the same value that is returned after the multi-sig creation process, while the `signatures` field should contain the list of `SingleSignature` that make up the multi-sig.

### Creating a chain link
#### 1. Create the ownership proofs
In order to create a valid chain link ownership proof, the following steps are needed:

1. get the address of the Desmos profile that should be connected to the external address;
2. sign the Desmos address using your external account private key;
3. assemble the signature, Desmos address and public key into a `Proof` object.

Here is an example of how to create a proof (in JavaScript) using a raw signature of the Desmos address. Please note that different value types might be used as described [above](#singlesignature):

```js
const desmosAddress = "cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm";
const hexEncodedDesmosAddress = hex.encode(utf8.decode(desmosAddress));
const signature = base64.encode(externalWallet.sign(utf8.decode(desmosAddress)));

const proof = {
    "pub_key": {
        "@type": "/cosmos.crypto.secp256k1.PubKey",
        "key": base64.encode(externalWallet.pubKeyBytes)
    },
   "signature": {
      "@type": "/desmos.profiles.v3.SingleSignature",
      "value_type": 1,
      "signature": signature,
   },
    "plain_text": hexEncodedDesmosAddress
}
```

Note that the `pub_key` field must be encoded using Protobuf and must be compatible with the public key types that are currently supported. These are the following:
- [Ed25519](https://github.com/cosmos/cosmos-sdk/blob/main/proto/cosmos/crypto/ed25519/keys.proto)
- [Secp256k1](https://github.com/cosmos/cosmos-sdk/blob/main/proto/cosmos/crypto/secp256k1/keys.proto)
- [Secp256r1](https://github.com/cosmos/cosmos-sdk/blob/main/proto/cosmos/crypto/secp256r1/keys.proto)
- [EthSecp256k1](https://github.com/evmos/ethermint/blob/main/proto/ethermint/crypto/v1/ethsecp256k1/keys.proto)

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
https://github.com/desmos-labs/desmos/blob/master/x/profiles/types/models_packets.pb.go#L28-L43
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


## Default External Address
Since each user can have multiple links to the same chain, we allow users to specify which address should be considered as the default one for each chain. By default, the first link created for each chain will be considered as the default one. Also, if a chain link was set as the default one and is later deleted, the oldest chain link for the same chain (if any) will be used as the default one.

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
