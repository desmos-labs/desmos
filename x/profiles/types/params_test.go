package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestDefaultParams(t *testing.T) {
	nameSurnameParams := types.NewMonikerParams(sdk.NewInt(2), sdk.NewInt(1000))
	monikerParams := types.NewDtagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(30))
	bioParams := sdk.NewInt(1000)

	params := types.NewParams(nameSurnameParams, monikerParams, bioParams)

	require.Equal(t, params, types.DefaultParams())
}

func TestValidateParams(t *testing.T) {
	invalidNameMin := sdk.NewInt(1)
	validNameMax := sdk.NewInt(1000)
	validDtagMin := sdk.NewInt(3)
	invalidDtagMax := sdk.NewInt(-30)

	tests := []struct {
		name   string
		params types.Params
		expErr error
	}{
		{
			name:   "Invalid min moniker param returns error",
			params: types.NewParams(types.NewMonikerParams(invalidNameMin, validNameMax), types.DefaultDtagParams(), types.DefaultMaxBioLength),
			expErr: fmt.Errorf("invalid minimum moniker length param: 1"),
		},
		{
			name:   "Invalid max dTag param return error",
			params: types.NewParams(types.DefaultMonikerParams(), types.NewDtagParams("regEx", validDtagMin, invalidDtagMax), types.DefaultMaxBioLength),
			expErr: fmt.Errorf("invalid max dTag length param: -30"),
		},
		{
			name:   "Invalid max param returns error",
			params: types.NewParams(types.DefaultMonikerParams(), types.DefaultDtagParams(), sdk.NewInt(-1000)),
			expErr: fmt.Errorf("invalid max bio length param: -1000"),
		},
		{
			name:   "Valid params return no error",
			params: types.NewParams(types.DefaultMonikerParams(), types.DefaultDtagParams(), types.DefaultMaxBioLength),
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

func TestDefaultMonikerParams(t *testing.T) {
	monikerParams := types.NewMonikerParams(sdk.NewInt(2), sdk.NewInt(1000))
	defaultMonikerParams := types.DefaultMonikerParams()
	require.Equal(t, defaultMonikerParams, monikerParams)
}

func TestDefaultDTagParams(t *testing.T) {
	dTagParams := types.NewDtagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(30))
	defaultDTagParams := types.DefaultDtagParams()
	require.Equal(t, defaultDTagParams, dTagParams)
}

func TestValidateMonikerParams(t *testing.T) {
	invalidMonikerMin := sdk.NewInt(1)
	invalidMonikerMax := sdk.NewInt(-10)
	validMonikerMin := sdk.NewInt(2)
	validMonikerMax := sdk.NewInt(1000)

	tests := []struct {
		name   string
		params interface{}
		expErr error
	}{
		{
			name:   "Invalid min param returns error",
			params: types.NewMonikerParams(invalidMonikerMin, validMonikerMax),
			expErr: fmt.Errorf("invalid minimum moniker length param: 1"),
		},
		{
			name:   "Invalid max param returns error",
			params: types.NewMonikerParams(validMonikerMin, invalidMonikerMax),
			expErr: fmt.Errorf("invalid max moniker length param: -10"),
		},
		{
			name:   "Valid params returns no error",
			params: types.NewMonikerParams(validMonikerMin, validMonikerMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, types.ValidateMonikerParams(test.params))
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
			params: types.NewDtagParams("", validMin, validMax),
			expErr: fmt.Errorf("empty dTag regEx param"),
		},
		{
			name:   "Invalid min param return error",
			params: types.NewDtagParams(regEx, invalidMin, validMax),
			expErr: fmt.Errorf("invalid minimum dTag length param: 1"),
		},
		{
			name:   "Invalid max param return error",
			params: types.NewDtagParams(regEx, validMin, invalidMax),
			expErr: fmt.Errorf("invalid max dTag length param: -30"),
		},
		{
			name:   "Valid params returns no error",
			params: types.NewDtagParams(regEx, validMin, validMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, types.ValidateDtagParams(test.params))
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
