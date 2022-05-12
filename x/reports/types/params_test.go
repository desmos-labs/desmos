package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

func TestParams_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.Params
		shouldErr bool
	}{
		{
			name: "invalid reasons return error",
			params: types.NewParams(
				types.NewStandardReasons(
					types.NewStandardReason(1, "Spam", "This content is spam"),
					types.NewStandardReason(1, "Harm", "This content contains self-harm/suicide images"),
				),
			),
			shouldErr: true,
		},
		{
			name: "valid params return no error",
			params: types.NewParams(
				types.NewStandardReasons(
					types.NewStandardReason(1, "Spam", "This content is spam"),
					types.NewStandardReason(2, "Harm", "This content contains self-harm/suicide images"),
				),
			),
			shouldErr: false,
		},
		{
			name:      "default params return no error",
			params:    types.DefaultParams(),
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
