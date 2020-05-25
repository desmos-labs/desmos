package msgs_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/stretchr/testify/require"
)

func TestMsgReportPost_Route(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	postID := posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	msgReport := types.NewMsgReportPost(postID, "type", "message", creator)
	actual := msgReport.Route()
	require.Equal(t, "reports", actual)
}

func TestMsgReportPost_Type(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	postID := posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	msgReport := types.NewMsgReportPost(postID, "type", "message", creator)
	actual := msgReport.Type()
	require.Equal(t, "report_post", actual)
}

func TestMsgReportPost_ValidateBasic(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	postID := posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	tests := []struct {
		name  string
		msg   types.MsgReportPost
		error error
	}{
		{
			name:  "invalid post ID returns error",
			msg:   types.NewMsgReportPost("123", "type", "message", creator),
			error: fmt.Errorf("invalid postID: 123"),
		},
		{
			name:  "invalid reports returns error",
			msg:   types.NewMsgReportPost(postID, "scam", "", creator),
			error: fmt.Errorf("reports's message cannot be empty"),
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
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	postID := posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	msgReport := types.NewMsgReportPost(postID, "type", "message", creator)
	actual := msgReport.GetSignBytes()

	expected := `{"type":"desmos/MsgReportPost","value":{"post_id":"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af","report":{"message":"message","type":"type","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}}`
	require.Equal(t, expected, string(actual))
}

func TestNewMsgReportPost_GetSigners(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	postID := posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	msgReport := types.NewMsgReportPost(postID, "type", "message", creator)
	actual := msgReport.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgReport.Report.User, actual[0])
}
