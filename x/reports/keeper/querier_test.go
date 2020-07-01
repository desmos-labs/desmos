package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func Test_queryReports(t *testing.T) {
	reports := types.Reports{types.NewReport("type", "message", creator)}
	tests := []struct {
		name          string
		path          []string
		storedReports types.Reports
		expErr        error
		expResponse   types.ReportsQueryResponse
	}{
		{
			name:          "Invalid ID",
			path:          []string{types.QueryReports, "1234"},
			storedReports: nil,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid postID: 1234"),
		},
		{
			name:          "Non empty reports and valid ID",
			path:          []string{types.QueryReports, postID.String()},
			storedReports: reports,
			expErr:        nil,
			expResponse:   types.NewReportResponse(postID, reports),
		},
		{
			name:          "Empty reports and valid ID",
			path:          []string{types.QueryReports, postID.String()},
			storedReports: nil,
			expErr:        nil,
			expResponse:   types.NewReportResponse(postID, types.Reports{}),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k, _ := SetupTestInput()

			for _, rep := range test.storedReports {
				k.SaveReport(ctx, postID, rep)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				require.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.expResponse)
				require.NoError(t, err)
				require.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
				require.Nil(t, result)
			}

		})
	}
}
