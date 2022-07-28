package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	resttypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/desmos-labs/desmos/v4/x/supply/types"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/supply/total/{%s}", DenomParam), queryTotalSupplyFn(clientCtx)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/supply/circulating/{%s}", DenomParam), queryCirculatingSupplyFn(clientCtx)).Methods("GET")
}

func queryTotalSupplyFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientCtx, ok := resttypes.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		dividerStr := r.URL.Query().Get(DividerExponentParam)
		if len(dividerStr) == 0 {
			dividerStr = "0"
		}

		divider, ok := resttypes.ParseUint64OrReturnBadRequest(w, dividerStr)
		if !ok {
			return
		}

		vars := mux.Vars(r)
		params := types.NewQueryTotalRequest(vars[DenomParam], divider)
		bz, err := clientCtx.Codec.Marshal(params)
		if resttypes.CheckBadRequestError(w, err) {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalSupply)
		res, height, err := clientCtx.QueryWithData(route, bz)
		if resttypes.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		resttypes.PostProcessResponseBare(w, clientCtx, res)
	}
}

func queryCirculatingSupplyFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientCtx, ok := resttypes.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		dividerStr := r.URL.Query().Get(DividerExponentParam)
		if len(dividerStr) == 0 {
			dividerStr = "0"
		}

		divider, ok := resttypes.ParseUint64OrReturnBadRequest(w, dividerStr)
		if !ok {
			return
		}

		vars := mux.Vars(r)
		params := types.NewQueryCirculatingRequest(vars[DenomParam], divider)
		bz, err := clientCtx.Codec.Marshal(params)
		if resttypes.CheckBadRequestError(w, err) {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCirculatingSupply)
		res, height, err := clientCtx.QueryWithData(route, bz)
		if resttypes.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		resttypes.PostProcessResponseBare(w, clientCtx, res)
	}
}
