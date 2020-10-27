package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryUserRelationships:
			return queryUserRelationships(ctx, path[1:], req, keeper, legacyQuerierCdc)
		case types.QueryRelationships:
			return queryRelationships(ctx, req, keeper, legacyQuerierCdc)
		case types.QueryUserBlocks:
			return queryUserBlocks(ctx, path[1:], req, keeper, legacyQuerierCdc)
		default:
			return nil, fmt.Errorf("unknown relationships query endpoint")
		}
	}
}

// queryRelationships handles the request of listing all the relationships in the given context
func queryRelationships(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	relationships, err := keeper.GetAllRelationships(ctx)
	if err != nil {
		return nil, err
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &relationships)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryUserRelationships handles the request of listing all the users' stored
func queryUserRelationships(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	relationships, err := keeper.GetUserRelationships(ctx, path[0])
	if err != nil {
		return nil, err
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &relationships)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryUserBlocks handles the request of listing all the users' blocked users
func queryUserBlocks(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	userBlocks, err := keeper.GetUserBlocks(ctx, path[0])
	if err != nil {
		return nil, err
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &userBlocks)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
