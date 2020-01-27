package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestPollUserAnswersQueryResponse_String(t *testing.T) {
	testOwner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	pollResponse := types.PollUserAnswersQueryResponse{
		PostID:  types.PostID(0),
		User:    testOwner,
		Answers: []uint64{1, 2},
	}

	assert.Equal(t, "Post ID [0] - User [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns] \n Answers: answerID [1] answerID [2]", pollResponse.String())
}

func TestPollAnswersAmountResponse_String(t *testing.T) {
	pollResponse := types.PollAnswersAmountResponse{
		PostID:        types.PostID(1),
		AnswersAmount: sdk.NewInt(1),
	}

	assert.Equal(t, "Post ID [1] - Answers Amount [1]", pollResponse.String())
}

func TestPollAnswerVotesResponse_String(t *testing.T) {
	pollResponse := types.PollAnswerVotesResponse{
		PostID:      types.PostID(1),
		AnswerID:    uint64(2),
		VotesAmount: sdk.NewInt(10),
	}

	assert.Equal(t, "Post ID [1] - Answers ID [2] Votes Amount [10]", pollResponse.String())
}
