package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestValidateGenesis(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	pubKey, err := sdk.GetPubKeyFromBech32(
		sdk.Bech32PubKeyTypeAccPub,
		"cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec",
	)
	require.NoError(t, err)

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
			name: "Invalid params returns error",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(-1), sdk.NewInt(10)),
					types.DefaultDTagParams(),
					types.DefaultMaxBioLength,
				),
				types.IBCPortID,
				nil,
				nil,
			),
			shouldError: true,
		},
		{
			name: "Invalid DTag requests returns error",
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"",
						addr.String(),
					),
				},
				nil,
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid relationship returns error",
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"",
						"",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid users blocks return error",
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid chain links return error",
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.DefaultParams(),
				types.IBCPortID,
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(pubKey, "sig_hex", "addr"),
						types.NewChainConfig(""),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
				nil,
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid application link returns error",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				[]types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid port ID returns error",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				types.DefaultParams(),
				"1235$512",
				nil,
				nil,
			),
			shouldError: true,
		},
		{
			name: "Valid genesis returns no errors",
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						addr.String(),
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.DefaultParams(),
				types.IBCPortID,
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(pubKey, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0"),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(pubKey, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0"),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
				[]types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
			),
			shouldError: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types.ValidateGenesis(test.genesis)
			if test.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
