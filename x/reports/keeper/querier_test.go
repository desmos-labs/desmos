package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) Test_queryReports() {
	reports := types.Reports{types.NewReport("type", "message", suite.testData.creator)}
	tests := []struct {
		name          string
		path          []string
		storedReports types.Reports
		expErr        error
		expResponse   types.ReportsQueryResponse
	}{
		{
			name:          "Invalid RelationshipID",
			path:          []string{types.QueryReports, "1234"},
			storedReports: nil,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid postID: 1234"),
		},
		{
			name:          "Non empty reports and valid RelationshipID",
			path:          []string{types.QueryReports, suite.testData.postID.String()},
			storedReports: reports,
			expErr:        nil,
			expResponse:   types.NewReportResponse(suite.testData.postID, reports),
		},
		{
			name:          "Empty reports and valid RelationshipID",
			path:          []string{types.QueryReports, suite.testData.postID.String()},
			storedReports: nil,
			expErr:        nil,
			expResponse:   types.NewReportResponse(suite.testData.postID, types.Reports{}),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			for _, rep := range test.storedReports {
				suite.keeper.SaveReport(suite.ctx, suite.testData.postID, rep)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResponse)
				suite.NoError(err)
				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
				suite.Nil(result)
			}

		})
	}
}
