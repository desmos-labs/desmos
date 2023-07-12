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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L80-L140
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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L159-L193
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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L206-L227
```

It's expected to fail if:
* the subspace does not exist;
* the post does not exist;
* the signer has no permission to delete posts within the subspace.

## Msg/AddPostAttachment
It's possible to add an attachment to an existing post with `MsgAddPostAttachment`. Attachment can be a [media](02-concepts.md#media) or a [poll](02-concepts.md#poll).

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L233-L262
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
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L281-L310
```

It's expected to fail if:
* the subspace does not exist;
* the post does not exist;
* the post editor is not the post author;
* the post editor has no permission to edit posts within the subspace;
* the attachment does not exist.

## Msg/AnswerPoll
It's possible to answer any active post's poll with `MsgAnswerPoll`.

```js reference 
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L324-L358
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

## Msg/UpdateParams
The `MsgUpdateParams` allows to update the posts module params.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L363-L381
```

It's expected to fail if:
* the authority is not the governance module;
* the params are not valid.

## Msg/MovePost
Posts can be moved from one subspace to another with `MsgMovePost`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L389-L425
```

It's expected to fail if:
* the subspace associated with the post does not exist;
* the post does not exist;
* the target subspace does not exist;
* the target section id does not exist;
* the signer does not have the permission to move the post to the target subspace and section.

## Msg/RequestPostOwnerTransfer
Users can transfer the ownership of their posts to another user with `MsgRequestPostOwnerTransfer`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L438-L469
```

It's expected to fail if:
* the subspace associated with the post does not exist;
* the post does not exist;
* the target user does not have a profile;
* the target user has blocked the request sender.

## Msg/CancelPostOwnerTransfer
Users can cancel a previously requested post ownership transfer with `MsgCancelPostOwnerTransfer`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L477-L503
```

It's expected to fail if:
* the transfer request does not exist.

## Msg/AcceptPostOwnerTransfer
Users can accept a previously requested post ownership transfer with `MsgAcceptPostOwnerTransfer`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L511-L536
```

It's expected to fail if:
* the transfer request does not exist.

## Msg/RefusePostOwnerTransfer
Users can refuse a previously requested post ownership transfer with `MsgRefusePostOwnerTransfer`.

```js reference
https://github.com/desmos-labs/desmos/blob/master/proto/desmos/posts/v3/msgs.proto#L544-L569
```

It's expected to fail if:
* the transfer request does not exist.