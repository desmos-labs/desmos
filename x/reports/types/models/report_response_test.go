package models_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/stretchr/testify/require"
)

func TestReportsQueryResponse_String(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	postID := posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	require.NoError(t, err)
	reports := types.Reports{
		types.NewReport("scam", "it's a trap", creator),
		types.NewReport("violence", "it's a trap", creator),
	}

	repQueryResp := types.NewReportResponse(postID, reports)

	require.Equal(t, "Post ID: 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af\n Reports: Type - Message - ReceivingUser\nscam - it's a trap - cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\nviolence - it's a trap - cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", repQueryResp.String())
}

func TestReportsQueryResponse_MarshalJSON(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	postID := posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	require.NoError(t, err)
	reports := types.Reports{
		types.NewReport("scam", "it's a trap", creator),
		types.NewReport("violence", "it's a trap", creator),
	}

	repQueryResp := types.NewReportResponse(postID, reports)

	tests := []struct {
		name        string
		response    types.ReportsQueryResponse
		expResponse string
	}{
		{
			name:        "Response with non-empty reports",
			response:    repQueryResp,
			expResponse: "{\"post_id\":\"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af\",\"reports\":[{\"type\":\"scam\",\"message\":\"it's a trap\",\"user\":\"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\"},{\"type\":\"violence\",\"message\":\"it's a trap\",\"user\":\"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\"}]}",
		},
		{
			name:        "Response with empty reports",
			response:    types.ReportsQueryResponse{PostID: postID, Reports: types.Reports{}},
			expResponse: "{\"post_id\":\"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af\",\"reports\":[]}",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			jsonData, err := json.Marshal(&test.response)
			require.NoError(t, err)
			require.Equal(t, test.expResponse, string(jsonData))
		})
	}

}
