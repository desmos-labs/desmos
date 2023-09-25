package types_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/x/supply/types"
)

func Test_NewDividerFromRawInt(t *testing.T) {
	testCases := []struct {
		name        string
		rawDivider  uint64
		expectedInt math.Int
	}{
		{
			name:        "raw divider equal to zero return 1",
			rawDivider:  0,
			expectedInt: math.NewInt(1),
		},
		{
			name:        "raw divider equal to 2 return 100",
			rawDivider:  2,
			expectedInt: math.NewInt(100),
		},
		{
			name:        "raw divider equal to 3 return 1000",
			rawDivider:  3,
			expectedInt: math.NewInt(1000),
		},
		{
			name:        "raw divider equal to 4 return 10000",
			rawDivider:  4,
			expectedInt: math.NewInt(10000),
		},
		{
			name:        "raw divider equal to 5 return 100000",
			rawDivider:  5,
			expectedInt: math.NewInt(100000),
		},
		{
			name:        "raw divider equal to 6 return 1000000",
			rawDivider:  6,
			expectedInt: math.NewInt(1000000),
		},
		{
			name:        "raw divider equal to 6 return 1000000",
			rawDivider:  7,
			expectedInt: math.NewInt(10000000),
		},
		{
			name:        "raw divider equal to 6 return 1000000",
			rawDivider:  8,
			expectedInt: math.NewInt(100000000),
		},
		{
			name:        "raw divider equal to 6 return 1000000",
			rawDivider:  9,
			expectedInt: math.NewInt(1000000000),
		},
		{
			name:        "raw divider equal to 6 return 1000000",
			rawDivider:  10,
			expectedInt: math.NewInt(10000000000),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			divider := types.NewDividerPoweredByExponent(tc.rawDivider)
			require.Equal(t, tc.expectedInt, divider)
		})
	}
}
