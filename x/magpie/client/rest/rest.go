package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

// REST Variable names
// nolint
const (
	RestParamsID = "session_id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

// --------------------------------------------------------------------------------------
// Tx Handler

type CreateSessionReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Messenger     string       `json:"messager"`
	Namespace     string       `json:"namespace"`
	ExternalOwner string       `json:"external_owner"`
	PublicKey     string       `json:"pubkey"`
	Signature     string       `json:"signature"`
}
