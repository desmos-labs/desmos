package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/gorilla/mux"
	"net/http"
)

func registerQueryRouters(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/subspaces/{subspace_id}",
		querySubspaceHandlerFn(cliCtx)).Methods("GET")
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
