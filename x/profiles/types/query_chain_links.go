package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryChainLinksRequest returns a new QueryChainLinksRequest instance
func NewQueryChainLinksRequest(
	user, chainName, target string, pageReq *query.PageRequest,
) *QueryChainLinksRequest {
	return &QueryChainLinksRequest{
		User:       user,
		ChainName:  chainName,
		Target:     target,
		Pagination: pageReq,
	}
}
