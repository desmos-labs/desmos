package types_test

import (
	"fmt"
	"testing"

	types2 "github.com/desmos-labs/desmos/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultParams(t *testing.T) {
	params := types2.NewParams(sdk.NewInt(500), sdk.NewInt(10), sdk.NewInt(200), sdk.NewInt(10))
	require.Equal(t, params, types2.DefaultParams())
}

func TestValidateParams(t *testing.T) {
	tests := []struct {
		name   string
		params types2.Params
		expErr error
	}{
		{
			name:   "invalid max post message length param returns error",
			params: types2.NewParams(sdk.NewInt(-1), sdk.NewInt(12), sdk.NewInt(200), sdk.NewInt(10)),
			expErr: fmt.Errorf("invalid max post message length param: -1"),
		},
		{
			name:   "invalid max additional attributes number param returns error",
			params: types2.NewParams(sdk.NewInt(500), sdk.NewInt(-1), sdk.NewInt(8), sdk.NewInt(10)),
			expErr: fmt.Errorf("invalid max additional attributes fields number param: -1"),
		},
		{
			name:   "invalid max additional attributes field value length returns error",
			params: types2.NewParams(sdk.NewInt(500), sdk.NewInt(8), sdk.NewInt(-1), sdk.NewInt(10)),
			expErr: fmt.Errorf("invalid max additional attributes fields value length param: -1"),
		},
		{
			name:   "invalid max additional attributes field key length returns error",
			params: types2.NewParams(sdk.NewInt(500), sdk.NewInt(8), sdk.NewInt(10), sdk.NewInt(-1)),
			expErr: fmt.Errorf("invalid max additional attributes fields key length param: -1"),
		},
		{
			name:   "valid params returns no error",
			params: types2.DefaultParams(),
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

func TestValidateMaxPostMessageLengthParam(t *testing.T) {
	tests := []struct {
		name             string
		maxPostMsgLength interface{}
		expErr           error
	}{
		{
			name:             "invalid param type returns error",
			maxPostMsgLength: "param",
			expErr:           fmt.Errorf("invalid parameters type: param"),
		},
		{
			name:             "invalid param returns error",
			maxPostMsgLength: sdk.NewInt(-1),
			expErr:           fmt.Errorf("invalid max post message length param: -1"),
		},
		{
			name:             "valid param returns no errors",
			maxPostMsgLength: sdk.NewInt(5000),
			expErr:           nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types2.ValidateMaxPostMessageLengthParam(test.maxPostMsgLength)
			require.Equal(t, test.expErr, err)
		})
	}
}

func TestValidateMaxAdditionalAttributesFieldNumberParam(t *testing.T) {
	tests := []struct {
		name            string
		maxOpDataNumber interface{}
		expErr          error
	}{
		{
			name:            "invalid param type returns error",
			maxOpDataNumber: "param",
			expErr:          fmt.Errorf("invalid parameters type: param"),
		},
		{
			name:            "invalid param returns error",
			maxOpDataNumber: sdk.NewInt(-1),
			expErr:          fmt.Errorf("invalid max additional attributes fields number param: -1"),
		},
		{
			name:            "valid param returns no errors",
			maxOpDataNumber: sdk.NewInt(50),
			expErr:          nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types2.ValidateMaxAdditionalAttributesFieldNumberParam(test.maxOpDataNumber)
			require.Equal(t, test.expErr, err)
		})
	}
}

func TestValidateMaxAdditionalAttributesFieldValueLengthParam(t *testing.T) {
	tests := []struct {
		name              string
		maxOpDataFieldLen interface{}
		expErr            error
	}{
		{
			name:              "invalid param type returns error",
			maxOpDataFieldLen: "param",
			expErr:            fmt.Errorf("invalid parameters type: param"),
		},
		{
			name:              "invalid param returns error",
			maxOpDataFieldLen: sdk.NewInt(-1),
			expErr:            fmt.Errorf("invalid max additional attributes fields value length param: -1"),
		},
		{
			name:              "valid param returns no errors",
			maxOpDataFieldLen: sdk.NewInt(500),
			expErr:            nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types2.ValidateMaxAdditionalAttributesFieldValueLengthParam(test.maxOpDataFieldLen)
			require.Equal(t, test.expErr, err)
		})
	}
}

func TestValidateMaxAdditionalAttributesFieldKeyLengthParam(t *testing.T) {
	tests := []struct {
		name              string
		maxOpDataFieldLen interface{}
		expErr            error
	}{
		{
			name:              "invalid param type returns error",
			maxOpDataFieldLen: "param",
			expErr:            fmt.Errorf("invalid parameters type: param"),
		},
		{
			name:              "invalid param returns error",
			maxOpDataFieldLen: sdk.NewInt(-1),
			expErr:            fmt.Errorf("invalid max additional attributes fields key length param: -1"),
		},
		{
			name:              "valid param returns no errors",
			maxOpDataFieldLen: sdk.NewInt(500),
			expErr:            nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types2.ValidateMaxAdditionalAttributesFieldKeyLengthParam(test.maxOpDataFieldLen)
			require.Equal(t, test.expErr, err)
		})
	}
}
