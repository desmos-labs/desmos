package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/x/tokenfactory/types"
)

func TestDeconstructDenom(t *testing.T) {

	for _, tc := range []struct {
		name             string
		denom            string
		expectedSubdenom string
		err              error
	}{
		{
			name:  "empty is invalid",
			denom: "",
			err:   types.ErrInvalidDenom,
		},
		{
			name:             "normal",
			denom:            "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
			expectedSubdenom: "bitcoin",
		},
		{
			name:             "multiple slashes in subdenom",
			denom:            "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin/1",
			expectedSubdenom: "bitcoin/1",
		},
		{
			name:             "no subdenom",
			denom:            "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/",
			expectedSubdenom: "",
		},
		{
			name:  "incorrect prefix",
			denom: "ibc/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/bitcoin",
			err:   types.ErrInvalidDenom,
		},
		{
			name:             "subdenom of only slashes",
			denom:            "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/////",
			expectedSubdenom: "////",
		},
		{
			name:  "too long name",
			denom: "factory/cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69/adsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsf",
			err:   types.ErrInvalidDenom,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			expectedCreator := "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"
			creator, subdenom, err := types.DeconstructDenom(tc.denom)
			if tc.err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, expectedCreator, creator)
				require.Equal(t, tc.expectedSubdenom, subdenom)
			}
		})
	}
}

func TestGetTokenDenom(t *testing.T) {
	for _, tc := range []struct {
		name      string
		creator   string
		subdenom  string
		shouldErr bool
	}{
		{
			name:     "normal returns no error",
			creator:  "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			subdenom: "bitcoin",
		},
		{
			name:     "multiple slashes in subdenom returns no error",
			creator:  "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			subdenom: "bitcoin/1",
		},
		{
			name:     "no subdenom returns no error",
			creator:  "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			subdenom: "",
		},
		{
			name:     "subdenom of only slashes returns no error",
			creator:  "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			subdenom: "/////",
		},
		{
			name:      "too long name returns error",
			creator:   "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			subdenom:  "adsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsf",
			shouldErr: true,
		},
		{
			name:     "subdenom is exactly max length returns no error",
			creator:  "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			subdenom: "bitcoinfsadfsdfeadfsafwefsefsefsdfsdafasefsf",
		},
		{
			name:     "creator is exactly max length return no error",
			creator:  "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69jhgjhgkhjklhkjhkjhgjhgjgjghelu",
			subdenom: "bitcoin",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := types.GetTokenDenom(tc.creator, tc.subdenom)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
