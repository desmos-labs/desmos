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

// GetQueryPrefix returns the prefix used for Query/ApplicationLinks
func (request *QueryApplicationLinksRequest) GetQueryPrefix() []byte {
	prefix := ApplicationLinkPrefix

	switch {
	case request.User != "" && request.Application != "" && request.Username != "":
		prefix = UserApplicationLinkKey(request.User, request.Application, request.Username)
	case request.User != "":
		prefix = UserApplicationLinksApplicationPrefix(request.User, request.Application)
	}

	return prefix
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

// GetQueryPrefix returns the prefix used for Query/ApplicationLinkOwners
func (request *QueryApplicationLinkOwnersRequest) GetQueryPrefix() []byte {
	prefix := ApplicationLinkAppPrefix

	switch {
	case request.Application != "" && request.Username != "":
		prefix = ApplicationLinkAppUsernameKey(request.Application, request.Username)
	case request.Application != "":
		prefix = ApplicationLinkAppKey(request.Application)
	}

	return prefix
}
