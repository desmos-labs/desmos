package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc("/sessions", createSessionHander(cliCtx)).Methods("POST")
	r.HandleFunc("/sessions/{sessionID}", getSessionHandler(cliCtx, storeName)).Methods("GET")
}

// --------------------------------------------------------------------------------------
// Tx Handler

type createSessionReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Messenger     string       `json:"messager"`
	Namespace     string       `json:"namespace"`
	Owner         string       `json:"owner"`
	ExternalOwner string       `json:"external_owner"`
	Pubkey        string       `json:"pubkey"`
	Signature     string       `json:"signature"`
}

func createSessionHander(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSessionReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the session
		msg := types.NewMsgCreateSession(addr, req.Namespace, req.ExternalOwner, req.Pubkey, req.Signature)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func getSessionHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sessionID := vars["sessionID"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/session/%s", storeName, sessionID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
