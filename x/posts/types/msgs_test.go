package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

var attachments = []types.AttachmentContent{
	types.NewMedia(
		"ftp://user:password@example.com/image.png",
		"image/png",
	),
	types.NewPoll(
		"What animal is best?",
		[]types.Poll_ProvidedAnswer{
			types.NewProvidedAnswer("Cat", []types.AttachmentContent{
				types.NewMedia(
					"ftp://user:password@example.com/cat.png",
					"image/png",
				),
			}),
			types.NewProvidedAnswer("Dog", nil),
		},
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
		false,
		false,
		nil,
	),
}

var msgCreatePost = types.NewMsgCreatePost(
	1,
	1,
	"External ID",
	"This is a text",
	1,
	types.REPLY_SETTING_EVERYONE,
	types.NewEntities(
		[]types.TextTag{
			types.NewTextTag(1, 3, "tag"),
		},
		[]types.TextTag{
			types.NewTextTag(4, 6, "tag"),
		},
		[]types.Url{
			types.NewURL(7, 9, "URL", "Display URL"),
		},
	),
	[]string{"general"},
	attachments,
	[]types.PostReference{
		types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
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
			name: "invalid subspace id returns error",
			msg: types.NewMsgCreatePost(
				0,
				msgCreatePost.SectionID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				msgCreatePost.Tags,
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
				msgCreatePost.SectionID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				types.REPLY_SETTING_UNSPECIFIED,
				msgCreatePost.Entities,
				msgCreatePost.Tags,
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
				msgCreatePost.SectionID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				types.NewEntities([]types.TextTag{
					types.NewTextTag(1, 1, "My tag"),
					types.NewTextTag(1, 1, "My tag"),
				}, nil, nil),
				msgCreatePost.Tags,
				attachments,
				msgCreatePost.ReferencedPosts,
				msgCreatePost.Author,
			),
			shouldErr: true,
		},
		{
			name: "invalid tags return error",
			msg: types.NewMsgCreatePost(
				msgCreatePost.SubspaceID,
				msgCreatePost.SectionID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				[]string{"   "},
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
				msgCreatePost.SectionID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				msgCreatePost.Tags,
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
				msgCreatePost.SectionID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				msgCreatePost.Tags,
				attachments,
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_UNSPECIFIED, 0, 1),
				},
				msgCreatePost.Author,
			),
			shouldErr: true,
		},
		{
			name: "invalid author returns error",
			msg: types.NewMsgCreatePost(
				msgCreatePost.SubspaceID,
				msgCreatePost.SectionID,
				msgCreatePost.ExternalID,
				msgCreatePost.Text,
				msgCreatePost.ConversationID,
				msgCreatePost.ReplySettings,
				msgCreatePost.Entities,
				msgCreatePost.Tags,
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
	expected := `{"type":"desmos/MsgCreatePost","value":{"attachments":[{"type":"desmos/Media","value":{"mime_type":"image/png","uri":"ftp://user:password@example.com/image.png"}},{"type":"desmos/Poll","value":{"end_date":"2020-01-01T12:00:00Z","provided_answers":[{"attachments":[{"type":"desmos/Media","value":{"mime_type":"image/png","uri":"ftp://user:password@example.com/cat.png"}}],"text":"Cat"},{"text":"Dog"}],"question":"What animal is best?"}}],"author":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","conversation_id":"1","entities":{"hashtags":[{"end":"3","start":"1","tag":"tag"}],"mentions":[{"end":"6","start":"4","tag":"tag"}],"urls":[{"display_url":"Display URL","end":"9","start":"7","url":"URL"}]},"external_id":"External ID","referenced_posts":[{"post_id":"1","type":2}],"reply_settings":1,"section_id":1,"subspace_id":"1","tags":["general"],"text":"This is a text"}}`
	require.Equal(t, expected, string(msgCreatePost.GetSignBytes()))
}

func TestMsgCreatePost_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCreatePost.Author)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreatePost.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgEditPost = types.NewMsgEditPost(
	1,
	1,
	"Edited text",
	types.NewEntities(
		[]types.TextTag{
			types.NewTextTag(1, 3, "tag"),
		},
		[]types.TextTag{
			types.NewTextTag(4, 6, "tag"),
		},
		[]types.Url{
			types.NewURL(7, 9, "URL", "Display URL"),
		},
	),
	[]string{"general"},
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgEditPost_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgEditPost.Route())
}

