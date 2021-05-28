package types_test

import (
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidateGenesis(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	require.NoError(t, err)

	tests := []struct {
		name      string
		genesis   *types.GenesisState
		shouldErr bool
	}{
		{
			name:      "Default genesis does not error",
			genesis:   types.DefaultGenesisState(),
			shouldErr: false,
		},
		{
			name: "Genesis with invalid subspaces returns error",
			genesis: types.NewGenesisState(
				[]types.Subspace{
					types.NewSubspace(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						"",
						"",
						"",
						types.Open,
						time.Time{},
					),
				},
			),
			shouldErr: true,
		},
		{
			name: "Genesis with duplicated subspaces returns error",
			genesis: types.NewGenesisState(
				[]types.Subspace{
					types.NewSubspace(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						"name",
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						types.Open,
						date,
					),
					types.NewSubspace(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						"name",
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						types.Open,
						date,
					),
				},
			),
			shouldErr: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.shouldErr {
				require.Error(t, types.ValidateGenesis(test.genesis))
			} else {
				require.NoError(t, types.ValidateGenesis(test.genesis))
			}
		})
	}
}
