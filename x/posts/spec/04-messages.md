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

The message is expected to fail when the following conditions are matched:
* The `SubspaceID` is equal to 0;
* The `Author` address is invalid;
* The `ReplySettings` are not specified;
* One or more `Entities` are invalid;
* One or more `Attachments` are invalid;
* One or more `PostReferences` are invalid.

## Msg/EditPost
A previously created post can be edited with the following `MsgEditPost`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L80
```
The message is expected to fail when the following conditions are matched:
* The `SubspaceID` is equal to 0;
* The `PostID` is equal to 0;
* One or more `Entities` are invalid;
* The `Editor` address is invalid.

## Msg/DeletePost
A post can be deleted with the following `MsgDeletePost`. Deleting a post will also delete all it's related `Attachment`s 
and `Reactions`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L107
```

The message is expected to fail when the following conditions are matched:
* The `SubspaceID` is equal to 0;
* The `PostID` is equal to 0;
* The `Signer` address is invalid.

## Msg/AddPostAttachment
With `MsgAddPostAttachment` it is possible to add an attachment to a post. Attachment can be a [media](02-concepts.md#media)
or a [poll](02-concepts.md#poll).

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L123
```

The message is expected to fail when the following conditions are matched:
* The `SubspaceID` is equal to 0;
* The `PostID` is equal to 0;
* The `Content` is equal to `nil`;
* The `Editor` address is invalid.

## Msg/RemovePostAttachment
A previously added attachment can be removed with `MsgRemovePostAttachment`.

```js reference
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L149
```

The message is expected to fail when the following conditions are matched:
* The `SubspaceID` is equal to 0;
* The `PostID` is equal to 0;
* The `AttachmentID` is equal to 0;
* The `Editor` address is invalid.

## Msg/AnswerPoll
With `MsgAnswerPoll` it is possible to answer any active post's poll.

```js reference 
https://github.com/desmos-labs/desmos/blob/6787823c96a29241aacfa96e4b0b21f782d059cd/proto/desmos/posts/v1/msgs.proto#L172
```

The message is expected to fail when the following conditions are matched:
* The `SubspaceID` is equal to 0;
* The `PostID` is equal to 0;
* The `PollID` is equal to 0;
* The `AnswerIndexes` array length is equal to 0;
* There are duplicated answers;
* The `signer` address is invalid.