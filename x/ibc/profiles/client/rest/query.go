package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/links/{source_address}",
		queryLinkHandler(cliCtx)).Methods("Get")
}

// HTTP request handler to query single link based on its source address
func queryLinkHandler(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["source_address"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryLink, address)
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
