package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryProfile:
			return queryProfile(ctx, path[1:], req, keeper)
		case types.QueryProfiles:
			return queryProfiles(ctx, req, keeper)
		case types.QueryParams:
			return queryProfileParams(ctx, req, keeper)
		case types.QueryDTagRequests:
			return queryDTagRequests(ctx, path[1:], req, keeper)
		default:
			return nil, fmt.Errorf("unknown profiles query endpoint")
		}
	}
}

// queryDTagRequests handles the request to get all the dTag requests of a user
func queryDTagRequests(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	user, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid bech32 address: %s", path[0]))
	}

	dTagRequests := keeper.GetUserDTagTransferRequests(ctx, user)

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &dTagRequests)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryProfile handles the request to get a profile having a dtag or an address
func queryProfile(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	if len(strings.TrimSpace(path[0])) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "DTag or address cannot be empty or blank")
	}

	address, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		address = keeper.GetDtagRelatedAddress(ctx, path[0])
		if address == nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("No address related to this dtag: %s", path[0]))
		}

	}

	account, found := keeper.GetProfile(ctx, address)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("Profile with address %s doesn't exists", path[0]))
	}

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &account)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryProfiles handles the request of listing all the profiles
func queryProfiles(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	accounts := keeper.GetProfiles(ctx)

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &accounts)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryProfileParams handles the request of listing all the profiles params
func queryProfileParams(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	profileParams := keeper.GetParams(ctx)

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &profileParams)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
