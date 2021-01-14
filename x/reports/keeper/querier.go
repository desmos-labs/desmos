package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	poststypes "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCodec *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryReports:
			return queryReports(ctx, path[1:], req, keeper, legacyQuerierCodec)
		default:
			return nil, fmt.Errorf("unknown post query endpoint")
		}
	}
}

// queryReports handles the request of listing all the stored related to the given id
func queryReports(
	ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	id := path[0]
	if !poststypes.IsValidPostID(id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("invalid postID: %s", id))
	}

	reports := keeper.GetPostReports(ctx, path[0])

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, &reports)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
