package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgCreateSession is the constructor of MsgCreateSession
func NewMsgCreateSession(
	owner string, namespace string, externalOwner string, pubkey string, signature string,
) *MsgCreateSession {
	return &MsgCreateSession{
		Owner:         owner,
		Namespace:     namespace,
		ExternalOwner: externalOwner,
		PubKey:        pubkey,
		Signature:     signature,
	}
}

// Route should return the name of the module
func (msg MsgCreateSession) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateSession) Type() string { return ActionCreationSession }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateSession) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, msg.Owner)
	}

	if len(strings.TrimSpace(msg.Namespace)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "session namespace cannot be empty")
	}

	if len(strings.TrimSpace(msg.PubKey)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "signer's public key cannot be empty")
	}

	// The external signer address doesn't have to exist on Desmos
	if len(strings.TrimSpace(msg.ExternalOwner)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "session external owner cannot be empty")
	}

	if len(strings.TrimSpace(msg.Signature)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "session signature cannot be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateSession) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateSession) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{owner}
}
