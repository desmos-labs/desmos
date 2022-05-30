# ADR 010: Posts module

## Changelog

- April 05th, 2022: Initial draft;
- April 06th, 2022: First review;
- April 11th, 2022: Second review;
- April 12th, 2022: Third review;
- May 26th, 2022: Fourth review;

## Status

PROPOSED Implemented

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
syntax = "proto3";

// Post contains all the information about a single post
message Post {
  // Id of the subspace inside which the post has been created 
  uint64 subspace_id = 1;

  // Unique id of the post
  uint64 id = 2;

  // External id for this post
  string external_id = 3;

  // Text of the post
  string text = 4;

  // Entities connected to this post
  Entities entities = 5;

  // Author of the post
  string author = 6;

  // Id of the original post of the conversation
  uint64 conversation_id = 7 [default = 0];

  // A list this posts references (either as a reply, repost or quote)
  repeated PostReference referenced_posts = 8;

  // Reply settings of this post
  ReplySetting reply_settings = 9;

  // Creation date of the post
  google.protobuf.Timestamp creation_date = 10;

  // Last edited time of the post
  google.protobuf.Timestamp last_edited_date = 11;
}

// PostReference contains the details of a post reference
message PostReference {
  // Type of reference 
  Type type = 1;
  
  // Id of the referenced post
  uint64 post_id = 2;
  
  enum Type {
    TYPE_UNSPECIFIED = 0;
    TYPE_REPLIED_TO = 1;
    TYPE_QUOTED = 2;
    TYPE_REPOSTED = 3;
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
  // No reply setting specified
  REPLY_SETTING_UNSPECIFIED = 0;
  
  // Everyone will be able to reply to this post
  REPLY_SETTING_EVERYONE = 1;
  
  // Only followers of the author will be able to reply to this post
  REPLY_SETTING_FOLLOWERS = 2;
  
  // Only the author mutual followers will be able to reply to this post 
  REPLY_SETTING_MUTUAL = 3;
  
  // Only people mentioned inside this post will be able to reply
  REPLY_SETTING_MENTIONS = 4;
}

// Tag represents a generic tag 
message Tag {
  // Index of the character inside the text at which the tag starts 
  uint64 start = 1;
  
  // Index of the character inside the text at which the tag ends
  uint64 end = 2;
  
  // Tag reference (user address, hashtag value, etc)
  string tag = 3;
}

// Url contains the details of a generic URL
message Url {
  // Index of the character inside the text at which the URL starts 
  uint64 start = 1;
  
  // Index of the character inside the text at which the URL ends
  uint64 end = 2;
  
  // Value of the URL where the user should be redirected to
  string url = 3;
  
  // Display value of the URL
  string display_url = 4;
}
```

#### Attachments

```protobuf
syntax = "proto3";

// Attachment contains the data of a single post attachment
message Attachment {
  // Id of the subspace inside which the post to which this attachment should be
  // connected is
  uint64 subspace_id = 1;

  // Id of the post to which this attachment should be connected
  uint64 post_id = 2;

  // If of this attachment
  uint32 id = 3;

  // Content of the attachment
  google.protobuf.Any content = 4;
}

// Media represents a media attachment
message Media {
  string uri = 2;
  string mime_type = 3;
}

// Poll represents a poll attachment
message Poll {
  // Question of the poll
  string question = 1;
  
  // Answers the users can choose from
  repeated ProvidedAnswer provided_answers = 2;
  
  // Date at which the poll will close
  google.protobuf.Timestamp end_date = 3;
  
  // Whether the poll allows multiple choices from the same user or not 
  bool allows_multiple_answers = 4;
  
  // Whether the poll allows to edit an answer or not
  bool allows_answer_edits = 5;

  // Final poll results
  PollTallyResults final_tally_results = 6;
  
  // Provided answer contains the details of a possible poll answer 
  message ProvidedAnswer {
    // Text of the answer
    string text = 1;
    
    // Attachments of the answer
    repeated Attachment attachments = 2;
  }
}

// UserAnswer represents a user answer to a poll
message UserAnswer {
  // Subspace id inside which the post related to this attachment is located
  uint64 subspace_id = 1;

  // Id of the post associated to this attachment
  uint64 post_id = 2;

  // Id of the poll to which this answer is associated
  uint32 poll_id = 3;

  // Indexes of the answers inside the ProvidedAnswers array
  repeated uint32 answers_indexes = 4;

  // Address of the user answering the poll
  string user = 5;
}

// PollTallyResults contains the tally results for a poll
message PollTallyResults {
  repeated AnswerResult results = 1;
  
  // AnswerResult contains the result of a single poll provided answer
  message AnswerResult {
    // Index of the answer inside the poll's ProvidedAnswers slice 
    uint32 answer_index = 1;
    
    // Number of votes the answer has received
    uint64 votes = 2;
  }
}
```

### `Params` 
```protobuf
// Params contains the parameters for the posts module
message Params {
  // Maximum length of the post text
  uint64 max_text_length = 1; 
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
syntax = "proto3";

service Msg {
  // CreatePost allows to create a new post
  rpc CreatePost(MsgCreatePost) returns (MsgCreatePostResponse);
  
  // EditPost allows to edit an existing post
  rpc EditPost(MsgEditPost) returns (MsgEditPostResponse);

  // DeletePost allows to delete an existing post
  rpc DeletePost(MsgDeletePost) returns (MsgDeletePostResponse);
  
  // AddPostAttachment allows to add a new attachment to a post
  rpc AddPostAttachment(MsgAddPostAttachment) returns (MsgAddPostAttachmentResponse);
  
  // RemovePostAttachment allows to remove an attachment from a post
  rpc RemovePostAttachment(MsgRemovePostAttachment) returns (MsgRemovePostAttachmentResponse);
  
  // AnswerPoll allows to answer a post poll
  rpc AnswerPoll(MsgAnswerPoll) returns (MsgAnswerPollResponse);
}

// MsgCreatePost represents the message to be used to create a post.
message MsgCreatePost {
  // Id of the subspace inside which the post must be created 
  uint64 subspace_id = 1;
  
  // External id for this post
  string external_id = 2;

  // Text of the post
  string text = 3;

  // Entities connected to this post
  Entities entities = 4;

  // Attachments of the post
  repeated google.protobuf.Any attachments = 5;
  
  // Author of the post
  string author = 6;

  // Id of the original post of the conversation
  uint64 conversation_id = 7;

  // Reply settings of this post
  ReplySetting reply_settings = 8;

  // A list this posts references (either as a reply, repost or quote)
  repeated PostReference referenced_posts = 9;
}

// MsgCreatePostResponse defines the Msg/CreatePost response type.
message MsgCreatePostResponse {
  // Id of the newly created post
  uint64 post_id = 1;
  
  // Creation date of the post
  google.protobuf.Timestamp creation_date = 2;
}

// MsgEditPost represents the message to be used to edit a post.
message MsgEditPost {
  // Id of the subspace inside which the post is
  uint64 subspace_id = 1;

  // Id of the post to edit
  uint64 id = 2;

  // New text of the post
  string text = 3;

  // Editor of the post
  string editor = 4;

  // New entities connected to this post
  Entities entities = 5;

  // Author of the post
  string author = 6;
}

// MsgCreatePostResponse defines the Msg/EditPost response type.
message MsgEditPostResponse {
  // Edit date of the post
  google.protobuf.Timestamp edit_date = 1;
}

// MsgDeletePost represents the message used when deleting a post.
message MsgDeletePost {
  // Id of the subspace containing the post
  uint64 subspace_id = 1;

  // Id of the post to be deleted
  uint64 post_id = 2;

  // User that is deleting the post
  string signer = 3;
}

// MsgDeletePostResponse represents the Msg/DeletePost response type
message MsgDeletePostResponse {}

// MsgAddPostAttachment represents the message that should be
// used when adding an attachment to post
message MsgAddPostAttachment {
  // Id of the subspace containing the post
  uint64 subspace_id = 1;
  
  // Id of the post to which to add the attachment
  uint64 post_id = 2;

  // Content of the attachment
  google.protobuf.Any content = 3;

  // Editor of the post
  string editor = 4;
}

// MsgAddPostAttachmentResponse defines the Msg/AddPostAttachment response type.
message MsgAddPostAttachmentResponse {
  // New id of the uploaded attachment 
  uint32 attachment_id = 1;

  // Edit date of the post
  google.protobuf.Timestamp edit_date = 2;
}

// MsgRemovePostAttachment represents the message to be used when 
// removing an attachment from a post
message MsgRemovePostAttachment {
  // Id of the subspace containing the post
  uint64 subspace_id = 1;

  // Id of the post from which to remove the attachment
  uint64 post_id = 2;
  
  // Id of the attachment to be removed
  uint32 attachment_id = 3;

  // User that is removing the attachment
  string editor = 4;
}

// MsgRemovePostAttachmentResponse defines the 
// Msg/RemovePostAttachment response type.
message MsgRemovePostAttachmentResponse {
  // Edit date of the post
  google.protobuf.Timestamp edit_date = 1;
}

// MsgAnswerPoll represents the message used to answer a poll
message MsgAnswerPoll {
  // Id of the subspace containing the post
  uint64 subspace_id = 1;

  // Id of the post that contains the poll to be answered
  uint64 post_id = 2;
  
  // Id of the poll to be answered 
  uint32 poll_id = 3;
  
  // Indexes of the answer inside the ProvidedAnswers array
  repeated uint32 answers_indexes = 4;
  
  // Address of the user answering the poll
  string signer = 5;
}

// MsgAnswerPollResponse represents the MSg/AnswerPoll response type
message MsgAnswerPollResponse {}
```

### `Query` Service
```protobuf
syntax = "proto3";

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
  
  // Params queries the module parameters
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse) {
    option (google.api.http).get = "/desmos/posts/v1/params";
  }
}

