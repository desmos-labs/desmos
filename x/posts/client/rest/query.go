package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	RestSortBy         = "sort_by"
	RestSortOrder      = "sort_order"
	RestParentID       = "parent_id"
	RestCreationTime   = "creation_time"
	RestAllowsComments = "allows_comments"
	RestSubspace       = "subspace"
	RestCreator        = "creator"
	RestHashtags       = "hashtags"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/posts/{postID}", queryPostHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/posts", queryPostsWithParameterHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/posts/{postID}/poll-answers", queryPostPollAnswersHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/registeredReactions", queryRegisteredReactions(cliCtx)).Methods("GET")
	r.HandleFunc("/posts/params", queryPostsParamsHandlerFn(cliCtx)).Methods("GET")
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

		// Default params
		params := types.DefaultQueryPostsParams(page, limit)

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		if v := r.URL.Query().Get(RestSortBy); len(v) != 0 {
			params.SortBy = v
		}

		if v := r.URL.Query().Get(RestSortOrder); len(v) != 0 {
			params.SortOrder = v
		}

		if v := r.URL.Query().Get(RestParentID); len(v) != 0 {
			parentID := types.PostID(v)
			if !parentID.Valid() {
				rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid postID: %s", parentID))
				return
			}
			params.ParentID = &parentID
		}

		if v := r.URL.Query().Get(RestCreationTime); len(v) != 0 {
			parsedTime, err := time.Parse(time.RFC3339, v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			params.CreationTime = &parsedTime
		}

		if v := r.URL.Query().Get(RestAllowsComments); len(v) != 0 {
			parsedAllowsComments, err := strconv.ParseBool(v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			params.AllowsComments = &parsedAllowsComments
		}

		if v := r.URL.Query().Get(RestSubspace); len(v) != 0 {
			params.Subspace = v
		}

		if v := r.URL.Query().Get(RestCreator); len(v) != 0 {
			creatorAddr, err := sdk.AccAddressFromBech32(v)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			params.Creator = creatorAddr
		}

		if v := r.URL.Query().Get(RestHashtags); len(v) != 0 {
			params.Hashtags = strings.Split(v, ",")
		}

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

func queryPostPollAnswersHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		postID := vars["postID"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryPollAnswers, postID)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryRegisteredReactions(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRegisteredReactions)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query list of posts' module params
func queryPostsParamsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
