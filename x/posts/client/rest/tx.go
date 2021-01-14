package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func registerTxRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/posts",
		createPostHandler(cliCtx)).Methods("POST")

	r.HandleFunc("/posts/reactions",
		addReactionToPostHandler(cliCtx)).Methods("POST")

	r.HandleFunc("/posts/reactions",
		removeReactionToPostHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc(fmt.Sprintf("/posts/{%s}/answers", ParamPostID),
		addAnswerToPostPollHandler(cliCtx)).Methods("POST")

	r.HandleFunc("/registeredReactions",
		registerReactionHandler(cliCtx)).Methods("POST")
}

func createPostHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreatePostReq

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
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

		parentID := req.ParentID
		if !types.IsValidPostID(parentID) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid postID: %s", parentID))
			return
		}

		msg := types.NewMsgCreatePost(
			req.Message,
			parentID,
			req.AllowsComments,
			req.Subspace,
			req.OptionalData,
			addr.String(),
			req.Medias,
			req.PollData,
		)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func addReactionToPostHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddReactionReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
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

		postID := req.PostID
		if !types.IsValidPostID(postID) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid postID: %s", postID))
			return
		}

		msg := types.NewMsgAddPostReaction(postID, req.Reaction, addr.String())
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func removeReactionToPostHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RemoveReactionReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
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

		postID := req.PostID
		if !types.IsValidPostID(postID) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid postID: %s", postID))
			return
		}

		msg := types.NewMsgRemovePostReaction(postID, addr.String(), req.Reaction)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func addAnswerToPostPollHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var req AnswerPollPostReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
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

		postID := vars[ParamPostID]
		if !types.IsValidPostID(postID) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid postID: %s", postID))
			return
		}

		msg := types.NewMsgAnswerPoll(postID, req.Answers, addr.String())
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func registerReactionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterReactionReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		creator, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgRegisterReaction(creator.String(), req.Shortcode, req.Value, req.Subspace)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
