package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func TestValidateGenesis(t *testing.T) {
	tests := []struct {
		name        string
		genesis     *types.GenesisState
		shouldError bool
	}{
		{
			name:        "DefaultGenesis does not error",
			genesis:     types.DefaultGenesisState(),
			shouldError: false,
		},
		{
			name: "Genesis with invalid post errors",
			genesis: types.NewGenesisState(
				[]types.Post{{PostID: ""}},
				nil,
				nil,
				nil,
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid post reaction errors",
			genesis: types.NewGenesisState(
				[]types.Post{},
				nil,
				[]types.PostReactionsEntry{
					types.NewPostReactionsEntry("1", []types.PostReaction{{Owner: ""}}),
				},
				nil,
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid poll answer errors",
			genesis: types.NewGenesisState(
				[]types.Post{},
				[]types.UserAnswersEntry{
					types.NewUserAnswersEntry("1", []types.UserAnswer{
						types.NewUserAnswer([]string{""}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
					}),
				},
				nil,
				nil,
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid registered reaction errors",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				[]types.RegisteredReaction{
					types.NewRegisteredReaction(
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						":smile",
						"smile.jpg",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					)},
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid params returns errors",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				nil,
				types.Params{
					MaxPostMessageLength:            sdk.NewInt(-1),
					MaxOptionalDataFieldsNumber:     sdk.NewInt(-1),
					MaxOptionalDataFieldValueLength: sdk.NewInt(-1),
				},
			),
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
