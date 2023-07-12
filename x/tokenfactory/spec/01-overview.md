---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---

# `x/tokenfactory`

## Abstract
This document specifies the tokenfactory module of Desmos.

This module allows subspace admins to create new token denominations that can later be used as fee tokens inside a subspace.

This module's implementation is derived from the [Osmosis `x/tokenfactory` module](https://github.com/osmosis-labs/osmosis/tree/main/x/tokenfactory). Instead of forking it, we chose to replicate the implementation. This decision was made because Osmosis is built on a custom version of the Cosmos SDK, which would have caused significant challenges in maintaining the codebase. However, if the module is added to a repository that allows for easier forking, we are open to considering it.

## Contents
1. **[Concepts](02-concepts.md)**
   - [DenomAuthorityMetadata](02-concepts.md#denomauthoritymetadata)
2. **[State](03-state.md)**
   - [DenomAuthority](03-state.md#denomauthority)
   - [DenomAuthorityMetadata](03-state.md#denomauthoritymetadata)
   - [Params](03-state.md#params)
3. **[Msg Service](04-messages.md)**
   - [Msg/CreateDenom](04-messages.md#msgcreatedenom)
   - [Msg/Mint](04-messages.md#msgmint)
   - [Msg/Burn](04-messages.md#msgburn)
   - [Msg/SetDenomMetadata](04-messages.md#msgsetdenommetadata)
   - [Msg/UpdateParams](04-messages.md#msgupdateparams)
4. **[Events](05-events.md)**
   - [Handlers](05-events.md#handlers)
5. **[Params](06-params.md)**
6. **[Permissions](07-permissions.md)**
7. **[Client](08-client.md)**
   - [CLI](08-client.md#cli)
   - [gRPC](08-client.md#grpc)
   - [REST](08-client.md#rest)