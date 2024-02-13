package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/relationships/types"

	"github.com/stretchr/testify/require"
)

var msgCreateRelationship = types.NewMsgCreateRelationship(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	1,
)

func TestMsgCreateRelationship_Route(t *testing.T) {
	require.Equal(t, types.ModuleName, msgCreateRelationship.Route())
}

func TestMsgCreateRelationship_Type(t *testing.T) {
	require.Equal(t, types.ActionCreateRelationship, msgCreateRelationship.Type())
}

func TestMsgCreateRelationship_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgCreateRelationship
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgCreateRelationship(
				"",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name: "equal sender and receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				0,
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgCreateRelationship,
			shouldErr: false,
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

func TestMsgCreateRelationship_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgCreateRelationship","value":{"counterparty":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","signer":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgCreateRelationship.GetSignBytes()))
}

func TestMsgCreateRelationship_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCreateRelationship.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreateRelationship.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgDeleteRelationships = types.NewMsgDeleteRelationship(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	1,
)

func TestMsgDeleteRelationships_Route(t *testing.T) {
	require.Equal(t, types.ModuleName, msgDeleteRelationships.Route())
}

func TestMsgDeleteRelationships_Type(t *testing.T) {
	require.Equal(t, types.ActionDeleteRelationship, msgDeleteRelationships.Type())
}

func TestMsgDeleteRelationships_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgDeleteRelationship
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgDeleteRelationship(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				0,
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name: "equal sender and receiver returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				0,
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgDeleteRelationships,
			shouldErr: false,
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

func TestMsgDeleteRelationships_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgDeleteRelationship","value":{"counterparty":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","signer":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace_id":"1"}}`
	require.Equal(t, expected, string(msgDeleteRelationships.GetSignBytes()))
}

func TestMsgDeleteRelationships_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgDeleteRelationships.Signer)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteRelationships.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgBlockUser = types.NewMsgBlockUser(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"reason",
	0,
)

func TestMsgBlockUser_Route(t *testing.T) {
	require.Equal(t, types.ModuleName, msgBlockUser.Route())
}

func TestMsgBlockUser_Type(t *testing.T) {
	require.Equal(t, types.ActionBlockUser, msgBlockUser.Type())
}

func TestMsgBlockUser_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgBlockUser
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgBlockUser(
				"",
				"",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name: "equal sender and receiver returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name:      "valid message returns no error",
			msg:       msgBlockUser,
			shouldErr: false,
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

func TestMsgBlockUser_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgBlockUser","value":{"blocked":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","blocker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","reason":"reason"}}`
	require.Equal(t, expected, string(msgBlockUser.GetSignBytes()))
}

func TestMsgBlockUser_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgBlockUser.Blocker)
	require.Equal(t, []sdk.AccAddress{addr}, msgBlockUser.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgUnblockUser = types.NewMsgUnblockUser(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	0,
)

func TestMsgUnblockUser_Route(t *testing.T) {
	require.Equal(t, types.ModuleName, msgUnblockUser.Route())
}

func TestMsgUnblockUser_Type(t *testing.T) {
	require.Equal(t, types.ActionUnblockUser, msgUnblockUser.Type())
}

func TestMsgUnblockUser_ValidateBasic(t *testing.T) {
	testCases := []struct {
		name      string
		msg       *types.MsgUnblockUser
		shouldErr bool
	}{
		{
			name: "empty sender returns error",
			msg: types.NewMsgUnblockUser(
				"",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name: "empty receiver returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name: "equal sender and receiver returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				0,
			),
			shouldErr: true,
		},
		{
			name: "valid message returs no error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				0,
			),
			shouldErr: false,
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

func TestMsgUnblockUser_GetSignBytes(t *testing.T) {
	expected := `{"type":"desmos/MsgUnblockUser","value":{"blocked":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","blocker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(msgUnblockUser.GetSignBytes()))
}

func TestMsgUnblockUser_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUnblockUser.Blocker)
	require.Equal(t, []sdk.AccAddress{addr}, msgUnblockUser.GetSigners())
}
