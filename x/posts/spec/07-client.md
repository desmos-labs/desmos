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
The `post` query command allows users to query a post with the given id inside a given subspace

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

