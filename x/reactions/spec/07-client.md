---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client

## CLI

A user can query and interact with the `reactions` module using the CLI.

### Query

The `query` commands allow users to query the `reactions` state.

```
desmos query reactions --help
```

#### reaction
The `reaction` query command allows users to query a specific reaction with the given ID to a post with the given ID, inside the subspace with
the given ID.

```bash
desmos query reaction reaction [subspace-id] [post-id] [reaction-id] [flags]
```

Example: 
```bash
desmos query reactions reaction 5 1 1
```

Example output: 
```yaml
reaction:
  author: desmos159axlj0mkvch02f95t5tkghychyeueaslk6r8f
  id: 1
  post_id: "1"
  subspace_id: "5"
  value:
    '@type': /desmos.reactions.v1.FreeTextValue
    text: "\U0001F680"
```

#### reactions
The `reactions` query command allows users to query all the reactions inside the subspace with the given ID. It's also possible
to get all the reactions made to a post by specifying also its ID. 

```bash
desmos query reactions reactions [subspace-id] [[post-id]] [flags]
```

Example:
```bash
desmos query reactions reactions 5 1
```

Example output:
```yaml
pagination:
  next_key: null
  total: "0"
reactions:
- author: desmos159axlj0mkvch02f95t5tkghychyeueaslk6r8f
  id: 1
  post_id: "1"
  subspace_id: "5"
  value:
    '@type': /desmos.reactions.v1.FreeTextValue
    text: "\U0001F680"
- author: desmos1dx6h75tkj0cuvyqf6cwn6usc9qynu39v0245m4
  id: 2
  post_id: "1"
  subspace_id: "5"
  value:
    '@type': /desmos.reactions.v1.FreeTextValue
    text: "\U0001F602"
```

#### registered-reaction
The `registered-reaction` query command allows users to query the registered-reaction with the given ID inside the subspace
with the given ID.

```bash
desmos query reactions registered-reaction [subspace-id] [reaction-id] [flags]
```

Example:
```bash
desmos query reactions registered-reaction 5 1
```

Example output:
```yaml
registered_reaction:
  display_value: https://example.com?image=hello.png
  id: 7
  shorthand_code: :hello
  subspace_id: "5"
```

#### registered-reactions
The `registered-reactions` query command allows users to query all the registered-reactions inside the subspace with the given ID.

```bash
desmos query reactions registered-reactions [subspace-id] [flags]
```

Example:
```bash
desmos query reactions registered-reactions 5
```

Example output:
```yaml
pagination:
  next_key: null
  total: "0"
registered_reactions:
- display_value: https://example.com?image=hello.png
  id: 7
  shorthand_code: ':hello:' 
  subspace_id: "5"
- display_value: https://example.com?image=bye.png
  id: 8
  shorthand_code: ':bye:'
  subspace_id: "5"
```

### params
The `params` query command allows users to query the reactions parameters for the subspace with the given ID. 

```bash
desmos query reactions params [subspace-id] [flags]
```

Example:
```bash
desmos query reactions params 5
```

Example output:
```yaml
params:
  free_text:
    enabled: true
    max_length: 30
    reg_ex: '[a-z]'
  registered_reaction:
    enabled: true
  subspace_id: "5"
```

## gRPC
A user can query the `reactions` module gRPC endpoints.

### Reaction
The `Reaction` endpoint allows users to query a specific reaction with the given ID to a post with the given ID, 
inside the subspace with the given ID.

```bash
desmos.reactions.v1.Query/Reaction
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":5, "post_id":1, "reaction_id":1}' localhost:9090 desmos.reactions.v1.Query/Reaction
```

Example output:
```json
{
  "reaction": {
    "subspaceId": "5",
    "postId": "1",
    "id": 1,
    "value": {"@type":"/desmos.reactions.v1.FreeTextValue","text":"🚀"},
    "author": "desmos159axlj0mkvch02f95t5tkghychyeueaslk6r8f"
  }
}
```

### Reactions
The `Reactions` endpoint allows users to query all the reactions inside the subspace with the given ID. It is possible
to filter this request and get only the reactions made to a post with the given ID.

