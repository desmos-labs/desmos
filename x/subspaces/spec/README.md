<!--
order: 0
title: Subspaces Overview
parent:
  title: "subspaces"
-->

# `x/subspaces`

## Abstract 
This document specifies the subspaces module of Desmos.  

This module allows users to create and manage the representation of different social networks inside which contents will be created.

## Contents
1. **[Concepts](01_concepts.md)**
    - [Subspace](01_concepts.md#subspace)
    - [Section](01_concepts.md#section)
    - [UserGroup](01_concepts.md#user-group)
    - [Permissions](01_concepts.md#permissions)
2. **[State](02_state.md)**
    - [Subspace](02_state.md#subsapce)
    - [Section](02_state.md#section)
    - [UserGroup](02_state.md#user-group)
    - [Permissions](02_state.md#permissions)
3. **[Msg Service](03_messages.md)**
    - [Msg/CreateSubspace](03_messages.md#msgcreatesubspace)
    - [Msg/EditSubspace](03_messages.md#msgeditsubspace)
    - [Msg/DeleteSubspace](03_messages.md#msgdeletesubspace)
    - [Msg/CreateSection](03_messages.md#msgcreatesection)
    - [Msg/EditSection](03_messages.md#msgeditsection)
    - [Msg/MoveSection](03_messages.md#msgmovesection)
    - [Msg/DeleteSection](03_messages.md#msgdeletesection)
    - [Msg/CreateUserGroup](03_messages.md#msgcreateusergroup)
    - [Msg/EditUserGroup](03_messages.md#msgeditusergroup)
    - [Msg/MoveUserGroup](03_messages.md#msgmoveusergroup)
    - [Msg/SetUserGroupPermissions](03_messages.md#msgsetusergrouppermissions)
    - [Msg/DeleteUserGroup](03_messages.md#msgdeleteusergroup)
    - [Msg/AddUserToUserGroup](03_messages.md#msgaddusertousergroup)
    - [Msg/RemoveUserFromUserGroup](03_messages.md#msgremoveuserfromusergroup)
    - [Msg/SetUserPermissions](03_messages.md#msgsetuserpermissions)
4. **[Events](04_events.md)**
    - [Handlers](04_events.md#handlers) 
5. **[Client](05_client.md)**
    - [CLI](05_client.md#cli)
    - [gRPC](05_client.md#grpc)
    - [REST](05_client.md#rest)