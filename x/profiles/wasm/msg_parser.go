package wasm

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	wasmTypes "github.com/CosmWasm/wasmvm/types"
	"github.com/desmos-labs/desmos/wasm"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

var _ wasm.MsgParserInterface = MsgParser{}

type MsgParser struct{}

func NewWasmMsgParser() MsgParser {
	return MsgParser{}
}

func (MsgParser) Parse(_ sdk.AccAddress, _ wasmTypes.CosmosMsg) ([]sdk.Msg, error) {
	return nil, nil
}

type ProfilesMsg struct {
	SaveProfile               *types.MsgSaveProfile
	DeleteProfile             *types.MsgDeleteProfile
	CreateRelationship        *types.MsgCreateRelationship
	DeleteRelationship        *types.MsgDeleteRelationship
	BlockUser                 *types.MsgBlockUser
	UnblockUser               *types.MsgUnblockUser
	RequestDTagTransfer       *types.MsgRequestDTagTransfer
	CancelDTagTransferRequest *types.MsgCancelDTagTransferRequest
	AcceptDTagTransferRequest *types.MsgAcceptDTagTransferRequest
	RefuseDTagTransferRequest *types.MsgRefuseDTagTransferRequest
	LinkChainAccount          *types.MsgLinkChainAccount
	UnlinkChainAccount        *types.MsgUnlinkChainAccount
	LinkApplication           *types.MsgLinkApplication
	UnlinkApplication         *types.MsgUnlinkApplication
}

func (MsgParser) ParseCustom(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var msg ProfilesMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to parse profiles message from contract %s", contractAddr.String())
	}

	switch {
	case msg.SaveProfile != nil:
		return []sdk.Msg{msg.SaveProfile}, msg.SaveProfile.ValidateBasic()
	case msg.DeleteProfile != nil:
		return []sdk.Msg{msg.DeleteProfile}, msg.DeleteProfile.ValidateBasic()
	case msg.RequestDTagTransfer != nil:
		return []sdk.Msg{msg.RequestDTagTransfer}, msg.RequestDTagTransfer.ValidateBasic()
	case msg.AcceptDTagTransferRequest != nil:
		return []sdk.Msg{msg.AcceptDTagTransferRequest}, msg.AcceptDTagTransferRequest.ValidateBasic()
	case msg.RefuseDTagTransferRequest != nil:
		return []sdk.Msg{msg.RefuseDTagTransferRequest}, msg.RefuseDTagTransferRequest.ValidateBasic()
	case msg.CancelDTagTransferRequest != nil:
		return []sdk.Msg{msg.CancelDTagTransferRequest}, msg.CancelDTagTransferRequest.ValidateBasic()
	case msg.CreateRelationship != nil:
		return []sdk.Msg{msg.CreateRelationship}, msg.CreateRelationship.ValidateBasic()
	case msg.DeleteRelationship != nil:
		return []sdk.Msg{msg.DeleteRelationship}, msg.DeleteRelationship.ValidateBasic()
	case msg.BlockUser != nil:
		return []sdk.Msg{msg.BlockUser}, msg.BlockUser.ValidateBasic()
	case msg.UnblockUser != nil:
		return []sdk.Msg{msg.UnblockUser}, msg.UnblockUser.ValidateBasic()
	case msg.LinkChainAccount != nil:
		return []sdk.Msg{msg.LinkChainAccount}, msg.LinkChainAccount.ValidateBasic()
	case msg.UnlinkChainAccount != nil:
		return []sdk.Msg{msg.UnlinkChainAccount}, msg.UnlinkChainAccount.ValidateBasic()
	case msg.LinkApplication != nil:
		return []sdk.Msg{msg.LinkApplication}, msg.LinkApplication.ValidateBasic()
	case msg.UnlinkApplication != nil:
		return []sdk.Msg{msg.UnlinkApplication}, msg.UnlinkApplication.ValidateBasic()
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"unrecognized %s message type: %v", types.ModuleName, msg)
	}
}
