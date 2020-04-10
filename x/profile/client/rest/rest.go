package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRoutes(cliCtx, r)
}

// CreateProfileReq defines the properties of a profile creation request's body
type CreateProfileReq struct {
	BaseReq  rest.BaseReq    `json:"base_req"`
	Moniker  string          `json:"moniker"`
	Name     *string         `json:"name,omitempty"`
	Surname  *string         `json:"surname,omitempty"`
	Bio      *string         `json:"bio,omitempty"`
	Pictures *types.Pictures `json:"pictures,omitempty"`
}

// EditProfileReq defines the properties of a profile edit request's body
type EditProfileReq struct {
	BaseReq    rest.BaseReq    `json:"base_req"`
	NewMoniker *string         `json:"new_moniker"`
	Name       *string         `json:"name,omitempty"`
	Surname    *string         `json:"surname,omitempty"`
	Bio        *string         `json:"bio,omitempty"`
	Pictures   *types.Pictures `json:"pictures,omitempty"`
}

// Delete defines the properties of a profile deletion request's body
type DeleteProfileReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
}
