# ADR 010: Posts module

## Changelog

- April 05th, 2022: Initial draft;
- April 06th, 2022: First review;
- April 11th, 2022: Second review;
- April 12th, 2022: Third review;

## Status

PROPOSED

## Abstract

This ADR contains the definition of the `x/posts` module which will allow users to post text contents inside Desmos  subspaces.

## Context

As Desmos is thought to be a protocol to build decentralized socially-enabled applications (or social networks), one of the main features that we MUST make sure exists is the ability for users to create content inside such social networks. When designing this feature, we SHOULD consider the following:

1. a post will always be submitted inside a _subspace_;
2. a post should allow to specify the minimum amount of information so that any DApp can make use of them;
3. it is responsibility of the DApp to make sure that the various fields are filled properly upon uploading the post.

## Decision

We will implement a new module named `x/posts` that allows users to perform the following operations inside subspaces that allow them to:

- create a new post
- edit an existing post
- participate inside a discussion with a comment or reply to a post
- delete an existing post

### Types
Each post MUST always have an _author_ which identifies the user that has created the content. It also MUST always reference a _subspace_ inside which it is created. In order to be valid, a post MUST either contain a _text_ or at least one _attachment_.

Optionally, a post can also have a series of _entities_ that have been parsed out of the text. These allow to identify particular content that should be displayed in custom ways (hashtags, mentions, urls, etc).

#### Post

```protobuf
// Post contains all the information about a single post
message Post {
  // Id of the subspace inside which the post has been created 
  required uint64 subspace_id = 1;

  // Unique id of the post
  required uint64 id = 2;

  // External id for this post
  optional string external_id = 3;

  // Text of the post
  optional string text = 4;

  // Entities connected to this post
  optional Entities entities = 5;

  // Author of the post
  required string author = 6;

  // Id of the original post of the conversation
  optional uint64 conversation_id = 7 [default = 0];

  // A list this posts references (either as a reply, repost or quote)
  repeated PostReference referenced_posts = 8;

  // Reply settings of this post
  required ReplySetting reply_settings = 9;

  // Creation date of the post
  required google.protobuf.Timestamp creation_date = 10;

  // Last edited time of the post
  optional google.protobuf.Timestamp last_edited_date = 11;
}

// PostReference contains the details of a post reference
message PostReference {
  // Type of reference 
  required Type type = 1;
  
  // Id of the referenced post
  required uint64 post_id = 2;
  
  enum Type {
    REPLIED_TO = 1;
    QUOTED = 2;
    REPOSTED = 3;
  }
}

// Contains the details of entities parsed out of the post text
message Entities {
  repeated Tag hashtags = 1;
  repeated Tag mentions = 2;
  repeated Url urls = 3;
}

// ReplySetting contains the possible reply settings that a post can have
enum ReplySetting {
  // Everyone will be able to reply to this post
  EVERYONE = 1;
  
  // Only followers of the author will be able to reply to this post
  FOLLOWERS = 2;
  
  // Only the author mutual followers will be able to reply to this post 
  MUTUALS = 3;
  
  // Only people mentioned inside this post will be able to reply
  MENTIONS = 4;
}

// Tag represents a generic tag 
message Tag {
  // Index of the character inside the text at which the tag starts 
  required uint64 start = 1;
  
  // Index of the character inside the text at which the tag ends
  required uint64 end = 2;
  
  // Tag reference (user address, hashtag value, etc)
  required string tag = 3;
}

// Url contains the details of a generic URL
message Url {
  // Index of the character inside the text at which the URL starts 
  required uint64 start = 1;
  
  // Index of the character inside the text at which the URL ends
  required uint64 end = 2;
  
  // Value of the URL where the user should be redirected to
  required string url = 3;
  
  // Display value of the URL
  optional string display_url = 4;
}
```

#### Attachments

