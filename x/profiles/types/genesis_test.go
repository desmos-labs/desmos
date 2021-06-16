package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestValidateGenesis(t *testing.T) {
	addr1, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
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
						addr1.String(),
					),
				},
				nil,
				nil,
				types.DefaultParams(),
				types.IBCPortID,
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
			),
			shouldError: true,
		},
		{
			name: "Valid genesis returns no errors",
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						addr1.String(),
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
			),
			shouldError: false,
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
