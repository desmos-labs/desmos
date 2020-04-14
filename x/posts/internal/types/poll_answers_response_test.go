package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestPollAnswersQueryResponse_String(t *testing.T) {
	testOwner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	answers := types.UserAnswer{
		Answers: []types.AnswerID{1, 2},
		User:    testOwner,
	}

	pollResponse := types.PollAnswersQueryResponse{
		PostID:         id,
		AnswersDetails: types.UserAnswers{answers},
	}

	require.Equal(t, "Post ID [3139646530326531303563363861363065343563323839626666] - Answers Details:\nUser: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns \nAnswers IDs: 1 2", pollResponse.String())
}