```protobuf
// Attachment contains the data of a single post attachment
message Attachment {
  required uint32 id = 1;
  oneof sum {
    Poll poll = 2;
    Media media = 3;
  }
}

// Media represents a media attachment
message Media {
  required string uri = 2;
  required string mime_type = 3;
}

// Poll represents a poll attachment
message Poll {
  // Question of the poll
  required string question = 2;
  
  // Answers the users can choose from
  repeated ProvidedAnswer provided_answers = 3;
  
  // Date at which the poll will close
  required google.protobuf.Timestamp end_date = 4;
  
  // Whether the poll allows multiple choices from the same user or not 
  optional bool allows_multiple_answers = 5 [default = false];
  
  // Whether the poll allows to edit an answer or not
  optional bool allows_answer_edits = 6 [default = false];
  
  // Provided answer contains the details of a possible poll answer 
  message ProvidedAnswer {
    // Text of the answer
    optional string text = 1;
    
    // Attachments of the answer
    repeated Attachment attachments = 2;
  }
}

// PollTallyResults contains the tally results for a poll
message PollTallyResults {
  repeated AnswerResult results = 1;
  
  // AnswerResult contains the result of a single poll provided answer
  message AnswerResult {
    // Index of the answer inside the poll's ProvidedAnswers slice 
    required uint32 answer_index = 1;
    
    // Number of votes the answer has received
    required uint64 votes = 2;
  }
}
```

### `Params` 
```protobuf
// Params contains the parameters for the posts module
message Params {
  // Maximum length of the post text
  required uint64 max_text_length = 1; 
}
```

### `Msg` Service
We will allow the following operations to be performed.

**Post management**
- Create a post 
- Edit an existing post
- Add an attachment to a post
- Remove an attachment from a post
- Delete a post   

**Post interaction**
- Answer a post poll

> NOTE  
> In order to make sure subspace moderators and admins can make sure the ToS of their application is always respected, both the removing of an attachment and the deletion of a post should be allowed also to such people.

> NOTE  
> As per their nature, _attachments_ are **immutable**. This means that the only operations allowed on a post attachments are either adding an attachment or deleting an existing attachment. No edit on the attachment itself is permitted.

