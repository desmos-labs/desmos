package keeper_test

import (
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_CheckReportValidity() {
	tests := []struct {
		name      string
		report    types.Report
		shouldErr bool
	}{
		{
			name: "Invalid report reason returns error",
			report: types.NewReport(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				[]string{"sdd"},
				"message",
				"user",
			),
			shouldErr: true,
		},
		{
			name: "Invalid report id returns error",
			report: types.NewReport(
				"123",
				[]string{"scam"},
				"message",
				"user",
			),
			shouldErr: true,
		},
		{
			name: "Valid report returns no error",
			report: types.NewReport(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				[]string{"scam"},
				"message",
				"user",
			),
			shouldErr: false,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.k.SetParams(suite.ctx, types.DefaultParams())
		suite.Run(test.name, func() {
			res := suite.k.CheckReportValidity(suite.ctx, test.report)
			if test.shouldErr {
				suite.Require().Error(res)
			} else {
				suite.Require().NoError(res)
			}
		})
	}
}

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
			report:        types.NewReport("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", []string{"scam"}, "message", "user"),
			expErr:        false,
			expReports: []types.Report{
				types.NewReport("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", []string{"scam"}, "message", "user"),
			},
		},
		{
			name: "report is stored properly when existing slice is not empty",
			storedReports: []types.Report{
				types.NewReport("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", []string{"scam"}, "message", "user"),
			},
			report: types.NewReport("b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9", []string{"nudity"}, "message", "user"),
			expErr: false,
			expReports: []types.Report{
				types.NewReport("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", []string{"scam"}, "message", "user"),
				types.NewReport("b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9", []string{"nudity"}, "message", "user"),
			},
		},
		{
			name: "trying to store double report returns error",
			storedReports: []types.Report{
				types.NewReport("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", []string{"scam"}, "message", "user"),
			},
			report: types.NewReport("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", []string{"scam"}, "message", "user"),
			expErr: true,
		},
		{
			name:   "trying to store invalid report returns error",
			report: types.NewReport("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", []string{"skam"}, "message", "user"),
			expErr: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest()
		suite.k.SetParams(suite.ctx, types.DefaultParams())
		suite.Run(test.name, func() {

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
					"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
					[]string{"scam"},
					"message",
					suite.testData.postOwner,
				),
				types.NewReport(
					"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
					[]string{"nudity"},
					"message",
					suite.testData.postOwner,
				),
			},
			postID: "b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
			expected: []types.Report{
				types.NewReport(
					"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
					[]string{"scam"},
					"message",
					suite.testData.postOwner,
				),
			},
		},
		{
			name:     "Returns an empty array",
			postID:   "b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
			stored:   nil,
			expected: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest()
		suite.k.SetParams(suite.ctx, types.DefaultParams())
		suite.Run(test.name, func() {

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
					[]string{"scam"},
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
		suite.SetupTest()
		suite.k.SetParams(suite.ctx, types.DefaultParams())
		suite.Run(test.name, func() {

			for _, report := range test.reports {
				err := suite.k.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			stored := suite.k.GetAllReports(suite.ctx)
			suite.Require().Equal(test.reports, stored)
		})
	}
}
