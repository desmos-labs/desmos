package types_test

import (
	"testing"

	types2 "github.com/desmos-labs/desmos/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestValidateGenesis(t *testing.T) {
	tests := []struct {
		name        string
		genesis     *types2.GenesisState
		shouldError bool
	}{
		{
			name:        "DefaultGenesis does not error",
			genesis:     types2.DefaultGenesisState(),
			shouldError: false,
		},
		{
			name: "Genesis with invalid post errors",
			genesis: types2.NewGenesisState(
				[]types2.Post{{PostID: ""}},
				nil,
				nil,
				nil,
				nil,
				types2.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid post reaction errors",
			genesis: types2.NewGenesisState(
				[]types2.Post{},
				nil,
				[]types2.PostReactionsEntry{
					types2.NewPostReactionsEntry("1", []types2.PostReaction{{Owner: ""}}),
				},
				nil,
				nil,
				types2.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid user answer errors",
			genesis: types2.NewGenesisState(
				[]types2.Post{},
				[]types2.UserAnswer{
					types2.NewUserAnswer("1", "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", []string{}),
				},
				nil,
				nil,
				nil,
				types2.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid registered reaction errors",
			genesis: types2.NewGenesisState(
				nil,
				nil,
				nil,
				[]types2.RegisteredReaction{
					types2.NewRegisteredReaction(
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						":smile",
						"smile.jpg",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					)},
				nil,
				types2.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid reports return error",
			genesis: types2.NewGenesisState(
				nil,
				nil,
				nil,
				[]types2.RegisteredReaction{
					types2.NewRegisteredReaction(
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						":smile",
						"smile.jpg",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					)},
				[]types2.Report{
					types2.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"scam",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types2.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				types2.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid params returns errors",
			genesis: types2.NewGenesisState(
				nil,
				nil,
				nil,
				nil,
				nil,
				types2.Params{
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
				require.Error(t, types2.ValidateGenesis(test.genesis))
			} else {
				require.NoError(t, types2.ValidateGenesis(test.genesis))
			}
		})
	}
}
