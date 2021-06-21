package keeper_test

import (
	types2 "github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveReport() {
	tests := []struct {
		name          string
		storedReports []types2.Report
		report        types2.Report
		expErr        bool
		expReports    []types2.Report
	}{
		{
			name:          "report is stored properly when existing slice is empty",
			storedReports: nil,
			report:        types2.NewReport("post_id", "type", "message", "user"),
			expErr:        false,
			expReports: []types2.Report{
				types2.NewReport("post_id", "type", "message", "user"),
			},
		},
		{
			name: "report is stored properly when existing slice is not empty",
			storedReports: []types2.Report{
				types2.NewReport("post_id", "type", "message", "user"),
			},
			report: types2.NewReport("post_id", "type_2", "message", "user"),
			expErr: false,
			expReports: []types2.Report{
				types2.NewReport("post_id", "type", "message", "user"),
				types2.NewReport("post_id", "type_2", "message", "user"),
			},
		},
		{
			name: "trying to store double report returns error",
			storedReports: []types2.Report{
				types2.NewReport("post_id", "type", "message", "user"),
			},
			report: types2.NewReport("post_id", "type", "message", "user"),
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
		stored   []types2.Report
		postID   string
		expected []types2.Report
	}{
		{
			name: "Returns a non-empty array",
			stored: []types2.Report{
				types2.NewReport(
					"post_id",
					"type",
					"message",
					suite.testData.postOwner,
				),
				types2.NewReport(
					"another_post_id",
					"type",
					"message",
					suite.testData.postOwner,
				),
			},
			postID: "post_id",
			expected: []types2.Report{
				types2.NewReport(
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
		reports []types2.Report
	}{
		{
			name: "Empty stored are returned properly",
			reports: []types2.Report{
				types2.NewReport(
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
