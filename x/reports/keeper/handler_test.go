package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/stretchr/testify/require"
)

func (suite *KeeperTestSuite) Test_handleMsgReportPost() {
	msgReport := types.NewMsgReportPost(suite.testData.postID, "type", "message", suite.testData.creator)
	existentPost := posts.NewPost(suite.testData.postID,
		"",
		"Post",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		suite.testData.postCreationDate,
		suite.testData.creator,
	)

	tests := []struct {
		name         string
		msg          types.MsgReportPost
		existentPost *posts.Post
		expErr       error
	}{
		{
			name:         "post not found",
			msg:          msgReport,
			existentPost: nil,
			expErr:       sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with ID: %s doesn't exist", suite.testData.postID)),
		},
		{
			name:         "message handled correctly",
			msg:          msgReport,
			existentPost: &existentPost,
			expErr:       nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset

			if test.existentPost != nil {
				// Save the post
				suite.postsKeeper.SavePost(suite.ctx, *test.existentPost)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}
			if res != nil {
				//Check the data
				suite.Equal([]byte(fmt.Sprintf("post with ID: %s reported correctly", suite.testData.postID)), res.Data)

				//Check the events
				createReportEv := sdk.NewEvent(
					types.EventTypePostReported,
					sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID.String()),
					sdk.NewAttribute(types.AttributeKeyReportOwner, test.msg.Report.User.String()),
				)

				suite.Len(res.Events, 1)
				suite.Contains(res.Events, createReportEv)
			}

		})
	}
}
