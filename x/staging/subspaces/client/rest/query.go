package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/x/staging/subspaces/client/cli"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/gorilla/mux"
)

func registerQueryRouters(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/subspaces/{subspace_id}",
		querySubspaceHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/subspaces",
		querySubspacesHandlerFn(cliCtx)).Methods("GET")
}

func querySubspacesHandlerFn(cliCtx client.Context) http.HandlerFunc {
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

		queryParams := cli.DefaultQuerySubspacesRequest(uint64(page), uint64(limit))
		bz, err := codec.MarshalJSONIndent(cliCtx.LegacyAmino, queryParams)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySubspaces)
		res, height, err := cliCtx.QueryWithData(route, bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func querySubspaceHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		SubspaceID := vars[SubspaceID]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QuerySubspace, SubspaceID)
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
