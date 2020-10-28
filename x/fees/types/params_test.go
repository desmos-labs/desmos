package types_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/fees/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultParams(t *testing.T) {
	params := types.NewParams(types.DefaultFeeDenom, types.DefaultRequiredFee)
	require.Equal(t, params, types.DefaultParams())
}

func TestParams_String(t *testing.T) {
	params := types.DefaultParams()
	require.Equal(t, "Fee parameters:\nFeeDenom: udaric\nRequiredFee: 0.010000000000000000", params.String())
}

func TestValidateParams(t *testing.T) {

	tests := []struct {
		name   string
		params types.Params
		expErr error
	}{
		{
			name:   "invalid fee denom length param returns error",
			params: types.NewParams("", types.DefaultRequiredFee),
			expErr: fmt.Errorf("invalid fee denom param, it shouldn't be empty"),
		},
		{
			name:   "negative required fee returns error",
			params: types.NewParams("udaric", sdk.NewDecWithPrec(-1, 2)),
			expErr: fmt.Errorf("invalid minimum fee required, it shouldn't be a negative number"),
		},
		{
			name:   "valid params returns no error",
			params: types.DefaultParams(),
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

func TestValidateFeeDenomParam(t *testing.T) {
	tests := []struct {
		name     string
		feeDenom interface{}
		expErr   error
	}{
		{
			name:     "invalid param type returns error",
			feeDenom: sdk.NewInt(10),
			expErr:   fmt.Errorf("invalid parameter type: 10"),
		},
		{
			name:     "invalid param returns error",
			feeDenom: "",
			expErr:   fmt.Errorf("invalid fee denom param, it shouldn't be empty"),
		},
		{
			name:     "valid param returns no errors",
			feeDenom: "udaric",
			expErr:   nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types.ValidateFeeDenomParam(test.feeDenom)
			require.Equal(t, test.expErr, err)
		})
	}
}

func TestRequiredFeeParam(t *testing.T) {
	tests := []struct {
		name        string
		requiredFee interface{}
		expErr      error
	}{
		{
			name:        "invalid param type returns error",
			requiredFee: "param",
			expErr:      fmt.Errorf("invalid parameter type: param"),
		},
		{
			name:        "invalid param returns error",
			requiredFee: sdk.NewDec(-1),
			expErr:      fmt.Errorf("invalid minimum fee required, it shouldn't be a negative number"),
		},
		{
			name:        "valid param returns no errors",
			requiredFee: sdk.NewDecWithPrec(1, 2),
			expErr:      nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types.ValidateRequiredFeeParam(test.requiredFee)
			require.Equal(t, test.expErr, err)
		})
	}
}
