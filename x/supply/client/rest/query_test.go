package rest_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/v3/x/supply/client/rest"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_parseQueryParams(t *testing.T) {
	testCases := []struct {
		name          string
		vars          map[string]string
		expDenom      string
		expMultiplier int64
		expErr        error
	}{
		{
			name:          "invalid denom parsing returns error",
			vars:          map[string]string{"denom": "", "divider": ""},
			expDenom:      "",
			expMultiplier: 0,
			expErr:        fmt.Errorf("invalid empty denom string"),
		},
		{
			name:          "zero divider parsing return 1",
			vars:          map[string]string{"denom": "udsm", "divider": "0"},
			expDenom:      "",
			expMultiplier: 1,
			expErr:        nil,
		},
		{
			name:          "empty divider parsing return 1",
			vars:          map[string]string{"denom": "udsm", "divider": ""},
			expDenom:      "",
			expMultiplier: 1,
			expErr:        nil,
		},
		{
			name:          "invalid divider parsing return 1",
			vars:          map[string]string{"denom": "udsm", "divider": "----"},
			expDenom:      "",
			expMultiplier: 1,
			expErr:        nil,
		},
		{
			name:          "valid vars are returned correctly",
			vars:          map[string]string{"denom": "udsm", "divider": "10"},
			expDenom:      "udsm",
			expMultiplier: 10,
			expErr:        nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			denom, divider, err := rest.ParseQueryParams(tc.vars)
			if tc.expErr != nil {
				require.Error(t, err)
			}
			if tc.expDenom != "" {
				require.Equal(t, tc.expDenom, denom)
			}
			if tc.expMultiplier != 0 {
				require.Equal(t, tc.expMultiplier, divider)
			}
		})
	}
}
