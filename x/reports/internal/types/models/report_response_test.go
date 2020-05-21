package models_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
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

	require.Equal(t, "Post ID: 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af\n Reports: Type - Message - User\nscam - it's a trap - cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\nviolence - it's a trap - cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", repQueryResp.String())
}
