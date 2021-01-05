package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) TestQueryPostReports() {
	tests := []struct {
		name          string
		storedReports []types.Report
		request       *types.QueryPostReportsRequest
		expError      bool
		expResponse   *types.QueryPostReportsResponse
	}{
		{
			name:          "empty reports return nil",
			storedReports: []types.Report{},
			request:       &types.QueryPostReportsRequest{PostId: "post_id"},
			expError:      false,
			expResponse:   &types.QueryPostReportsResponse{Reports: nil},
		},
		{
			name: "reports are returned properly",
			storedReports: []types.Report{
				types.NewReport("post_id_1", "type_1", "message_1", "user_1"),
				types.NewReport("post_id_2", "type_2", "message_2", "user_2"),
			},
			request:  &types.QueryPostReportsRequest{PostId: "post_id_1"},
			expError: false,
			expResponse: &types.QueryPostReportsResponse{Reports: []types.Report{
				types.NewReport("post_id_1", "type_1", "message_1", "user_1"),
			}},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, report := range test.storedReports {
				err := suite.keeper.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			ctx := sdk.WrapSDKContext(suite.ctx)
			response, err := suite.keeper.PostReports(ctx, test.request)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expResponse, response)
			}
		})
	}
}
