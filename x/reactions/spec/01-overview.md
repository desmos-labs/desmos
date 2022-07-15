---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---

# `x/reactions`

## Abstract 
This document specifies the reactions module of Desmos.  

This module gives the possibility to react to contents and customise the reactions experience for you social dApp.

## Contents
1. **[Concepts](02-concepts.md)**
    - [Reaction](02-concepts.md#reaction)
    - [Registered Reaction Value](02-concepts.md#registered-reaction-value)
    - [Free Text Value](02-concepts.md#free-text-value)
    - [Subspace Reactions Params](02-concepts.md#subspace-reactions-params)
    - [Registered Reactions Value Params](02-concepts.md#registered-reaction-value-params)
    - [Free Text Value Params](02-concepts.md#free-text-value-params)
2. **[State](03-state.md)**
    - [Next Registered Reaction ID](03-state.md#next-registered-reaction-id)
    - [Registered Reaction](03-state.md#registered-reaction)
    - [Next Reaction ID](03-state.md#next-reaction-id)
    - [Reaction](03-state.md#reaction)
    - [Reactions Subspace Params](03-state.md#reactions-subspace-params)
3. **[Msg Service](04-messages.md)**
    - [Msg/AddReaction](04-messages.md#msgaddreaction)
    - [Msg/RemoveReaction](04-messages.md#msgremovereaction)
    - [Msg/AddRegisteredReaction](04-messages.md#msgaddregisteredreaction)
    - [Msg/EditRegisteredReaction](04-messages.md#msgeditregisteredreaction)
    - [Msg/RemoveRegisteredReaction](04-messages.md#msgremoveregisteredreaction)
    - [Msg/SetReactionsParams](04-messages.md#msgsetreactionsparams)
4. **[Events](05-events.md)**
    - [Handlers](05-events.md#handlers)
5. **[Permissions](06-permissions.md)** 
6. **[Client](07-client.md)**
    - [CLI](07-client.md#cli)
    - [gRPC](07-client.md#grpc)
    - [REST](07-client.md#rest)