func TestMsgEditPost_Type(t *testing.T) {
	require.Equal(t, types.ActionEditPost, msgEditPost.Type())
}

func TestMsgEditPost_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgEditPost
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgEditPost(
				0,
				msgEditPost.PostID,
				msgEditPost.Text,
				msgEditPost.Entities,
				msgEditPost.Tags,
				msgEditPost.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgEditPost(
				msgEditPost.SubspaceID,
				0,
				msgEditPost.Text,
				msgEditPost.Entities,
				msgEditPost.Tags,
				msgEditPost.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid entities returns error",
			msg: types.NewMsgEditPost(
				msgEditPost.SubspaceID,
				msgEditPost.PostID,
				msgEditPost.Text,
				types.NewEntities([]types.TextTag{
					types.NewTextTag(1, 1, "My tag"),
					types.NewTextTag(1, 1, "My tag"),
				}, nil, nil),
				msgEditPost.Tags,
				msgEditPost.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid tags return error",
			msg: types.NewMsgEditPost(
				msgEditPost.SubspaceID,
				msgEditPost.PostID,
				msgEditPost.Text,
				msgEditPost.Entities,
				[]string{"   "},
				msgEditPost.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			msg: types.NewMsgEditPost(
				msgEditPost.SubspaceID,
				msgEditPost.PostID,
				msgEditPost.Text,
				msgEditPost.Entities,
				msgEditPost.Tags,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgEditPost,
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

func TestMsgEditPost_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgEditPost","value":{"editor":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","entities":{"hashtags":[{"end":"3","start":"1","tag":"tag"}],"mentions":[{"end":"6","start":"4","tag":"tag"}],"urls":[{"display_url":"Display URL","end":"9","start":"7","url":"URL"}]},"post_id":"1","subspace_id":"1","tags":["general"],"text":"Edited text"}}`
	require.Equal(t, expected, string(msgEditPost.GetSignBytes()))
}

func TestMsgEditPost_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditPost.Editor)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditPost.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var attachmentContent = types.NewMedia(
	"ftp://user:password@example.com/image.png",
	"image/png",
)

var msgAddPostAttachment = types.NewMsgAddPostAttachment(
	1,
	1,
	attachmentContent,
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgAddPostAttachment_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgAddPostAttachment.Route())
}

func TestMsgAddPostAttachment_Type(t *testing.T) {
	require.Equal(t, types.ActionAddPostAttachment, msgAddPostAttachment.Type())
}

func TestMsgAddPostAttachment_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAddPostAttachment
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAddPostAttachment(
				0,
				msgAddPostAttachment.PostID,
				attachmentContent,
				msgAddPostAttachment.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgAddPostAttachment(
				msgAddPostAttachment.SubspaceID,
				0,
				attachmentContent,
				msgAddPostAttachment.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid attachment content returns error",
			msg: &types.MsgAddPostAttachment{
				SubspaceID: msgAddPostAttachment.SubspaceID,
				PostID:     msgAddPostAttachment.PostID,
				Content:    nil,
				Editor:     msgAddPostAttachment.Editor,
			},
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			msg: types.NewMsgAddPostAttachment(
				msgAddPostAttachment.SubspaceID,
				msgAddPostAttachment.PostID,
				attachmentContent,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgAddPostAttachment,
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

func TestMsgAddPostAttachment_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgAddPostAttachment","value":{"content":{"type":"desmos/Media","value":{"mime_type":"image/png","uri":"ftp://user:password@example.com/image.png"}},"editor":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","post_id":"1","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgAddPostAttachment.GetSignBytes()))
}

func TestMsgAddPostAttachment_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAddPostAttachment.Editor)
	require.Equal(t, []sdk.AccAddress{addr}, msgAddPostAttachment.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRemovePostAttachment = types.NewMsgRemovePostAttachment(
	1,
	1,
	1,
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgRemovePostAttachment_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRemovePostAttachment.Route())
}

func TestMsgRemovePostAttachment_Type(t *testing.T) {
	require.Equal(t, types.ActionRemovePostAttachment, msgRemovePostAttachment.Type())
}

func TestMsgRemovePostAttachment_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRemovePostAttachment
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRemovePostAttachment(
				0,
				msgRemovePostAttachment.PostID,
				msgRemovePostAttachment.AttachmentID,
				msgRemovePostAttachment.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgRemovePostAttachment(
				msgRemovePostAttachment.SubspaceID,
				0,
				msgRemovePostAttachment.AttachmentID,
				msgRemovePostAttachment.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid attachment id returns error",
			msg: types.NewMsgRemovePostAttachment(
				msgRemovePostAttachment.SubspaceID,
				msgRemovePostAttachment.PostID,
				0,
				msgRemovePostAttachment.Editor,
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			msg: types.NewMsgRemovePostAttachment(
				msgRemovePostAttachment.SubspaceID,
				msgRemovePostAttachment.PostID,
				msgRemovePostAttachment.AttachmentID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRemovePostAttachment,
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

func TestMsgRemovePostAttachment_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRemovePostAttachment","value":{"attachment_id":1,"editor":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","post_id":"1","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgRemovePostAttachment.GetSignBytes()))
}

func TestMsgRemovePostAttachment_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRemovePostAttachment.Editor)
	require.Equal(t, []sdk.AccAddress{addr}, msgRemovePostAttachment.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgDeletePost = types.NewMsgDeletePost(
	1,
	1,
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgDeletePost_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgDeletePost.Route())
}

func TestMsgDeletePost_Type(t *testing.T) {
	require.Equal(t, types.ActionDeletePost, msgDeletePost.Type())
}

func TestMsgDeletePost_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgDeletePost
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgDeletePost(
				0,
				msgDeletePost.PostID,
				msgDeletePost.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgDeletePost(
				msgDeletePost.SubspaceID,
				0,
				msgDeletePost.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			msg: types.NewMsgDeletePost(
				msgDeletePost.SubspaceID,
				msgDeletePost.PostID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgDeletePost,
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

func TestMsgDeletePost_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgDeletePost","value":{"post_id":"1","signer":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgDeletePost.GetSignBytes()))
}

func TestMsgDeletePost_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgDeletePost.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeletePost.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgAnswerPoll = types.NewMsgAnswerPoll(
	1,
	1,
	1,
	[]uint32{1, 2, 3},
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgAnswerPoll_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgAnswerPoll.Route())
}

func TestMsgAnswerPoll_Type(t *testing.T) {
	require.Equal(t, types.ActionAnswerPoll, msgAnswerPoll.Type())
}

func TestMsgAnswerPoll_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgAnswerPoll
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgAnswerPoll(
				0,
				msgAnswerPoll.PostID,
				msgAnswerPoll.PollID,
				msgAnswerPoll.AnswersIndexes,
				msgAnswerPoll.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgAnswerPoll(
				msgAnswerPoll.SubspaceID,
				0,
				msgAnswerPoll.PollID,
				msgAnswerPoll.AnswersIndexes,
				msgAnswerPoll.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid poll id returns error",
			msg: types.NewMsgAnswerPoll(
				msgAnswerPoll.SubspaceID,
				msgAnswerPoll.PostID,
				0,
				msgAnswerPoll.AnswersIndexes,
				msgAnswerPoll.Signer,
			),
			shouldErr: true,
		},
		{
			name: "empty answers returns error",
			msg: types.NewMsgAnswerPoll(
				msgAnswerPoll.SubspaceID,
				msgAnswerPoll.PostID,
				msgAnswerPoll.PollID,
				nil,
				msgAnswerPoll.Signer,
			),
			shouldErr: true,
		},
		{
			name: "duplicated answers returns error",
			msg: types.NewMsgAnswerPoll(
				msgAnswerPoll.SubspaceID,
				msgAnswerPoll.PostID,
				msgAnswerPoll.PollID,
				[]uint32{1, 2, 3, 4, 1},
				msgAnswerPoll.Signer,
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			msg: types.NewMsgAnswerPoll(
				msgAnswerPoll.SubspaceID,
				msgAnswerPoll.PostID,
				msgAnswerPoll.PollID,
				msgAnswerPoll.AnswersIndexes,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgAnswerPoll,
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

func TestMsgAnswerPoll_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgAnswerPoll","value":{"answers_indexes":[1,2,3],"poll_id":1,"post_id":"1","signer":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgAnswerPoll.GetSignBytes()))
}

func TestMsgAnswerPoll_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAnswerPoll.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgAnswerPoll.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgMovePost = types.NewMsgMovePost(
	1,
	1,
	1,
	0,
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgMovePost_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgMovePost.Route())
}

func TestMsgMovePost_Type(t *testing.T) {
	require.Equal(t, types.ActionMovePost, msgMovePost.Type())
}

func TestMsgMovePost_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgMovePost
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgMovePost(
				0,
				msgMovePost.PostID,
				msgMovePost.TargetSubspaceID,
				msgMovePost.TargetSectionID,
				msgMovePost.Owner,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			msg: types.NewMsgMovePost(
				msgMovePost.SubspaceID,
				0,
				msgMovePost.TargetSubspaceID,
				msgMovePost.TargetSectionID,
				msgMovePost.Owner,
			),
			shouldErr: true,
		},
		{
			name: "invalid target subspace id returns error",
			msg: types.NewMsgMovePost(
				msgMovePost.SubspaceID,
				msgMovePost.PostID,
				0,
				msgMovePost.TargetSectionID,
				msgMovePost.Owner,
			),
			shouldErr: true,
		},
		{
			name: "invalid owner returns error",
			msg: types.NewMsgMovePost(
				msgMovePost.SubspaceID,
				msgMovePost.PostID,
				msgMovePost.TargetSubspaceID,
				msgMovePost.TargetSectionID,
				"",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgMovePost,
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

func TestMsgMovePost_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgMovePost","value":{"owner":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","post_id":"1","subspace_id":"1","target_subspace_id":"1"}}`
	require.Equal(t, expected, string(msgMovePost.GetSignBytes()))
}

func TestMsgMovePost_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgMovePost.Owner)
	require.Equal(t, []sdk.AccAddress{addr}, msgMovePost.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgUpdateParams = types.NewMsgUpdateParams(
	types.DefaultParams(),
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgUpdateParams_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgUpdateParams.Route())
}

func TestMsgUpdateParams_Type(t *testing.T) {
	require.Equal(t, types.ActionUpdateParams, msgUpdateParams.Type())
}

func TestMsgUpdateParams_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgUpdateParams
		shouldErr bool
	}{
		{
			name: "invalid authority returns error",
			msg: types.NewMsgUpdateParams(
				types.DefaultParams(),
				"invalid",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgUpdateParams,
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

func TestMsgUpdateParams_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/x/posts/MsgUpdateParams","value":{"authority":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","params":{"max_text_length":500}}}`
	require.Equal(t, expected, string(msgUpdateParams.GetSignBytes()))
}

func TestMsgUpdateParams_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUpdateParams.Authority)
	require.Equal(t, []sdk.AccAddress{addr}, msgUpdateParams.GetSigners())
}
