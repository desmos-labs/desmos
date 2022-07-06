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
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reactions/v1/msgs.proto#L39-L60
```

The message is expected to fail if any of the following situations occur:
* The user does not have a profile;
* The subspace associated with the reaction does not exist;
* The post to react to does not exist;
* The post author blocked the user not allowing the interactions with their posts;
* The user has no permission to react to posts inside the subspace;
* The reaction's value is not valid;
* The reaction already exists;
* The system is unable to retrieve the next reaction ID;
* The reaction's validation fails.

## Msg/RemoveReaction

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reactions/v1/msgs.proto#L73-L94
```

The message is expected to fail if any of the following situations occur:
* The subspace associated with the reaction does not exist;
* The post to react to does not exist;
* The reaction does not exist;
* The user is not equal to the reaction author;
* The user has no permission to remove reactions within the subspace.

## Msg/AddRegisteredReaction

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reactions/v1/msgs.proto#L101-117
```

The message is expected to fail if any of the following situations occur:
* The subspace associated with the registered reaction does not exist;
* The user has no permission to register a reaction within the subspace;
* The system is unable to retrieve the next registered reaction ID;
* The reaction validation fails.


## Msg/EditRegisteredReaction

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reactions/v1/msgs.proto#L131-L153
```

The message is expected to fail if any of the following situations occur:
* The subspace associated with the registered reaction does not exist;
* The registered reaction does not exist;
* The user has no permission to manage registered reactions;
* The reaction validation fails.

## Msg/RemoveRegisteredReaction

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reactions/v1/msgs.proto#L161-L176
```

The message is expected to fail if any of the following situations occur:
* The subspace associated with the registered reaction does not exist;
* The registered reaction does not exist;
* The user has no permission to manage registered reactions.

## Msg/SetReactionsParams

```js reference
https://github.com/desmos-labs/desmos/blob/020cf82788b667924d0f71f9d8f1fd87efa5b340/proto/desmos/reactions/v1/msgs.proto#L184-L205
```

The message is expected to fail if any of the following situations occur:
* The specified subspace does not exist;
* The user has no permission to manage the reactions params;
* The params validation fails.