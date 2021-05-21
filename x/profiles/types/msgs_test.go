package types_test

import (
	"testing"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

var addr, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
var testProfile, _ = types.NewProfile(
	"dtag",
	"nickname",
	"biography",
	types.NewPictures(
		"https://shorturl.at/adnX3",
		"https://shorturl.at/cgpyF",
	),
	time.Unix(100, 0),
	authtypes.NewBaseAccountWithAddress(addr),
)

// ___________________________________________________________________________________________________________________

var msgEditProfile = types.NewMsgSaveProfile(
	"monk",
	testProfile.Nickname,
	testProfile.Bio,
	testProfile.Pictures.Profile,
	testProfile.Pictures.Cover,
	testProfile.GetAddress().String(),
)

func TestMsgSaveProfile_Route(t *testing.T) {
	require.Equal(t, "profiles", msgEditProfile.Route())
}

func TestMsgSaveProfile_Type(t *testing.T) {
	require.Equal(t, "save_profile", msgEditProfile.Type())
}

func TestMsgSaveProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgSaveProfile
		error error
	}{
		{
			name: "Empty owner returns error",
			msg: types.NewMsgSaveProfile(
				testProfile.DTag,
				testProfile.Nickname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator: "),
		},
		{
			name:  "Invalid empty dtag returns error",
			msg:   types.NewMsgSaveProfile("", "", "", "", "", testProfile.GetAddress().String()),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "profile dtag cannot be empty or blank"),
		},
		{
			name: "No error message",
			msg: types.NewMsgSaveProfile(
				"_crazy_papa_21",
				"custom-nickname",
				"custom-bio",
				"https://test.com/my-custom-profile-pic",
				"https://test.com/my-custom-cover-pic",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
	expected := `{"type":"desmos/MsgSaveProfile","value":{"bio":"biography","cover_picture":"https://shorturl.at/cgpyF","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","dtag":"monk","nickname":"nickname","profile_picture":"https://shorturl.at/adnX3"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgSaveProfile_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgEditProfile.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgEditProfile.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgDeleteProfile = types.NewMsgDeleteProfile(
	testProfile.GetAddress().String(),
)

func TestMsgDeleteProfile_Route(t *testing.T) {
	require.Equal(t, "profiles", msgDeleteProfile.Route())
}

func TestMsgDeleteProfile_Type(t *testing.T) {
	require.Equal(t, "delete_profile", msgDeleteProfile.Type())
}

func TestMsgDeleteProfile_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgDeleteProfile
		error error
	}{
		{
			name:  "Empty owner returns error",
			msg:   types.NewMsgDeleteProfile(""),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator: "),
		},
		{
			name:  "Valid message returns no error",
			msg:   types.NewMsgDeleteProfile(testProfile.GetAddress().String()),
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
	addr, _ := sdk.AccAddressFromBech32(msgDeleteProfile.Creator)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteProfile.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgRequestTransferDTag = types.NewMsgRequestDTagTransfer(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
)

func TestMsgRequestDTagTransfer_Route(t *testing.T) {
	require.Equal(t, "profiles", msgRequestTransferDTag.Route())
}

func TestMsgRequestDTagTransfer_Type(t *testing.T) {
	require.Equal(t, "request_dtag", msgRequestTransferDTag.Type())
}

func TestMsgRequestDTagTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgRequestDTagTransfer
		error error
	}{
		{
			name:  "Empty current owner returns error",
			msg:   types.NewMsgRequestDTagTransfer("", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name:  "Empty receiving user returns error",
			msg:   types.NewMsgRequestDTagTransfer("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", ""),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals current owner and receiving user returns error",
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgRequestDTagTransfer(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
	expected := `{"type":"desmos/MsgRequestDTagTransfer","value":{"receiver":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRequestDTagTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRequestTransferDTag.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgRequestTransferDTag.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgAcceptDTagTransfer = types.NewMsgAcceptDTagTransfer(
	"dtag",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgAcceptDTagTransfer_Route(t *testing.T) {
	require.Equal(t, "profiles", msgAcceptDTagTransfer.Route())
}

func TestMsgAcceptDTagTransfer_Type(t *testing.T) {
	require.Equal(t, "accept_dtag_request", msgAcceptDTagTransfer.Type())
}

func TestMsgAcceptDTagTransfer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgAcceptDTagTransfer
		error error
	}{
		{
			name: "Empty sender user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals current owner and receiving user returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "Empty newDTag returns error",
			msg: types.NewMsgAcceptDTagTransfer(
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new DTag can't be empty"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgAcceptDTagTransfer(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
	expected := `{"type":"desmos/MsgAcceptDTagTransfer","value":{"new_dtag":"dtag","receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgAcceptDTagTransfer_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgAcceptDTagTransfer.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgAcceptDTagTransfer.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgRejectDTagTransfer = types.NewMsgRefuseDTagTransferRequest(
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
)

func TestMsgRejectDTagRequest_Route(t *testing.T) {
	require.Equal(t, "profiles", msgRejectDTagTransfer.Route())
}

func TestMsgRejectDTagRequest_Type(t *testing.T) {
	require.Equal(t, "refuse_dtag_request", msgRejectDTagTransfer.Type())
}

func TestMsgRejectDTagRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgRefuseDTagTransfer
		error error
	}{
		{
			name: "Empty sender returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "No error message",
			msg: types.NewMsgRefuseDTagTransferRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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

func TestMsgRejectDTagRequest_GetSignBytes(t *testing.T) {
	actual := msgRejectDTagTransfer.GetSignBytes()
	expected := `{"type":"desmos/MsgRefuseDTagTransfer","value":{"receiver":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","sender":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgRejectDTagRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgRejectDTagTransfer.Receiver)
	require.Equal(t, []sdk.AccAddress{addr}, msgRejectDTagTransfer.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgCancelDTagTransferReq = types.NewMsgCancelDTagTransferRequest(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
)

func TestMsgCancelDTagRequest_Route(t *testing.T) {
	require.Equal(t, "profiles", msgCancelDTagTransferReq.Route())
}

func TestMsgCancelDTagRequest_Type(t *testing.T) {
	require.Equal(t, "cancel_dtag_request", msgCancelDTagTransferReq.Type())
}

func TestMsgCancelDTagRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name  string
		msg   *types.MsgCancelDTagTransfer
		error error
	}{
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address: "),
		},
		{
			name: "Empty sender returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address: "),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different"),
		},
		{
			name: "No error message",
			msg: types.NewMsgCancelDTagTransferRequest(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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

func TestMsgCancelDTagRequest_GetSignBytes(t *testing.T) {
	actual := msgCancelDTagTransferReq.GetSignBytes()
	expected := `{"type":"desmos/MsgCancelDTagTransfer","value":{"receiver":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","sender":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgCancelDTagRequest_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgCancelDTagTransferReq.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCancelDTagTransferReq.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgCreateRelationship = types.NewMsgCreateRelationship(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgCreateRelationship_Route(t *testing.T) {
	require.Equal(t, "profiles", msgCreateRelationship.Route())
}

func TestMsgCreateRelationship_Type(t *testing.T) {
	require.Equal(t, "create_relationship", msgCreateRelationship.Type())
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
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address"),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address"),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgCreateRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender and receiver must be different"),
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
	addr, _ := sdk.AccAddressFromBech32(msgCreateRelationship.Sender)
	require.Equal(t, []sdk.AccAddress{addr}, msgCreateRelationship.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgDeleteRelationships = types.NewMsgDeleteRelationship(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgDeleteRelationships_Route(t *testing.T) {
	require.Equal(t, "profiles", msgDeleteRelationships.Route())
}

func TestMsgDeleteRelationships_Type(t *testing.T) {
	require.Equal(t, "delete_relationship", msgDeleteRelationships.Type())
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
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address"),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid counterparty address"),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user and counterparty must be different"),
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
	expected := `{"type":"desmos/MsgDeleteRelationship","value":{"counterparty":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgDeleteRelationships_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgDeleteRelationships.User)
	require.Equal(t, []sdk.AccAddress{addr}, msgDeleteRelationships.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgBlockUser = types.NewMsgBlockUser(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"reason",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgBlockUser_Route(t *testing.T) {
	require.Equal(t, "profiles", msgBlockUser.Route())
}

func TestMsgBlockUser_Type(t *testing.T) {
	require.Equal(t, "block_user", msgBlockUser.Type())
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
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker address"),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked address"),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgBlockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "blocker and blocked must be different"),
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
	expected := `{"type":"desmos/MsgBlockUser","value":{"blocked":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47","blocker":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","reason":"reason","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgBlockUser_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgBlockUser.Blocker)
	require.Equal(t, []sdk.AccAddress{addr}, msgBlockUser.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgUnblockUser = types.NewMsgUnblockUser(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
)

func TestMsgUnblockUser_Route(t *testing.T) {
	require.Equal(t, "profiles", msgUnblockUser.Route())
}

func TestMsgUnblockUser_Type(t *testing.T) {
	require.Equal(t, "unblock_user", msgUnblockUser.Type())
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
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker"),
		},
		{
			name: "Empty receiver returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked"),
		},
		{
			name: "Equals sender and receiver returns error",
			msg: types.NewMsgUnblockUser(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "blocker and blocked must be different"),
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
	addr, _ := sdk.AccAddressFromBech32(msgUnblockUser.Blocker)
	require.Equal(t, []sdk.AccAddress{addr}, msgUnblockUser.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgLink = types.NewMsgLink(
	"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
	"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
	"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
	"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
)

func TestMsgLink_Route(t *testing.T) {
	require.Equal(t, "profiles", msgLink.Route())
}

func TestMsgLink_Type(t *testing.T) {
	require.Equal(t, "link", msgLink.Type())
}

func TestMsgLink_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgLink
		err  error
	}{
		{
			name: "Invalid source address returns error",
			msg: types.NewMsgLink(
				"",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid source address"),
		},
		{
			name: "Invalid destination address returns error",
			msg: types.NewMsgLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid destination address"),
		},
		{
			name: "Source address is same as destination address returns error",
			msg: types.NewMsgLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "source address and destination must be different"),
		},
		{
			name: "Invalid source signature returns error",
			msg: types.NewMsgLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"=",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid source signature"),
		},
		{
			name: "Invalid destination signature returns error",
			msg: types.NewMsgLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"=",
			),
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid destination signature"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgLink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
				"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c",
				"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0",
			),
			err: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.err == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.err.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgLink_GetSignBytes(t *testing.T) {
	actual := msgLink.GetSignBytes()
	expected := `{"type":"desmos/MsgLink","value":{"destination_address":"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq","destination_signature":"42dd1f8d98c5de91a12259cf46098104132f69b61eaa24e112bf504d17e1a0b71274dad981bbb4a13dc440905a19be92eaf4497940751f431c530cc4d68e78b0","source_address":"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70","source_signature":"82b1a7005a04b8863fee46af0663d33704dab037f077527f51383b1de09e388a4354c9791a7ceb765d6f6b71e758232cb1d0fd1c82bdef7dfd30e1722a493b1c"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgLink_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgLink.SourceAddress)
	require.Equal(t, []sdk.AccAddress{addr}, msgLink.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgUnlink = types.NewMsgUnlink(
	"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
	"test-net",
	"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
)

func TestMsgUnlink_Route(t *testing.T) {
	require.Equal(t, "profiles", msgUnlink.Route())
}

func TestMsgUnlink_Type(t *testing.T) {
	require.Equal(t, "unlink", msgUnlink.Type())
}

func TestMsgUnlink_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgUnlink
		err  error
	}{
		{
			name: "Invalid owner returns error",
			msg: types.NewMsgUnlink(
				"",
				"test-net",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
			),
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner"),
		},
		{
			name: "Invalid chain id returns error",
			msg: types.NewMsgUnlink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
			),
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "chain id cannot be empty or blank"),
		},
		{
			name: "Invalid target returns error",
			msg: types.NewMsgUnlink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"test-net",
				"",
			),
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid target"),
		},
		{
			name: "No errors message",
			msg: types.NewMsgUnlink(
				"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70",
				"test-net",
				"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq",
			),
			err: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.err == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.err.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgUnlink_GetSignBytes(t *testing.T) {
	actual := msgUnlink.GetSignBytes()
	expected := `{"type":"desmos/MsgUnlink","value":{"chain_id":"test-net","owner":"cosmos1yt7rqhj0hjw92ed0948r2pqwtp9smukurqcs70","target":"cosmos13rzf5gph4drs3qnf63jmuyf4g9q7a4cv9n0uqq"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgUnlink_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUnlink.Owner)
	require.Equal(t, []sdk.AccAddress{addr}, msgUnlink.GetSigners())
}
