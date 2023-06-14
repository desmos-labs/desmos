package types

// DONTCOVER

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
)

// NewQuerySubspacePostsRequest returns a new QuerySubspacePostsRequest instance
func NewQuerySubspacePostsRequest(subspaceID uint64, pagination *query.PageRequest) *QuerySubspacePostsRequest {
	return &QuerySubspacePostsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQuerySectionPostsRequest returns a new QuerySectionPostsRequest instance
func NewQuerySectionPostsRequest(subspaceID uint64, sectionID uint32, pagination *query.PageRequest) *QuerySectionPostsRequest {
	return &QuerySectionPostsRequest{
		SubspaceId: subspaceID,
		SectionId:  sectionID,
		Pagination: pagination,
	}
}

// NewQueryPostRequest returns a new QueryPostRequest instance
func NewQueryPostRequest(subspaceID uint64, postID uint64) *QueryPostRequest {
	return &QueryPostRequest{
		SubspaceId: subspaceID,
		PostId:     postID,
	}
}

// NewQueryPostAttachmentsRequest returns a new QueryPostAttachmentsRequest instance
func NewQueryPostAttachmentsRequest(
	subspaceID uint64, postID uint64, pagination *query.PageRequest,
) *QueryPostAttachmentsRequest {
	return &QueryPostAttachmentsRequest{
		SubspaceId: subspaceID,
		PostId:     postID,
		Pagination: pagination,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryPostAttachmentsResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, a := range r.Attachments {
		var content AttachmentContent
		err := unpacker.UnpackAny(a.Content, &content)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewQueryPollAnswersRequest returns a new QueryPollAnswersRequest instance
func NewQueryPollAnswersRequest(
	subspaceID uint64, postID uint64, pollID uint32, user string, pagination *query.PageRequest,
) *QueryPollAnswersRequest {
	return &QueryPollAnswersRequest{
		SubspaceId: subspaceID,
		PostId:     postID,
		PollId:     pollID,
		User:       user,
		Pagination: pagination,
	}
}

// NewQueryParamsRequest returns a new QueryParamsRequest instance
func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}

// QueryIncomingPostOwnerTransferRequestsRequest returns a new QueryIncomingPostOwnerTransferRequestsRequest instance
func NewQueryIncomingPostOwnerTransferRequestsRequest(
	subspaceID uint64, receiver string, pagination *query.PageRequest,
) *QueryIncomingPostOwnerTransferRequestsRequest {
	return &QueryIncomingPostOwnerTransferRequestsRequest{
		SubspaceId: subspaceID,
		Receiver:   receiver,
		Pagination: pagination,
	}
}
