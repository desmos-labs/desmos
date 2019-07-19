package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/kwunyeung/desmos/x/magpie/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/gorilla/mux"
)

const (
	restName = "magpie"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc("/magpie/posts", createPostHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/magpie/posts/{postID}", getPostHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc("/magpie/like", likePostHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/magpie/like/{likeID}", getLikeHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc("/magpie/session", createSessionHander(cliCtx)).Methods("POST")
	r.HandleFunc("/magpie/session/{sessionID}", getSessionHandler(cliCtx, storeName)).Methods("GET")
}

// --------------------------------------------------------------------------------------
// Tx Handler

type createPostReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Message       string       `json:"message"`
	ParentID      string       `json:"parent_id"`
	Owner         string       `json:"owner"`
	Namespace     string       `json:"namespace"`
	ExternalOwner string       `json:"external_owner"`
}

func createPostHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPostReq

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

		// coins, err := sdk.ParseCoins(req.Amount)
		// if err != nil {
		// 	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }

		// create the message
		msg := types.NewMsgCreatePost(req.Message, req.ParentID, time.Now(), addr, req.Namespace, req.ExternalOwner)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type addLikeReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	PostID        string       `json:"post_id"`
	Owner         string       `json:"owner"`
	Namespace     string       `json:"namespace"`
	ExternalOwner string       `json:"external_owner"`
}

func likePostHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addLikeReq
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

		// create the message
		msg := types.NewMsgLike(req.PostID, time.Now(), addr, req.Namespace, req.ExternalOwner)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

//--------------------------------------------------------------------------------------
// Query Handlers

func getPostHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["postID"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/post/%s", storeName, postID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getLikeHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		likeID := vars["likeID"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/like/%s", storeName, likeID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

type createSessionReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Messenger     string       `json:"messager"`
	Namespace     string       `json:"namespace"`
	Owner         string       `json:"owner"`
	ExternalOwner string       `json:"external_owner"`
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
		msg := types.NewMsgCreateSession(time.Now(), addr, req.Namespace, req.ExternalOwner, req.Signature)
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
