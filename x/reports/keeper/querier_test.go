package keeper_test

import (
	"encoding/json"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) Test_queryReports() {
	tests := []struct {
		name          string
		path          []string
		storedReports []types.Report
		expErr        error
		expResponse   types.QueryPostReportsResponse
	}{
		{
			name:          "Invalid ID",
			path:          []string{types.QueryReports, "1234"},
			storedReports: nil,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid postID: 1234"),
		},
		{
			name: "Non empty stored and valid ID",
			path: []string{types.QueryReports, "1234"},
			storedReports: []types.Report{
				types.NewReport(
					"1234",
					"type",
					"message",
					suite.testData.creator.String(),
				),
				types.NewReport(
					"other_post",
					"type",
					"message",
					suite.testData.creator.String(),
				),
			},
			expErr: nil,
			expResponse: types.QueryPostReportsResponse{Reports: []types.Report{
				types.NewReport(
					"1234",
					"type",
					"message",
					suite.testData.creator.String(),
				),
			}},
		},
		{
			name:          "Empty stored and valid ID",
			path:          []string{types.QueryReports, suite.testData.postID.String()},
			storedReports: nil,
			expErr:        nil,
			expResponse:   types.QueryPostReportsResponse{Reports: []types.Report{}},
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

			if result != nil {
				suite.Require().Nil(err)

				expectedIndented, err := json.MarshalIndent(&test.expResponse, "", "")
				suite.Require().NoError(err)
				suite.Require().Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
				suite.Require().Nil(result)
			}
		})
	}
}
