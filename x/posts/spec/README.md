<!--
order: 0
title: Profiles Overview
parent:
  title: "profiles"
-->

# `x/profiles`

## Abstract 
This document specifies the profiles module of Desmos.  

This module allows the creation and management of an on-chain social profile that can be connected to external chains and services.

## Contents
1. **[Concepts](01_concepts.md)**
    - [Profile](01_concepts.md#profile)
    - [DTag Transfer Request](01_concepts.md#dtag-transfer-request)
    - [Chain Link](01_concepts.md#chain-link)
    - [Application Link](01_concepts.md#application-link)
2. **[State](02_state.md)**
    - [Profile](02_state.md#profile)
    - [DTag Transfer Request](02_state.md#dtag-transfer-request)
    - [Chain Link](02_state.md#chain-link)
    - [Application Link](02_state.md#application-link)
    - [IBC Port](02_state.md#ibc-port)
3. **[Msg Service](03_messages.md)**
    - [Msg/SaveProfile](03_messages.md#msgsaveprofile)
    - [Msg/DeleteProfile](03_messages.md#msgdeleteprofile)
    - [Msg/RequestDTagTransfer](03_messages.md#msgrequestdtagtransfer)
    - [Msg/CancelDTagTransferRequest](03_messages.md#msgcanceldtagtransferrequest)
    - [Msg/AcceptDTagTransferRequest](03_messages.md#msgacceptdtagtransferrequest)
    - [Msg/RefuseDTagTransferRequest](03_messages.md#msgrefusedtagtransferrequest)
    - [Msg/LinkChainAccount](03_messages.md#msglinkchainaccount)
    - [Msg/UnlinkChainAccount](03_messages.md#msgunlinkchainaccount)
    - [Msg/LinkApplication](03_messages.md#msglinkapplication)
    - [Msg/UnlinkApplication](03_messages.md#msgunlinkapplication)
4. **[Events](04_events.md)**
    - [Handlers](04_events.md#handlers) 
    - [Keeper](04_events.md#keeper)
5. **[Parameters](05_params.md)**
6. **[Client](06_client.md)**
    - [CLI](06_client.md#cli)
    - [gRPC](06_client.md#grpc)
    - [REST](06_client.md#rest)