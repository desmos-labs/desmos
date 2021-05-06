package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/desmos-labs/desmos/x/links/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

func registerTxRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/links/ibclink",
		createIBCAccountLinkHandler(cliCtx)).Methods("Post")
	r.HandleFunc("/links/ibconnection",
		createIBCAccountConnectionHandler(cliCtx)).Methods("Post")
}

func createIBCAccountLinkHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateIBCAccountLinkReq

		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		pubKey := req.SourcePubKey

		msg := types.NewMsgCreateIBCAccountLink(
			req.Port,
			req.ChannelID,
			types.NewIBCAccountLinkPacketData(
				sdk.GetConfig().GetBech32AccountAddrPrefix(),
				addr.String(),
				pubKey,
				req.Signature,
			),
			1000,
		)

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}

func createIBCAccountConnectionHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateIBCAccountConnectionReq

		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		pubKey := req.SourcePubKey

		msg := types.NewMsgCreateIBCAccountConnection(
			req.Port,
			req.ChannelID,
			types.NewIBCAccountConnectionPacketData(
				sdk.GetConfig().GetBech32AccountAddrPrefix(),
				addr.String(),
				pubKey,
				req.DestinationAddress,
				req.SourceSignature,
				req.DestinationSignature,
			),
			1000,
		)

		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}
