package types_test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"

	"github.com/cosmos/cosmos-sdk/types/bech32"
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

func generateMsgLinkChainAccount(t *testing.T) (*types.MsgLinkChainAccount, types.AddressData) {
	srcPrivKey := &secp256k1.PrivKey{Key: []byte{26, 15, 45, 205, 181, 29, 11, 13, 171, 161, 135, 61, 94, 174, 82, 9, 220, 10, 66, 180, 9, 49, 96, 179, 16, 189, 143, 132, 152, 111, 59, 30}}
	srcPubKey := srcPrivKey.PubKey()
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	require.NoError(t, err)

	srcPlainText := srcAddr
	srcSig, err := srcPrivKey.Sign([]byte(srcPlainText))
	require.NoError(t, err)
	srcSigHex := hex.EncodeToString(srcSig)

	destPrivKey := &secp256k1.PrivKey{Key: []byte{25, 15, 45, 205, 181, 29, 11, 13, 171, 161, 135, 61, 94, 174, 82, 9, 220, 10, 66, 180, 9, 49, 96, 179, 16, 189, 143, 132, 152, 111, 59, 30}}
	destPubKey := destPrivKey.PubKey()
	destAddr, err := bech32.ConvertAndEncode("cosmos", destPubKey.Address().Bytes())
	require.NoError(t, err)

	destPlainText := destAddr
	destSig, err := destPrivKey.Sign([]byte(destPlainText))
	require.NoError(t, err)
	destSigHex := hex.EncodeToString(destSig)

	return types.NewMsgLinkChainAccount(
		types.NewBech32Address(srcAddr, "cosmos"),
		types.NewProof(srcPubKey, srcSigHex, srcPlainText),
		types.NewChainConfig("cosmos"),
		destAddr,
		types.NewProof(destPubKey, destSigHex, destPlainText),
	), types.NewBech32Address(srcAddr, "cosmos")
}

func TestMsgLinkChainAccount_Route(t *testing.T) {
	msg, _ := generateMsgLinkChainAccount(t)
	require.Equal(t, "profiles", msg.Route())
}

func TestMsgLinkChainAccount_Type(t *testing.T) {
	msg, _ := generateMsgLinkChainAccount(t)
	require.Equal(t, "link_chain_account", msg.Type())
}

