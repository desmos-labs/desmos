package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryReactionsRequest returns a new QueryReactionsRequest instance
func NewQueryReactionsRequest(subspaceID uint64, postID uint64, pagination *query.PageRequest) *QueryReactionsRequest {
	return &QueryReactionsRequest{
		SubspaceId: subspaceID,
		PostId:     postID,
		Pagination: pagination,
	}
}

// NewQueryRegisteredReactionsRequest returns a new QueryRegisteredReactionsRequest instance
func NewQueryRegisteredReactionsRequest(subspaceID uint64, pagination *query.PageRequest) *QueryRegisteredReactionsRequest {
	return &QueryRegisteredReactionsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQueryReactionsParamsRequest returns a new QueryReactionsParamsRequest instance
func NewQueryReactionsParamsRequest(subspaceID uint64) *QueryReactionsParamsRequest {
	return &QueryReactionsParamsRequest{
		SubspaceId: subspaceID,
	}
}
