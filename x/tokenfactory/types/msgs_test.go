package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

var msgCreateDenom = types.NewMsgCreateDenom(
	1,
	"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
	"minttoken",
)

func TestMsgCreateDenom_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgCreateDenom.Route())
}

func TestMsgCreateDenom_Type(t *testing.T) {
	require.Equal(t, types.ActionCreateDenom, msgCreateDenom.Type())
}

func TestMsgCreateDenom_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCreateDenom
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgCreateDenom(
				0,
				msgCreateDenom.Sender,
				msgCreateDenom.Subdenom,
			),
			shouldErr: true,
		},
		{
			name: "invalid sender returns error",
			msg: types.NewMsgCreateDenom(
				msgCreateDenom.SubspaceID,
				"",
				msgCreateDenom.Subdenom,
			),
			shouldErr: true,
		},
		{
			name: "invalid subdenom returns error",
			msg: types.NewMsgCreateDenom(
				msgCreateDenom.SubspaceID,
				msgCreateDenom.Sender,
				"%invalid%",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgCreateDenom,
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

func TestMsgCreateDenom_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/x/tokenfactory/MsgCreateDenom","value":{"sender":"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69","subdenom":"minttoken","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgCreateDenom.GetSignBytes()))
}

func TestMsgCreateDenom_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCreateDenom.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreateDenom.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgMint = types.NewMsgMint(
	1,
	"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
	sdk.NewCoin("uminttoken", sdk.NewInt(100)),
)

func TestMsgMint_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgMint.Route())
}

func TestMsgMint_Type(t *testing.T) {
	require.Equal(t, types.ActionMint, msgMint.Type())
}

func TestMsgMint_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgMint
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgMint(
				0,
				msgMint.Sender,
				msgMint.Amount,
			),
			shouldErr: true,
		},
		{
			name: "invalid sender returns error",
			msg: types.NewMsgMint(
				msgMint.SubspaceID,
				"",
				msgMint.Amount,
			),
			shouldErr: true,
		},
		{
			name: "invalid amount returns error",
			msg: types.NewMsgMint(
				msgMint.SubspaceID,
				msgMint.Sender,
				sdk.Coin{Denom: "%invalid%", Amount: sdk.NewInt(100)},
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgMint,
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

func TestMsgMint_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/x/tokenfactory/MsgMint","value":{"amount":{"amount":"100","denom":"uminttoken"},"sender":"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgMint.GetSignBytes()))
}

func TestMsgMint_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgMint.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgMint.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgBurn = types.NewMsgBurn(
	1,
	"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
	sdk.NewCoin("uminttoken", sdk.NewInt(100)),
)

func TestMsgBurn_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgBurn.Route())
}

func TestMsgBurn_Type(t *testing.T) {
	require.Equal(t, types.ActionBurn, msgBurn.Type())
}

func TestMsgBurn_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgBurn
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgBurn(
				0,
				msgBurn.Sender,
				msgBurn.Amount,
			),
			shouldErr: true,
		},
		{
			name: "invalid sender returns error",
			msg: types.NewMsgBurn(
				msgBurn.SubspaceID,
				"",
				msgBurn.Amount,
			),
			shouldErr: true,
		},
		{
			name: "invalid amount returns error",
			msg: types.NewMsgBurn(
				msgBurn.SubspaceID,
				msgBurn.Sender,
				sdk.Coin{Denom: "%invalid%", Amount: sdk.NewInt(100)},
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgBurn,
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

func TestMsgBurn_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/x/tokenfactory/MsgBurn","value":{"amount":{"amount":"100","denom":"uminttoken"},"sender":"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgBurn.GetSignBytes()))
}

func TestMsgBurn_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgBurn.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgBurn.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgSetDenomMetadata = types.NewMsgSetDenomMetadata(
	1,
	"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
	banktypes.Metadata{
		Name:        "Mint Token",
		Symbol:      "MTK",
		Description: "The custom token of the test subspace.",
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", Exponent: uint32(0), Aliases: nil},
			{Denom: "minttoken", Exponent: uint32(6), Aliases: []string{"minttoken"}},
		},
		Base:    "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
		Display: "minttoken",
	},
)

func TestMsgSetDenomMetadata_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgSetDenomMetadata.Route())
}

func TestMsgSetDenomMetadata_Type(t *testing.T) {
	require.Equal(t, types.ActionSetDenomMetadata, msgSetDenomMetadata.Type())
}

func TestMsgSetDenomMetadata_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgSetDenomMetadata
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgSetDenomMetadata(
				0,
				msgSetDenomMetadata.Sender,
				msgSetDenomMetadata.Metadata,
			),
			shouldErr: true,
		},
		{
			name: "invalid sender returns error",
			msg: types.NewMsgSetDenomMetadata(
				msgSetDenomMetadata.SubspaceID,
				"",
				msgSetDenomMetadata.Metadata,
			),
			shouldErr: true,
		},
		{
			name: "invalid metadata returns error",
			msg: types.NewMsgSetDenomMetadata(
				msgSetDenomMetadata.SubspaceID,
				msgSetDenomMetadata.Sender,
				banktypes.Metadata{},
			),
			shouldErr: true,
		},
		{
			name: "invalid metadata base returns error",
			msg: types.NewMsgSetDenomMetadata(
				msgSetDenomMetadata.SubspaceID,
				msgSetDenomMetadata.Sender,
				banktypes.Metadata{
					Name:        "Mint Token",
					Symbol:      "MINTTOKEN",
					Description: "The custom token of the test subspace.",
					DenomUnits: []*banktypes.DenomUnit{
						{Denom: "uminttoken", Exponent: uint32(0), Aliases: nil},
					},
					Base:    "uminttoken",
					Display: "uminttoken",
				},
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgSetDenomMetadata,
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

func TestMsgSetDenomMetadata_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/x/tokenfactory/MsgSetDenomMetadata","value":{"metadata":{"base":"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken","denom_units":[{"denom":"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken"},{"aliases":["minttoken"],"denom":"minttoken","exponent":6}],"description":"The custom token of the test subspace.","display":"minttoken","name":"Mint Token","symbol":"MTK"},"sender":"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgSetDenomMetadata.GetSignBytes()))
}

func TestMsgSetDenomMetadata_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgSetDenomMetadata.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgSetDenomMetadata.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

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
	expected := `{"type":"desmos/x/tokenfactory/MsgUpdateParams","value":{"authority":"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd","params":{"denom_creation_fee":[{"amount":"10000000","denom":"stake"}]}}}`
	require.Equal(t, expected, string(msgUpdateParams.GetSignBytes()))
}

func TestMsgUpdateParams_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUpdateParams.Authority)
	require.Equal(t, []sdk.AccAddress{addr}, msgUpdateParams.GetSigners())
}
