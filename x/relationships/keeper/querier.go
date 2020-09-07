package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryUserRelationships:
			return queryUserRelationships(ctx, path[1:], req, keeper)
		case types.QueryRelationships:
			return queryRelationships(ctx, req, keeper)
		case types.QueryUserBlocks:
			return queryUserBlocks(ctx, path[1:], req, keeper)
		default:
			return nil, fmt.Errorf("unknown relationships query endpoint")
		}
	}
}

// queryRelationships handles the request of listing all the relationships in the given context
func queryRelationships(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	relationships := keeper.GetUsersRelationships(ctx)
	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &relationships)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryUserRelationships handles the request of listing all the users' storedRelationships
func queryUserRelationships(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	user, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid bech32 address: %s", path[0]))
	}

	relationships := types.NewRelationshipResponse(keeper.GetUserRelationships(ctx, user))

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &relationships)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryUserBlocks handles the request of listing all the users' blocked users
func queryUserBlocks(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	user, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid bech32 address: %s", path[0]))
	}

	userBlocks := keeper.GetUserBlocks(ctx, user)

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &userBlocks)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
