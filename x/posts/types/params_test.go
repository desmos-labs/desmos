package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func TestParams_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.Params
		shouldErr bool
	}{
		{
			name:      "invalid max text length returns error",
			params:    types.NewParams(sdk.NewInt(-1)),
			shouldErr: true,
		},
		{
			name:      "default params returns no error",
			params:    types.DefaultParams(),
			shouldErr: false,
		},
		{
			name:      "valid params returns no error",
			params:    types.NewParams(sdk.NewInt(100)),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
