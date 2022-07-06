---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---

# `x/reports`

## Abstract 
This document specifies the reports module of Desmos.  

This module handles the creation and management of a reporting system toward content and users.

## Contents
1. **[Concepts](02-concepts.md)**
    - [Report](02-concepts.md#report)
    - [User Target](02-concepts.md#user-target)
    - [Post Target](02-concepts.md#post-target)
    - [Reason](02-concepts.md#reason)
2. **[State](03-state.md)**
    - [Next Report ID](03-state.md#next-report-id)
    - [Report](03-state.md#report)
    - [Post Report](03-state.md#posts-report)
    - [User Report](03-state.md#user-report)
    - [Next Reason ID](03-state.md#next-reason-id)
    - [Reason](03-state.md#reason)
3. **[Msg Service](04-messages.md)**
    - [Msg/CreateReport](04-messages.md#msgcreatereport)
    - [Msg/DeleteReport](04-messages.md#msgdeletereport)
    - [Msg/SupportStandardReason](04-messages.md#msgsupportstandardreason)
    - [Msg/AddReason](04-messages.md#msgaddreason)
    - [Msg/RemoveReason](04-messages.md#msgremovereason)
4. **[Events](05-events.md)**
    - [Handlers](05-events.md#handlers)
5. **[Permissions](06-permissions.md)**
6. **[Parameters](07-params.md)**
7. **[Client](08-client.md)**
    - [CLI](08-client.md#cli)
    - [gRPC](08-client.md#grpc)
    - [REST](08-client.md#rest)