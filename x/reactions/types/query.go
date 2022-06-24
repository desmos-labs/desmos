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

// NewQueryReactionRequest returns a new QueryReactionRequest request
func NewQueryReactionRequest(subspaceID uint64, postID uint64, reactionID uint32) *QueryReactionRequest {
	return &QueryReactionRequest{
		SubspaceId: subspaceID,
		PostId:     postID,
		ReactionId: reactionID,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryReactionResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return r.Reaction.UnpackInterfaces(unpacker)
}

// NewQueryRegisteredReactionsRequest returns a new QueryRegisteredReactionsRequest instance
func NewQueryRegisteredReactionsRequest(subspaceID uint64, pagination *query.PageRequest) *QueryRegisteredReactionsRequest {
	return &QueryRegisteredReactionsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQueryRegisteredReactionRequest returns a new QueryRegisteredReactionRequest instance
func NewQueryRegisteredReactionRequest(subspaceID uint64, registeredReactionID uint32) *QueryRegisteredReactionRequest {
	return &QueryRegisteredReactionRequest{
		SubspaceId: subspaceID,
		ReactionId: registeredReactionID,
	}
}

// NewQueryReactionsParamsRequest returns a new QueryReactionsParamsRequest instance
func NewQueryReactionsParamsRequest(subspaceID uint64) *QueryReactionsParamsRequest {
	return &QueryReactionsParamsRequest{
		SubspaceId: subspaceID,
	}
}
