package types_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestLinkChainAccountPacketData_Validate(t *testing.T) {
	tests := []struct {
		name      string
		packet    types.LinkChainAccountPacketData
		shouldErr bool
	}{
		{
			name: "Null source address returns error",
			packet: types.LinkChainAccountPacketData{
				SourceProof: types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
				SourceChainConfig:  types.NewChainConfig("cosmos"),
				DestinationAddress: "cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				DestinationProof: types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
			},
			shouldErr: true,
		},
		{
			name: "Invalid source proof returns error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "wrong"),
				types.NewChainConfig("cosmos"),
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
			),
			shouldErr: true,
		},
		{
			name: "Invalid chain config returns error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
				types.NewChainConfig(""),
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
			),
			shouldErr: true,
		},
		{
			name: "Invalid destination address returns error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
				types.NewChainConfig("cosmos"),
				"",
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
			),
			shouldErr: true,
		},
		{
			name: "Invalid destination proof returns error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
				types.NewChainConfig("cosmos"),
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "wrong"),
			),
			shouldErr: true,
		},
		{
			name: "Valid packet returns no error",
			packet: types.NewLinkChainAccountPacketData(
				types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
				types.NewChainConfig("cosmos"),
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					secp256k1.GenPrivKey().PubKey(),
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"plain_text",
				),
			),
			shouldErr: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.packet.Validate()
			if test.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestLinkChainAccountPacketData_GetBytes(t *testing.T) {
	p := types.NewLinkChainAccountPacketData(
		types.NewBech32Address("cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70", "cosmos"),
		types.NewProof(secp256k1.GenPrivKey().PubKey(), "032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561", "plain_text"),
		types.NewChainConfig("cosmos"),
		"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
		types.NewProof(secp256k1.GenPrivKey().PubKey(), "032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561", "plain_text"),
	)
	_, err := p.GetBytes()
	require.NoError(t, err)
}
