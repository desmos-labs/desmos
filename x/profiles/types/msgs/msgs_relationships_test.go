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

	msgRequestBiDirectionalRelationship = msgs.MsgRequestBidirectionalRelationship{
		Sender:   user,
		Receiver: user,
		Message:  "",
	}

	msgAcceptBiDirectionalRelationship = msgs.MsgAcceptBidirectionalRelationship{Receiver: user, ID: "1234"}

	msgDenyBiDirectionalRelationship = msgs.MsgDenyBidirectionalRelationship{Receiver: user, ID: "1234"}

	msgDeleteRelationships = msgs.MsgDeleteRelationship{
		User: user,
		ID:   "1234",
	}
)

// MsgCreateMonoDirectionalRelationship

func TestMsgCreateMonoDirectionalRelationship_Route(t *testing.T) {
	actual := msgCreateMonoDirectionalRelationship.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgCreateMonoDirectionalRelationship_Type(t *testing.T) {
	actual := msgCreateMonoDirectionalRelationship.Type()
	require.Equal(t, "create_mono_directional_relationship", actual)
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
			name: "No errors message",
			msg: msgs.NewMsgCreateMonoDirectionalRelationship(
				user, user,
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

// MsgRequestBiDirectionalRelationship

func TestMsgRequestBidirectionalRelationship_Route(t *testing.T) {
	actual := msgRequestBiDirectionalRelationship.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgRequestBidirectionalRelationship_Type(t *testing.T) {
	actual := msgRequestBiDirectionalRelationship.Type()
	require.Equal(t, "request_bi_directional_relationship", actual)
}

func TestMsgRequestBidirectionalRelationship_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   msgs.MsgRequestBidirectionalRelationship
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: msgs.NewMsgRequestBidirectionalRelationship(
				nil, nil, "",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: msgs.NewMsgRequestBidirectionalRelationship(
				user, nil, "",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "No errors message",
			msg: msgs.NewMsgRequestBidirectionalRelationship(
				user, user, "",
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

func TestMsgRequestBidirectionalRelationship_GetSignBytes(t *testing.T) {
	actual := msgRequestBiDirectionalRelationship.GetSignBytes()
	expected := `{"type":"desmos/MsgRequestBidirectionalRelationship","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRequestBidirectionalRelationship_GetSigners(t *testing.T) {
	actual := msgRequestBiDirectionalRelationship.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgRequestBiDirectionalRelationship.Sender, actual[0])
}

// MsgAcceptBiDirectionalRelationship

func TestMsgAcceptBidirectionalRelationship_Route(t *testing.T) {
	actual := msgAcceptBiDirectionalRelationship.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgAcceptBidirectionalRelationship_Type(t *testing.T) {
	actual := msgAcceptBiDirectionalRelationship.Type()
	require.Equal(t, "accept_bi_directional_relationship", actual)
}

func TestMsgAcceptBidirectionalRelationship_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   msgs.MsgAcceptBidirectionalRelationship
		error error
	}{
		{
			name: "Empty ID returns error",
			msg: msgs.NewMsgAcceptBidirectionalRelationship(
				"", user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid relationship's id: "),
		},
		{
			name: "Empty receiver returns error",
			msg: msgs.NewMsgAcceptBidirectionalRelationship(
				"1234", nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "No errors message",
			msg: msgs.NewMsgAcceptBidirectionalRelationship(
				"1234", user,
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

func TestMsgAcceptBidirectionalRelationship_GetSignBytes(t *testing.T) {
	actual := msgAcceptBiDirectionalRelationship.GetSignBytes()
	expected := `{"type":"desmos/MsgAcceptBidirectionalRelationship","value":{"id":"1234","receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgAcceptBidirectionalRelationship_GetSigners(t *testing.T) {
	actual := msgAcceptBiDirectionalRelationship.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgAcceptBiDirectionalRelationship.Receiver, actual[0])
}

// MsgDenyBidirectionalRelationship

func TestMsgDenyBidirectionalRelationship_Route(t *testing.T) {
	actual := msgDenyBiDirectionalRelationship.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgDenyBidirectionalRelationship_Type(t *testing.T) {
	actual := msgDenyBiDirectionalRelationship.Type()
	require.Equal(t, "deny_bi_directional_relationship", actual)
}

func TestMsgDenyBidirectionalRelationship_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   msgs.MsgDenyBidirectionalRelationship
		error error
	}{
		{
			name: "Empty id returns error",
			msg: msgs.NewMsgDenyBidirectionalRelationship(
				"", user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid relationship's id: "),
		},
		{
			name: "Empty receiver returns error",
			msg: msgs.NewMsgDenyBidirectionalRelationship(
				"1234", nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "No errors message",
			msg: msgs.NewMsgDenyBidirectionalRelationship(
				"1234", user,
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

func TestMsgDenyBidirectionalRelationship_GetSignBytes(t *testing.T) {
	actual := msgDenyBiDirectionalRelationship.GetSignBytes()
	expected := `{"type":"desmos/MsgDenyBidirectionalRelationship","value":{"id":"1234","receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgDenyBidirectionalRelationship_GetSigners(t *testing.T) {
	actual := msgDenyBiDirectionalRelationship.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgDenyBiDirectionalRelationship.Receiver, actual[0])
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
			name: "Empty id returns error",
			msg: msgs.NewMsgDeleteRelationship(
				"", user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid relationship's id: "),
		},
		{
			name: "Empty receiver returns error",
			msg: msgs.NewMsgDeleteRelationship(
				"1234", nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address: "),
		},
		{
			name: "No errors message",
			msg: msgs.NewMsgDeleteRelationship(
				"1234", user,
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
	expected := `{"type":"desmos/MsgDeleteRelationship","value":{"id":"1234","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgDeleteRelationships_GetSigners(t *testing.T) {
	actual := msgDeleteRelationships.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgDeleteRelationships.User, actual[0])
}
