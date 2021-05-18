package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/gorilla/mux"
)

func registerTxRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}", SubspaceID),
		createSubspaceHandler(clientCtx)).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/subspaces/edit/{%s}", SubspaceID),
		editSubspaceHandler(clientCtx)).Methods("PUT")

	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}/add-admin", SubspaceID),
		addSubspaceAdminHandler(clientCtx)).Methods("POST")

	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}/remove-admin", SubspaceID),
		removeSubspaceAdminHandler(clientCtx)).Methods("DELETE")

	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}/register-user", SubspaceID),
		registerUserHandler(clientCtx)).Methods("PUT")

	r.HandleFunc(fmt.Sprintf("/subspaces/{%s}/block-user", SubspaceID),
		blockUserHandler(clientCtx)).Methods("PUT")
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

		msg := types.NewMsgCreateSubspace(subspaceID, req.Name, req.BaseReq.From, req.Open)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func editSubspaceHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var req EditSubspaceReq

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		subspaceID := vars[SubspaceID]

		msg := types.NewMsgEditSubspace(subspaceID, req.NewOwner, req.NewName, req.BaseReq.From)
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

func registerUserHandler(clientCtx client.Context) http.HandlerFunc {
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

		msg := types.NewMsgRegisterUser(req.Address, subspaceID, req.BaseReq.From)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func blockUserHandler(clientCtx client.Context) http.HandlerFunc {
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

		msg := types.NewMsgBlockUser(req.Address, subspaceID, req.BaseReq.From)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
