---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client

## CLI

A user can query and interact with the `posts` module using the CLI.

### Query

The `query` commands allow users to query the `posts` state.

```
desmos query posts --help
```

#### post
The `post` query command allows users to query a post with the given id inside a subspace with the given id.

```bash
desmos query posts post [subspace-id] [post-id] [flags]
```

Example:
```bash
desmos query posts post 1 1
```

Example output:
```yaml
post:
  author: desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
  conversation_id: "0"
  creation_date: "2022-06-20T15:13:10.751262Z"
  entities:
    hashtags: []
    mentions: []
    urls:
    - display_url: This
      end: "3"
      start: "0"
      url: https://example.com
  external_id: This is my external id
  id: "1"
  last_edited_date: null
  referenced_posts: []
  reply_settings: REPLY_SETTING_EVERYONE
  section_id: 1
  subspace_id: "1"
  text: This is my post text
```

#### posts
The `posts` query command allows users to query all the posts for a given subspace id optionally specifying also a section id.

```bash
desmos query posts posts [subspace-id] [[section-id]] [flags]
```

Examples:
```bash
desmos query posts posts 1 --page=1 --limit=100
desmos query posts posts 1 1 --page=1 --limit=100
```

Example output:
```bash
pagination:
  next_key: null
  total: "0"
posts:
- author: desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
  conversation_id: "0"
  creation_date: "2022-06-20T15:13:10.751262Z"
  entities:
    hashtags: []
    mentions: []
    urls:
    - display_url: This
      end: "3"
      start: "0"
      url: https://example.com
  external_id: This is my external id
  id: "1"
  last_edited_date: null
  referenced_posts: []
  reply_settings: REPLY_SETTING_EVERYONE
  section_id: 1
  subspace_id: "1"
  text: This is my post text
- author: desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
  conversation_id: "0"
  creation_date: "2022-06-21T09:19:12.343428Z"
  entities:
    hashtags: []
    mentions: []
    urls:
    - display_url: This
      end: "3"
      start: "0"
      url: https://example.com
  external_id: This is my external id
  id: "2"
  last_edited_date: null
  referenced_posts: []
  reply_settings: REPLY_SETTING_EVERYONE
  section_id: 1
  subspace_id: "1"
  text: This is my second post text
```

#### attachments
The `attachments` query command allow users to query all the attachments for the post with the given id inside the subspace with the
given id.

```bash
desmos query posts attachments [subspace-id] [post-id] [flags]
```

Example:
```bash
desmos query posts attachments 1 1 --page=1 --limit=100
```

Example output: 
```bash
attachments:
- content:
    '@type': /desmos.posts.v2.media
    mime_type: image/png
    uri: ftp://user:password@example.com/image.png
  id: 1
  post_id: "1"
  section_id: 0
  subspace_id: "1"
pagination:
  next_key: null
  total: "0"
```

#### answers
The `answers` query command allows users to query all the answers for a given poll attached to the given post living on the given subspace.
It is also possible to specify an optional user.

```bash
desmos query posts answers [subspace-id] [post-id] [poll-id] [[user]] [flags]
```

Examples:
```bash
desmos query posts answers 1 1 1
desmos query posts answers 1 1 1 desmos1mc0mrx23aawryc6gztvdyrupph00yz8lk42v40 --page=1 --limit=100
```

Examples output:
```bash
answers:
- answers_indexes:
  - 0
  - 1
  poll_id: 1
  post_id: "1"
  section_id: 0
  subspace_id: "1"
  user: desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3
pagination:
  next_key: null
  total: "0"
```

#### params
The `params` query command allows users to get the currently set parameters. 

```bash 
desmos query posts params [flags]
```

Examples:
```bash
desmos query posts params
```

Example output:
```bash
params:
  max_text_length: 500
```

## gRPC
A user can query the `posts` module gRPC endpoints. 

### Post
The `Post` endpoint allows users to query a post with the given id inside a subspace with the given id.

```bash
desmos.posts.v2.Query/Post
```

Example:
```bash
grpcurl -plaintext \
-d '{"subspace_id":1, "post_id":1}' localhost:9090 desmos.posts.v2.Query/Post
```

Example output:
```json
{
  "post": {
    "subspaceId": "1",
    "sectionId": 1,
    "id": "1",
    "externalId": "This is my external id",
    "text": "This is my post text",
    "entities": {
      "urls": [
        {
          "end": "3",
          "url": "https://example.com",
          "displayUrl": "This"
        }
      ]
    },
    "author": "desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
    "replySettings": "REPLY_SETTING_EVERYONE",
    "creationDate": "2022-06-20T15:13:10.751262Z",
    "lastEditedDate": "2022-06-21T15:04:05.722967Z"
  }
}
```

### SubspacePosts
The `SubspacePosts` endpoint allows users to query all the posts of the subspace with the given id.

```bash
desmos.posts.v2.Query/SubspacePosts
```

Example:
```bash
grpcurl -plaintext \
-d '{"subspace_id":1}' localhost:9090 desmos.posts.v2.Query/SubspacePosts
```

