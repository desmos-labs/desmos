package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultParams(t *testing.T) {
	params := types.NewParams(sdk.NewInt(500), sdk.NewInt(10), sdk.NewInt(200))
	require.Equal(t, params, types.DefaultParams())
}

func TestParams_String(t *testing.T) {
	params := types.DefaultParams()
	require.Equal(t, "Posts parameters:\nMaxPostMessageLength: 500\nMaxOptionalDataFieldsNumber: 10\nMaxOptionalDataFieldValueLength: 200", params.String())
}

func TestValidateParams(t *testing.T) {

	tests := []struct {
		name   string
		params types.Params
		expErr error
	}{
		{
			name:   "invalid max post message length param returns error",
			params: types.NewParams(sdk.NewInt(-1), sdk.NewInt(12), sdk.NewInt(200)),
			expErr: fmt.Errorf("invalid max post message length param: -1"),
		},
		{
			name:   "invalid max optional data number param returns error",
			params: types.NewParams(sdk.NewInt(500), sdk.NewInt(8), sdk.NewInt(200)),
			expErr: fmt.Errorf("invalid max optional data fields number param: 8"),
		},
		{
			name:   "invalid max optional data field value length returns error",
			params: types.NewParams(sdk.NewInt(500), sdk.NewInt(10), sdk.NewInt(10)),
			expErr: fmt.Errorf("invalid max optional data fields value length param: %s", sdk.NewInt(10)),
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
			err := types.ValidateMaxPostMessageLengthParam(test.maxPostMsgLength)
			require.Equal(t, test.expErr, err)
		})
	}
}

func TestValidateMaxOptionalDataFieldNumberParam(t *testing.T) {
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
			expErr:          fmt.Errorf("invalid max optional data fields number param: -1"),
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
			err := types.ValidateMaxOptionalDataFieldNumberParam(test.maxOpDataNumber)
			require.Equal(t, test.expErr, err)
		})
	}
}

func TestValidateMaxOptionalDataFieldValueLengthParam(t *testing.T) {
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
			expErr:            fmt.Errorf("invalid max optional data fields value length param: -1"),
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
			err := types.ValidateMaxOptionalDataFieldValueLengthParam(test.maxOpDataFieldLen)
			require.Equal(t, test.expErr, err)
		})
	}
}
