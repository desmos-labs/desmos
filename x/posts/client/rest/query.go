package rest

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/desmos-labs/desmos/x/posts/client/cli"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/posts/parameters", queryPostsParamsHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/posts/{postID}", queryPostHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/posts", queryPostsWithParameterHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/posts/{postID}/poll-answers", queryPostPollAnswersHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/registered-reactions", queryRegisteredReactions(cliCtx)).Methods("GET")
}

// HTTP request handler to query a single post based on its ID
func queryPostHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["postID"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryPost, postID)
		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query a list of posts
func queryPostsWithParameterHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, limit, err := rest.ParseHTTPArgsWithLimit(r, 0)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// Default params
		params := cli.DefaultQueryPostsRequest(uint64(page), uint64(limit))

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		if v := r.URL.Query().Get(ParamSortBy); len(v) != 0 {
			params.SortBy = v
		}

		if v := r.URL.Query().Get(ParamSortOrder); len(v) != 0 {
			params.SortOrder = v
		}

		if v := r.URL.Query().Get(ParamParentID); len(v) != 0 {
			parentID := v
			if !types.IsValidPostID(parentID) {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid postID: %s", parentID))
				return
			}
			params.ParentID = parentID
		}

		if v := r.URL.Query().Get(ParamCreationTime); len(v) != 0 {
			parsedTime, err := time.Parse(time.RFC3339, v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			params.CreationTime = &parsedTime
		}

		if v := r.URL.Query().Get(ParamSubspace); len(v) != 0 {
			params.Subspace = v
		}

		if v := r.URL.Query().Get(ParamCreator); len(v) != 0 {
			creatorAddr, err := sdk.AccAddressFromBech32(v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			params.Creator = creatorAddr.String()
		}

		if v := r.URL.Query().Get(ParamHashtags); len(v) != 0 {
			params.Hashtags = strings.Split(v, ",")
		}

		bz, err := codec.MarshalJSONIndent(cliCtx.LegacyAmino, params)
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

// HTTP request handler to query the answers of a poll
func queryPostPollAnswersHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["postID"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryPollAnswers, postID)
		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query the registered reactions
func queryRegisteredReactions(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRegisteredReactions)
		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query list of posts' module params
func queryPostsParamsHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)
		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
