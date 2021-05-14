package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

const (
	SubspaceID = "subspace_id"
)

// RegisterRouters - Central function to define routes that get registered by the main application
func RegisterRouters(cliCtx client.Context, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRouters(cliCtx, r)
}

// CreateSubspaceReq defines the properties of a create subspace request's body
type CreateSubspaceReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Name    string       `json:"name"`
}

// CommonSubspaceAdminReq defines the properties request's body of add/remove admin and enable/disable user posts
type CommonSubspaceReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Address string       `json:"address"`
}
