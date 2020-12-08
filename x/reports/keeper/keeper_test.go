package keeper_test

import (
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) TestKeeper_CheckPostExistence() {
	tests := []struct {
		name         string
		existentPost *posts.Post
		postID       string
		expBool      bool
	}{
		{
			name:         "post does not exist",
			existentPost: nil,
			postID:       "post_id",
			expBool:      false,
		},
		{
			name: "post exists",
			existentPost: &posts.Post{
				PostID:       suite.testData.postID,
				Message:      "Post",
				Created:      suite.testData.creationDate,
				Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: nil,
				Creator:      suite.testData.creator,
			},
			postID:  suite.testData.postID,
			expBool: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			if test.existentPost != nil {
				suite.postsKeeper.SavePost(suite.ctx, *test.existentPost)
			}

			actualBool := suite.keeper.CheckPostExistence(suite.ctx, suite.testData.postID)
			suite.Require().Equal(test.expBool, actualBool)
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
				err := suite.keeper.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			err := suite.keeper.SaveReport(suite.ctx, test.report)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored := suite.keeper.GetAllReports(suite.ctx)
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
					suite.testData.creator,
				),
				types.NewReport(
					"another_post_id",
					"type",
					"message",
					suite.testData.creator,
				),
			},
			postID: "post_id",
			expected: []types.Report{
				types.NewReport(
					"post_id",
					"type",
					"message",
					suite.testData.creator,
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
				err := suite.keeper.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			stored := suite.keeper.GetPostReports(suite.ctx, test.postID)
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
					suite.testData.creator,
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
				err := suite.keeper.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			stored := suite.keeper.GetAllReports(suite.ctx)
			suite.Require().Equal(test.reports, stored)
		})
	}
}
