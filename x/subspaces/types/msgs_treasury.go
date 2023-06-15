package types

import (
	"fmt"
	"strings"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

var (
	_ sdk.Msg = &MsgGrantTreasuryAuthorization{}
	_ sdk.Msg = &MsgRevokeTreasuryAuthorization{}

	_ legacytx.LegacyMsg = &MsgGrantTreasuryAuthorization{}
	_ legacytx.LegacyMsg = &MsgRevokeTreasuryAuthorization{}

	_ ManageSubspaceMsg = &MsgGrantTreasuryAuthorization{}
	_ ManageSubspaceMsg = &MsgRevokeTreasuryAuthorization{}

	_ codectypes.UnpackInterfacesMessage = &MsgGrantTreasuryAuthorization{}
)

func NewMsgGrantTreasuryAuthorization(subspaceID uint64, granter string, grantee string, authorization authz.Authorization, expiration *time.Time) *MsgGrantTreasuryAuthorization {
	grant, err := authz.NewGrant(time.Time{}, authorization, expiration)
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

// Route implements legacytx.LegacyMsg
func (msg *MsgGrantTreasuryAuthorization) Route() string { return RouterKey }

// Type implements legacytx.LegacyMsg
func (msg *MsgGrantTreasuryAuthorization) Type() string { return ActionGrantTreasuryAuthorization }

// ValidateBasic implements sdk.Msg
func (msg *MsgGrantTreasuryAuthorization) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("invalid subspace id")
	}

	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid granter address: %s", msg.Granter)
	}

	_, err = sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", msg.Grantee)
	}

	return msg.Grant.ValidateBasic()
}

// GetSignBytes implements legacytx.LegacyMsg
func (msg *MsgGrantTreasuryAuthorization) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg *MsgGrantTreasuryAuthorization) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Granter)
	return []sdk.AccAddress{addr}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg *MsgGrantTreasuryAuthorization) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return msg.Grant.UnpackInterfaces(unpacker)
}

// IsManageSubspaceMsg implements subspacestypes.ManageSubspaceMsg
func (msg *MsgGrantTreasuryAuthorization) IsManageSubspaceMsg() {}

// --------------------------------------------------------------------------------------------------------------------

func NewMsgRevokeTreasuryAuthorization(subspaceID uint64, granter string, grantee string, msgTypeUrl string) *MsgRevokeTreasuryAuthorization {
	return &MsgRevokeTreasuryAuthorization{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    grantee,
		MsgTypeUrl: msgTypeUrl,
	}
}

// GetSignBytes implements legacytx.LegacyMsg
func (msg *MsgRevokeTreasuryAuthorization) Route() string { return RouterKey }

// GetSignBytes implements legacytx.LegacyMsg
func (msg *MsgRevokeTreasuryAuthorization) Type() string { return ActionRevokeTreasuryAuthorization }

// ValidateBasic implements sdk.Msg
func (msg *MsgRevokeTreasuryAuthorization) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id")
	}

	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid granter address: %s", msg.Granter)
	}

	_, err = sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid grantee address: %s", msg.Grantee)
	}

	if strings.TrimSpace(msg.MsgTypeUrl) == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("missing method name")
	}

	return nil
}

// GetSignBytes implements legacytx.LegacyMsg
func (msg *MsgRevokeTreasuryAuthorization) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg *MsgRevokeTreasuryAuthorization) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Granter)
	return []sdk.AccAddress{addr}
}

// IsManageSubspaceMsg implements subspacestypes.ManageSubspaceMsg
func (msg *MsgRevokeTreasuryAuthorization) IsManageSubspaceMsg() {}
