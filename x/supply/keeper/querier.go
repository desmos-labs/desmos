package keeper

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v4/x/supply/types"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryTotalSupply:
			return queryTotalSupply(ctx, req, k)
		case types.QueryCirculatingSupply:
			return queryCirculatingSupply(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}
	}
}

// queryTotalSupply queries the total supply of the given params.Denom
func queryTotalSupply(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var request types.QueryTotalRequest
	err := k.cdc.Unmarshal(req.Data, &request)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	res, err := k.Total(sdk.WrapSDKContext(ctx), &request)
	if err != nil {
		return nil, err
	}

	supply, err := res.TotalSupply.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return supply, nil
}

// queryCirculatingSupply queries the current circulating supply of the given params.Denom
func queryCirculatingSupply(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var request types.QueryCirculatingRequest
	err := k.cdc.Unmarshal(req.Data, &request)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	res, err := k.Circulating(sdk.WrapSDKContext(ctx), &request)
	if err != nil {
		return nil, err
	}

	supply, err := res.CirculatingSupply.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return supply, nil
}
