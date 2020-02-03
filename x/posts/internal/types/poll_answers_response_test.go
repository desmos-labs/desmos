package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestPollAnswersQueryResponse_String(t *testing.T) {
	testOwner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	answers := types.AnswersDetails{
		Answers: []uint64{1, 2},
		User:    testOwner,
	}

	pollResponse := types.PollAnswersQueryResponse{
		PostID:         types.PostID(0),
		AnswersDetails: types.UsersAnswersDetails{answers},
	}

	assert.Equal(t, "Post ID [0] - Answers Details:\nUser: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns \nAnswers IDs: 1 2", pollResponse.String())
}
