---
id: link-chain-account
title: Link chain account
sidebar_label: Link chain account
slug: link-chain-account
---

# `MsgLinkChainAccount`
This message allows you to link an external chain address 
and add the [chain link](../../02-types/profiles/chain-link.md) of it to your Desmos profile.

## Structure

```json
{
    "@type":"desmos/MsgLinkChainAccount",
    "chain_address": {
        "@type": "/desmos.profiles.v1beta1.Bech32Address",
        "prefix": "<Bech32 prefix of the external chain account>",
        "value": "<Address of the external chain account>"
    },
    "chain_config": {
        "name": "<Name of the target external chain>"
    },
    "proof": {
        "pub_key": {
            "@type": "/cosmos.crypto.secp256k1.PubKey",
            "value": "<Base64 encoded public key>"
        },
        "signature": "<Hex encoded signature created with private key associated to the given public key>",
        "plain_text": "<Text signed with private key>"
    },
    "signer": "<Desmos address of the profile linking the chain account>"
}
```

### Attributes

| Attribute |                                         Type                                         | Description | Required |
| :-------: |:------------------------------------------------------------------------------------:| :-------- | :------- |
| `chain_address` | [AddressData](../../02-types/profiles/chain-link.md#Contained-data) | Address data of the external chain account | yes |
| `chain_config` | [ChainConfig](../../02-types/profiles/chain-link.md#Contained-data) | Details of the target external chain | yes |
| `proof` |    [Proof](../../02-types/profiles/chain-link.md#Contained-data)    | Data proving the ownership of the external chain account | yes |
| `signer` |                                        String                                        | Desmos address of the profile with which the link will be associated | yes |


## Example

```json
{
    "@type": "desmos/MsgLinkChainAccount",
    "chain_address": {
        "@type": "/desmos.profiles.v1beta1.Bech32Address",
        "prefix": "cosmos",
        "value": "cosmos13j7p6faa9jr8ty6lvqv0prldprr6m5xenmafnt"
    },
    "chain_config":{
        "name": "cosmos"
    },
    "proof": {
         "pub_key": {
            "@type": "/cosmos.crypto.secp256k1.PubKey",
            "key": "AjUIjuahImftpkEAKHBsTsGSsc4Eopje+NrRwUYlcBLr"
        },
        "plain_text": "cosmos13j7p6faa9jr8ty6lvqv0prldprr6m5xenmafnt",
        "signature": "c3bd014b2178d63d94b9c28e628bfcf56736de28f352841b0bb27d6fff2968d62c13a10aeddd1ebfe3b13f3f8e61f79a2c63ae6ff5cb78cb0d64e6b0a70fae57",
    },
    "signer": "desmos1qchdngxk8zkl4c4mheqdlpgcegkdrtucmwllpx"
}
```

## Message action
The action associated to this message is the following:

```
link_chain_account
```