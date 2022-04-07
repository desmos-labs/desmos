package types_test

import (
	"testing"

	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
)

func TestValidateParams(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.Params
		shouldErr bool
	}{
		{
			name: "invalid params returns error",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1)))),
			}),
			shouldErr: true,
		},
		{
			name:      "default params returns no error",
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

func TestValidateMinFeesParam(t *testing.T) {
	testCases := []struct {
		name        string
		requiredFee interface{}
		shouldErr   bool
	}{
		{
			name:        "invalid param type returns error",
			requiredFee: "param",
			shouldErr:   true,
		},
		{
			name: "invalid param value returns error",
			requiredFee: []types.MinFee{
				types.NewMinFee("", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1)))),
			},
			shouldErr: true,
		},
		{
			name: "valid params returns no errors",
			requiredFee: []types.MinFee{
				types.NewMinFee(
					sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
					sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000))),
				),
			},
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateMinFeesParam(tc.requiredFee)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
