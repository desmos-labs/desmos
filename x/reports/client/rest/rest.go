package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

const (
	ParamPostID = "post_id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	registerTxRoutes(cliCtx, r)
	registerQueryRoutes(cliCtx, r)
}

type ReportPostReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	ReportType    string       `json:"report_type"`
	ReportMessage string       `json:"report_message"`
}
