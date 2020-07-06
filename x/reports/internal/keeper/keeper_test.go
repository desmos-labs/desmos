package keeper_test

import (
	"github.com/desmos-labs/desmos/x/posts"

	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models"
)

func (suite *KeeperTestSuite) TestKeeper_CheckExistence() {
	existentPost := posts.NewPost(postID,
		"",
		"Post",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		testPostCreationDate,
		creator,
	)

	tests := []struct {
		name         string
		existentPost *posts.Post
		postID       posts.PostID
		expBool      bool
	}{
		{
			name:         "Post not exist",
			existentPost: nil,
			postID:       postID,
			expBool:      false,
		},
		{
			name:         "Post exist",
			existentPost: &existentPost,
			postID:       postID,
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

			actualBool := suite.keeper.CheckPostExistence(suite.ctx, postID)
			suite.Equal(test.expBool, actualBool)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveReport() {
	expReports := models.Reports{models.NewReport("type", "message", creator)}
	report := models.NewReport("type", "message", creator)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)

	suite.keeper.SaveReport(suite.ctx, postID, report)

	var reports models.Reports
	suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.ReportStoreKey(postID)), &reports)
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
				{Type: "type", Message: "message", User: creator},
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
				store.Set(types.ReportStoreKey(postID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.expReports))
			}

			actualRep := suite.keeper.GetPostReports(suite.ctx, postID)
			suite.Equal(test.expReports, actualRep)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetReportsMap() {
	reports := models.Reports{
		{Type: "type", Message: "message", User: creator},
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
				postID.String(): reports,
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
				store.Set(types.ReportStoreKey(postID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.existingReports))
			}

			actualRep := suite.keeper.GetReportsMap(suite.ctx)
			suite.Equal(test.expReportsMap, actualRep)
		})
	}
}
