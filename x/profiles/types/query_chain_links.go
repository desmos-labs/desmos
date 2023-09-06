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

// GetQueryPrefix returns the prefix used for Query/ChainLinks
func (request *QueryChainLinksRequest) GetQueryPrefix() []byte {
	prefix := ChainLinksPrefix

	switch {
	case request.User != "" && request.ChainName != "" && request.Target != "":
		prefix = ChainLinksStoreKey(request.User, request.ChainName, request.Target)
	case request.User != "":
		prefix = UserChainLinksChainPrefix(request.User, request.ChainName)
	}

	return prefix
}

// NewQueryChainLinkOwnersRequest returns a new QueryChainLinkOwnersRequest instance
func NewQueryChainLinkOwnersRequest(chainName, target string, pageReq *query.PageRequest) *QueryChainLinkOwnersRequest {
	return &QueryChainLinkOwnersRequest{
		ChainName:  chainName,
		Target:     target,
		Pagination: pageReq,
	}
}

// GetQueryPrefix returns the prefix used for Query/ChainLinkOwners
func (request *QueryChainLinkOwnersRequest) GetQueryPrefix() []byte {
	prefix := ChainLinkChainPrefix

	switch {
	case request.ChainName != "" && request.Target != "":
		prefix = ChainLinkChainAddressKey(request.ChainName, request.Target)
	case request.ChainName != "":
		prefix = ChainLinkChainKey(request.ChainName)
	}

	return prefix
}

// NewQueryDefaultExternalAddressesRequest returns a new QueryDefaultExternalAddressesRequest instance
func NewQueryDefaultExternalAddressesRequest(owner, chainName string, pageReq *query.PageRequest) *QueryDefaultExternalAddressesRequest {
	return &QueryDefaultExternalAddressesRequest{
		Owner:      owner,
		ChainName:  chainName,
		Pagination: pageReq,
	}
}

// GetQueryPrefix returns the prefix used for Query/DefaultExternalAddresses
func (request *QueryDefaultExternalAddressesRequest) GetQueryPrefix() []byte {
	prefix := DefaultExternalAddressPrefix

	switch {
	case request.Owner != "" && request.ChainName != "":
		prefix = DefaultExternalAddressKey(request.Owner, request.ChainName)
	case request.Owner != "":
		prefix = OwnerDefaultExternalAddressPrefix(request.Owner)
	}

	return prefix
}
