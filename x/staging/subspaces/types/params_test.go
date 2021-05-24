package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultParams(t *testing.T) {
	nameParams := types.NewNameParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(10))
	params := types.NewParams(nameParams)
	require.Equal(t, params, types.DefaultParams())
}

func TestValidateParams(t *testing.T) {
	tests := []struct {
		name   string
		params types.Params
		expErr error
	}{
		{
			name:   "Invalid max name param return error",
			params: types.NewParams(types.NewNameParams("regEx", sdk.NewInt(3), sdk.NewInt(-30))),
			expErr: fmt.Errorf("invalid max name length param: -30"),
		},
		{
			name:   "Valid params return no error",
			params: types.NewParams(types.DefaultNameParams()),
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

func TestDefaultDTagParams(t *testing.T) {
	dTagParams := types.NewNameParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(10))
	defaultDTagParams := types.DefaultNameParams()
	require.Equal(t, defaultDTagParams, dTagParams)
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
			params: types.NewNameParams("", validMin, validMax),
			expErr: fmt.Errorf("empty name regEx param"),
		},
		{
			name:   "Invalid min param return error",
			params: types.NewNameParams(regEx, invalidMin, validMax),
			expErr: fmt.Errorf("invalid minimum name length param: 1"),
		},
		{
			name:   "Invalid max param return error",
			params: types.NewNameParams(regEx, validMin, invalidMax),
			expErr: fmt.Errorf("invalid max name length param: -30"),
		},
		{
			name:   "Valid params returns no error",
			params: types.NewNameParams(regEx, validMin, validMax),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, types.ValidateNameParams(test.params))
		})
	}
}
