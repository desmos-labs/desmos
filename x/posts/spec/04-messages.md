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
https://github.com/desmos-labs/desmos/blob/v4.1.0/proto/desmos/posts/v2/msgs.proto#L36-L89
```

It's to fail if:
* the post author does not have a profile;
* the subspace does not exist;
* the section does not exist;
* the post author does not have the permission to create content within the subspace;
* the post contents are invalid.

## Msg/EditPost
A previously created post can be edited with the `MsgEditPost`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.1.0/proto/desmos/posts/v2/msgs.proto#L107-L135
```
It's expected to fail if:
* the subspace does not exist;
* the post does not exist;
* the post editor is not the post author;
* the post editor does not have the permission to edit the posts within the subspace;
* the updated post contents are invalid.

## Msg/DeletePost
A post can be deleted with the following `MsgDeletePost`. Deleting a post will also delete all it's related `Attachment`s 
and `Reactions`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.1.0/proto/desmos/posts/v2/msgs.proto#L147-L163
```

The message is expected to fail if any of the following situations occur:
* The subspace associated with the post does not exist;
* The post does not exist;
* The signer has no permission to delete the post in the subspace;

## Msg/AddPostAttachment
With `MsgAddPostAttachment` it is possible to add an attachment to a post. Attachment can be a [media](02-concepts.md#media)
or a [poll](02-concepts.md#poll).

```js reference
https://github.com/desmos-labs/desmos/blob/v4.1.0/proto/desmos/posts/v2/msgs.proto#L170-L191
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
https://github.com/desmos-labs/desmos/blob/v4.1.0/proto/desmos/posts/v2/msgs.proto#L209-L232
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
https://github.com/desmos-labs/desmos/blob/v4.1.0/proto/desmos/posts/v2/msgs.proto#L245-271
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