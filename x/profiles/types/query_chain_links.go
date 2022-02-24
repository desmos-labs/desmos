package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryChainLinksRequest returns a new QueryChainLinksRequest instance
func NewQueryChainLinksRequest(user string, pageReq *query.PageRequest) *QueryChainLinksRequest {
	return &QueryChainLinksRequest{
		User:       user,
		Pagination: pageReq,
	}
}

// NewQueryUserChainLinkRequest returns a new QueryUserChainLinkRequest instance
func NewQueryUserChainLinkRequest(user string, chainName string, target string) *QueryUserChainLinkRequest {
	return &QueryUserChainLinkRequest{
		User:      user,
		ChainName: chainName,
		Target:    target,
	}
}
