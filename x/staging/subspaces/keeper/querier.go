package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/commons"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCodec *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QuerySubspaces:
			return querySubspaces(ctx, path[1:], req, keeper, legacyQuerierCodec)

		case types.QuerySubspaceAdmins:
			return querySubspacesAdmins(ctx, path[1:], req, keeper, legacyQuerierCodec)

		case types.QuerySubspaceBlockedUsers:
			return querySubspacesBlockedUsers(ctx, path[1:], req, keeper, legacyQuerierCodec)

		default:
			return nil, fmt.Errorf("unknown Subspaces query endpoint")
		}
	}
}

// querySubspaces handles the request to get all the subspaces inside the given context
func querySubspaces(
	ctx sdk.Context, _ []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {

	requests := keeper.GetAllSubspaces(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &requests)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// checkSubspaceValidity ensure that the subspaceId given is valid and correspond to an existent subspace
func checkSubspaceValidity(ctx sdk.Context, keeper Keeper, subspaceId string) error {
	if !commons.IsValidSubspace(subspaceId) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %s", subspaceId)
	}

	if !keeper.DoesSubspaceExists(ctx, subspaceId) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id: %s not found", subspaceId)
	}

	return nil
}

// querySubspacesAdmins handles the request to get all the admins of the subspace with the given id
func querySubspacesAdmins(
	ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	subspaceId := path[0]
	if err := checkSubspaceValidity(ctx, keeper, subspaceId); err != nil {
		return nil, err
	}

	admins := keeper.GetAllSubspaceAdmins(ctx, subspaceId)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &admins)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func querySubspacesBlockedUsers(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	subspaceId := path[0]
	if err := checkSubspaceValidity(ctx, keeper, subspaceId); err != nil {
		return nil, err
	}

	admins := keeper.GetAllSubspaceAdmins(ctx, subspaceId)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &admins)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
