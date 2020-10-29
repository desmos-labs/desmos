package types_test

import (
	"github.com/desmos-labs/desmos/x/profiles/types"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types/common"
	"github.com/stretchr/testify/require"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

var user, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var otherUser, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
var testProfile = Profile{
	DTag:    "dtag",
	Moniker: common.NewStrPtr("moniker"),
	Bio:     common.NewStrPtr("biography"),
	Pictures: NewPictures(
		common.NewStrPtr("https://shorturl.at/adnX3"),
		common.NewStrPtr("https://shorturl.at/cgpyF"),
	),
	Creator: user,
}

var msgEditProfile = types.NewMsgSaveProfile(
	"monk",
	testProfile.Moniker,
	testProfile.Bio,
	testProfile.Pictures.Profile,
	testProfile.Pictures.Cover,
	testProfile.Creator,
)

var msgDeleteProfile = types.NewMsgDeleteProfile(
	testProfile.Creator,
)

var msgRequestTransferDTag = types.NewMsgRequestDTagTransfer(
	user,
	otherUser,
)

var msgAcceptDTagTransfer = types.NewMsgAcceptDTagTransfer(
	"dtag",
	user,
	otherUser,
)

var msgRejectDTagTransfer = types.NewMsgRefuseDTagTransferRequest(
	user,
	otherUser,
)

var msgCancelDTagTransferReq = types.NewMsgCancelDTagTransferRequest(
	user,
	otherUser,
)

func TestMsgSaveProfile_Route(t *testing.T) {
	actual := msgEditProfile.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgSaveProfile_Type(t *testing.T) {
	actual := msgEditProfile.Type()
	require.Equal(t, "save_profile", actual)
}

func TestMsgSaveProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgSaveProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgSaveProfile(
				testProfile.DTag,
				testProfile.Moniker,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address: "),
		},
		{
			name:  "Invalid empty dtag returns error",
			msg:   types.NewMsgSaveProfile("", nil, nil, nil, nil, testProfile.Creator),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "profile dtag cannot be empty or blank"),
		},
		{
			name: "No error message",
			msg: types.NewMsgSaveProfile(
				"_crazy_papa_21",
				common.NewStrPtr("custom-moniker"),
				common.NewStrPtr("custom-bio"),
				common.NewStrPtr("https://test.com/my-custom-profile-pic"),
				common.NewStrPtr("https://test.com/my-custom-cover-pic"),
				user,
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

func TestMsgSaveProfile_GetSignBytes(t *testing.T) {
	actual := msgEditProfile.GetSignBytes()
	expected := `{"type":"desmos/MsgSaveProfile","value":{"bio":"biography","cover_picture":"https://shorturl.at/cgpyF","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","dtag":"monk","moniker":"moniker","profile_picture":"https://shorturl.at/adnX3"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgSaveProfile_GetSigners(t *testing.T) {
	actual := msgEditProfile.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgEditProfile.Creator, actual[0])
}

// ----------------------
// --- MsgDeleteProfile
// ----------------------

func TestMsgDeleteProfile_Route(t *testing.T) {
	actual := msgDeleteProfile.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgDeleteProfile_Type(t *testing.T) {
	actual := msgDeleteProfile.Type()
	require.Equal(t, "delete_profile", actual)
}

func TestMsgDeleteProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgDeleteProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgDeleteProfile(
				nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address: "),
		},
		{
			name: "Valid message returns no error",
			msg: types.NewMsgDeleteProfile(
				testProfile.Creator,
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

func TestMsgDeleteProfile_GetSignBytes(t *testing.T) {
	actual := msgDeleteProfile.GetSignBytes()
	expected := `{"type":"desmos/MsgDeleteProfile","value":{"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgDeleteProfile_GetSigners(t *testing.T) {
	actual := msgDeleteProfile.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgDeleteProfile.Creator, actual[0])
}

// ----------------------
// --- MsgRequestDTagTransfer
// ----------------------

func TestMsgRequestDTagTransfer_Route(t *testing.T) {
	actual := msgRequestTransferDTag.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgRequestDTagTransfer_Type(t *testing.T) {
	actual := msgRequestTransferDTag.Type()
	require.Equal(t, "request_dtag", actual)
}

func TestMsgRequestDTagTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgRequestDTagTransfer
		error error
	}{
		{
			name: "Empty current owner returns error",
			msg: types.NewMsgRequestDTagTransfer(
				nil, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Empty receiving user returns error",
			msg: types.NewMsgRequestDTagTransfer(
				user, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Equals current owner and receiving user returns error",
			msg: types.NewMsgRequestDTagTransfer(
				user, user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgRequestDTagTransfer(
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

func TestMsgRequestDTagTransfer_GetSignBytes(t *testing.T) {
	actual := msgRequestTransferDTag.GetSignBytes()
	expected := `{"type":"desmos/MsgRequestDTagTransfer","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRequestDTagTransfer_GetSigners(t *testing.T) {
	actual := msgRequestTransferDTag.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgRequestTransferDTag.Sender, actual[0])
}

// ----------------------
// --- MsgAcceptDTagTransfer
// ----------------------

func TestMsgAcceptDTagTransfer_Route(t *testing.T) {
	actual := msgAcceptDTagTransfer.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgAcceptDTagTransfer_Type(t *testing.T) {
	actual := msgAcceptDTagTransfer.Type()
	require.Equal(t, "accept_dtag_request", actual)
}

func TestMsgAcceptDTagTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgAcceptDTagTransfer
		error error
	}{
		{
			name: "Empty current owner returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag", nil, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid current owner address: "),
		},
		{
			name: "Empty receiving user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag", user, nil,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiving user address: "),
		},
		{
			name: "Equals current owner and receiving user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag", user, user,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "Empty newDTag returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"", user, otherUser,
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new dTag can't be empty"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag", user, otherUser,
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

func TestMsgAcceptDTagTransfer_GetSignBytes(t *testing.T) {
	actual := msgAcceptDTagTransfer.GetSignBytes()
	expected := `{"type":"desmos/MsgAcceptDTagTransfer","value":{"new_d_tag":"dtag","receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgAcceptDTagTransfer_GetSigners(t *testing.T) {
	actual := msgAcceptDTagTransfer.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgRequestTransferDTag.Receiver, actual[0])
}

func TestMsgRejectDTagRequest_Route(t *testing.T) {
	actual := msgRejectDTagTransfer.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgRejectDTagRequest_Type(t *testing.T) {
	actual := msgRejectDTagTransfer.Type()
	require.Equal(t, "refuse_dtag_request", actual)
}

func TestMsgRejectDTagRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgRefuseDTagTransfer
		error error
	}{
		{
			name:  "Empty sender returns error",
			msg:   types.NewMsgRefuseDTagTransferRequest(user, nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name:  "Empty receiver returns error",
			msg:   types.NewMsgRefuseDTagTransferRequest(nil, user),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name:  "Equals sender and receiver returns error",
			msg:   types.NewMsgRefuseDTagTransferRequest(user, user),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name:  "No error message",
			msg:   types.NewMsgRefuseDTagTransferRequest(user, otherUser),
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

func TestMsgRejectDTagRequest_GetSignBytes(t *testing.T) {
	actual := msgRejectDTagTransfer.GetSignBytes()
	expected := `{"type":"desmos/MsgRefuseDTagTransfer","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRejectDTagRequest_GetSigners(t *testing.T) {
	actual := msgRejectDTagTransfer.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgRejectDTagTransfer.Sender, actual[0])
}

func TestMsgCancelDTagRequest_Route(t *testing.T) {
	actual := msgCancelDTagTransferReq.Route()
	require.Equal(t, "profiles", actual)
}

func TestMsgCancelDTagRequest_Type(t *testing.T) {
	actual := msgCancelDTagTransferReq.Type()
	require.Equal(t, "cancel_dtag_request", actual)
}

func TestMsgCancelDTagRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   types.MsgCancelDTagTransfer
		error error
	}{
		{
			name:  "Empty receiver returns error",
			msg:   types.NewMsgCancelDTagTransferRequest(user, nil),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name:  "Empty sender returns error",
			msg:   types.NewMsgCancelDTagTransferRequest(nil, user),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name:  "Equals sender and receiver returns error",
			msg:   types.NewMsgCancelDTagTransferRequest(user, user),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name:  "No error message",
			msg:   types.NewMsgCancelDTagTransferRequest(user, otherUser),
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

func TestMsgCancelDTagRequest_GetSignBytes(t *testing.T) {
	actual := msgCancelDTagTransferReq.GetSignBytes()
	expected := `{"type":"desmos/MsgCancelDTagTransfer","value":{"receiver":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCancelDTagRequest_GetSigners(t *testing.T) {
	actual := msgCancelDTagTransferReq.GetSigners()
	require.Equal(t, 1, len(actual))
	require.Equal(t, msgCancelDTagTransferReq.Sender, actual[0])
}
