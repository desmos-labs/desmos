---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---

# `x/relationships`

This document specifies the Relationships module of Desmos. 

The module allows Desmos users to establish relationships between them, and also to block misbehaving users from harassing them in the future.  

## Contents
1. **[Concepts](02-concepts.md)**
   - [Relationship](02-concepts.md#relationship)
   - [User Block](02-concepts.md#user-block)
2. **[State](03-state.md)**
   - [Relationship](03-state.md#relationships)
   - [User Blocks](03-state.md#user-blocks)
3. **[Msg Server](04-messages.md)**
   - [Msg/CreateRelationship](04-messages.md#msgcreaterelationship)
   - [Msg/DeleteRelationship](04-messages.md#msgdeleterelationship)
   - [Msg/BlockUser](04-messages.md#msgblockuser)
   - [Msg/UnblockUser](04-messages.md#msgunblockuser)
4. **[Events](05-events.md)**
   - [Handlers](05-events.md#handlers)
5. **[Client](06-client.md)**
   - [CLI](06-client.md#cli)
   - [gRPC](06-client.md#grpc)
   - [REST](06-client.md#rest)