```protobuf
service Msg {
  // CreatePost allows to create a new post
  rpc CreatePost(MsgCreatePost) returns (MsgCreatePostResponse);
  
  // EditPost allows to edit an existing post
  rpc EditPost(MsgEditPost) returns (MsgEditPostResponse);
  
  // AddPostAttachment allows to add a new attachment to a post
  rpc AddPostAttachment(MsgAddPostAttachment) returns (MsgAddPostAttachmentResponse);
  
  // RemovePostAttachment allows to remove an attachment from a post
  rpc RemovePostAttachment(MsgRemovePostAttachment) returns (MsgRemovePostAttachmentResponse);
  
  // DeletePost allows to delete an existing post
  rpc DeletePost(MsgDeletePost) returns (MsgDeletePostResponse);
  
  // AnswerPoll allows to answer a post poll
  rpc AnswerPoll(MsgAnswerPoll) returns (MsgAnswerPollResponse);
}

// MsgCreatePost represents the message to be used to create a post.
message MsgCreatePost {
  // Id of the subspace inside which the post must be created 
  required uint64 subspace_id = 1;
  
  // External id for this post
  optional string external_id = 2;

  // Text of the post
  optional string text = 3;

  // Entities connected to this post
  optional Entities entities = 4;

  // Attachments of the post
  repeated Attachment attachments = 5;
  
  // Author of the post
  required string author = 6;

  // Id of the original post of the conversation
  optional uint64 conversation_id = 7 [default = 0];

  // Reply settings of this post
  required ReplySetting reply_settings = 8;

  // A list this posts references (either as a reply, repost or quote)
  repeated PostReference referenced_posts = 9;

  // Attachment contains the data of a single post attachment
  message Attachment {
    oneof sum {
      Poll poll = 1;
      Media media = 2;
    }
  }
}

// MsgCreatePostResponse defines the Msg/CreatePost response type.
message MsgCreatePostResponse {
  // Id of the newly created post
  required uint64 post_id = 1;
  
  // Creation date of the post
  required google.protobuf.Timestamp creation_date = 2;
}

// MsgEditPost represents the message to be used to edit a post.
message MsgEditPost {
  // Id of the subspace inside which the post is
  required uint64 subspace_id = 1;

  // Id of the post to edit
  required uint64 id = 2;

  // New text of the post
  required string text = 3;

  // Editor of the post
  required string editor = 4;

  // New entities connected to this post
  optional Entities entities = 5;

  // Author of the post
  required string author = 6;
}

// MsgCreatePostResponse defines the Msg/EditPost response type.
message MsgEditPostResponse {
  // Edit date of the post
  required google.protobuf.Timestamp edit_date = 1;
}

// MsgAddPostAttachment represents the message that should be
// used when adding an attachment to post
message MsgAddPostAttachment {
  // Id of the subspace containing the post
  required uint64 subspace_id = 1;
  
  // Id of the post to which to add the attachment
  required uint64 post_id = 2;
  
  // Attachment to be added to the post 
  required Attachment attachment = 3;

  // Author of the post
  required string author = 4;
  
  message Attachment {
    oneof sum {
      Poll poll = 1;
      Media media = 2;
    }
  }
}

// MsgAddPostAttachmentResponse defines the Msg/AddPostAttachment response type.
message MsgAddPostAttachmentResponse {
  // New id of the uploaded attachment 
  required uint32 attachment_id = 1;

  // Edit date of the post
  required google.protobuf.Timestamp edit_date = 2;
}

// MsgRemovePostAttachment represents the message to be used when 
// removing an attachment from a post
message MsgRemovePostAttachment {
  // Id of the subspace containing the post
  required uint64 subspace_id = 1;

  // Id of the post from which to remove the attachment
  required uint64 post_id = 2;
  
  // Id of the attachment to be removed
  required uint32 attachment_id = 3;

  // User that is removing the attachment
  required string signer = 4;
}

// MsgRemovePostAttachmentResponse defines the 
// Msg/RemovePostAttachment response type.
message MsgRemovePostAttachmentResponse {
  // Edit date of the post
  required google.protobuf.Timestamp edit_date = 1;
}

// MsgDeletePost represents the message used when deleting a post.
message MsgDeletePost {
  // Id of the subspace containing the post
  required uint64 subspace_id = 1;

  // Id of the post to be deleted
  required uint64 post_id = 2;
  
  // User that is deleting the post
  required string signer = 3;
}

// MsgDeletePostResponse represents the Msg/DeletePost response type
message MsgDeletePostResponse {}

// MsgAnswerPoll represents the message used to answer a poll
message MsgAnswerPoll {
  // Id of the subspace containing the post
  required uint64 subspace_id = 1;

  // Id of the post that contains the poll to be answered
  required uint64 post_id = 2;
  
  // Id of the poll to be answered 
  required uint32 poll_id = 3;
  
  // Indexes of the answer inside the ProvidedAnswers array
  repeated uint32 answers_indexes = 4;
  
  // Address of the user answering the poll
  required string signer = 5;
}

// MsgAnswerPollResponse represents the MSg/AnswerPoll response type
message MsgAnswerPollResponse {}
```

