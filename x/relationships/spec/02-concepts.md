---
id: concepts
title: Concepts
sidebar_label: Concepts
slug: concepts
---

# Concepts

## Relationship
A relationships between two users is a **mono-directional** link between a _creator_ and a _recipient_. This is equivalent to the concept of "follow" that is present on traditional social networks (i.e. Twitter). 

When a user A creates a relationships towards a user B, it means that they are interested in being notified about what user B does. 

_Friendship_ can be represented by a mutual relationship, which consists of two mono-directional relationships. If user A creates a relationships towards user B (`A -> B`), and user B creates a relationship towards user A (`B -> A`), then user A and B can be considered to be _friends_ since a mutual relationship (`A <-> B`) exists.

## User Block
A user block from one user (_blocker_) to another (_blocked_) represents the willingness of the first to block any future interaction that the latter might have with them. This concept is used to allow users to block misbehaving users from future harassment or unwanted interactions. 

When a user A creates a user block towards a user B, they can specify inside which subspace they want to block the user. If no particular subspace is provided, this means the B will not be allowed to have Desmos-level interactions with A in the future (e.g. requesting A to exchange their DTag). Blocking a user on subspace with id `0` (default value) **will not** block such user from interacting inside other subspaces.
