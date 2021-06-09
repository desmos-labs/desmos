package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/commons"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCodec *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QuerySubspace:
			return querySubspace(ctx, path[1:], req, keeper, legacyQuerierCodec)
		case types.QuerySubspaces:
			return querySubspaces(ctx, req, keeper, legacyQuerierCodec)
		default:
			return nil, fmt.Errorf("unknown Subspaces query endpoint")
		}
	}
}

func querySubspaces(ctx sdk.Context, _ abci.RequestQuery, keeper Keeper, legacyQuerierCodec *codec.LegacyAmino,
) ([]byte, error) {
	subspaces := keeper.GetAllSubspaces(ctx)

	subspacesResponse := types.QuerySubspacesResponse{Subspaces: subspaces}
	bytes, err := codec.MarshalJSONIndent(legacyQuerierCodec, &subspacesResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bytes, nil
}

// querySubspace handles the request to get the subspaces with the given id
func querySubspace(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	subspaceID := path[0]

	if !commons.IsValidSubspace(subspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspaces id: %s", subspaceID)
	}

	subspace, exists := keeper.GetSubspace(ctx, subspaceID)
	if !exists {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspaces with id: %s not found", subspaceID)
	}

	response := types.QuerySubspaceResponse{
		Subspace: subspace,
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &response)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
