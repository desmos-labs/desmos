package types_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func TestValidateParams(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.Params
		shouldErr bool
	}{
		{
			name: "invalid nickname param returns error",
			params: types.NewParams(
				types.NewNicknameParams(math.NewInt(1), math.NewInt(1000)),
				types.DefaultDTagParams(),
				types.DefaultBioParams(),
				types.DefaultOracleParams(),
				types.DefaultAppLinksParams(),
			),
			shouldErr: true,
		},
		{
			name: "invalid DTag param returns error",
			params: types.NewParams(
				types.DefaultNicknameParams(),
				types.NewDTagParams("regEx", math.NewInt(3), math.NewInt(-30)),
				types.DefaultBioParams(),
				types.DefaultOracleParams(),
				types.DefaultAppLinksParams(),
			),
			shouldErr: true,
		},
		{
			name: "invalid bio param returns error",
			params: types.NewParams(
				types.DefaultNicknameParams(),
				types.DefaultDTagParams(),
				types.NewBioParams(math.NewInt(-1000)),
				types.DefaultOracleParams(),
				types.DefaultAppLinksParams(),
			),
			shouldErr: true,
		},
		{
			name: "invalid oracle params return error",
			params: types.NewParams(
				types.DefaultNicknameParams(),
				types.DefaultDTagParams(),
				types.DefaultBioParams(),
				types.NewOracleParams(
					0,
					0,
					0,
					0,
					0,
					sdk.NewCoins()...,
				),
				types.DefaultAppLinksParams()),
			shouldErr: true,
		},
		{
			name: "invalid app links params return error",
			params: types.NewParams(
				types.DefaultNicknameParams(),
				types.DefaultDTagParams(),
				types.DefaultBioParams(),
				types.DefaultOracleParams(),
				types.NewAppLinksParams(time.Duration(0)),
			),
			shouldErr: true,
		},
		{
			name: "valid params return no error",
			params: types.NewParams(
				types.DefaultNicknameParams(),
				types.DefaultDTagParams(),
				types.DefaultBioParams(),
				types.DefaultOracleParams(),
				types.DefaultAppLinksParams(),
			),
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

func TestValidateNicknameParams(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.NicknameParams
		shouldErr bool
	}{
		{
			name:      "invalid max returns error - nil min",
			params:    types.NicknameParams{MinLength: sdkmath.Int{}, MaxLength: math.NewInt(10)},
			shouldErr: true,
		},
		{
			name:      "invalid min returns error",
			params:    types.NewNicknameParams(math.NewInt(1), math.NewInt(1000)),
			shouldErr: true,
		},
		{
			name:      "invalid max returns error - nil max",
			params:    types.NicknameParams{MinLength: math.NewInt(2), MaxLength: sdkmath.Int{}},
			shouldErr: true,
		},
		{
			name:      "invalid max returns error - lower than min",
			params:    types.NewNicknameParams(math.NewInt(2), math.NewInt(-10)),
			shouldErr: true,
		},
		{
			name:      "valid values return no error",
			params:    types.NewNicknameParams(math.NewInt(2), math.NewInt(10)),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateNicknameParams(tc.params)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateDTagParams(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.DTagParams
		shouldErr bool
	}{
		{
			name:      "empty regex returns error",
			params:    types.NewDTagParams("", math.NewInt(3), math.NewInt(30)),
			shouldErr: true,
		},
		{
			name:      "invalid min returns error",
			params:    types.NewDTagParams("regExParam", math.NewInt(1), math.NewInt(30)),
			shouldErr: true,
		},
		{
			name:      "invalid max returns error - lower than min",
			params:    types.NewDTagParams("regExParam", math.NewInt(3), math.NewInt(-30)),
			shouldErr: true,
		},
		{
			name:      "valid params return no error",
			params:    types.NewDTagParams("regExParam", math.NewInt(3), math.NewInt(30)),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateDTagParams(tc.params)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateBioParams(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.BioParams
		shouldErr bool
	}{
		{
			name:      "invalid value returns error",
			params:    types.NewBioParams(math.NewInt(-1000)),
			shouldErr: true,
		},
		{
			name:      "valid value returns no error",
			params:    types.NewBioParams(math.NewInt(1000)),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateBioParams(tc.params)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateOracleParams(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.OracleParams
		shouldErr bool
	}{
		{
			name: "invalid ask count returns error",
			params: types.NewOracleParams(
				32,
				0,
				6,
				50_000,
				200_000,
				sdk.NewCoin("band", math.NewInt(10)),
			),
			shouldErr: true,
		},
		{
			name: "invalid min count returns error",
			params: types.NewOracleParams(
				32,
				10,
				0,
				50_000,
				200_000,
				sdk.NewCoin("band", math.NewInt(10)),
			),
			shouldErr: true,
		},
		{
			name: "invalid prepare gas returns error",
			params: types.NewOracleParams(
				32,
				10,
				6,
				0,
				200_000,
				sdk.NewCoin("band", math.NewInt(10)),
			),
			shouldErr: true,
		},
		{
			name: "invalid execute gas returns error",
			params: types.NewOracleParams(
				32,
				10,
				6,
				50_000,
				0,
				sdk.NewCoin("band", math.NewInt(10)),
			),
			shouldErr: true,
		},
		{
			name: "invalid fee coins returns error",
			params: types.NewOracleParams(
				32,
				10,
				6,
				50_000,
				200_000,
				sdk.NewCoin("bank", math.NewInt(0)),
			),
			shouldErr: true,
		},
		{
			name: "valid value returns no error",
			params: types.NewOracleParams(
				32,
				10,
				6,
				50_000,
				200_000,
				sdk.NewCoin("band", math.NewInt(10)),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateOracleParams(tc.params)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateAppLinksParams(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.AppLinksParams
		shouldErr bool
	}{
		{
			name:      "time duration zero returns error",
			params:    types.NewAppLinksParams(time.Duration(0)),
			shouldErr: true,
		},
		{
			name:      "invalid duration returns error",
			params:    types.NewAppLinksParams(types.FourteenDaysCorrectionFactor - 1),
			shouldErr: true,
		},
		{
			name:      "minimum duration returns no error",
			params:    types.NewAppLinksParams(types.FourteenDaysCorrectionFactor),
			shouldErr: false,
		},
		{
			name:      "default params return no error",
			params:    types.DefaultAppLinksParams(),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateAppLinksParams(tc.params)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
