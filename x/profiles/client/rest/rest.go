package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

const (
	ParamsAddress = "address"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRoutes(cliCtx, r)
}

// SaveProfileReq defines the properties of a profile save request's body
type SaveProfileReq struct {
	BaseReq  rest.BaseReq   `json:"base_req"`
	DTag     string         `json:"dtag"`
	Moniker  string         `json:"moniker,omitempty"`
	Bio      string         `json:"bio,omitempty"`
	Pictures types.Pictures `json:"pictures,omitempty"`
}

// DeleteProfileReq defines the properties of a profile deletion request body
type DeleteProfileReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
}

// TransferDTagReq defines the properties of a DTag transferring request body
type TransferDTagReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Receiver string       `json:"receiver"`
}

// AcceptDTagTransferReq contains the data that should be sent to accept a DTag transfer request
type AcceptDTagTransferReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	NewDTag string       `json:"new_dtag"`
}
