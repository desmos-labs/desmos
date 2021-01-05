package keeper_test

import (
	"github.com/desmos-labs/desmos/x/reports/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) Test_handleMsgReportPost() {
	tests := []struct {
		name          string
		storedPosts   []poststypes.Post
		storedReports []types.Report
		msg           *types.MsgReportPost
		expErr        bool
		expEvents     sdk.Events
		expReports    []types.Report
	}{
		{
			name: "invalid post id",
			msg: types.NewMsgReportPost(
				"post_id",
				"type",
				"message",
				suite.testData.creator,
			),
			expErr: true,
		},
		{
			name:        "post not found",
			storedPosts: nil,
			msg: types.NewMsgReportPost(
				suite.testData.postID,
				"type",
				"message",
				suite.testData.creator,
			),
			expErr: true,
		},
		{
			name: "double report",
			storedPosts: []poststypes.Post{
				{
					PostID:       suite.testData.postID,
					Message:      "Post",
					Created:      suite.testData.creationDate,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.creator,
				},
			},
			storedReports: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.creator,
				),
			},
			msg: types.NewMsgReportPost(
				suite.testData.postID,
				"type",
				"message",
				suite.testData.creator,
			),
			expErr: true,
		},
		{
			name: "message handled correctly",
			storedPosts: []poststypes.Post{
				{
					PostID:       suite.testData.postID,
					Message:      "Post",
					Created:      suite.testData.creationDate,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.creator,
				},
			},
			msg: types.NewMsgReportPost(
				suite.testData.postID,
				"type",
				"message",
				suite.testData.creator,
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypePostReported,
					sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types.AttributeKeyReportOwner, suite.testData.creator),
				),
			},
			expReports: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.creator,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, post := range test.storedPosts {
				suite.postsKeeper.SavePost(suite.ctx, post)
			}

			for _, report := range test.storedReports {
				err := suite.keeper.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.keeper)
			_, err := server.ReportPost(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				reports := suite.keeper.GetAllReports(suite.ctx)
				suite.Require().Equal(test.expReports, reports)
			}
		})
	}
}
