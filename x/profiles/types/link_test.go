package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/stretchr/testify/require"
)

func TestLink_Validate(t *testing.T) {
	tests := []struct {
		name   string
		link   types.Link
		expErr error
	}{
		{
			name: "Valid proof",
			link: types.NewLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				),
				types.NewChainConfig("test-chain", "cosmos"),
				time.Now(),
			),
			expErr: nil,
		},
		{
			name: "Empty src address returns error",
			link: types.NewLink(
				"",
				types.NewProof(
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				),
				types.NewChainConfig("test-chain", "cosmos"),
				time.Now(),
			),
			expErr: fmt.Errorf("source address cannot be empty"),
		},
		{
			name: "Invalid proof returns error",
			link: types.NewLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					"=",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				),
				types.NewChainConfig("test-chain", "cosmos"),
				time.Now(),
			),
			expErr: fmt.Errorf("failed to decode hex string of pubkey"),
		},
		{
			name: "Invalid chain config returns error",
			link: types.NewLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				),
				types.NewChainConfig("", "cosmos"),
				time.Now(),
			),
			expErr: fmt.Errorf("chain config id cannot be empty"),
		},
		{
			name: "Invalid signature returns error",
			link: types.NewLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				types.NewProof(
					"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
					"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
				),
				types.NewChainConfig("test-chain", "cosmos"),
				time.Now(),
			),
			expErr: fmt.Errorf("failed to verify signature"),
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.link.Validate())
		})
	}
}

func TestProof_Validate(t *testing.T) {
	tests := []struct {
		name   string
		proof  types.Proof
		expErr error
	}{
		{
			name: "Valid proof",
			proof: types.NewProof(
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
			),
			expErr: nil,
		},
		{
			name: "Invalid hex string of pub key",
			proof: types.NewProof(
				"=",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
			),
			expErr: fmt.Errorf("failed to decode hex string of pubkey"),
		},
		{
			name: "Invalid hex string of signature",
			proof: types.NewProof(
				"032086ede8d4bce29fe364a94744ca71dbeaf370221ba20f9716a165c54b079561",
				"=",
			),
			expErr: fmt.Errorf("failed to decode hex string of signature"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.proof.Validate())
		})
	}
}

func TestChainConfig_Validate(t *testing.T) {
	tests := []struct {
		name        string
		chainConfig types.ChainConfig
		expErr      error
	}{
		{
			name: "Valid chain config",
			chainConfig: types.NewChainConfig(
				"testnet",
				"cosmos",
			),
			expErr: nil,
		},
		{
			name: "Empty id returns error",
			chainConfig: types.NewChainConfig(
				"",
				"cosmos",
			),
			expErr: fmt.Errorf("chain config id cannot be empty"),
		},
		{
			name: "Empty prefix returns error",
			chainConfig: types.NewChainConfig(
				"testnet",
				"",
			),
			expErr: fmt.Errorf("bech32 addr prefix config id cannot be empty"),
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.chainConfig.Validate())
		})
	}
}
