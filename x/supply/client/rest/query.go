package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	resttypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		fmt.Sprintf("/supply/circulating-supply/{%s}", DenomParam),
		queryCirculatingSupplyFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		fmt.Sprintf("/supply/total-supply/{%s}", DenomParam),
		queryTotalSupplyFn(clientCtx),
	).Methods("GET")
}

func queryCirculatingSupplyFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientCtx, ok := resttypes.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCirculatingSupply)

		var req SupplyReq
		if !resttypes.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		divider, err := strconv.ParseUint(req.DividerExponent, 10, 0)
		if err != nil {
			resttypes.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		vars := mux.Vars(r)
		params := types.NewQueryCirculatingSupplyRequest(vars[DenomParam], divider)
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
		clientCtx, ok := resttypes.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalSupply)

		var req SupplyReq
		if !resttypes.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		divider, err := strconv.ParseUint(req.DividerExponent, 10, 0)
		if err != nil {
			resttypes.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		vars := mux.Vars(r)
		params := types.NewQueryTotalSupplyRequest(vars[DenomParam], divider)
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
