package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/v3/x/supply/types"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryCirculatingSupply:
			return queryCirculatingSupply(ctx, req, k)
		case types.QueryTotalSupply:
			return queryTotalSupply(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

// queryCirculatingSupply queries the current circulating supply of the given params.Denom
func queryCirculatingSupply(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var request types.QueryCirculatingSupplyRequest
	err := k.cdc.Unmarshal(req.Data, &request)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	res, err := k.CirculatingSupply(sdk.WrapSDKContext(ctx), &request)
	if err != nil {
		return nil, err
	}

	supply, err := res.CirculatingSupply.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return supply, nil
}

// queryTotalSupply queries the total supply of the given params.Denom
func queryTotalSupply(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var request types.QueryTotalSupplyRequest
	err := k.cdc.Unmarshal(req.Data, &request)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	res, err := k.TotalSupply(sdk.WrapSDKContext(ctx), &request)
	if err != nil {
		return nil, err
	}

	supply, err := res.TotalSupply.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return supply, nil
}
