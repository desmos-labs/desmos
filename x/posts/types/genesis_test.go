package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
	"github.com/stretchr/testify/require"
)

func TestValidateGenesis(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name        string
		genesis     types.GenesisState
		shouldError bool
	}{
		{
			name:        "DefaultGenesis does not error",
			genesis:     types.DefaultGenesisState(),
			shouldError: false,
		},
		{
			name: "Genesis with invalid post errors",
			genesis: types.GenesisState{
				Posts:         types.Posts{types.Post{PostID: types.PostID("")}},
				PostReactions: map[string]types.PostReactions{},
				Params:        types.DefaultParams(),
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid post reaction errors",
			genesis: types.GenesisState{
				Posts: types.Posts{},
				PostReactions: map[string]types.PostReactions{
					"1": {types.PostReaction{Owner: nil}},
				},
				Params: types.DefaultParams(),
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid poll answers errors",
			genesis: types.GenesisState{
				Posts: types.Posts{},
				UsersPollAnswers: map[string]types.UserAnswers{
					"1": {
						types.NewUserAnswer([]types.AnswerID{}, user),
					},
				},
				PostReactions: map[string]types.PostReactions{},
				Params:        types.DefaultParams(),
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid registered reaction errors",
			genesis: types.GenesisState{
				Posts: types.Posts{},
				RegisteredReactions: types.Reactions{types.NewReaction(user, ":smile", "smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
				Params: types.DefaultParams(),
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid params returns errors",
			genesis: types.GenesisState{
				Posts:               types.Posts{},
				RegisteredReactions: types.Reactions{},
				PostReactions:       map[string]types.PostReactions{},
				UsersPollAnswers:    map[string]types.UserAnswers{},
				Params: types.Params{
					MaxPostMessageLength:            sdk.NewInt(-1),
					MaxOptionalDataFieldsNumber:     sdk.NewInt(-1),
					MaxOptionalDataFieldValueLength: sdk.NewInt(-1),
				},
			},
			shouldError: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.shouldError {
				require.Error(t, types.ValidateGenesis(test.genesis))
			} else {
				require.NoError(t, types.ValidateGenesis(test.genesis))
			}
		})
	}
}
