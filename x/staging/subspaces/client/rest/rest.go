package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

const (
	SubspaceID = "subspace_id"
)

// RegisterRestRouters - Central function to define routes that get registered by the main application
func RegisterRestRoutes(cliCtx client.Context, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRouters(cliCtx, r)
}

// CreateSubspaceReq defines the properties of a create subspace request's body
type CreateSubspaceReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Name    string       `json:"name"`
	Open    bool         `json:"open"`
}

// EditSubspaceReq defines the properties of a edit subspace request's body
type EditSubspaceReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	NewName  string       `json:"new_name"`
	NewOwner string       `json:"new_owner"`
}

// CommonSubspaceAdminReq defines the properties request's body of add/remove admin and register/block users
type CommonSubspaceReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Address string       `json:"address"`
}
