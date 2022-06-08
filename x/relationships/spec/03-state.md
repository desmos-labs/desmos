---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Relationships
Relationships are stored tied to the subspace in which they were created as well as the creator and counterparty addresses:

* `0x01 | Subspace ID | Creator address | Counterparty address | -> ProtocolBuffer(Relationship)`

## User Blocks 
User blocks are stored tied to the subspace for which they were created, the blocker and blocked addresses:

* `0x02 | Subspace ID | Blocker address | Blocked address | -> ProtocolBuffer(UserBlock)`