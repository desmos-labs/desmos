package types

// DONTCOVER

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryPostsRequest returns a new QueryPostsRequest instance
func NewQueryPostsRequest(subspaceID uint64, pagination *query.PageRequest) *QueryPostsRequest {
	return &QueryPostsRequest{
		SubspaceId: subspaceID,
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
