package keeper_test

import (
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveReport() {
	tests := []struct {
		name          string
		storedReports []types.Report
		report        types.Report
		expErr        bool
		expReports    []types.Report
	}{
		{
			name:          "report is stored properly when existing slice is empty",
			storedReports: nil,
			report:        types.NewReport("post_id", "type", "message", "user"),
			expErr:        false,
			expReports: []types.Report{
				types.NewReport("post_id", "type", "message", "user"),
			},
		},
		{
			name: "report is stored properly when existing slice is not empty",
			storedReports: []types.Report{
				types.NewReport("post_id", "type", "message", "user"),
			},
			report: types.NewReport("post_id", "type_2", "message", "user"),
			expErr: false,
			expReports: []types.Report{
				types.NewReport("post_id", "type", "message", "user"),
				types.NewReport("post_id", "type_2", "message", "user"),
			},
		},
		{
			name: "trying to store double report returns error",
			storedReports: []types.Report{
				types.NewReport("post_id", "type", "message", "user"),
			},
			report: types.NewReport("post_id", "type", "message", "user"),
			expErr: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, report := range test.storedReports {
				err := suite.k.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			err := suite.k.SaveReport(suite.ctx, test.report)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored := suite.k.GetAllReports(suite.ctx)
				suite.Require().Equal(test.expReports, stored)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPostReports() {
	tests := []struct {
		name     string
		stored   []types.Report
		postID   string
		expected []types.Report
	}{
		{
			name: "Returns a non-empty array",
			stored: []types.Report{
				types.NewReport(
					"post_id",
					"type",
					"message",
					suite.testData.postOwner,
				),
				types.NewReport(
					"another_post_id",
					"type",
					"message",
					suite.testData.postOwner,
				),
			},
			postID: "post_id",
			expected: []types.Report{
				types.NewReport(
					"post_id",
					"type",
					"message",
					suite.testData.postOwner,
				),
			},
		},
		{
			name:     "Returns an empty array",
			postID:   "post_id",
			stored:   nil,
			expected: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, report := range test.stored {
				err := suite.k.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			stored := suite.k.GetPostReports(suite.ctx, test.postID)
			suite.Require().Equal(test.expected, stored)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllReports() {
	tests := []struct {
		name    string
		reports []types.Report
	}{
		{
			name: "Empty stored are returned properly",
			reports: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.postOwner,
				),
			},
		},
		{
			name:    "Returns an empty array",
			reports: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, report := range test.reports {
				err := suite.k.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			stored := suite.k.GetAllReports(suite.ctx)
			suite.Require().Equal(test.reports, stored)
		})
	}
}
