---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---

# `x/posts`

## Abstract 
This document specifies the posts module of Desmos.  

This module allows the creation and management of on-chain contents in the form of posts that can be enriched with different
kind of data.

## Contents
1. **[Concepts](02-concepts.md)**
    - [Post](02-concepts.md#post)
    - [Attchment](02-concepts.md#attachment)
    - [Media](02-concepts.md#media)
    - [Poll](02-concepts.md#poll)
2. **[State](03-state.md)**
    - [Next Post ID](03-state.md#next-post-id)
    - [Post](03-state.md#post)
    - [Next Attachment ID](03-state.md#post-section)
    - [Attachment](03-state.md#attachment)
    - [User Answer](03-state.md#user-answer)
    - [Active Poll Queue](03-state.md#active-poll-queue)
3. **[Msg Service](04-messages.md)**
    - [Msg/CreatePost](04-messages.md#msgcreatepost)
    - [Msg/EditPost](04-messages.md#msgeditpost)
    - [Msg/DeletePost](04-messages.md#msgdeletepost)
    - [Msg/AddPostAttachment](04-messages.md#msgaddpostattachment)
    - [Msg/RemovePostAttachment](04-messages.md#msgremovepostattachment)
    - [Msg/AnswerPoll](04-messages.md#msganswerpoll)
4. **[Events](05-events.md)**
    - [Handlers](05-events.md#handlers)
    - [Keeper](05-events.md#keeper)
5. **[Permissions](06-permissions.md)**
6. **[Parameters](07-params.md)**
7. **[Client](08-client.md)**
   - [CLI](08-client.md#cli)
   - [gRPC](08-client.md#grpc)
   - [REST](08-client.md#rest)
