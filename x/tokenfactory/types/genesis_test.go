package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		name      string
		genState  *types.GenesisState
		shouldErr bool
	}{
		{
			name:     "default is valid returns no error",
			genState: types.DefaultGenesis(),
		},
		{
			name: "valid genesis state returns no error",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						},
					},
				},
			},
		},
		{
			name: "different admin from creator returns no error",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						},
					},
				},
			},
		},
		{
			name: "empty admin returns no error",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
		},
		{
			name: "no admin returns no error",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
					},
				},
			},
		},
		{
			name: "invalid admin returns error",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "moose",
						},
					},
				},
			},
			shouldErr: true,
		},
		{
			name: "multiple denoms returns no error",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/litecoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
		},
		{
			name: "duplicate denoms returns error",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
					{
						Denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
			shouldErr: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
