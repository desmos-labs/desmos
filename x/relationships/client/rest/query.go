package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/relationships", queryRelationships(cliCtx)).Methods("GET")
	r.HandleFunc("/relationships/{address}", queryUserRelationships(cliCtx)).Methods("GET")
	r.HandleFunc("/blacklist/{address}", queryUserBlocks(cliCtx)).Methods("GET")
}

// HTTP request handler to query list of user's relationships
func queryUserRelationships(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryRelationships, address)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query list of all relationships
func queryRelationships(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRelationships)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryUserBlocks(cliCtx context.CLIContext) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryUserBlocks, address)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
