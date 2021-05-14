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
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				types.NewProof(
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
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
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				types.NewProof(
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				),
				types.NewChainConfig("test-chain", "cosmos"),
				time.Now(),
			),
			expErr: fmt.Errorf("source address cannot be empty"),
		},
		{
			name: "Empty dest address returns error",
			link: types.NewLink(
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"",
				types.NewProof(
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				),
				types.NewChainConfig("test-chain", "cosmos"),
				time.Now(),
			),
			expErr: fmt.Errorf("destination address cannot be empty"),
		},
		{
			name: "Invalid proof returns error",
			link: types.NewLink(
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				types.NewProof(
					"---",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				),
				types.NewChainConfig("test-chain", "cosmos"),
				time.Now(),
			),
			expErr: fmt.Errorf("failed to decode hex string of pubkey"),
		},
		{
			name: "Invalid chain config returns error",
			link: types.NewLink(
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				types.NewProof(
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
					"143bbd3131d76232f973f84a9ea7be751243044315056dbecb968942da97e474401caa1c8f2c4ce5e48052cf44066717f2166c21a7277de9911d75c57eca598d",
				),
				types.NewChainConfig("", "cosmos"),
				time.Now(),
			),
			expErr: fmt.Errorf("chain config id cannot be empty"),
		},
		{
			name: "Invalid signature returns error",
			link: types.NewLink(
				"cosmos1u33w3u4ler4654phrpt6xqvh92ch0v6mcjrj97",
				"cosmos1qdasq0mzpajknaaj32kf9lk5nmcy8g65mddd4p",
				types.NewProof(
					"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
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
				"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: nil,
		},
		{
			name: "Invalid hex string of pub key",
			proof: types.NewProof(
				"=",
				"1a18d5f012ce0e8258fd3455c01b48249bb019231e416c4323ab2bb170b4ad0951b370138d2ea69a376feb942d3c619c9152d63a6d2e0232aaff77162df66636",
			),
			expErr: fmt.Errorf("failed to decode hex string of pubkey"),
		},
		{
			name: "Invalid hex string of signature",
			proof: types.NewProof(
				"02bcd0738e3b7e0f6650c8e6eb10bd4266fd6818c92a0283b4cb0884f046051c3e",
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
