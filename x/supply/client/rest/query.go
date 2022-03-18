package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	resttypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		"/circulating-supply/{denom}",
		queryCirculatingSupplyFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/total-supply/{denom}",
		queryTotalSupplyFn(clientCtx),
	).Methods("GET")
}

func queryCirculatingSupplyFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		clientCtx, ok := resttypes.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		denom := strings.TrimSpace(vars["denom"])
		if denom == "" {
			resttypes.WriteErrorResponse(w, http.StatusBadRequest, "Invalid empty denom string")
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCirculatingSupply)
		params := types.NewQueryCirculatingSupplyRequest(denom)
		bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if resttypes.CheckBadRequestError(w, err) {
			return
		}

		res, height, err := clientCtx.QueryWithData(route, bz)
		if resttypes.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		resttypes.PostProcessResponse(w, clientCtx, res)
	}
}

func queryTotalSupplyFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		clientCtx, ok := resttypes.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		denom := strings.TrimSpace(vars["denom"])
		if denom == "" {
			resttypes.WriteErrorResponse(w, http.StatusBadRequest, "Invalid empty denom string")
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalSupply)
		params := types.QueryTotalSupplyRequest{Denom: denom}
		bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if resttypes.CheckBadRequestError(w, err) {
			return
		}

		res, height, err := clientCtx.QueryWithData(route, bz)
		if resttypes.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		resttypes.PostProcessResponse(w, clientCtx, res)
	}
}
