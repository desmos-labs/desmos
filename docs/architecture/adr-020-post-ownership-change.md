# ADR 020: Post ownership transfer

## Changelog
- April 21th, 2023: First draft;
- April 26th, 2023: First review;

## Status

Accepted Not Implemented

## Abstract

This ADR introduces a new feature that enables users to transfer post ownership to another person.

## Context

Desmos is a social network protocol that allows users to create, share, and engage with content on a decentralized platform. As of now, Desmos does not provide a feature that allows users to transfer the ownership of their posts to other users. This has caused inconvenience for users who wish to transfer ownership of their posts, as they have to create new posts and lose the engagement history and feedback of the original post. Therefore, the introduction of the new feature that enables users to transfer post ownership to another person aims to address this issue and provide users with more control over their content. The proposed feature is expected to improve user experience and enhance the functionality of the Desmos protocol.

## Decision

We will add an __owner__ field to the `Post` structure and implement post ownership transfer functionality. Simple transfer functionality could result in someone receiving spam messages that damage their reputation. For example, scammers can transfer scam messages to victims, leading to reputational harm. Therefore, the functionality we will implement should include the following features:
1. Permit the sender to initiate a request to transfer post ownership to the receiver
2. Allow the receiver to accept or refuse the transfer request from the sender
3. Allow the sender to cancel the transfer request if the receiver has not deal with the request

### Type

The upcoming changes of `Post` is as follows:

```proto
message Post {
  uint64 subspace_id = 1;
  uint32 section_id = 2;
  
  ...skip

  google.protobuf.Timestamp last_edited_date = 13;
  
  // Owner of the post
  string owner = 14;
}

In order to handle the post owner transfer process easily, we will define a new structure as follows:

// PostOwnerTransferRequest represents a request to transfer the ownership of a post from the sender to the receiver
message PostOwnerTransferRequest {

  // Id of the subspace that holds the post to transfer
  uint64 subspace_id = 1;
  
  // Id of the post which will be transferred
  uint32 post_id = 2;

  // Address of the sender
  string sender = 3;

  // Address of the receiver
  string receiver = 4;
}
```

### Store

#### PostOwnerTransferRequest

To simplify the management of post owner transfer requests, we will store each request using the following key format:

```
IncomingPostOwnerTransferRequestPrefix | SubspaceID | ReceiverAddress | PostID | -> Protobuf(PostOwnerTransferRequest)
```

This structure enables the receiver to easily manage incoming requests by iterating over all requests with a given subspace ID and receiver address, which will be the most frequently used query.

### `Msg` Service

  To ensure the safe handling of post owner transfer requests, we will have the following operations:
  1. request a post owner transfer to a receiver
  2. cancel a post owner transfer request
  3. accept a post owner transfer request
  4. refuse a post owner transfer request

