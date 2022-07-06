---
id: overview
title: Overview
sidebar_label: Overview
slug: overview
---

# `x/subspaces`

## Abstract 
This document specifies the subspaces module of Desmos.  

This module allows users to create and manage the representation of different social networks inside which contents will be created.

## Contents
1. **[Concepts](02-concepts.md)**
    - [Subspace](02-concepts.md#subspace)
    - [Section](02-concepts.md#section)
    - [User Group](02-concepts.md#user-group)
    - [User Permission](02-concepts.md#user-permission)
2. **[State](03-state.md)**
    - [Next Subspace ID](03-state.md#next-subspace-id)
    - [Subspace](03-state.md#subspace)
    - [Next Section ID](03-state.md#next-section-id)
    - [Section](03-state.md#section)
    - [Next Group ID](03-state.md#next-group-id)
    - [User Group](03-state.md#user-group)
    - [User Group Member](03-state.md#user-group-member)
    - [User Permission](03-state.md#user-permission)
3. **[Msg Service](04-messages.md)**
    - [Msg/CreateSubspace](04-messages.md#msgcreatesubspace)
    - [Msg/EditSubspace](04-messages.md#msgeditsubspace)
    - [Msg/DeleteSubspace](04-messages.md#msgdeletesubspace)
    - [Msg/CreateSection](04-messages.md#msgcreatesection)
    - [Msg/EditSection](04-messages.md#msgeditsection)
    - [Msg/MoveSection](04-messages.md#msgmovesection)
    - [Msg/DeleteSection](04-messages.md#msgdeletesection)
    - [Msg/CreateUserGroup](04-messages.md#msgcreateusergroup)
    - [Msg/EditUserGroup](04-messages.md#msgeditusergroup)
    - [Msg/MoveUserGroup](04-messages.md#msgmoveusergroup)
    - [Msg/SetUserGroupPermissions](04-messages.md#msgsetusergrouppermissions)
    - [Msg/DeleteUserGroup](04-messages.md#msgdeleteusergroup)
    - [Msg/AddUserToUserGroup](04-messages.md#msgaddusertousergroup)
    - [Msg/RemoveUserFromUserGroup](04-messages.md#msgremoveuserfromusergroup)
    - [Msg/SetUserPermissions](04-messages.md#msgsetuserpermissions)
4. **[Events](05-events.md)**
    - [Handlers](05-events.md#handlers) 
5. **[Permissions](06-permissions.md)**
6. **[Client](07-client.md)**
    - [CLI](07-client.md#cli)
    - [gRPC](07-client.md#grpc)
    - [REST](07-client.md#rest)