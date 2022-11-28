---
id: state
title: State
sidebar_label: State
slug: state
---

# State

## Next Registered Reaction ID
The next registered reaction id is stored using the associated subspace id as the key:

- `0x01 | Subspace ID | -> bytes(NextRegisteredReactionID)`

## Registered Reaction
A registered reaction is stored using the associated subspace id and its id as the key. This allows to easily
query:
- all the registered reactions of a subspace;
- a specific registered reaction.

- `0x02 | Subspace ID | Reaction ID | -> ProtocolBuffer(RegisteredReaction)`

## Next Reaction ID
The next reaction id is stored using the associated subspace id as the key:

- `0x10 | Subspace ID | -> bytes(NextReactionID)`

## Reaction
A post reaction is stored using the subspace id, the associated post id and its id combined as the key. This allows to easily query:
- all the reactions of a subspace;
- all the reactions of a post;
- a specific post's reaction.

- `0x11 | Subspace ID | Post ID | Reaction ID | -> ProtocolBuffer(Reaction)`

## Reactions Subspace Params
The reactions' subspace params are stored using the associated subspace id as the key:

- `0x20 | Subspace ID | -> ProtocolBuffer(SubspaceReactionsParams)`