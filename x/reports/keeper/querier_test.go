package keeper_test

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) Test_queryReports() {
	tests := []struct {
		name          string
		path          []string
		storedReports []types.Report
		expErr        bool
		expResponse   []types.Report
	}{
		{
			name:          "Invalid Post ID",
			path:          []string{types.QueryReports, "1234"},
			storedReports: nil,
			expErr:        true,
		},
		{
			name: "Valid request returns correctly",
			path: []string{types.QueryReports, suite.testData.postID},
			storedReports: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.creator,
				),
				types.NewReport(
					"other_post",
					"type",
					"message",
					suite.testData.creator,
				),
			},
			expErr: false,
			expResponse: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.creator,
				),
			},
		},
		{
			name:          "Empty reports and valid id",
			path:          []string{types.QueryReports, suite.testData.postID},
			storedReports: nil,
			expErr:        false,
			expResponse:   nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, rep := range test.storedReports {
				err := suite.keeper.SaveReport(suite.ctx, rep)
				suite.Require().NoError(err)
			}

			querier := keeper.NewQuerier(suite.keeper, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				var reports []types.Report
				suite.Require().NoError(suite.legacyAminoCdc.UnmarshalJSON(result, &reports))
				suite.Require().Equal(test.expResponse, reports)
			}
		})
	}
}
