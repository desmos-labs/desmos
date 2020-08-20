package models_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/types/models"
	"github.com/desmos-labs/desmos/x/posts/types/models/polls"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestPollAnswersQueryResponse_String(t *testing.T) {
	testOwner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	pollResponse := models.PollAnswersQueryResponse{
		PostID: "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
		AnswersDetails: polls.NewUserAnswers(
			polls.NewUserAnswer([]polls.AnswerID{1, 2}, testOwner),
		),
	}

	require.Equal(t, "Post ID [dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1] - Answers Details:\nUser: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns \nAnswers IDs: 1 2", pollResponse.String())
}
