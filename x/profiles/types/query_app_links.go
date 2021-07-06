package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryUserApplicationLinksRequest returns a new QueryUserApplicationLinksRequest instance
func NewQueryUserApplicationLinksRequest(user string, pageReq *query.PageRequest) *QueryUserApplicationLinksRequest {
	return &QueryUserApplicationLinksRequest{
		User:       user,
		Pagination: pageReq,
	}
}
