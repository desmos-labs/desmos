package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestValidateGenesis(t *testing.T) {
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
