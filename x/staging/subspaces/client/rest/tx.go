package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/gorilla/mux"
	"net/http"
)

func registerTxRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}", SubspaceID),
		createSubspaceHandler(clientCtx)).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}/add-admin", SubspaceID),
		addSubspaceAdminHandler(clientCtx)).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}/remove-admin", SubspaceID),
		removeSubspaceAdminHandler(clientCtx)).Methods("DELETE")

	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}/enable-user-posts", SubspaceID),
		enablePostsForUserHandler(clientCtx)).Methods("PUT")

	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}/disable-user-posts", SubspaceID),
		disablePostsForUserHandler(clientCtx)).Methods("PUT")
}

func createSubspaceHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var req CreateSubspaceReq

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		subspaceID := vars[SubspaceID]

		msg := types.NewMsgCreateSubspace(subspaceID, req.BaseReq.From)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func addSubspaceAdminHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var req CommonSubspaceReq

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		subspaceID := vars[SubspaceID]

		msg := types.NewMsgAddAdmin(subspaceID, req.Address, req.BaseReq.From)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func removeSubspaceAdminHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var req CommonSubspaceReq

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		subspaceID := vars[SubspaceID]

		msg := types.NewMsgRemoveAdmin(subspaceID, req.Address, req.BaseReq.From)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func enablePostsForUserHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var req CommonSubspaceReq

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		subspaceID := vars[SubspaceID]

		msg := types.NewMsgEnableUserPosts(req.Address, subspaceID, req.BaseReq.From)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func disablePostsForUserHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var req CommonSubspaceReq

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		subspaceID := vars[SubspaceID]

		msg := types.NewMsgDisableUserPosts(req.Address, subspaceID, req.BaseReq.From)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
