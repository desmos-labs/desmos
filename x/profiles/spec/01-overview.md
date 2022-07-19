---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---

# `x/profiles`

## Abstract 
This document specifies the profiles module of Desmos.  

This module allows the creation and management of an on-chain social profile that can be connected to external chains and services.

## Contents
1. **[Concepts](02-concepts.md)**
    - [Profile](02-concepts.md#profile)
    - [DTag Transfer Request](02-concepts.md#dtag-transfer-request)
    - [Chain Link](02-concepts.md#chain-link)
    - [Default External Address](02-concepts.md#default-external-address)
    - [Application Link](02-concepts.md#application-link)
2. **[State](03-state.md)**
    - [Profile](03-state.md#profile)
    - [DTag Transfer Request](03-state.md#dtag-transfer-request)
    - [Chain Link](03-state.md#chain-link)
    - [Default External Address](03-state.md#default-external-address)
    - [Application Link](03-state.md#application-link)
    - [IBC Port](03-state.md#ibc-port)
3. **[Msg Service](04-messages.md)**
    - [Msg/SaveProfile](04-messages.md#msgsaveprofile)
    - [Msg/DeleteProfile](04-messages.md#msgdeleteprofile)
    - [Msg/RequestDTagTransfer](04-messages.md#msgrequestdtagtransfer)
    - [Msg/CancelDTagTransferRequest](04-messages.md#msgcanceldtagtransferrequest)
    - [Msg/AcceptDTagTransferRequest](04-messages.md#msgacceptdtagtransferrequest)
    - [Msg/RefuseDTagTransferRequest](04-messages.md#msgrefusedtagtransferrequest)
    - [Msg/LinkChainAccount](04-messages.md#msglinkchainaccount)
    - [Msg/UnlinkChainAccount](04-messages.md#msgunlinkchainaccount)
    - [Msg/SetDefaultExternalAddress](04-messages.md#msgsetdefaultexternaladdress)
    - [Msg/LinkApplication](04-messages.md#msglinkapplication)
    - [Msg/UnlinkApplication](04-messages.md#msgunlinkapplication)
4. **[Events](05-events.md)**
    - [Handlers](05-events.md#handlers) 
    - [Keeper](05-events.md#keeper)
5. **[Parameters](06-params.md)**
6. **[Client](07-client.md)**
    - [CLI](07-client.md#cli)
    - [gRPC](07-client.md#grpc)
    - [REST](07-client.md#rest)