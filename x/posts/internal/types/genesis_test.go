package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestValidateGenesis(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")

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
				Posts:     types.Posts{types.Post{PostID: types.PostID(0)}},
				Reactions: map[string]types.Reactions{},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid reaction errors",
			genesis: types.GenesisState{
				Posts: types.Posts{},
				Reactions: map[string]types.Reactions{
					"1": {types.Reaction{Owner: nil}},
				},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid poll answers errors",
			genesis: types.GenesisState{
				Posts: types.Posts{},
				PollAnswers: map[string]types.UsersAnswersDetails{
					"1": {
						types.NewAnswersDetails([]uint64{}, user),
					},
				},
				Reactions: map[string]types.Reactions{},
			},
			shouldError: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.shouldError {
				assert.Error(t, types.ValidateGenesis(test.genesis))
			} else {
				assert.NoError(t, types.ValidateGenesis(test.genesis))
			}
		})
	}
}
