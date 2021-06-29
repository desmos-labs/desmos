package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
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
				nil,
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid user answer errors",
			genesis: types.NewGenesisState(
				[]types.Post{},
				[]types.UserAnswer{
					types.NewUserAnswer("1", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", []string{}),
				},
				nil,
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
				nil,
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid reports return error",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				nil,
				[]types.Report{
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						[]string{"scam"},
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						[]string{""},
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
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
				nil,
				types.Params{
					MaxPostMessageLength:                    sdk.NewInt(-1),
					MaxAdditionalAttributesFieldsNumber:     sdk.NewInt(-1),
					MaxAdditionalAttributesFieldValueLength: sdk.NewInt(-1),
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
