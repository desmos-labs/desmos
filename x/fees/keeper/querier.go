package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/fees/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParams:
			return queryParams(ctx, req, keeper)
		default:
			return nil, fmt.Errorf("unknown fees endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	params := keeper.GetParams(ctx)

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &params)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
