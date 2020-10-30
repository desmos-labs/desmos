package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/reports/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	postserrors "github.com/desmos-labs/desmos/x/posts/types/errors"

	"github.com/stretchr/testify/require"
)

func TestMsgReportPost_Route(t *testing.T) {
	msg := types.NewMsgReportPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"type",
		"message",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "reports", msg.Route())
}

func TestMsgReportPost_Type(t *testing.T) {
	msg := types.NewMsgReportPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"type",
		"message",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, "report_post", msg.Type())
}

func TestMsgReportPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgReportPost
		error error
	}{
		{
			name: "invalid post ID returns error",
			msg: types.NewMsgReportPost(
				"123",
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(postserrors.ErrInvalidPostID, "123"),
		},
		{
			name: "invalid reports returns error",
			msg: types.NewMsgReportPost(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "report message cannot be empty"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgReportPost_GetSignBytes(t *testing.T) {
	msg := types.NewMsgReportPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"type",
		"message",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	expected := `{"type":"desmos/MsgReportPost","value":{"post_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","report":{"message":"message","type":"type","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestNewMsgReportPost_GetSigners(t *testing.T) {
	msg := types.NewMsgReportPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"type",
		"message",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	)
	require.Equal(t, 1, len(msg.GetSigners()))
	require.Equal(t, msg.User, msg.GetSigners()[0])
}
