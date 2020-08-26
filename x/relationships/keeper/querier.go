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
		case types.QueryRelationships:
			return queryUserRelationships(ctx, path[1:], req, keeper)
		default:
			return nil, fmt.Errorf("unknown profiles query endpoint")
		}
	}
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
