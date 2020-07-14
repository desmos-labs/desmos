package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryReports:
			return queryReports(ctx, path[1:], req, keeper)
		default:
			return nil, fmt.Errorf("unknown post query endpoint")
		}
	}
}

// getReportsResponse returns a ReportsQueryResponse using the information retrieved with id
func getReportsResponse(ctx sdk.Context, keeper Keeper, id posts.PostID) types.ReportsQueryResponse {
	reports := keeper.GetPostReports(ctx, id)
	if reports == nil {
		reports = types.Reports{}
	}

	return types.NewReportResponse(id, reports)
}

// queryReports handles the request of listing all the reports related to the given id
func queryReports(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	id := posts.PostID(path[0])
	if !id.Valid() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("invalid postID: %s", id))
	}

	reportsResponse := getReportsResponse(ctx, keeper, id)

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &reportsResponse)

	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
