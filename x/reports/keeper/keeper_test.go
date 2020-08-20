package keeper_test

import (
	posts "github.com/desmos-labs/desmos/x/posts/types"

	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/desmos-labs/desmos/x/reports/types/models"
)

func (suite *KeeperTestSuite) TestKeeper_CheckExistence() {
	existentPost := posts.Post{
		PostID:       suite.testData.postID,
		Message:      "Post",
		Created:      suite.testData.postCreationDate,
		Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData: map[string]string{},
		Creator:      suite.testData.creator,
	}

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
			name:         "Post exist",
			existentPost: &existentPost,
			postID:       suite.testData.postID,
			expBool:      true,
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
			suite.Equal(test.expBool, actualBool)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveReport() {
	expReports := models.Reports{models.NewReport("type", "message", suite.testData.creator)}
	report := models.NewReport("type", "message", suite.testData.creator)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)

	suite.keeper.SaveReport(suite.ctx, suite.testData.postID, report)

	var reports models.Reports
	suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.ReportStoreKey(suite.testData.postID)), &reports)
	suite.Equal(expReports, reports)

}

func (suite *KeeperTestSuite) TestKeeper_GetPostReports() {
	tests := []struct {
		name       string
		expReports models.Reports
	}{
		{
			name: "Returns a non-empty reports array",
			expReports: models.Reports{
				{Type: "type", Message: "message", User: suite.testData.creator},
			},
		},
		{
			name:       "Returns an empty reports array",
			expReports: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.expReports != nil {
				store.Set(types.ReportStoreKey(suite.testData.postID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.expReports))
			}

			actualRep := suite.keeper.GetPostReports(suite.ctx, suite.testData.postID)
			suite.Equal(test.expReports, actualRep)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetReportsMap() {
	reports := models.Reports{
		{Type: "type", Message: "message", User: suite.testData.creator},
	}
	tests := []struct {
		name            string
		existingReports models.Reports
		expReportsMap   map[string]models.Reports
	}{
		{
			name:            "Returns a non-empty reports map",
			existingReports: reports,
			expReportsMap: map[string]models.Reports{
				suite.testData.postID.String(): reports,
			},
		},
		{
			name:            "Returns an empty reports map",
			existingReports: nil,
			expReportsMap:   map[string]models.Reports{},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.existingReports != nil {
				store.Set(types.ReportStoreKey(suite.testData.postID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.existingReports))
			}

			actualRep := suite.keeper.GetReportsMap(suite.ctx)
			suite.Equal(test.expReportsMap, actualRep)
		})
	}
}