```bash
desmos.reactions.v1.Query/Reactions
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":5, "post_id":1}' localhost:9090 desmos.reactions.v1.Query/Reactions 
```

Example output:
```json
{
  "reactions": [
    {
      "subspaceId": "5",
      "postId": "1",
      "id": 1,
      "value": {"@type":"/desmos.reactions.v1.FreeTextValue","text":"🚀"},
      "author": "desmos159axlj0mkvch02f95t5tkghychyeueaslk6r8f"
    },
    {
      "subspaceId": "5",
      "postId": "1",
      "id": 2,
      "value": {"@type":"/desmos.reactions.v1.FreeTextValue","text":"😂"},
      "author": "desmos1dx6h75tkj0cuvyqf6cwn6usc9qynu39v0245m4"
    }
  ],
  "pagination": {
    "total": "2"
  }
}
```

### RegisteredReaction
The `RegisteredReaction` endpoint allows users to query a specific registered reaction with the given ID inside a subspace
with the given ID.

```bash
desmos.reactions.v1.Query/RegisteredReaction
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":5, "reaction_id":7}' localhost:9090 desmos.reactions.v1.Query/RegisteredReaction
```

Example output:
```json
{
  "registeredReaction": {
    "subspaceId": "5",
    "id": 7,
    "shorthandCode": ":hello:",
    "displayValue": "https://example.com?image=hello.png"
  }
}
```

### RegisteredReactions
The `RegisteredReactions` endpoint allows users to query all the registered reactions within the subspace with the given ID.

```bash
desmos.reactions.v1.Query/RegisteredReactions
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":5}' localhost:9090 desmos.reactions.v1.Query/RegisteredReactions
```

Example output:
```json
{
  "registeredReactions": [
    {
      "subspaceId": "5",
      "id": 7,
      "shorthandCode": ":hello:",
      "displayValue": "https://example.com?image=hello.png"
    },
    {
      "subspaceId": "5",
      "id": 8,
      "shorthandCode": ":bye:",
      "displayValue": "https://example.com?image=bye.png"
    }
  ],
  "pagination": {
    "total": "2"
  }
}
```

### ReactionsParams
The `ReactionParams` endpoint allows users to query the reaction parameters of a subspace with the given ID.

```bash
desmos.reactions.v1.Query/ReactionsParams
```

Example:
```bash
grpcurl -plaintext -d '{"subspace_id":5}' localhost:9090 desmos.reactions.v1.Query/ReactionsParams
```

Example output:
```json
{
  "params": {
    "subspaceId": "5",
    "registeredReaction": {
      "enabled": true
    },
    "freeText": {
      "enabled": true,
      "maxLength": 30,
      "regEx": "[a-z]"
    }
  }
}
```

## REST
A user can query the `reactions` module using REST endpoints.

### Reaction
The `Reaction` endpoint allows users to query a specific reaction with the given ID to a post with the given ID,
inside the subspace with the given ID.

```
/desmos/reactions/v1/subspaces/{subspace_id}/posts/{post_id}/reactions/{reaction_id}
```

### Reactions
The `Reactions` endpoint allows users to query all the reactions inside the subspace with the given ID. It is possible
to filter this request and get only the reactions made to a post with the given ID.

```
/desmos/reactions/v1/subspaces/{subspace_id}/posts/{post_id}/reactions
```

### RegisteredReaction
The `RegisteredReaction` endpoint allows users to query a specific registered reaction with the given ID inside a subspace
with the given ID.

```
/desmos/reactions/v1/subspaces/{subspace_id}/registered-reactions/{reaction_id}
```

### RegisteredReactions
The `RegisteredReactions` endpoint allows users to query all the registered reactions within the subspace with the given ID.

```
/desmos/reactions/v1/subspaces/{subspace_id}/registered-reactions
```

### ReactionsParams
The `ReactionParams` endpoint allows users to query the reaction parameters of a subspace with the given ID.

```
/desmos/reactions/v1/subspaces/{subspace_id}/params
```