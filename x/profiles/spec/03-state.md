---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Profile 
Since a `Profile` extends the Cosmos SDK `AccountI` interface, we store profiles inside the `x/auth` module using the `AccountKeeper`. For this reason, there will only be a single profile for each on-chain account and having multiple profiles require to have different on-chain accounts for each one.

In order to make it possible for users to search a profile based on the DTag, we also store the following reference: 

* DTag: `0x10 | DTag | -> Address`

## DTag Transfer Request
As DTag transfer requests are more important for those who receive them rather for those who send them, we have decided to store them so that they can be searched by the recipient: 

* DTag Transfer Request: `0x11 | Recipient address | Sender address | -> ProtocolBuffer(DTag Transfer Request)`

## Chain Link
To make it possible to query chain links given a user address or given a chain name and an external address, we are using the following keys: 

* Chain Link: `0x12 | User address | Chain name | External address | -> ProtocolBuffer(ChainLink)`
* Chain Link Owner: `0x15 | ChainName | 0x00 | External address | 0x00 | User address | -> 0x01 `

## Default External Address
A chain external address is stored using the owner address and the chain name as the key:

* Default External Address: `0x18 | Owner address | Chain name | -> bytes(ExternalAddress)`

## Application Link
Storing a single application link requires the usage of three different keys to allow for the following queries: 
* application links of a user, given their address;
* application links given an application name and username (reverse search);
* application link given a client id (used during the verification process).

So, we use the following keys: 

* Application Link: `0x13 | User address | Application | Username | -> ProtocolBuffer(ApplicationLink)`
* Application Link Client ID: `0x14 | Client ID | -> ApplicationLinkKey`
* Application Link Owner: `0x16 | Application | 0x00 | Username | 0x00 | User | -> 0x01`

## IBC Port
In order to properly initialize the IBC support for the module, we also store the following data:

* IBC Port: `0x01 | -> Port ID`