package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultParams(t *testing.T) {
	nsMin := sdk.NewInt(2)
	nsMax := sdk.NewInt(1000)
	nameSurnameParams := types.NewMonikerLenParams(nsMin, nsMax)
	mMin := sdk.NewInt(2)
	mMax := sdk.NewInt(30)
	monikerParams := types.NewDtagLenParams(mMin, mMax)
	bioParams := sdk.NewInt(1000)

	params := types.NewParams(nameSurnameParams, monikerParams, bioParams)

	require.Equal(t, params, types.DefaultParams())
}

func TestParams_String(t *testing.T) {
	params := types.DefaultParams()
	require.Equal(t, "Profiles parameters:\nName and Surname params lengths:\nMin accepted length: 2\nMax accepted length: 1000\nMoniker params lengths:\nMin accepted length: 2\nMax accepted length: 30\nBiography params lengths:\nMax accepted length: 1000", params.String())
}

func TestValidateParams(t *testing.T) {
	invalidNameMin := sdk.NewInt(1)
	validNameMax := sdk.NewInt(1000)
	validMonikerMin := sdk.NewInt(2)
	invalidMonikerMax := sdk.NewInt(-30)

	tests := []struct {
		name   string
		params types.Params
		expErr error
	}{
		{
			name:   "Invalid min name/surname param returns error",
			params: types.NewParams(types.NewMonikerLenParams(invalidNameMin, validNameMax), types.DefaultDtagLenParams(), types.DefaultMaxBioLength),
			expErr: fmt.Errorf("invalid minimum name/surname length param: 1"),
		},
		{
			name:   "Invalid max param return error",
			params: types.NewParams(types.DefaultMonikerLenParams(), types.NewDtagLenParams(validMonikerMin, invalidMonikerMax), types.DefaultMaxBioLength),
			expErr: fmt.Errorf("invalid max moniker length param: -30"),
		},
		{
			name:   "Invalid max param returns error",
			params: types.NewParams(types.DefaultMonikerLenParams(), types.DefaultDtagLenParams(), sdk.NewInt(-1000)),
			expErr: fmt.Errorf("invalid max bio length param: -1000"),
		},
		{
			name:   "Valid params return no error",
			params: types.NewParams(types.DefaultMonikerLenParams(), types.DefaultDtagLenParams(), types.DefaultMaxBioLength),
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

func TestDefaultNameSurnameLenParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	nameSurnameParams := types.NewMonikerLenParams(min, max)
	defaultNSParams := types.DefaultMonikerLenParams()
	require.Equal(t, defaultNSParams, nameSurnameParams)
}

func TestDefaultMonikerLenParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(30)
	monikerParams := types.NewDtagLenParams(min, max)
	defaultMonikerParams := types.DefaultDtagLenParams()
	require.Equal(t, defaultMonikerParams, monikerParams)
}

func TestNameSurnameLenParams_String(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	nameSurnameParams := types.NewMonikerLenParams(min, max)
	actual := nameSurnameParams.String()
	require.Equal(t, "Name and Surname params lengths:\nMin accepted length: 2\nMax accepted length: 1000", actual)
}

func TestMonikerLenParams_String(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(30)
	monikerParams := types.NewDtagLenParams(min, max)
	actual := monikerParams.String()
	require.Equal(t, "Moniker params lengths:\nMin accepted length: 2\nMax accepted length: 30", actual)
}

func TestValidateNameSurnameLenParams(t *testing.T) {
	invalidNameMin := sdk.NewInt(1)
	invalidNameMax := sdk.NewInt(-10)
	validNameMin := sdk.NewInt(2)
	validNameMax := sdk.NewInt(1000)

	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid min param returns error",
			params: types.NewMonikerLenParams(invalidNameMin, validNameMax),
			expErr: fmt.Errorf("invalid minimum name/surname length param: 1"),
		},
		{
			name:   "Invalid max param returns error",
			params: types.NewMonikerLenParams(validNameMin, invalidNameMax),
			expErr: fmt.Errorf("invalid max name/surname length param: -10"),
		},
		{
			name:   "Valid params returns no error",
			params: types.NewMonikerLenParams(validNameMin, validNameMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, types.ValidateMonikerLenParams(test.params))
		})
	}
}

func TestValidateMonikerLenParams(t *testing.T) {
	validMin := sdk.NewInt(2)
	validMax := sdk.NewInt(30)

	invalidMin := sdk.NewInt(1)
	invalidMax := sdk.NewInt(-30)

	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid min param return error",
			params: types.NewDtagLenParams(invalidMin, validMax),
			expErr: fmt.Errorf("invalid minimum moniker length param: 1"),
		},
		{
			name:   "Invalid max param return error",
			params: types.NewDtagLenParams(validMin, invalidMax),
			expErr: fmt.Errorf("invalid max moniker length param: -30"),
		},
		{
			name:   "Valid params returns no error",
			params: types.NewDtagLenParams(validMin, validMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, types.ValidateDtagLenParams(test.params))
		})
	}
}

func TestValidateBioLenParams(t *testing.T) {
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
			require.Equal(t, test.expErr, types.ValidateBioLenParams(test.params))
		})
	}
}
