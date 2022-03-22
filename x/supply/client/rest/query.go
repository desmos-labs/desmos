package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	resttypes "github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		"/circulating-supply/{denom}/{multiplier}",
		queryCirculatingSupplyFn(clientCtx),
	).Methods("GET")
	r.HandleFunc(
		"/total-supply/{denom}/{multiplier}",
		queryTotalSupplyFn(clientCtx),
	).Methods("GET")
}

func queryCirculatingSupplyFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientCtx, ok := resttypes.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		denom, multiplier, err := ParseQueryParams(mux.Vars(r))
		if err != nil {
			resttypes.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCirculatingSupply)
		params := types.NewQueryCirculatingSupplyRequest(denom, multiplier)
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

		denom, multiplier, err := ParseQueryParams(mux.Vars(r))
		if err != nil {
			resttypes.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalSupply)
		params := types.NewQueryTotalSupplyRequest(denom, multiplier)
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

// ParseQueryParams parses the query parameters of the given request
func ParseQueryParams(vars map[string]string) (string, int64, error) {
	denom := strings.TrimSpace(vars["denom"])
	if denom == "" {
		return "", 1, fmt.Errorf("invalid empty denom string")
	}

	multiplier := int64(1)
	parsedMultiplier := strings.TrimSpace(vars["multiplier"])
	if parsedMultiplier != "0" && parsedMultiplier != "" {
		var err error
		multiplier, err = strconv.ParseInt(parsedMultiplier, 10, 0)
		if err != nil {
			return "", 1, fmt.Errorf("invalid multiplier factor")
		}
	}

	return denom, multiplier, nil
}
