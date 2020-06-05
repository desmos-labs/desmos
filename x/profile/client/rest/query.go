package rest

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/profiles/{address_or_moniker}", queryProfileHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/profiles", queryProfilesHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/profiles/params", queryProfilesParamsHandlerFn(cliCtx)).Methods("GET")
}

// HTTP request handler to query a single profile based on its moniker
func queryProfileHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address_or_moniker"]

		route := fmt.Sprintf("custom/%s/%s/%s", models.QuerierRoute, models.QueryProfile, address)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query list of profiles
func queryProfilesHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", models.QuerierRoute, models.QueryProfiles)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// HTTP request handler to query list of profiles' module params
func queryProfilesParamsHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", models.QuerierRoute, models.QueryParams)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
