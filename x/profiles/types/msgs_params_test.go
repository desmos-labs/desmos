package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

var msgUpdateParams = types.NewMsgUpdateParams(
	types.DefaultParams(),
	"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
)

func TestMsgUpdateParams_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgUpdateParams.Route())
}

func TestMsgUpdateParams_Type(t *testing.T) {
	require.Equal(t, types.ActionUpdateParams, msgUpdateParams.Type())
}

func TestMsgUpdateParams_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgUpdateParams
		shouldErr bool
	}{
		{
			name: "invalid authority returns error",
			msg: types.NewMsgUpdateParams(
				types.DefaultParams(),
				"invalid",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgUpdateParams,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgUpdateParams_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/x/profiles/MsgUpdateParams","value":{"authority":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","params":{"app_links":{"validity_duration":"31536000000000000"},"bio":{"max_length":"1000"},"dtag":{"max_length":"30","min_length":"3","reg_ex":"^[A-Za-z0-9_]+$"},"nickname":{"max_length":"1000","min_length":"2"},"oracle":{"ask_count":"1","execute_gas":"200000","fee_amount":[{"amount":"10","denom":"band"}],"min_count":"1","prepare_gas":"50000"}}}}`
	require.Equal(t, expected, string(msgUpdateParams.GetSignBytes()))
}

func TestMsgUpdateParams_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUpdateParams.Authority)
	require.Equal(t, []sdk.AccAddress{addr}, msgUpdateParams.GetSigners())
}