func TestMsgLinkChainAccount_ValidateBasic(t *testing.T) {
	validMsg, srcAddr := generateMsgLinkChainAccount(t)
	tests := []struct {
		name     string
		msg      *types.MsgLinkChainAccount
		expError error
	}{
		{
			name: "Empty source address returns error",
			msg: &types.MsgLinkChainAccount{
				nil,
				validMsg.SourceProof,
				validMsg.SourceChainConfig,
				validMsg.DestinationAddress,
				validMsg.DestinationProof,
			},
			expError: fmt.Errorf("source address cannot be nil"),
		},
		{
			name: "Invalid source proof returns error",
			msg: types.NewMsgLinkChainAccount(
				srcAddr,
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "wrong"),
				validMsg.SourceChainConfig,
				validMsg.DestinationAddress,
				validMsg.DestinationProof,
			),
			expError: fmt.Errorf("failed to decode hex string of signature"),
		},
		{
			name: "Invalid chain config returns error",
			msg: types.NewMsgLinkChainAccount(
				srcAddr,
				validMsg.SourceProof,
				types.NewChainConfig(""),
				validMsg.DestinationAddress,
				validMsg.DestinationProof,
			),
			expError: fmt.Errorf("chain name cannot be empty or blank"),
		},
		{
			name: "Invalid destination address returns error",
			msg: types.NewMsgLinkChainAccount(
				srcAddr,
				validMsg.SourceProof,
				validMsg.SourceChainConfig,
				"",
				validMsg.DestinationProof,
			),
			expError: fmt.Errorf("invalid destination address: %s", ""),
		},
		{
			name: "Invalid destination proof config returns error",
			msg: types.NewMsgLinkChainAccount(
				srcAddr,
				validMsg.SourceProof,
				validMsg.SourceChainConfig,
				validMsg.DestinationAddress,
				types.NewProof(secp256k1.GenPrivKey().PubKey(), "=", "wrong"),
			),
			expError: fmt.Errorf("failed to decode hex string of signature"),
		},
		{
			name:     "No error message",
			msg:      validMsg,
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.expError == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.expError.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgLinkChainAccount_GetSignBytes(t *testing.T) {
	msg, _ := generateMsgLinkChainAccount(t)
	actual := msg.GetSignBytes()
	expected := `{"type":"desmos/MsgLinkChainAccount","value":{"destination_address":"cosmos1u9hgsqfpe3snftr7p7fsyja3wtlmj2sgf2w9yl","destination_proof":{"plain_text":"cosmos1u9hgsqfpe3snftr7p7fsyja3wtlmj2sgf2w9yl","pub_key":{"type":"tendermint/PubKeySecp256k1","value":"AkbOmF4y2laQnlJ+1clOOkCt799+eEKa16yG0l3zdD7W"},"signature":"0cc2f168d580dcaa5894def8f62594a8bfc2d591e36ae613fe194ea17aa3dd0d0a66f7256dc68d9403adb8975af1405ee8674d20702f67f9652b23906cdc275a"},"source_address":{"prefix":"cosmos","value":"cosmos1ma346arwsqpmjmkctwxa5uxdx66le3nty0jeax"},"source_chain_config":{"name":"cosmos"},"source_proof":{"plain_text":"cosmos1ma346arwsqpmjmkctwxa5uxdx66le3nty0jeax","pub_key":{"type":"tendermint/PubKeySecp256k1","value":"A7v3HEjiNO2jXJA+2gcBtO2VQ6Vsirs7GODz7dN39H7Q"},"signature":"ad112abb30e5240c7b9d21b4cc5421d76cfadfcd5977cca262523b5f5bc759457d4aa6d5c1eb6223db104b47aa1f222468be8eb5bb2762b971622ac5b96351b5"}}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgLinkChainAccount_GetSigners(t *testing.T) {
	msg, _ := generateMsgLinkChainAccount(t)
	addr, _ := sdk.AccAddressFromBech32(msg.DestinationAddress)
	require.Equal(t, []sdk.AccAddress{addr}, msg.GetSigners())
}

// ___________________________________________________________________________________________________________________

var msgUnlinkChainAccount = types.NewMsgUnlinkChainAccount(
	"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
	"cosmos",
	"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
)

func TestMsgUnlinkChainAccount_Route(t *testing.T) {
	require.Equal(t, "profiles", msgUnlinkChainAccount.Route())
}

func TestMsgUnlinkChainAccount_Type(t *testing.T) {
	require.Equal(t, "unlink_chain_account", msgUnlinkChainAccount.Type())
}

func TestMsgUnlinkChainAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name     string
		msg      *types.MsgUnlinkChainAccount
		expError error
	}{
		{
			name: "Invalid owner returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"",
				"cosmos",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner"),
		},
		{
			name: "Invalid chain name returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "chain name cannot be empty or blank"),
		},
		{
			name: "Invalid target returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos",
				"",
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid target"),
		},
		{
			name:     "No error message",
			msg:      msgUnlinkChainAccount,
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			returnedError := test.msg.ValidateBasic()
			if test.expError == nil {
				require.Nil(t, returnedError)
			} else {
				require.NotNil(t, returnedError)
				require.Equal(t, test.expError.Error(), returnedError.Error())
			}
		})
	}
}

func TestMsgUnlinkChainAccount_GetSignBytes(t *testing.T) {
	actual := msgUnlinkChainAccount.GetSignBytes()
	expected := `{"type":"desmos/MsgUnlinkChainAccount","value":{"chain_name":"cosmos","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","target":"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}}`
	require.Equal(t, expected, string(actual))
}

func TestMsgUnlinkChainAccount_GetSigners(t *testing.T) {
	addr, _ := sdk.AccAddressFromBech32(msgUnlinkChainAccount.Owner)
	require.Equal(t, []sdk.AccAddress{addr}, msgUnlinkChainAccount.GetSigners())
}
