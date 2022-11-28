---
id: messages
title: Messages
sidebar_label: Messages
slug: messages
---

# Msg Service

## Msg/AddReaction
A post reaction can be added with the `MsgAddReaction`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/reactions/v1/msgs.proto#L38-L60
```

It's expected to fail if:
* the user does not have a profile;
* the subspace does not exist;
* the post does not exist;
* the post author has blocked the user within the subspace;
* the user has no permission to react to posts inside the subspace;
* the reaction's value is not valid;
* the reaction already exists.

## Msg/RemoveReaction
A reaction can be removed with the `MsgRemoveReaction`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/reactions/v1/msgs.proto#L71-L94
```

It's expected to fail if:
* the subspace does not exist;
* the post does not exist;
* the reaction does not exist;
* the user has no permission to remove reactions within the subspace.

## Msg/AddRegisteredReaction
A registered reaction can be added to a subspace with the `MsgAddRegisteredReaction`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/reactions/v1/msgs.proto#L99-L117
```

It's expected to fail if:
* the subspace does not exist;
* the user has no permission to register a reaction within the subspace;
* the provided shorthand code is either blank or empty; 
* the provided display value is either blank or empty.

## Msg/EditRegisteredReaction
A registered reaction can be edited with the `MsgEditRegisteredReaction`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/reactions/v1/msgs.proto#L129-L153
```

it's expected to fail if:
* the subspace does not exist;
* the registered reaction does not exist;
* the user has no permission to manage registered reactions;
* the new shorthand code or display value are invalid.

## Msg/RemoveRegisteredReaction
A registered reaction ca be removed with the `MsgRemoveRegisteredReaction`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/reactions/v1/msgs.proto#L159-L176
```

It's expected to fail if:
* the subspace does not exist;
* the registered reaction does not exist;
* the user has no permission to manage registered reactions.

## Msg/SetReactionsParams
A subspace's reaction params can be set with the `MsgSetReactionsParams`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/reactions/v1/msgs.proto#L182-L205
```

It's expected to fail if:
* the specified subspace does not exist;
* the user has no permission to manage the reactions params;
* the provided params are invalid.