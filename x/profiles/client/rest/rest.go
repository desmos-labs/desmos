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
	Username string         `json:"username,omitempty"`
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
