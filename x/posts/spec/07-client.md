gr---
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

Example Output:
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
The `attachments` query allow users to query all the attachments for the post with the given id inside the subspace with the
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
    '@type': /desmos.posts.v1.Media
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
The `answers` query allows users to query all the answers for a given poll attached to the given post living on the given subspace.
It is also possible to specify an optional user.

```bash
desmos query posts answers [subspace-id] [post-id] [poll-id] [[user]] [flags]
```

Examples:
```bash
desmos query posts answers 1 1 1 --page=2 --limit=100
desmos query posts answers 1 1 1 desmos1mc0mrx23aawryc6gztvdyrupph00yz8lk42v40 --page=2 --limit=100
```

