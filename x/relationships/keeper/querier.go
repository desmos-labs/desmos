package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryUserRelationships:
			return queryUserRelationships(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case types.QueryUserBlocks:
			return queryUserBlocks(ctx, path[1:], req, keeper, legacyQuerierCdc)
		default:
			return nil, fmt.Errorf("unknown relationships query endpoint")
		}
	}
}

// queryUserRelationships handles the request of listing all the users' stored relationships
func queryUserRelationships(
	ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	relationships := keeper.GetUserRelationships(ctx, path[0])

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &relationships)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryUserBlocks handles the request of listing all the users' blocked users
func queryUserBlocks(
	ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	userBlocks := keeper.GetUserBlocks(ctx, path[0])

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &userBlocks)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
