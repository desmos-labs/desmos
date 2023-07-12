---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## DenomAuthorityMetadata
The authority metadata of various denominations are stored on the chain as follows:

* `denoms | <Denom> | authoritymetadata -> ProtocolBuffer(DenomAuthorityMetadata)`

## Creator
The creator of a denomination is stored on the chain as follows:

* `creator | <Address> | <Denom> -> <Denom>`

## Params
The params of the module are stored on the chain as follows:

* `params -> ProtocolBuffer(Params)`

