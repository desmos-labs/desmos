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
			name: "invalid type returns error",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1)))),
			}),
			shouldErr: true,
		},
		{
			name: "invalid amount returns error",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("/desmos.profiles.v2.SaveProfile", sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1)}}),
			}),
			shouldErr: true,
		},
		{
			name: "duplicated min fee returns error",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("/desmos.profiles.v2.SaveProfile", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1)))),
				types.NewMinFee("/desmos.profiles.v2.SaveProfile", sdk.NewCoins(sdk.NewCoin("photino", sdk.NewInt(1)))),
			}),
			shouldErr: true,
		},
		{
			name:      "default params returns no error",
			params:    types.DefaultParams(),
			shouldErr: false,
		},
		{
			name: "custom params returns no error",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("/desmos.profiles.v2.SaveProfile", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1)))),
				types.NewMinFee("/desmos.profiles.v2.CreateChainLink", sdk.NewCoins(sdk.NewCoin("photino", sdk.NewInt(1)))),
			}),
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
