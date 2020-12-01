package types_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/fees/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultParams(t *testing.T) {
	params := types.NewParams(types.DefaultMinFees)
	require.Equal(t, params, types.DefaultParams())
}

func TestValidateParams(t *testing.T) {

	tests := []struct {
		name   string
		params types.Params
		expErr error
	}{
		{
			name: "invalid min fees param returns error",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1))))},
			),
			expErr: fmt.Errorf("invalid minimum fee message type"),
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

func TestValidateMinFeesParam(t *testing.T) {
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
			name: "invalid param returns error",
			requiredFee: types.MinFees{types.NewMinFee("",
				sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1))))},
			expErr: fmt.Errorf("invalid minimum fee message type"),
		},
		{
			name: "valid param returns no errors",
			requiredFee: types.MinFees{types.NewMinFee("desmos/createPost",
				sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000))))},
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types.ValidateMinFeesParam(test.requiredFee)
			require.Equal(t, test.expErr, err)
		})
	}
}
