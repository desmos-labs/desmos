package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryProfile:
			return queryAccount(ctx, path[1:], req, keeper)
		case types.QueryProfiles:
			return queryAccounts(ctx, req, keeper)
		default:
			return nil, fmt.Errorf("unknown post query endpoint")
		}
	}
}

// queryAccount handles the request to get an account having a moniker
func queryAccount(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	account, found := keeper.GetProfile(ctx, path[0])

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("Profile with moniker %s doesn't exists", path[0]))
	}

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &account)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryPosts handles the request of listing all the accounts
func queryAccounts(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	accounts := keeper.GetProfiles(ctx)

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &accounts)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
