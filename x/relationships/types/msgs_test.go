package types_test

import (
	"github.com/desmos-labs/desmos/x/relationships/types"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

var (
	msgCreateRelationship = types.NewMsgCreateRelationship(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	msgDeleteRelationships = types.NewMsgDeleteRelationship(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	msgBlockUser = types.NewMsgBlockUser(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"reason",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	msgUnblockUser = types.NewMsgUnblockUser(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
)

// ___________________________________________________________________________________________________________________

func TestMsgCreateRelationship_Route(t *testing.T) {
	actual := msgCreateRelationship.Route()
	require.Equal(t, "relationships", actual)
}

func TestMsgCreateRelationship_Type(t *testing.T) {
	actual := msgCreateRelationship.Type()
	require.Equal(t, "create_relationship", actual)
}

func TestMsgCreateRelationship_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgCreateRelationship
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: types.NewMsgCreateRelationship(
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different"),
		},
		{
			name: "Invalid subspace returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"1234",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgCreateRelationship_GetSignBytes(t *testing.T) {
	actual := msgCreateRelationship.GetSignBytes()
	expected := `{"type":"desmos/MsgCreateRelationship","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCreateRelationship_GetSigners(t *testing.T) {
	actual := msgCreateRelationship.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgCreateRelationship.Sender, actual[0])
}

// ___________________________________________________________________________________________________________________

func TestMsgDeleteRelationships_Route(t *testing.T) {
	actual := msgDeleteRelationships.Route()
	require.Equal(t, "relationships", actual)
}

func TestMsgDeleteRelationships_Type(t *testing.T) {
	actual := msgDeleteRelationships.Type()
	require.Equal(t, "delete_relationship", actual)
}

func TestMsgDeleteRelationships_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgDeleteRelationship
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: types.NewMsgDeleteRelationship(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid counterparty address: "),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different"),
		},
		{
			name: "Invalid subspace returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"1234",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgDeleteRelationships_GetSignBytes(t *testing.T) {
	actual := msgDeleteRelationships.GetSignBytes()
	expected := `{"type":"desmos/MsgDeleteRelationship","value":{"counterparty":"","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgDeleteRelationships_GetSigners(t *testing.T) {
	actual := msgDeleteRelationships.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgDeleteRelationships.Sender, actual[0])
}

// ___________________________________________________________________________________________________________________

func TestMsgBlockUser_Route(t *testing.T) {
	actual := msgBlockUser.Route()
	require.Equal(t, "relationships", actual)
}

func TestMsgBlockUser_Type(t *testing.T) {
	actual := msgBlockUser.Type()
	require.Equal(t, "block_user", actual)
}

func TestMsgBlockUser_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgBlockUser
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: types.NewMsgBlockUser(
				"",
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked address: "),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "blocker and blocked must be different"),
		},
		{
			name: "Invalid subspace returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				"yeah",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"mobbing",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgBlockUser_GetSignBytes(t *testing.T) {
	actual := msgBlockUser.GetSignBytes()
	expected := `{"type":"desmos/MsgBlockUser","value":{"blocked":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","blocker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgBlockUser_GetSigners(t *testing.T) {
	actual := msgBlockUser.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgBlockUser.Blocker, actual[0])
}

// ___________________________________________________________________________________________________________________

func TestMsgUnblockUser_Route(t *testing.T) {
	actual := msgUnblockUser.Route()
	require.Equal(t, "relationships", actual)
}

func TestMsgUnblockUser_Type(t *testing.T) {
	actual := msgUnblockUser.Type()
	require.Equal(t, "unblock_user", actual)
}

func TestMsgUnblockUser_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgUnblockUser
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: types.NewMsgUnblockUser(
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked address: "),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "blocker and blocked must be different"),
		},
		{
			name: "Invalid subspace returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"yeah",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.error == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.error.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgUnblockUser_GetSignBytes(t *testing.T) {
	actual := msgUnblockUser.GetSignBytes()
	expected := `{"type":"desmos/MsgUnblockUser","value":{"blocked":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","blocker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgUnblockUser_GetSigners(t *testing.T) {
	actual := msgUnblockUser.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgUnblockUser.Blocker, actual[0])
}
