package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/kwunyeung/desmos/x/dwitter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/gorilla/mux"
)

const (
	restName = "name"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/posts", storeName), createPostHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/posts/{postID}", storeName), getPostHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/like", storeName), likePostHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/like/{likeID}", storeName), getLikeHandler(cliCtx, storeName)).Methods("GET")
}

// --------------------------------------------------------------------------------------
// Tx Handler

type createPostReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Message string       `json:"message"`
	Owner   string       `json:"owner"`
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
		msg := types.NewMsgCreatePost(req.Message, time.Now(), addr)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type addLikeReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	PostID  string       `json:"post_id"`
	Owner   string       `json:"owner"`
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
		msg := types.NewMsgLike(req.PostID, time.Now(), addr)
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

// func namesHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/names", storeName), nil)
// 		if err != nil {
// 			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
// 			return
// 		}
// 		rest.PostProcessResponse(w, cliCtx, res)
// 	}
// }
