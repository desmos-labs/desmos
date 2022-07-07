---
id: messages
title: Messages
sidebar_label: Messages
slug: messages
---

# Msg Service

## Msg/CreatePost
A post can be created using the `MsgCreatePost`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L37-L68
```

The message is expected to fail if any of the following situations occur:
* The post author does not have a profile;
* The subspace associated with the post does not exist;
* The section associated with the post does not exist;
* The post author does not have the permission to create content in the subspace;
* The initial post-ID has not been set for the subspace;
* The post's validation fails.

## Msg/EditPost
A previously created post can be edited with the following `MsgEditPost`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L80-L98
```
The message is expected to fail if any of the following situations occur:
* The subspace associated with the post does not exist;
* The post does not exist;
* The post editor is not the post author;
* The post editor does not have the permission to edit content in the subspace;
* The updated post's validation fails.

## Msg/DeletePost
A post can be deleted with the following `MsgDeletePost`. Deleting a post will also delete all it's related `Attachment`s 
and `Reactions`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L108-L117
```

The message is expected to fail if any of the following situations occur:
* The subspace associated with the post does not exist;
* The post does not exist;
* The signer has no permission to delete the post in the subspace;

## Msg/AddPostAttachment
With `MsgAddPostAttachment` it is possible to add an attachment to a post. Attachment can be a [media](02-concepts.md#media)
or a [poll](02-concepts.md#poll).

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L124-L137
```

The message is expected to fail if any of the following situations occur:
* The subspace associated with the post does not exist;
* The post does not exist;
* The post editor is not the post author;
* The post editor has no permission to edit the post in the subspace;
* The attachment's validation fails;
* The post's validation fails.

## Msg/RemovePostAttachment
A previously added attachment can be removed with `MsgRemovePostAttachment`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L151-L163
```

The message is expected to fail if any of the following situations occur:
* The subspace associated with the post does not exist;
* The post does not exist;
* The post editor is not the post author;
* The post editor has no permission to edit the post in the subspace;
* The attachment does not exist;
* The post's validation fails.

## Msg/AnswerPoll
With `MsgAnswerPoll` it is possible to answer any active post's poll.

```js reference 
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L174-L189
```

The message is expected to fail if any of the following situations occur:
* The signer does not have a profile;
* The subspace associated with the post does not exist;
* The poll's associated post does not exist;
* The signer does not have the permission to interact with content in the subspace;
* The poll does not exist;
* The signer try to edit its own answer but the poll does not allow answers edits;
* The signer try to give multiple answers but the poll does not allow multiple answers;
* The answer given does not correspond to any answer index (the answer does not exist).