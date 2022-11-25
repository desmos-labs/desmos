package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	proto "github.com/gogo/protobuf/proto"
)

var (
	_ sdk.Msg = &MsgGrantUserAllowance{}
	_ sdk.Msg = &MsgGrantGroupAllowance{}
)

func NewMsgGrantUserAllowance(subspaceID uint64, granter string, grantee string, allowance feegranttypes.FeeAllowanceI) *MsgGrantUserAllowance {
	msg, ok := allowance.(proto.Message)
	if !ok {
		panic("cannot proto marshal allowance")
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		panic("failed to pack allowance to any type")
	}
	return &MsgGrantUserAllowance{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    grantee,
		Allowance:  any,
	}
}

func (msg MsgGrantUserAllowance) ValidateBasic() error {
	return nil
}

func (msg MsgGrantUserAllowance) GetSigners() []sdk.AccAddress {
	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{granter}
}

func (msg MsgGrantUserAllowance) Type() string {
	return sdk.MsgTypeURL(&msg)
}

func (msg MsgGrantUserAllowance) Route() string {
	return sdk.MsgTypeURL(&msg)
}

func (msg MsgGrantUserAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

func (msg MsgGrantUserAllowance) GetFeeAllowanceI() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := msg.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, sdkerrors.Wrap(ErrNoAllowance, "failed to get allowance")
	}

	return allowance, nil
}

func (msg MsgGrantUserAllowance) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(msg.Allowance, &allowance)
}

// --------------------------------------------------------------------------------------------------------------------

func NewMsgGrantGroupAllowance(subspaceID uint64, granter string, groupID uint32, allowance feegranttypes.FeeAllowanceI) *MsgGrantGroupAllowance {
	msg, ok := allowance.(proto.Message)
	if !ok {
		panic("cannot proto marshal allowance")
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		panic("failed to pack allowance to any type")
	}
	return &MsgGrantGroupAllowance{
		SubspaceID: subspaceID,
		Granter:    granter,
		GroupID:    groupID,
		Allowance:  any,
	}
}

func (msg MsgGrantGroupAllowance) ValidateBasic() error {
	return nil
}

func (msg MsgGrantGroupAllowance) GetSigners() []sdk.AccAddress {
	granter, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{granter}
}

func (msg MsgGrantGroupAllowance) Type() string {
	return sdk.MsgTypeURL(&msg)
}

func (msg MsgGrantGroupAllowance) Route() string {
	return sdk.MsgTypeURL(&msg)
}

func (msg MsgGrantGroupAllowance) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

func (msg MsgGrantGroupAllowance) GetFeeAllowanceI() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := msg.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, sdkerrors.Wrap(ErrNoAllowance, "failed to get allowance")
	}

	return allowance, nil
}

func (msg MsgGrantGroupAllowance) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(msg.Allowance, &allowance)
}
