package types_test

import (
	"github.com/desmos-labs/desmos/x/subspaces/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTokenomics_Validate(t *testing.T) {
	tests := []struct {
		name       string
		tokenomics types.Tokenomics
		shouldErr  bool
	}{
		{
			name: "Invalid subspace returns error",
			tokenomics: types.NewTokenomics(
				"",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				[]byte("message"),
			),
			shouldErr: true,
		},
		{
			name: "Invalid contract address returns error",
			tokenomics: types.NewTokenomics(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"",
				[]byte("message"),
			),
			shouldErr: true,
		},
		{
			name: "Invalid contract message returns error",
			tokenomics: types.NewTokenomics(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				nil,
			),
			shouldErr: true,
		},
		{
			name: "Valid tokenomics returns no error",
			tokenomics: types.NewTokenomics(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				[]byte("message"),
			),
			shouldErr: false,
		},
	}

	for _, testCase := range tests {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldErr {
				require.Error(t, tc.tokenomics.Validate())
			} else {
				require.NoError(t, tc.tokenomics.Validate())
			}
		})
	}
}
