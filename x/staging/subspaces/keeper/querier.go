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
		case types.QuerySubspace:
			return querySubspace(ctx, path[1:], req, keeper, legacyQuerierCodec)
		default:
			return nil, fmt.Errorf("unknown Subspaces query endpoint")
		}
	}
}

// querySubspace handles the request to get the subspace with the given id
func querySubspace(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	subspaceId := path[0]

	if !commons.IsValidSubspace(subspaceId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %s", subspaceId)
	}

	subspace, exists := keeper.GetSubspace(ctx, subspaceId)
	if !exists {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id: %s not found", subspaceId)
	}

	admins := keeper.GetAllSubspaceAdmins(ctx, subspaceId)

	blockedUsers := keeper.GetSubspaceBlockedUsers(ctx, subspaceId)

	response := types.QuerySubspaceResponse{
		Subspace:           subspace,
		Admins:             admins,
		BlockedToPostUsers: blockedUsers,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &response)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
