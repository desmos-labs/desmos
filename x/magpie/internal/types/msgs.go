package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgCreateSession defines the MsgCreateSession message
type MsgCreateSession struct {
	Owner         sdk.AccAddress `json:"owner"`
	Created       time.Time      `json:"created"`
	Namespace     string         `json:"namespace"`
	ExternalOwner string         `json:"external_owner"`
	PubKey        string         `json:"pub_key"`
	Signature     string         `json:"signature"`
}

// NewMsgCreateSession is the constructor of MsgCreateSession
func NewMsgCreateSession(created time.Time, owner sdk.AccAddress, namespace string, externalOwner string, pubkey string, signature string) MsgCreateSession {
	return MsgCreateSession{
		Created:       created,
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
func (msg MsgCreateSession) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrUnknownRequest("Session owner cannot be empty.")
	}

	if msg.Created.IsZero() {
		return sdk.ErrUnknownRequest("Session created time cannot be empty")
	}

	if len(msg.Namespace) == 0 {
		return sdk.ErrUnknownRequest("Session namespace cannot be empty")
	}

	if len(msg.PubKey) == 0 {
		return sdk.ErrUnknownRequest("Signer pubkey cannot be empty")
	}

	// The external signer address doesn't have to exist on Desmos
	if msg.ExternalOwner == "" {
		return sdk.ErrUnknownRequest("Session external owner cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateSession) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateSession) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
