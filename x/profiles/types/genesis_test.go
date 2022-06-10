package types_test

import (
	"testing"
	"time"

	"github.com/desmos-labs/desmos/v3/testutil/profilestesting"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name      string
		genesis   *types.GenesisState
		shouldErr bool
	}{
		{
			name:      "default genesis does not error",
			genesis:   types.DefaultGenesisState(),
			shouldErr: false,
		},
		{
			name: "invalid params returns error",
			genesis: types.NewGenesisState(
				nil,
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(-1), sdk.NewInt(10)),
					types.DefaultDTagParams(),
					types.DefaultBioParams(),
					types.DefaultOracleParams(),
				),
				types.IBCPortID,
				nil,
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid DTag requests returns error",
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					),
				},
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid chain links return error",
			genesis: types.NewGenesisState(
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(
							profilestesting.PubKeyFromBech32("cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec"),
							&types.SingleSignatureData{},
							"addr",
						),
						types.NewChainConfig(""),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid default external address return error",
			genesis: types.NewGenesisState(
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				[]types.DefaultExternalAddressEntry{types.NewDefaultExternalAddressEntry("", "", "")},
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid application link returns error",
			genesis: types.NewGenesisState(
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
				[]types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
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
			shouldErr: true,
		},
		{
			name: "invalid port id returns error",
			genesis: types.NewGenesisState(
				nil,
				types.DefaultParams(),
				"1235$512",
				nil,
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "valid data returns no errors",
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				types.DefaultParams(),
				types.IBCPortID,
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(
							profilestesting.PubKeyFromBech32("cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec"),
							profilestesting.SingleSignatureProtoFromHex("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
							"636f736d6f7331786d717563393434687a75366e3671746c6a6365786b7568687a37366d75637874676d357830",
						),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(
							profilestesting.PubKeyFromBech32("cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec"),
							profilestesting.SingleSignatureProtoFromHex("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
							"636f736d6f7331786d717563393434687a75366e3671746c6a6365786b7568687a37366d75637874676d357830",
						),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
				nil,
				[]types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
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
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateGenesis(tc.genesis)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
