package types

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authz "github.com/cosmos/cosmos-sdk/x/authz"
)

func NewMsgGrantTreasuryAuthorization(subspaceID uint64, granter string, grantee string, authorization authz.Authorization, expiration time.Time) *MsgGrantTreasuryAuthorization {
	grant, err := authz.NewGrant(authorization, expiration)
	if err != nil {
		panic("failed to pack authorization to grant")
	}
	return &MsgGrantTreasuryAuthorization{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    grantee,
		Grant:      grant,
	}
}

// Route implements sdk.Msg
func (msg MsgGrantTreasuryAuthorization) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgGrantTreasuryAuthorization) Type() string { return ActionCreateSubspace }

// ValidateBasic implements sdk.Msg
func (msg MsgGrantTreasuryAuthorization) ValidateBasic() error {

	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid granter address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgGrantTreasuryAuthorization) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgGrantTreasuryAuthorization) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Granter)
	return []sdk.AccAddress{addr}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg MsgGrantTreasuryAuthorization) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return msg.Grant.UnpackInterfaces(unpacker)
}
