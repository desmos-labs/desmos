package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

const (
	ParamAddress = "address"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

// CommonRelationshipReq defines the properties of a create relationship operation request's body
type CommonRelationshipReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Receiver string       `json:"receiver"`
	Subspace string       `json:"subspace"`
}

// UserBlockReq defines the properties of a block user operation request's body
type UserBlockReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Blocked  string       `json:"blocked"`
	Reason   string       `json:"reason,omitempty"`
	Subspace string       `json:"subspace"`
}

// UserUnblockReq defines the properties of an unblock user operation request's body
type UserUnblockReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Blocked  string       `json:"blocked"`
	Subspace string       `json:"subspace"`
}
