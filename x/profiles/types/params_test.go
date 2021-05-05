package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestDefaultParams(t *testing.T) {
	nameSurnameParams := types.NewNicknameParams(sdk.NewInt(2), sdk.NewInt(1000))
	nicknameParams := types.NewDTagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(30))
	bioParams := sdk.NewInt(1000)

	params := types.NewParams(nameSurnameParams, nicknameParams, bioParams)

	require.Equal(t, params, types.DefaultParams())
}

func TestValidateParams(t *testing.T) {
	invalidNameMin := sdk.NewInt(1)
	validNameMax := sdk.NewInt(1000)
	validDTagMin := sdk.NewInt(3)
	invalidDTagMax := sdk.NewInt(-30)

	tests := []struct {
		name   string
		params types.Params
		expErr error
	}{
		{
			name:   "Invalid min nickname param returns error",
			params: types.NewParams(types.NewNicknameParams(invalidNameMin, validNameMax), types.DefaultDTagParams(), types.DefaultMaxBioLength),
			expErr: fmt.Errorf("invalid minimum nickname length param: 1"),
		},
		{
			name:   "Invalid max dTag param return error",
			params: types.NewParams(types.DefaultNicknameParams(), types.NewDTagParams("regEx", validDTagMin, invalidDTagMax), types.DefaultMaxBioLength),
			expErr: fmt.Errorf("invalid max dTag length param: -30"),
		},
		{
			name:   "Invalid max param returns error",
			params: types.NewParams(types.DefaultNicknameParams(), types.DefaultDTagParams(), sdk.NewInt(-1000)),
			expErr: fmt.Errorf("invalid max bio length param: -1000"),
		},
		{
			name:   "Valid params return no error",
			params: types.NewParams(types.DefaultNicknameParams(), types.DefaultDTagParams(), types.DefaultMaxBioLength),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.params.Validate())
		})
	}
}

func TestDefaultNicknameParams(t *testing.T) {
	nicknameParams := types.NewNicknameParams(sdk.NewInt(2), sdk.NewInt(1000))
	defaultnicknameParams := types.DefaultNicknameParams()
	require.Equal(t, defaultnicknameParams, nicknameParams)
}

func TestDefaultDTagParams(t *testing.T) {
	dTagParams := types.NewDTagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(30))
	defaultDTagParams := types.DefaultDTagParams()
	require.Equal(t, defaultDTagParams, dTagParams)
}

func TestValidateNicknameParams(t *testing.T) {
	invalidNicknameMin := sdk.NewInt(1)
	invalidNicknameMax := sdk.NewInt(-10)
	validNicknameMin := sdk.NewInt(2)
	validNicknameMax := sdk.NewInt(1000)

	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid min param returns error",
			params: types.NewNicknameParams(invalidNicknameMin, validNicknameMax),
			expErr: fmt.Errorf("invalid minimum nickname length param: 1"),
		},
		{
			name:   "Invalid max param returns error",
			params: types.NewNicknameParams(validNicknameMin, invalidNicknameMax),
			expErr: fmt.Errorf("invalid max nickname length param: -10"),
		},
		{
			name:   "Valid params returns no error",
			params: types.NewNicknameParams(validNicknameMin, validNicknameMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, types.ValidateNicknameParams(test.params))
		})
	}
}

func TestValidateDTagParams(t *testing.T) {
	regEx := "regExParam"
	validMin := sdk.NewInt(3)
	validMax := sdk.NewInt(30)

	invalidMin := sdk.NewInt(1)
	invalidMax := sdk.NewInt(-30)

	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid empty regEx return error",
			params: types.NewDTagParams("", validMin, validMax),
			expErr: fmt.Errorf("empty dTag regEx param"),
		},
		{
			name:   "Invalid min param return error",
			params: types.NewDTagParams(regEx, invalidMin, validMax),
			expErr: fmt.Errorf("invalid minimum dTag length param: 1"),
		},
		{
			name:   "Invalid max param return error",
			params: types.NewDTagParams(regEx, validMin, invalidMax),
			expErr: fmt.Errorf("invalid max dTag length param: -30"),
		},
		{
			name:   "Valid params returns no error",
			params: types.NewDTagParams(regEx, validMin, validMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, types.ValidateDTagParams(test.params))
		})
	}
}

func TestValidateBioParams(t *testing.T) {
	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid max param returns error",
			params: sdk.NewInt(-1000),
			expErr: fmt.Errorf("invalid max bio length param: -1000"),
		},
		{
			name:   "Valid params returns no error",
			params: sdk.NewInt(1000),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, types.ValidateBioParams(test.params))
		})
	}
}
