package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
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
				types.NewNicknameParams(sdk.NewInt(1), sdk.NewInt(1000)),
				types.DefaultDTagParams(),
				types.DefaultMaxBioLength,
				types.DefaultOracleParams(),
			),
			shouldErr: true,
		},
		{
			name: "invalid DTag param returns error",
			params: types.NewParams(
				types.DefaultNicknameParams(),
				types.NewDTagParams("regEx", sdk.NewInt(3), sdk.NewInt(-30)),
				types.DefaultMaxBioLength,
				types.DefaultOracleParams(),
			),
			shouldErr: true,
		},
		{
			name:      "invalid bio param returns error",
			params:    types.NewParams(types.DefaultNicknameParams(), types.DefaultDTagParams(), sdk.NewInt(-1000), types.DefaultOracleParams()),
			shouldErr: true,
		},
		{
			name: "invalid oracle params return no error",
			params: types.NewParams(
				types.DefaultNicknameParams(),
				types.DefaultDTagParams(),
				types.DefaultMaxBioLength,
				types.NewOracleParams(
					types.DefaultOracleScriptID,
					types.DefaultOracleAskCount,
					types.DefaultOracleMinCount,
					types.DefaultOraclePrepareGas,
					types.DefaultOracleExecuteGas,
					"",
					types.DefaultOracleFeeCoins...,
				)),
			shouldErr: false,
		},
		{
			name:      "valid params return no error",
			params:    types.NewParams(types.DefaultNicknameParams(), types.DefaultDTagParams(), types.DefaultMaxBioLength, types.DefaultOracleParams()),
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
			name:      "invalid min returns error",
			params:    types.NewNicknameParams(sdk.NewInt(1), sdk.NewInt(1000)),
			shouldErr: true,
		},
		{
			name:      "invalid max returns error",
			params:    types.NewNicknameParams(sdk.NewInt(2), sdk.NewInt(-10)),
			shouldErr: true,
		},
		{
			name:      "valid values return no error",
			params:    types.NewNicknameParams(sdk.NewInt(2), sdk.NewInt(1000)),
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
			params:    types.NewDTagParams("", sdk.NewInt(3), sdk.NewInt(30)),
			shouldErr: true,
		},
		{
			name:      "invalid min returns error",
			params:    types.NewDTagParams("regExParam", sdk.NewInt(1), sdk.NewInt(30)),
			shouldErr: true,
		},
		{
			name:      "invalid max returns error",
			params:    types.NewDTagParams("regExParam", sdk.NewInt(3), sdk.NewInt(-30)),
			shouldErr: true,
		},
		{
			name:      "valid params return no error",
			params:    types.NewDTagParams("regExParam", sdk.NewInt(3), sdk.NewInt(30)),
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
		params    sdk.Int
		shouldErr bool
	}{
		{
			name:      "invalid value returns error",
			params:    sdk.NewInt(-1000),
			shouldErr: true,
		},
		{
			name:      "valid value returns no error",
			params:    sdk.NewInt(1000),
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
				"desmos-ibc-profiles",
				sdk.NewCoin("band", sdk.NewInt(10)),
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
				"desmos-ibc-profiles",
				sdk.NewCoin("band", sdk.NewInt(10)),
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
				"desmos-ibc-profiles",
				sdk.NewCoin("band", sdk.NewInt(10)),
			),
			shouldErr: true,
		},
		{
			name: "invalid excute gas returns error",
			params: types.NewOracleParams(
				32,
				10,
				6,
				50_000,
				0,
				"desmos-ibc-profiles",
				sdk.NewCoin("band", sdk.NewInt(10)),
			),
			shouldErr: true,
		},
		{
			name: "empty fee payer returns error",
			params: types.NewOracleParams(
				32,
				10,
				6,
				50_000,
				200_000,
				"",
				sdk.NewCoin("band", sdk.NewInt(10)),
			),
			shouldErr: true,
		},
		{
			name: "blank fee payer returns error",
			params: types.NewOracleParams(
				32,
				10,
				6,
				50_000,
				200_000,
				" ",
				sdk.NewCoin("band", sdk.NewInt(10)),
			),
			shouldErr: true,
		},
		{
			name: "invalid fee amount returns error",
			params: types.NewOracleParams(
				32,
				10,
				6,
				50_000,
				200_000,
				"desmos-ibc-profiles",
				sdk.NewCoin("bank", sdk.NewInt(0)),
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
				"desmos-ibc-profiles",
				sdk.NewCoin("band", sdk.NewInt(10)),
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
