package types

import (
	"strings"

	"cosmossdk.io/errors"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgAddReaction{}
	_ sdk.Msg = &MsgRemoveReaction{}
	_ sdk.Msg = &MsgAddRegisteredReaction{}
	_ sdk.Msg = &MsgEditRegisteredReaction{}
	_ sdk.Msg = &MsgRemoveRegisteredReaction{}
	_ sdk.Msg = &MsgSetReactionsParams{}
)

// NewMsgAddReaction returns a new MsgAddReaction instance
func NewMsgAddReaction(subspaceID uint64, postID uint64, value ReactionValue, user string) *MsgAddReaction {
	valueAny, err := codectypes.NewAnyWithValue(value)
	if err != nil {
		panic("failed to pack value to any type")
	}

	return &MsgAddReaction{
		SubspaceID: subspaceID,
		PostID:     postID,
		Value:      valueAny,
		User:       user,
	}
}

// Route implements sdk.Msg
func (msg *MsgAddReaction) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg *MsgAddReaction) Type() string { return ActionAddReaction }

// ValidateBasic implements sdk.Msg
func (msg *MsgAddReaction) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.Value == nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid value: %s", msg.Value)
	}

	err := msg.Value.GetCachedValue().(ReactionValue).Validate()
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid value: %s", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg *MsgAddReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg *MsgAddReaction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg *MsgAddReaction) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var target ReactionValue
	return unpacker.UnpackAny(msg.Value, &target)
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgRemoveReaction returns a new MsgRemoveReaction instance
func NewMsgRemoveReaction(subspaceID uint64, postID uint64, reactionID uint32, user string) *MsgRemoveReaction {
	return &MsgRemoveReaction{
		SubspaceID: subspaceID,
		PostID:     postID,
		ReactionID: reactionID,
		User:       user,
	}
}

// Route implements sdk.Msg
func (msg *MsgRemoveReaction) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg *MsgRemoveReaction) Type() string { return ActionRemoveReaction }

// ValidateBasic implements sdk.Msg
func (msg *MsgRemoveReaction) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.ReactionID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid reaction id: %d", msg.ReactionID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg *MsgRemoveReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg *MsgRemoveReaction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgAddRegisteredReaction returns a new MsgAddRegisteredReaction instance
func NewMsgAddRegisteredReaction(
	subspaceID uint64,
	shorthandCode string,
	displayValue string,
	user string,
) *MsgAddRegisteredReaction {
	return &MsgAddRegisteredReaction{
		SubspaceID:    subspaceID,
		ShorthandCode: shorthandCode,
		DisplayValue:  displayValue,
		User:          user,
	}
}

// Route implements sdk.Msg
func (msg *MsgAddRegisteredReaction) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg *MsgAddRegisteredReaction) Type() string { return ActionAddRegisteredReaction }

// ValidateBasic implements sdk.Msg
func (msg *MsgAddRegisteredReaction) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if strings.TrimSpace(msg.ShorthandCode) == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid shorthand code: %s", msg.ShorthandCode)
	}

	if strings.TrimSpace(msg.DisplayValue) == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid display value: %s", msg.DisplayValue)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg *MsgAddRegisteredReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg *MsgAddRegisteredReaction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}

// IsManageSubspaceMsg implements subspacestypes.ManageSubspaceMsg
func (msg *MsgAddRegisteredReaction) IsManageSubspaceMsg() {}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgEditRegisteredReaction returns a new MsgEditRegisteredReaction instance
func NewMsgEditRegisteredReaction(
	subspaceID uint64,
	registeredReactionID uint32,
	shorthandCode string,
	displayValue string,
	user string,
) *MsgEditRegisteredReaction {
	return &MsgEditRegisteredReaction{
		SubspaceID:           subspaceID,
		RegisteredReactionID: registeredReactionID,
		ShorthandCode:        shorthandCode,
		DisplayValue:         displayValue,
		User:                 user,
	}
}

// Route implements sdk.Msg
func (msg *MsgEditRegisteredReaction) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg *MsgEditRegisteredReaction) Type() string { return ActionEditRegisteredReaction }

// ValidateBasic implements sdk.Msg
func (msg *MsgEditRegisteredReaction) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.RegisteredReactionID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid registered reaction id: %d", msg.RegisteredReactionID)
	}

	if strings.TrimSpace(msg.ShorthandCode) == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid shorthand code: %s", msg.ShorthandCode)
	}

	if strings.TrimSpace(msg.DisplayValue) == "" {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid display value: %s", msg.DisplayValue)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg *MsgEditRegisteredReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg *MsgEditRegisteredReaction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}

// IsManageSubspaceMsg implements subspacestypes.ManageSubspaceMsg
func (msg *MsgEditRegisteredReaction) IsManageSubspaceMsg() {}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgRemoveRegisteredReaction returns a new MsgRemoveRegisteredReaction instance
func NewMsgRemoveRegisteredReaction(
	subspaceID uint64,
	registeredReactionID uint32,
	user string,
) *MsgRemoveRegisteredReaction {
	return &MsgRemoveRegisteredReaction{
		SubspaceID:           subspaceID,
		RegisteredReactionID: registeredReactionID,
		User:                 user,
	}
}

// Route implements sdk.Msg
func (msg *MsgRemoveRegisteredReaction) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg *MsgRemoveRegisteredReaction) Type() string { return ActionRemoveRegisteredReaction }

// ValidateBasic implements sdk.Msg
func (msg *MsgRemoveRegisteredReaction) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.RegisteredReactionID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid registered reaction id: %d", msg.RegisteredReactionID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg *MsgRemoveRegisteredReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg *MsgRemoveRegisteredReaction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}

// IsManageSubspaceMsg implements subspacestypes.ManageSubspaceMsg
func (msg *MsgRemoveRegisteredReaction) IsManageSubspaceMsg() {}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgSetReactionsParams returns a new MsgSetReactionsParams instance
func NewMsgSetReactionsParams(
	subspaceID uint64,
	registeredReaction RegisteredReactionValueParams,
	freeText FreeTextValueParams,
	user string,
) *MsgSetReactionsParams {
	return &MsgSetReactionsParams{
		SubspaceID:         subspaceID,
		RegisteredReaction: registeredReaction,
		FreeText:           freeText,
		User:               user,
	}
}

// Route implements sdk.Msg
func (msg *MsgSetReactionsParams) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg *MsgSetReactionsParams) Type() string { return ActionSetReactionParams }

// ValidateBasic implements sdk.Msg
func (msg *MsgSetReactionsParams) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	err := msg.FreeText.Validate()
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid free text params: %s", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg *MsgSetReactionsParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(msg))
}

// GetSigners implements sdk.Msg
func (msg *MsgSetReactionsParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}

// IsManageSubspaceMsg implements subspacestypes.ManageSubspaceMsg
func (msg *MsgSetReactionsParams) IsManageSubspaceMsg() {}
