package keeper_test

import (
	posts "github.com/desmos-labs/desmos/x/posts/types"

	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) TestKeeper_CheckExistence() {
	tests := []struct {
		name         string
		existentPost *posts.Post
		postID       posts.PostID
		expBool      bool
	}{
		{
			name:         "Post not exist",
			existentPost: nil,
			postID:       suite.testData.postID,
			expBool:      false,
		},
		{
			name: "Post exist",
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
	store := suite.ctx.KVStore(suite.storeKey)

	report := types.NewReport(
		suite.testData.postID.String(),
		"type",
		"message",
		suite.testData.creator.String(),
	)
	err := suite.keeper.SaveReport(suite.ctx, report)
	suite.Require().NoError(err)

	reports, err := suite.keeper.UnmarshalReports(store.Get(types.ReportStoreKey(suite.testData.postID.String())))
	suite.Require().NoError(err)

	suite.Require().Equal(reports, []types.Report{report})
}

func (suite *KeeperTestSuite) TestKeeper_GetPostReports() {
	tests := []struct {
		name     string
		postID   string
		stored   []types.Report
		expected []types.Report
	}{
		{
			name: "Returns a non-empty stored array",
			stored: []types.Report{
				types.NewReport(
					"post_id",
					"type",
					"message",
					suite.testData.creator.String(),
				),
				types.NewReport(
					"another_post_id",
					"type",
					"message",
					suite.testData.creator.String(),
				),
			},
			postID: "post_id",
			expected: []types.Report{
				types.NewReport(
					suite.testData.postID.String(),
					"type",
					"message",
					suite.testData.creator.String(),
				),
			},
		},
		{
			name:     "Returns an empty stored array",
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

			stored, err := suite.keeper.GetPostReports(suite.ctx, test.postID)
			suite.Require().NoError(err)
			suite.Require().Equal(test.expected, stored)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetReports() {
	tests := []struct {
		name    string
		reports []types.Report
	}{
		{
			name: "Empty stored are returned properly",
			reports: []types.Report{
				types.NewReport(
					suite.testData.postID.String(),
					"type",
					"message",
					suite.testData.creator.String(),
				),
			},
		},
		{
			name:    "Returns an empty stored map",
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

			stored, err := suite.keeper.GetReports(suite.ctx)
			suite.Require().NoError(err)
			suite.Require().Equal(test.reports, stored)
		})
	}
}
