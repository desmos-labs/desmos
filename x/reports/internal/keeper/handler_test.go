package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/keeper"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/stretchr/testify/require"
)

func Test_handleMsgReportPost(t *testing.T) {
	msgReport := types.NewMsgReportPost(postID, "type", "message", creator)
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
		msg          types.MsgReportPost
		existentPost *posts.Post
		expErr       error
	}{
		{
			name:         "post not found",
			msg:          msgReport,
			existentPost: nil,
			expErr:       sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with ID: %s doesn't exist", postID)),
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
		t.Run(test.name, func(t *testing.T) {
			ctx, k, pk := SetupTestInput()

			if test.existentPost != nil {
				// Save the post
				pk.SavePost(ctx, *test.existentPost)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {
				//Check the data
				require.Equal(t, []byte(fmt.Sprintf("post with ID: %s reported correctly", postID)), res.Data)

				//Check the events
				createReportEv := sdk.NewEvent(
					types.EventTypePostReported,
					sdk.NewAttribute(types.AttributeKeyPostID, postID.String()),
					sdk.NewAttribute(types.AttributeKeyReportOwner, test.msg.Report.User.String()),
				)

				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, createReportEv)
			}

		})
	}
}
