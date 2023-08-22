package wasm_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/posts/types"
	"github.com/desmos-labs/desmos/v6/x/posts/wasm"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func TestMsgsParser_ParseCustomMsgs(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	parser := wasm.NewWasmMsgParser(cdc)
	contractAddr, err := sdk.AccAddressFromBech32("cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr")
	require.NoError(t, err)

	wrongMsgBz, err := json.Marshal(subspacestypes.SubspacesMsg{DeleteSubspace: nil})
	require.NoError(t, err)

	testCases := []struct {
		name      string
		msg       json.RawMessage
		shouldErr bool
		expMsgs   []sdk.Msg
	}{
		{
			name:      "wrong module message returns error",
			msg:       wrongMsgBz,
			shouldErr: true,
			expMsgs:   nil,
		},
		{
			name: "create post json message is parsed correctly",
			msg: buildCreatePostRequest(cdc, types.NewMsgCreatePost(
				1,
				2,
				"External ID",
				"test",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				[]string{"general"},
				nil,
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
				},
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgCreatePost(
				1,
				2,
				"External ID",
				"test",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				[]string{"general"},
				nil,
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
				},
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			)},
		},
		{
			name: "edit post json message is parsed correctly",
			msg: buildEditPostRequest(cdc, types.NewMsgEditPost(
				1,
				1,
				"Edited text",
				nil,
				[]string{"general"},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgEditPost(
				1,
				1,
				"Edited text",
				nil,
				[]string{"general"},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)},
		},
		{
			name: "delete post json message is parsed correctly",
			msg: buildDeletePostRequest(cdc, types.NewMsgDeletePost(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgDeletePost(
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)},
		},
		{
			name: "add post attachment json message is parsed correctly",
			msg: buildAddPostAttachmentRequest(cdc, types.NewMsgAddPostAttachment(
				1,
				1,
				types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgAddPostAttachment(
				1,
				1,
				types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)},
		},
		{
			name: "remove post attachment json message is parsed correctly",
			msg: buildRemovePostAttachmentRequest(cdc, types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)},
		},
		{
			name: "answer poll json message is parsed correctly",
			msg: buildAnswerPollRequest(cdc, types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{0, 1},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)),
			shouldErr: false,
			expMsgs: []sdk.Msg{types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{0, 1},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msgs, err := parser.ParseCustomMsgs(contractAddr, tc.msg)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expMsgs, msgs)
		})
	}
}
