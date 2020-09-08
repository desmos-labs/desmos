package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRoutes(cliCtx, r)
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

type UserUnblockReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Blocked  string       `json:"blocked"`
	Subspace string       `json:"subspace"`
}
