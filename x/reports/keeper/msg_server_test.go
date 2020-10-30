package keeper_test

import (
	"context"
	"fmt"

	"github.com/desmos-labs/desmos/x/reports/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	poststypes "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) Test_handleMsgReportPost() {
	tests := []struct {
		name        string
		storedPosts []poststypes.Post
		msg         *types.MsgReportPost
		expErr      error
	}{
		{
			name:        "post not found",
			storedPosts: nil,
			msg: types.NewMsgReportPost(
				suite.testData.postID.String(),
				"type",
				"message",
				suite.testData.creator.String(),
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with ID: %s doesn't exist", suite.testData.postID)),
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
				suite.testData.postID.String(),
				"type",
				"message",
				suite.testData.creator.String(),
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset

			for _, post := range test.storedPosts {
				suite.postsKeeper.SavePost(suite.ctx, post)
			}

			server := keeper.NewMsgServerImpl(suite.keeper)
			_, err := server.ReportPost(context.Background(), test.msg)

			if test.expErr == nil {
				// Check the events
				createReportEv := sdk.NewEvent(
					types.EventTypePostReported,
					sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID.String()),
					sdk.NewAttribute(types.AttributeKeyReportOwner, test.msg.User),
				)

				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), createReportEv)
			}

			if test.expErr != nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
		})
	}
}