// QueryPostsRequest is the request type for the Query/Posts RPC method
message QueryPostsRequest {
  // Id of the subspace to query the posts for
  uint64 subspace_id = 1;
  
  // pagination defines an pagination for the request.
  cosmos.base.query.v1beta1.PageRequest pagination = 2;
}

// QueryPostsResponse is the response type for the Query/Posts RPC method
message QueryPostsResponse {
  repeated Post posts = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryPostRequest is the request type for the Query/Post RPC method
message QueryPostRequest {
  uint64 subspace_id = 1;
  uint64 post_id = 2;
}

// QueryPostResponse is the response type for the Query/Post RPC method
message QueryPostResponse {
  Post post = 1;
}

// QueryPostsRequest is the request type for the Query/PostAttachments RPC method
message QueryPostAttachmentsRequest {
  // Id of the subspace where the post is stored 
  uint64 subspace_id = 1;
  
  // Id of the post to query the attachments for
  uint64 post_id = 2;

  // pagination defines an pagination for the request.
  cosmos.base.query.v1beta1.PageRequest pagination = 3;
}

// QueryPostAttachmentsResponse is the response type for the Query/PostAttachments RPC method
message QueryPostAttachmentsResponse {
  repeated Attachment attachments = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryPollAnswersRequest is the request type for the Query/PollAnswers RPC method
message QueryPollAnswersRequest {
  // Id of the subspace where the post is stored 
  uint64 subspace_id = 1;

  // Id of the post that holds the poll
  uint64 post_id = 2;
  
  // Id of the poll to query the answers for
  uint32 poll_id = 3;

  // (Optional) Address of the user to query the responses for
  string user = 4;

  // pagination defines an pagination for the request.
  cosmos.base.query.v1beta1.PageRequest pagination = 5;
}

// QueryPollAnswersResponse is the response type for the Query/PollAnswers RPC method
message QueryPollAnswersResponse {
  repeated Answer answers = 1;
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
  
  // Answer contains the details about a single user answer to a poll
  message Answer {
    // Address of the user that input this 
    string user = 1;
    
    // Indexes of the answers inside the ProvidedAnswers array
    repeated uint32 answers_indexes = 2;
  }
}

// QueryParamsRequest is the request type for the Query/Params RPC method
message QueryParamsRequest {}

// QueryParamsResponse is the response type for the Query/Params RPC method
message QueryParamsResponse {
  Params params = 1;
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
   - `PermissionEditContent` to allow users to edit a content inside a subspace.

## Further Discussions

## References

- {reference link}