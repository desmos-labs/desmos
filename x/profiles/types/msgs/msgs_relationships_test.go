package msgs_test

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types/msgs"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	msgCreateMonoDirectionalRelationship = msgs.MsgCreateMonoDirectionalRelationship{
		Sender:   user,
		Receiver: user,
	}

	msgDeleteRelationships = msgs.MsgDeleteRelationship{
		Sender: user,
	}
)

// MsgCreateMonoDirectionalRelationship
func TestMsgCreateMonoDirectionalRelationship_Route(t *testing.T) {
	actual := msgCreateMonoDirectionalRelationship.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgCreateMonoDirectionalRelationship_Type(t *testing.T) {
	actual := msgCreateMonoDirectionalRelationship.Type()
	require.Equal(t, "create_relationship", actual)
}

func TestMsgCreateMonoDirectionalRelationship_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   msgs.MsgCreateMonoDirectionalRelationship
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: msgs.NewMsgCreateMonoDirectionalRelationship(
				nil, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: msgs.NewMsgCreateMonoDirectionalRelationship(
				user, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals sender and receiver",
			msg: msgs.NewMsgCreateMonoDirectionalRelationship(
				user, user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different"),
		},
		{
			name: "No errors message",
			msg: msgs.NewMsgCreateMonoDirectionalRelationship(
				user, otherUser,
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

func TestMsgCreateMonoDirectionalRelationship_GetSignBytes(t *testing.T) {
	actual := msgCreateMonoDirectionalRelationship.GetSignBytes()
	expected := `{"type":"desmos/MsgCreateMonoDirectionalRelationship","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCreateMonoDirectionalRelationship_GetSigners(t *testing.T) {
	actual := msgCreateMonoDirectionalRelationship.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgCreateMonoDirectionalRelationship.Sender, actual[0])
}

// MsgDeleteRelationship

func TestMsgDeleteRelationships_Route(t *testing.T) {
	actual := msgDeleteRelationships.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgDeleteRelationships_Type(t *testing.T) {
	actual := msgDeleteRelationships.Type()
	require.Equal(t, "delete_relationship", actual)
}

func TestMsgDeleteRelationships_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   msgs.MsgDeleteRelationship
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: msgs.NewMsgDeleteRelationship(
				nil, user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: msgs.NewMsgDeleteRelationship(
				user, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid counterparty address: "),
		},
		{
			name: "Equals sender and receiver",
			msg: msgs.NewMsgDeleteRelationship(
				user, user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different"),
		},
		{
			name: "No errors message",
			msg: msgs.NewMsgDeleteRelationship(
				user, otherUser,
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
	expected := `{"type":"desmos/MsgDeleteRelationship","value":{"counterparty":"","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgDeleteRelationships_GetSigners(t *testing.T) {
	actual := msgDeleteRelationships.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgDeleteRelationships.Sender, actual[0])
}
