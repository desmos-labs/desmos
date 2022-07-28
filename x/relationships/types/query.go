package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryRelationshipsRequest returns a new QueryRelationshipsRequest instance
func NewQueryRelationshipsRequest(
	subspaceID uint64, user string, counterparty string, pagination *query.PageRequest,
) *QueryRelationshipsRequest {
	return &QueryRelationshipsRequest{
		SubspaceId:   subspaceID,
		User:         user,
		Counterparty: counterparty,
		Pagination:   pagination,
	}
}

// NewQueryBlocksRequest returns a new QueryBlocksRequest instance
func NewQueryBlocksRequest(
	subspaceID uint64, blocker string, blocked string, pagination *query.PageRequest,
) *QueryBlocksRequest {
	return &QueryBlocksRequest{
		SubspaceId: subspaceID,
		Blocker:    blocker,
		Blocked:    blocked,
		Pagination: pagination,
	}
}
