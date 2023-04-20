package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/x/fees/types"
)

var msgUpdateParams = types.NewMsgUpdateParams(
	types.NewParams([]types.MinFee{
		types.NewMinFee(
			"/test.v1beta1.MsgTest",
			sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
		)},
	),
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
			name: "invalid params returns error",
			msg: types.NewMsgUpdateParams(
				types.NewParams([]types.MinFee{
					types.NewMinFee(
						"test.v1beta1.MsgTest",
						sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
					)},
				),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid authority returns error",
			msg: types.NewMsgUpdateParams(
				types.NewParams([]types.MinFee{
					types.NewMinFee(
						"/test.v1beta1.MsgTest",
						sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
					)},
				),
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
	expected := `{"type":"desmos/x/fees/MsgUpdateParams","value":{"authority":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","params":{"min_fees":[{"amount":[{"amount":"10000","denom":"stake"}],"message_type":"/test.v1beta1.MsgTest"}]}}}`
	require.Equal(t, expected, string(msgUpdateParams.GetSignBytes()))
}

func TestMsgUpdateParams_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUpdateParams.Authority)
	require.Equal(t, []sdk.AccAddress{addr}, msgUpdateParams.GetSigners())
}
