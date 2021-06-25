package types

import (
	query "github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryUserApplicationLinksRequest returns a new QueryUserApplicationLinksRequest instance
func NewQueryUserApplicationLinksRequest(user string, pageReq *query.PageRequest) *QueryUserApplicationLinksRequest {
	return &QueryUserApplicationLinksRequest{
		User:       user,
		Pagination: pageReq,
	}
}

// NewQueryUserApplicationLinkRequest returns a new QueryUserApplicationLinkRequest instance
func NewQueryUserApplicationLinkRequest(user, application, username string) *QueryUserApplicationLinkRequest {
	return &QueryUserApplicationLinkRequest{
		User:        user,
		Application: application,
		Username:    username,
	}
}

// NewQueryApplicationLinkByClientIDRequest returns a new QueryApplicationLinkByClientIDRequest instance
func NewQueryApplicationLinkByClientIDRequest(clientID string) *QueryApplicationLinkByClientIDRequest {
	return &QueryApplicationLinkByClientIDRequest{
		ClientId: clientID,
	}
}
