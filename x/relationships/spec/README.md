<!--
order: 0
title: Relationships Overview
parent:
  title: "relationships"
-->

# `relationships`

This document specifies the Relationships module of Desmos. 

The module allows Desmos users to establish relationships between them, and also to block misbehaving users from harassing them in the future.  

## Contents
1. **[Concepts](01_concepts.md)**
   - [Relationship](01_concepts.md#relationship)
   - [User Block](01_concepts.md#user-block)
2. **[State](02_state.md)**
   - [Relationship](02_state.md#relationships)
   - [User Blocks](02_state.md#user-blocks)
3. **[Msg Server](03_messages.md)**
   - [Msg/CreateRelationship](03_messages.md#msgcreaterelationship)
   - [Msg/DeleteRelationship](03_messages.md#msgdeleterelationship)
   - [Msg/BlockUser](03_messages.md#msgblockuser)
   - [Msg/UnblockUser](03_messages.md#msgunblockuser)
4. **[Events](04_events.md)**
   - [Handlers](04_events.md#handlers)
5. **[Client](05_client.md)**
   - [CLI](05_client.md#cli)
   - [gRPC](05_client.md#grpc)
   - [REST](05_client.md#rest)