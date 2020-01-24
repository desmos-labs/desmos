package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/gorilla/mux"
)

// REST Variable names
// nolint
const (
	RestParentID       = "parent_id"
	RestCreationTime   = "creation_time"
	RestAllowsComments = "allows_comments"
	RestSubspace       = "subspace"
	RestCreator        = "creator"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/posts/{postID}", queryPostHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/posts", queryPostsWithParameterHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/poll-answers/{postID}/{user-address}", queryPollPostUserAnswersHandleFn(cliCtx)).Methods("GET")
	r.HandleFunc("/poll-answers-amount/{postID}", queryPollPostAnswersAmountHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/poll-answer-votes/{postID}/{answerID}", queryPollAnswerVotesHandlerFn(cliCtx)).Methods("GET")
}

// HTTP request handler to query a single post based on its ID
func queryPostHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["postID"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryPost, postID)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query list of posts
func queryPostsWithParameterHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		var (
			parentID       *types.PostID
			creationTime   *time.Time
			allowsComments *bool
			subspace       string
			creatorAddr    sdk.AccAddress
		)

		if v := r.URL.Query().Get(RestParentID); len(v) != 0 {
			parsedParentID, err := types.ParsePostID(v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			parentID = &parsedParentID
		}

		if v := r.URL.Query().Get(RestCreationTime); len(v) != 0 {
			parsedTime, err := time.Parse(time.RFC3339, v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			creationTime = &parsedTime
		}

		if v := r.URL.Query().Get(RestAllowsComments); len(v) != 0 {
			parsedAllowsComments, err := strconv.ParseBool(v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			allowsComments = &parsedAllowsComments
		}

		if v := r.URL.Query().Get(RestSubspace); len(v) != 0 {
			subspace = v
		}

		if v := r.URL.Query().Get(RestCreator); len(v) != 0 {
			creatorAddr, err = sdk.AccAddressFromBech32(v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		params := types.NewQueryPostsParams(page, limit, parentID, creationTime, allowsComments, subspace, creatorAddr)
		bz, err := cliCtx.Codec.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryPosts)
		res, height, err := cliCtx.QueryWithData(route, bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query poll post's answer made by a user
func queryPollPostUserAnswersHandleFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["postID"]
		answerer, err := sdk.AccAddressFromBech32(vars["user-address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s/%s/%s", types.QuerierRoute, types.QueryPollUserAnswer, postID, answerer)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query poll post's answers amount
func queryPollPostAnswersAmountHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["postID"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryAnswersAmount, postID)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query answer's votes amount
func queryPollAnswerVotesHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["postID"]
		answerID := vars["answerID"]

		route := fmt.Sprintf("custom/%s/%s/%s/%s", types.QuerierRoute, types.QueryAnswerVotes, postID, answerID)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
