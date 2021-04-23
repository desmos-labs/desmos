package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryProfile:
			return queryProfile(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case types.QueryIncomingDTagRequests:
			return queryIncomingDTagRequests(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case types.QueryParams:
			return queryProfileParams(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryUserRelationships:
			return queryUserRelationships(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case types.QueryUserBlocks:
			return queryUserBlocks(ctx, path[1:], req, keeper, legacyQuerierCdc)

		default:
			return nil, fmt.Errorf("unknown Profiles query endpoint")
		}
	}
}

// queryProfile handles the request to get a profile having a DTag or an address
func queryProfile(
	ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	dTagOrAddress := path[0]
	if strings.TrimSpace(dTagOrAddress) == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "DTag or address cannot be empty or blank")
	}

	sdkAddress, err := sdk.AccAddressFromBech32(dTagOrAddress)
	if err != nil {
		addr := keeper.GetAddressFromDTag(ctx, dTagOrAddress)
		if addr == "" {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"No address related to this DTag: %s", dTagOrAddress)
		}

		sdkAddress, err = sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, addr)
		}
	}

	account, found, err := keeper.GetProfile(ctx, sdkAddress.String())
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Profile with address %s doesn't exists", dTagOrAddress)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &account)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryIncomingDTagRequests handles the request to get all the incoming DTag requests of a user
func queryIncomingDTagRequests(
	ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	user, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid bech32 address: %s", path[0]))
	}

	requests := keeper.GetUserIncomingDTagTransferRequests(ctx, user.String())

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &requests)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryProfileParams handles the request of listing all the Profiles params
func queryProfileParams(
	ctx sdk.Context, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	profileParams := keeper.GetParams(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &profileParams)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
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
