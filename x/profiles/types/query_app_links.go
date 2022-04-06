package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryApplicationLinksRequest returns a new QueryApplicationLinksRequest instance
func NewQueryApplicationLinksRequest(
	user, application, username string, pageReq *query.PageRequest,
) *QueryApplicationLinksRequest {
	return &QueryApplicationLinksRequest{
		User:        user,
		Application: application,
		Username:    username,
		Pagination:  pageReq,
	}
}

// NewQueryApplicationLinkByClientIDRequest returns a new QueryApplicationLinkByClientIDRequest instance
func NewQueryApplicationLinkByClientIDRequest(clientID string) *QueryApplicationLinkByClientIDRequest {
	return &QueryApplicationLinkByClientIDRequest{
		ClientId: clientID,
	}
}

// NewQueryApplicationLinkOwnersRequest returns a new QueryApplicationLinkOwnersRequest instance
func NewQueryApplicationLinkOwnersRequest(
	application, username string, pageReq *query.PageRequest,
) *QueryApplicationLinkOwnersRequest {
	return &QueryApplicationLinkOwnersRequest{
		Application: application,
		Username:    username,
		Pagination:  pageReq,
	}
}
