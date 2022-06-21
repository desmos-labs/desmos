package types

// DONTCOVER

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryReactionsRequest returns a new QueryReactionsRequest instance
func NewQueryReactionsRequest(subspaceID uint64, postID uint64, user string, pagination *query.PageRequest) *QueryReactionsRequest {
	return &QueryReactionsRequest{
		SubspaceId: subspaceID,
		PostId:     postID,
		User:       user,
		Pagination: pagination,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryReactionsResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, reaction := range r.Reactions {
		err := reaction.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
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
