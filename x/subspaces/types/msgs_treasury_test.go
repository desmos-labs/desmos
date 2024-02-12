package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

var expiration = time.Date(2100, 1, 11, 0, 0, 0, 0, time.UTC)

var msgGrantTreasuryAuthorization = types.NewMsgGrantTreasuryAuthorization(
	1,
	"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
	"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
	&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
	&expiration,
)

func TestMsgGrantTreasuryAuthorization_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgGrantTreasuryAuthorization.Route())
}

func TestMsgGrantTreasuryAuthorization_Type(t *testing.T) {
	require.Equal(t, types.ActionGrantTreasuryAuthorization, msgGrantTreasuryAuthorization.Type())
}

func TestMsgGrantTreasuryAuthorization_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgGrantTreasuryAuthorization
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgGrantTreasuryAuthorization(
				0,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
				&expiration,
			),
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			msg: types.NewMsgGrantTreasuryAuthorization(
				1,
				"",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
				&expiration,
			),
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			msg: types.NewMsgGrantTreasuryAuthorization(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"",
				&banktypes.SendAuthorization{SpendLimit: sdk.NewCoins(sdk.NewInt64Coin("steak", 100))},
				&expiration,
			),
			shouldErr: true,
		},
		{
			name: "invalid authorization returns error",
			msg: types.NewMsgGrantTreasuryAuthorization(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				&banktypes.SendAuthorization{SpendLimit: nil},
				&expiration,
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgGrantTreasuryAuthorization,
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

func TestMsgGrantTreasuryAuthorization_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgGrantTreasuryAuthorization","value":{"grant":{"authorization":{"spend_limit":[{"amount":"100","denom":"steak"}]},"expiration":"2100-01-11T00:00:00Z"},"grantee":"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69","granter":"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgGrantTreasuryAuthorization.GetSignBytes()))
}

func TestMsgGrantTreasuryAuthorization_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgGrantTreasuryAuthorization.Granter)
	require.Equal(t, []sdk.AccAddress{addr}, msgGrantTreasuryAuthorization.GetSigners())
}

// --------------------------------------------------------------------------------------------------------------------

var msgRevokeTreasuryAuthorization = types.NewMsgRevokeTreasuryAuthorization(
	1,
	"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez",
	"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
	"/cosmos.bank.v1betat1.MsgSend",
)

func TestMsgRevokeTreasuryAuthorization_Route(t *testing.T) {
	require.Equal(t, types.RouterKey, msgRevokeTreasuryAuthorization.Route())
}

func TestMsgRevokeTreasuryAuthorization_Type(t *testing.T) {
	require.Equal(t, types.ActionRevokeTreasuryAuthorization, msgRevokeTreasuryAuthorization.Type())
}

func TestMsgRevokeTreasuryAuthorization_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgRevokeTreasuryAuthorization
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			msg: types.NewMsgRevokeTreasuryAuthorization(
				0,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"/cosmos.bank.v1betat1.MsgSend",
			),
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"/cosmos.bank.v1betat1.MsgSend",
			),
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error",
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"",
				"/cosmos.bank.v1betat1.MsgSend",
			),
			shouldErr: true,
		},
		{
			name: "invalid msg type url returns error - empty",
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"",
			),
			shouldErr: true,
		},
		{
			name: "invalid msg type url returns error - blank",
			msg: types.NewMsgRevokeTreasuryAuthorization(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"  ",
			),
			shouldErr: true,
		},
		{
			name: "valid message returns no error",
			msg:  msgRevokeTreasuryAuthorization,
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

func TestMsgRevokeTreasuryAuthorization_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgRevokeTreasuryAuthorization","value":{"grantee":"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69","granter":"cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez","msg_type_url":"/cosmos.bank.v1betat1.MsgSend","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgRevokeTreasuryAuthorization.GetSignBytes()))
}

func TestMsgRevokeTreasuryAuthorization_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRevokeTreasuryAuthorization.Granter)
	require.Equal(t, []sdk.AccAddress{addr}, msgRevokeTreasuryAuthorization.GetSigners())
}
