package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	// this line is used by starport scaffolding # 1
)

// RegisterRoutes registers links-related REST handlers to a router
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRoutes(cliCtx, r)
}

// CreateIBCAccountLinkReq defines the properties of the request's body of creating ibc link with same key
type CreateIBCAccountLinkReq struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Port         string       `json:"port"`
	ChannelId    string       `json:"channel_id"`
	SourcePubKey string       `json:"source_pub_key"`
	Signature    string       `json:"signature"`
}

// CreateIBCAccountConnectionReq defines the properties of the request's body of creating ibc link with different keys
type CreateIBCAccountConnectionReq struct {
	BaseReq              rest.BaseReq `json:"base_req"`
	Port                 string       `json:"port"`
	ChannelId            string       `json:"channel_id"`
	SourcePubKey         string       `json:"source_pub_key"`
	SourceSignature      string       `json:"source_signature"`
	DestinationAddress   string       `json:"destination_address"`
	DestinationSignature string       `json:"destination_signature"`
}