```proto
service Msg {
  // RequestPostOwnerTransfer allows sender to send a request to transfer a post ownership to receiver
  rpc RequestPostOwnerTransfer(MsgRequestPostOwnerTransfer) returns (MsgRequestPostOwnerTransferResponse);

  // CancelPostOwnerTransfer allows sender to cancel an outgoing post owner transfer request
  rpc CancelPostOwnerTransfer(MsgCancelPostOwnerTransfer) returns (MsgCancelPostOwnerTransferResponse);

  // AcceptPostOwnerTransfer allows receiver to accept an incoming post transfer request
  rpc AcceptPostOwnerTransfer(MsgAcceptPostOwnerTransfer) returns (MsgAcceptPostOwnerTransferResponse);

  // RefusePostOwnerTransfer allows receiver to refuse an incoming post transfer request
  rpc RefusePostOwnerTransfer(MsgRefusePostOwnerTransfer) returns (MsgRefusePostOwnerTransferResponse);
}

// MsgRequestPostOwnerTransfer represent a message used to transfer a post ownership to receiver
message MsgRequestPostOwnerTransfer {
  // Id of the subspace that holds the post which ownership should be transfered
  uint64 subspace_id = 1;
    
  // Id of the post which will be transferred
  uint64 post_id = 2;

  // Address of the post ownership receiver
  string receiver = 3;
    
  // Address of the sender who is creating a transfer request
  string sender = 4;
}
// MsgRequestPostOwnerTransferResponse defines the Msg/RequestPostOwnerTransfer response type
message MsgRequestPostOwnerTransferResponse {}

// MsgCancelPostOwnerTransfer represents a message used to cancel a outgoing post transfer request
message MsgCancelPostOwnerTransfer {
  // Id of the subspace that holds the post for which the request should be canceled
  uint64 subspace_id = 1;
    
  // Id of the post whose request will be cancelled
  uint64 post_id = 2;

  // Address of the sender who is cancelling the request
  string sender = 3;
}
// MsgCancelPostOwnerTransferResponse defines the Msg/CancelPostOwnerTransfer response type
message MsgRequestPostOwnerTransferResponse {}

// MsgAcceptPostOwnerTransfer represents a message used to accept a incoming post transfer request
message MsgAcceptPostOwnerTransfer {
  // Id of the subspace where the request will be accepted
  uint64 subspace_id = 1;
    
  // Id of the post whose request will be accepted
  uint64 post_id = 2;

  // Address of the receiver who is accepting the request
  string receiver = 3;
}

// MsgAcceptPostOwnerTransferResponse defines the Msg/AcceptPostOwnerTransfer response type
message MsgAcceptPostOwnerTransferResponse {}

// MsgRefusePostOwnerTransfer represents a message used to refuse a incoming post transfer request
message MsgRefusePostOwnerTransfer {
  // Id of the subspace holding the post for which the request will be refused
  uint64 subspace_id = 1;
    
  // Id of the post for which the request will be refused
  uint64 post_id = 2;

  // Address of the request receiver
  string receiver = 3;
}

// MsgRefusePostOwnerTransfer defines the Msg/RefusePostOwnerTransfer response type
message MsgRefusePostOwnerTransferResponse {}
```

### `Query` service

We will also implement a query service that enables the user to manage the incoming post owner requests by the following queries.

```proto
service Query {
  // IncomingPostOwnerTransferRequests queries all the post owner transfers requests that
  // have been made towards the receiver with the given address
  rpc IncomingPostOwnerTransferRequests(QueryIncomingPostOwnerTransferRequestsRequest) returns (QueryIncomingPostOwnerTransferRequestResponse) {
    option (google.api.http).get = "/desmos/posts/v4/subspaces/{subspace_id}/post-owner-transfer-requests";
  }
}

// QueryIncomingPostOwnerTransferRequestsRequest is the request type for the
// Query/IncomingPostOwnerTransferRequests RPC endpoint
message QueryIncomingPostOwnerTransferRequestsRequest {
  option (gogoproto.equal) = false;
  option (gogoproto.goproto_getters) = false;

  // (optional) Receiver represents the address of the user to which query the
  // incoming requests for
  string receiver = 1;

  // Pagination defines an optional pagination for the request
  cosmos.base.query.v1beta1.PageRequest pagination = 2;
}

// QueryIncomingPostOwnerTransferRequestsResponse is the response type for the
// Query/IncomingPostOwnerTransferRequests RPC method.
message QueryIncomingPostOwnerTransferRequestsResponse {
  // Requests represent the list of all the post owner transfer requests made towards
  // the receiver
  repeated desmos.posts.v4.PostOwnerTransferRequest requests = 1
      [ (gogoproto.nullable) = false ];

  // Pagination defines the pagination response
  cosmos.base.query.v1beta1.PageResponse pagination = 2;
}
```

## Consequences

### Backwards Compatibility

The solution outlined above is **not** backwards compatible and will require a migration script to update all existing posts to the new version. This script will handle the following tasks:
- migrate all posts to have a new __owner__ field.

### Positive

- Enable users to transfer the ownership of a post to another user

### Negative

- Storing extra post transfer requests info takes up more storage space

### Neutral

(none known)

## References