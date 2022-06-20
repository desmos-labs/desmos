---
id: messages
title: Messages
sidebar_label: Messages
slug: messages
---

# Msg Service

## Msg/CreatePost
A post can be created using the `MsgCreatePost`, specifying the targets `Subspace ID` and `Section ID` and the information
needed to compose a post showed below.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L37
```

## Msg/EditPost
A previously created post can be edited with the following `MsgEditPost`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L80
```

## Msg/DeletePost
A post can be deleted with the following `MsgDeletePost`. Deleting a post will also delete all it's related `Attachment`s 
and `Reactions`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L107
```

## Msg/AddPostAttachment
With `MsgAddPostAttachment` it is possible to add an attachment to a post. Attachment can be a [media](02-concepts.md#media)
or a [poll](02-concepts.md#poll).

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L123
```

## Msg/RemovePostAttachment
A previously added attachment can be removed with `MsgRemovePostAttachment`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L149
```

## Msg/AnswerPoll
With `MsgAnswerPoll` it is possible to answer any active post's poll.

```js reference 
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L172
```