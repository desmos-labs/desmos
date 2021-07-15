package types_test

import (
	"testing"
	"time"

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
				nil,
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with post reaction to non existent post errors",
			genesis: types.NewGenesisState(
				[]types.Post{},
				nil,
				[]types.PostReaction{
					{PostID: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", Owner: ""},
				},
				nil,
				nil,
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid post reaction errors",
			genesis: types.NewGenesisState(
				[]types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "",
						"Message",
						types.CommentsStateBlocked,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				nil,
				[]types.PostReaction{
					{PostID: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", Owner: ""},
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
			name: "Genesis with user answer to non existent post errors",
			genesis: types.NewGenesisState(
				[]types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "",
						"Message",
						types.CommentsStateBlocked,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				[]types.UserAnswer{
					types.NewUserAnswer(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						[]string{},
					),
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
				[]types.RegisteredReaction{
					types.NewRegisteredReaction(
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						":smile",
						"smile.jpg",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					)},
				[]types.Report{
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"scam",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"",
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
