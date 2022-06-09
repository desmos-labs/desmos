package types_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func TestLinkChainAccountPacketData_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		packet    types.LinkChainAccountPacketData
		shouldErr bool
	}{
		{
			name: "null source address returns error",
			packet: types.LinkChainAccountPacketData{
				SourceProof: types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
				SourceChainConfig:  types.NewChainConfig("cosmos"),
				DestinationAddress: "cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				DestinationProof: types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
			},
			shouldErr: true,
		},
		{
			name: "invalid source proof returns error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), &types.SingleSignatureData{}, "wrong"),
				types.NewChainConfig("cosmos"),
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
			),
			shouldErr: true,
		},
		{
			name: "invalid chain config returns error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
				types.NewChainConfig(""),
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
			),
			shouldErr: true,
		},
		{
			name: "invalid destination address returns error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
				types.NewChainConfig("cosmos"),
				"",
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
			),
			shouldErr: true,
		},
		{
			name: "invalid destination proof returns error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
				types.NewChainConfig("cosmos"),
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(secp256k1.GenPrivKey().PubKey(), &types.SingleSignatureData{}, "wrong"),
			),
			shouldErr: true,
		},
		{
			name: "valid packet returns no error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
				types.NewChainConfig("cosmos"),
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"),
					"706c61696e5f74657874",
				),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.packet.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestLinkChainAccountPacketData_GetBytes(t *testing.T) {
	packetData := types.NewLinkChainAccountPacketData(
		types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
		types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"), "plain_text"),
		types.NewChainConfig("cosmos"),
		"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
		types.NewProof(secp256k1.GenPrivKey().PubKey(), profilestesting.SingleSignatureProtoFromHex("032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561"), "plain_text"),
	)
	_, err := packetData.GetBytes()
	require.NoError(t, err)
}