### `Query` Service
```protobuf
// Query defines the gRPC querier service
service Query {
  // Posts queries all the posts inside a given subspace
  rpc Posts(QueryPostsRequest) returns (QueryPostsResponse) {
    option (google.api.http).get = "/desmos/posts/v1/{subspace_id}/posts";
  }
  
  // Post queries for a single post inside a given subspace
  rpc Post(QueryPostRequest) returns (QueryPostResponse) {
    option (google.api.http).get = "/desmos/posts/v1/{subspace_id}/posts/{post_id}";
  }
  
  // PostAttachments queries the attachments of the post having the given id
  rpc PostAttachments(QueryPostAttachmentsRequest) returns (QueryPostAttachmentsResponse) {
    option (google.api.http).get = "/desmos/posts/v1/{subspace_id}/posts/{post_id}/attachments";
  }
  
  // PollAnswers queries the answers for the poll having the given id
  rpc PollAnswers(QueryPollAnswersRequest) returns (QueryPollAnswersResponse) {
    option (google.api.http).get = "/desmos/posts/v1/{subspace_id}/posts/{post_id}/polls/{poll_id}/answers";
  }
  
  // PollTallyResults queries the tally results for an ended poll
  rpc PollTallyResults(QueryPollTallyResultRequest) returns (QueryPollTallyResultResponse) {
    option (google.api.http).get = "/desmos/posts/v1/{subspace_id}/posts/{post_id}/polls/{poll_id}/results";
  }
  
  // Params queries the module parameters
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse) {
    option (google.api.http).get = "/desmos/posts/v1/params";
  }
}

// QueryPostsRequest is the request type for the Query/Posts RPC method
message QueryPostsRequest {
  // Id of the subspace to query the posts for
  required uint64 subspace_id = 1;
  
  // pagination defines an optional pagination for the request.
  optional cosmos.base.query.v1beta1.PageRequest pagination = 2;
}

// QueryPostsResponse is the response type for the Query/Posts RPC method
message QueryPostsResponse {
  repeated Post posts = 1;
  optional cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryPostRequest is the request type for the Query/Post RPC method
message QueryPostRequest {
  required uint64 subspace_id = 1;
  required uint64 post_id = 2;
}

// QueryPostResponse is the response type for the Query/Post RPC method
message QueryPostResponse {
  required Post post = 1;
}

// QueryPostsRequest is the request type for the Query/PostAttachments RPC method
message QueryPostAttachmentsRequest {
  // Id of the subspace where the post is stored 
  required uint64 subspace_id = 1;
  
  // Id of the post to query the attachments for
  required uint64 post_id = 2;

  // pagination defines an optional pagination for the request.
  cosmos.base.query.v1beta1.PageRequest pagination = 3;
}

// QueryPostAttachmentsResponse is the response type for the Query/PostAttachments RPC method
message QueryPostAttachmentsResponse {
  repeated Attachment attachments = 1;
  optional cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryPollAnswersRequest is the request type for the Query/PollAnswers RPC method
message QueryPollAnswersRequest {
  // Id of the subspace where the post is stored 
  required uint64 subspace_id = 1;

  // Id of the post that holds the poll
  required uint64 post_id = 2;
  
  // Id of the poll to query the answers for
  required uint32 poll_id = 3;

  // (Optional) Address of the user to query the responses for
  optional string user = 4;

  // pagination defines an optional pagination for the request.
  cosmos.base.query.v1beta1.PageRequest pagination = 5;
}

// QueryPollAnswersResponse is the response type for the Query/PollAnswers RPC method
message QueryPollAnswersResponse {
  repeated Answer answers = 1;
  optional cosmos.base.query.v1beta1.PageResponse pagination = 2;
  
  // Answer contains the details about a single user answer to a poll
  message Answer {
    // Address of the user that input this 
    required string user = 1;
    
    // Indexes of the answers inside the ProvidedAnswers array
    repeated uint32 answers_indexes = 2;
  }
}

// QueryPollTallyResultRequest is the request type for the Query/PollTallyResults RPC method
message QueryPollTallyResultRequest {
  required uint64 subspace_id = 1;
  required uint64 post_id = 2;
  required uint32 poll_id = 3;
}

// QueryPollTallyResultResponse is the response type for the Query/PollTallyResults RPC method
message QueryPollTallyResultResponse {
  repeated PollTallyResults results = 1;
}

// QueryParamsRequest is the request type for the Query/Params RPC method
message QueryParamsRequest {}

// QueryParamsResponse is the response type for the Query/Params RPC method
message QueryParamsResponse {
  required Params params = 1;
}
```

## Consequences

### Backwards Compatibility

The changes described inside this ADR are **not** backward compatible. To solve this, we will rely on the `x/upgrade` module in order to properly add these new features inside a running chain. If necessary, to make sure no extra operation is performed, we should make sure that `fromVm[poststypes.ModuleName]` is set to `1` before running the migrations, so that the `InitGenesis` method does not get called.

### Positive

- Allows users to create content inside an application

### Negative

- Not known

### Neutral

- Required the `x/subspaces` to implement two new permissions: 
   - `PermissionManageContent` to allow moderators to remove post attachment and posts from a subspace; 
   - `PermissionCreateContent` to allow users to create a content inside a subspace;
   - `PermissionEditContent` to allow users to edit a content inside a subspace.

## Further Discussions

## References

- {reference link}