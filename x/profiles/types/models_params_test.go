package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestParams_Validate(t *testing.T) {
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
			),
			shouldErr: true,
		},
		{
			name: "invalid DTag param returns error",
			params: types.NewParams(
				types.DefaultNicknameParams(),
				types.NewDTagParams("regEx", sdk.NewInt(3), sdk.NewInt(-30)),
				types.DefaultMaxBioLength,
			),
			shouldErr: true,
		},
		{
			name:      "invalid bio param returns error",
			params:    types.NewParams(types.DefaultNicknameParams(), types.DefaultDTagParams(), sdk.NewInt(-1000)),
			shouldErr: true,
		},
		{
			name:      "valid params return no error",
			params:    types.NewParams(types.DefaultNicknameParams(), types.DefaultDTagParams(), types.DefaultMaxBioLength),
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
