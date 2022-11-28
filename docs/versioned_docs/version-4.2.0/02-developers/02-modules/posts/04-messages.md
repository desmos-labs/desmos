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
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/posts/v2/msgs.proto#L36-L89
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
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/posts/v2/msgs.proto#L107-L135
```
It's expected to fail if:
* the subspace does not exist;
* the post does not exist;
* the post editor is not the post author;
* the post editor does not have the permission to edit the posts within the subspace;
* the updated post contents are invalid.

## Msg/DeletePost
A post can be deleted with the `MsgDeletePost`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/posts/v2/msgs.proto#L147-L163
```

It's expected to fail if:
* the subspace does not exist;
* the post does not exist;
* the signer has no permission to delete posts within the subspace.

## Msg/AddPostAttachment
It's possible to add an attachment to an existing post with `MsgAddPostAttachment`. Attachment can be a [media](02-concepts.md#media) or a [poll](02-concepts.md#poll).

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/posts/v2/msgs.proto#L168-L191
```

It's expected to fail if:
* the subspace does not exist;
* the post does not exist;
* the post editor is not the post author;
* the post editor has no permission to edit posts within the subspace;
* the attachment is invalid.

## Msg/RemovePostAttachment
A post attachment can be removed with `MsgRemovePostAttachment`.

```js reference
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/posts/v2/msgs.proto#L209-L232
```

It's expected to fail if:
* the subspace does not exist;
* the post does not exist;
* the post editor is not the post author;
* the post editor has no permission to edit posts within the subspace;
* the attachment does not exist.

## Msg/AnswerPoll
It's possible to answer any active post's poll With `MsgAnswerPoll`.

```js reference 
https://github.com/desmos-labs/desmos/blob/v4.2.0/proto/desmos/posts/v2/msgs.proto#L245-L271
```

It's expected to fail if:
* the signer does not have a profile;
* the subspace associated with the post does not exist;
* the poll does not exist;
* the poll voting period already ended;
* the signer does not have the permission to interact with contents within the subspace;
* the signer is trying to edit their own answer but the poll does not allow answers edits;
* the signer is trying to give multiple answers but the poll does not allow multiple answers;
* the answer is invalid.