Example output:
```json
{
  "posts": [
    {
      "subspaceId": "1",
      "sectionId": 1,
      "id": "1",
      "externalId": "This is my external id",
      "text": "This is my post text",
      "entities": {
        "urls": [
          {
            "end": "3",
            "url": "https://example.com",
            "displayUrl": "This"
          }
        ]
      },
      "author": "desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
      "replySettings": "REPLY_SETTING_EVERYONE",
      "creationDate": "2022-06-20T15:13:10.751262Z",
      "lastEditedDate": "2022-06-21T15:04:05.722967Z"
    },
    {
      "subspaceId": "1",
      "sectionId": 1,
      "id": "2",
      "externalId": "This is my external id",
      "text": "This is my second post text",
      "entities": {
        "urls": [
          {
            "end": "3",
            "url": "https://example.com",
            "displayUrl": "This"
          }
        ]
      },
      "author": "desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
      "replySettings": "REPLY_SETTING_EVERYONE",
      "creationDate": "2022-06-21T09:19:12.343428Z"
    }
  ],
  "pagination": {
    "total": "2"
  }
}
```


### SectionPosts
The `SectionPosts` endpoint allows users to return all the posts associated with the section with the given id.

```bash
desmos.posts.v2.Query/SectionPosts
```

Example:
```bash
grpcurl -plaintext \
-d '{"subspace_id":1, "section_id":1}' localhost:9090 desmos.posts.v2.Query/SectionPosts
```

Example output:
```json
{
  "posts": [
    {
      "subspaceId": "1",
      "sectionId": 1,
      "id": "1",
      "externalId": "This is my external id",
      "text": "This is my post text",
      "entities": {
        "urls": [
          {
            "end": "3",
            "url": "https://example.com",
            "displayUrl": "This"
          }
        ]
      },
      "author": "desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
      "replySettings": "REPLY_SETTING_EVERYONE",
      "creationDate": "2022-06-20T15:13:10.751262Z",
      "lastEditedDate": "2022-06-21T15:04:05.722967Z"
    },
    {
      "subspaceId": "1",
      "sectionId": 1,
      "id": "2",
      "externalId": "This is my external id",
      "text": "This is my second post text",
      "entities": {
        "urls": [
          {
            "end": "3",
            "url": "https://example.com",
            "displayUrl": "This"
          }
        ]
      },
      "author": "desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3",
      "replySettings": "REPLY_SETTING_EVERYONE",
      "creationDate": "2022-06-21T09:19:12.343428Z"
    }
  ],
  "pagination": {
    "total": "2"
  }
}
```

### PostAttachments
The `PostAttachments` endpoint allows users to query all the attachment associated with the post id given.

```bash
desmos.posts.v2.Query/PostAttachments
```

Example:
```bash
grpcurl -plaintext \
-d '{"subspace_id":1, "post_id":1}' localhost:9090 desmos.posts.v2.Query/PostAttachments
```

Example output:
```json
{
  "attachments": [
    {
      "subspaceId": "1",
      "postId": "1",
      "id": 1,
      "content": {"@type":"/desmos.posts.v2.media","mimeType":"image/png","uri":"ftp://user:password@example.com/image.png"}
    },
    {
      "subspaceId": "1",
      "postId": "1",
      "id": 2,
      "content": {"@type":"/desmos.posts.v2.poll","allowsAnswerEdits":true,"allowsMultipleAnswers":true,"endDate":"2025-01-01T12:00:00Z","providedAnswers":[{"text":"yes"},{"text":"no"}],"question":"A question"}
    }
  ],
  "pagination": {
    "total": "2"
  }
}

```

### PollAnswers
The `PollAnswers` endpoint allows users to query al the poll answer associated with the given poll id attached to the post
with the given post id.

```bash
desmos.posts.v2.Query/PollAnswers
```

Examples:
```bash
grpcurl -plaintext \
-d '{"subspace_id":1, "post_id":1, "poll_id":2}' localhost:9090 desmos.posts.v2.Query/PollAnswers
grpcurl -plaintext \
-d '{"subspace_id":1, "post_id":1, "poll_id":2, "user":"desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3"}' localhost:9090 desmos.posts.v2.Query/PollAnswers
```

Example output:
```json
{
  "answers": [
    {
      "subspaceId": "1",
      "postId": "1",
      "pollId": 2,
      "answersIndexes": [
        0,
        1
      ],
      "user": "desmos1rfv0f7mx7w9d3jv3h803u38vqym9ygg344asm3"
    }
  ],
  "pagination": {
    "total": "1"
  }
}

```

### Params
The `Params` endpoint allows users to query the module's parameters.

```bash
desmos.posts.v2.Query/Params
```

Example:
```bash
grpcurl -plaintext localhost:9090 desmos.posts.v2.Query/Params
```

Example output:
```json
{
  "params": {
    "maxTextLength": 500
  }
}
```

## REST
A user can query the `posts` module using REST endpoints.

### Post
The `Post` endpoint allows users to query a post with the given id inside a subspace with the given id.

```
/desmos/posts/v2/subspaces/{subspace_id}/posts/{post_id}
```

### SubspacePosts
The `SubspacePosts` endpoint allows users to query all the posts of the subspace with the given id.

```
/desmos/posts/v2/subspaces/{subspace_id}/posts
```

### SectionPosts
The `SectionPosts` endpoint allows users to return all the posts associated with the section with the given id associated
to the subspace with the given id.

```
/desmos/posts/v2/subspaces/{subspace_id}/sections/{section_id}/posts
```

### PostAttachments
The `PostAttachments` endpoint allows users to query all the attachment associated with the post id given living inside
the subspace with the given id.

```
/desmos/posts/v2/subspaces/{subspace_id}/posts/{post_id}/attachments
```

### PollAnswers
The `PollAnswers` endpoint allows users to query al the poll answer associated with the given poll id attached to the post
with the given post id inside the subspace with the given id.

```
/desmos/posts/v2/subspaces/{subspace_id}/posts/{post_id}/polls/{poll_id}/answers
```

### Params
The `Params` endpoint allows users to query the module's parameters.

```
/desmos/posts/v2/params
```