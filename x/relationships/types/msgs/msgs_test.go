package msgs_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	msgs "github.com/desmos-labs/desmos/x/relationships/types/msgs"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	user, _               = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	otherUser, _          = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	msgCreateRelationship = msgs.MsgCreateRelationship{
		Sender:   user,
		Receiver: user,
	}

	msgDeleteRelationships = msgs.MsgDeleteRelationship{
		Sender: user,
	}
)

// MsgCreateRelationship
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
		msg   msgs.MsgCreateRelationship
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: msgs.NewMsgCreateRelationship(
				nil, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: msgs.NewMsgCreateRelationship(
				user, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals sender and receiver",
			msg: msgs.NewMsgCreateRelationship(
				user, user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender and receiver must be different"),
		},
		{
			name: "No errors message",
			msg: msgs.NewMsgCreateRelationship(
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

func TestMsgCreateRelationship_GetSignBytes(t *testing.T) {
	actual := msgCreateRelationship.GetSignBytes()
	expected := `{"type":"desmos/MsgCreateRelationship","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCreateRelationship_GetSigners(t *testing.T) {
	actual := msgCreateRelationship.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgCreateRelationship.Sender, actual[0])
}

// MsgDeleteRelationship

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
