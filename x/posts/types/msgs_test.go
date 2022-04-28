package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

var attachments = []types.AttachmentContent{
	types.NewMedia(
		"ftp://user:password@example.com/image.png",
		"image/png",
	),
	types.NewPoll(
		"What animal is best?",
		[]types.Poll_ProvidedAnswer{
			types.NewProvidedAnswer("Cat", nil),
			types.NewProvidedAnswer("Dog", nil),
		},
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
		false,
		false,
	),
}

var msgCreatePost = types.NewMsgCreatePost(
	1,
	"External ID",
	"This is a text",
	1,
	types.REPLY_SETTING_EVERYONE,
	types.NewEntities(
		[]types.Tag{
			types.NewTag(1, 3, "tag"),
		},
		[]types.Tag{
			types.NewTag(4, 6, "tag"),
		},
		[]types.Url{
			types.NewURL(7, 9, "URL", "Display URL"),
		},
	),
	attachments,
	[]types.PostReference{
		types.NewPostReference(types.TYPE_QUOTED, 1),
	},
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgCreatePost_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgCreatePost.Route())
}

func TestMsgCreatePost_Type(t *testing.T) {
	require.Equal(t, types.ActionCreatePost, msgCreatePost.Type())
}

func TestMsgCreatePost_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCreatePost
		shouldErr bool
	}{
		{
			name: "invalid subspace returns error",
			msg: types.NewMsgCreatePost(
				0,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				attachments,
				msgCreatePost.ReferencedPosts,
				msgCreatePost.Author,
			),
			shouldErr: true,
		},
		{
			name: "invalid reply settings returns error",
			msg: types.NewMsgCreatePost(
				msgCreatePost.SubspaceID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				types.REPLY_SETTING_UNSPECIFIED,
				msgCreatePost.Entities,
				attachments,
				msgCreatePost.ReferencedPosts,
				msgCreatePost.Author,
			),
			shouldErr: true,
		},
		{
			name: "invalid entities returns error",
			msg: types.NewMsgCreatePost(
				msgCreatePost.SubspaceID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				types.NewEntities([]types.Tag{
					types.NewTag(1, 1, "My tag"),
					types.NewTag(1, 1, "My tag"),
				}, nil, nil),
				attachments,
				msgCreatePost.ReferencedPosts,
				msgCreatePost.Author,
			),
			shouldErr: true,
		},
		{
			name: "invalid attachments returns error",
			msg: types.NewMsgCreatePost(
				msgCreatePost.SubspaceID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				[]types.AttachmentContent{
					types.NewMedia("", ""),
				},
				msgCreatePost.ReferencedPosts,
				msgCreatePost.Author,
			),
			shouldErr: true,
		},
		{
			name: "invalid post reference returns error",
			msg: types.NewMsgCreatePost(
				msgCreatePost.SubspaceID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				attachments,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_UNSPECIFIED, 0),
				},
				msgCreatePost.Author,
			),
			shouldErr: true,
		},
		{
			name: "invalid author returns error",
			msg: types.NewMsgCreatePost(
				msgCreatePost.SubspaceID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				attachments,
				msgCreatePost.ReferencedPosts,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgCreatePost,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCreatePost_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreatePost","value":{"attachments":[{"type":"desmos/Media","value":{"mime_type":"image/png","uri":"ftp://user:password@example.com/image.png"}},{"type":"desmos/Poll","value":{"end_date":"2020-01-01T12:00:00Z","provided_answers":[{"attachments":null,"text":"Cat"},{"attachments":null,"text":"Dog"}],"question":"What animal is best?"}}],"author":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","conversation_id":"1","entities":{"hashtags":[{"end":"3","start":"1","tag":"tag"}],"mentions":[{"end":"6","start":"4","tag":"tag"}],"urls":[{"display_url":"Display URL","end":"9","start":"7","url":"URL"}]},"external_id":"External ID","referenced_posts":[{"post_id":"1","type":2}],"reply_settings":1,"subspace_id":"1","text":"This is a text"}}`
	require.Equal(t, expected, string(msgCreatePost.GetSignBytes()))
}

func TestMsgCreatePost_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCreatePost.Author)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreatePost.GetSigners())
}
