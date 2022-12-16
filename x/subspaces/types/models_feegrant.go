package types

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	proto "github.com/gogo/protobuf/proto"
)

var _ codectypes.UnpackInterfacesMessage = &UserGrant{}
var _ codectypes.UnpackInterfacesMessage = &GroupGrant{}

// NewUserGrant is a constructor for the UserGrant type
func NewUserGrant(subspaceID uint64, granter, grantee string, feeAllowance feegranttypes.FeeAllowanceI) (UserGrant, error) {
	msg, ok := feeAllowance.(proto.Message)
	if !ok {
		return UserGrant{}, sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", feeAllowance)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return UserGrant{}, err
	}

	return UserGrant{
		SubspaceID: subspaceID,
		Granter:    granter,
		Grantee:    grantee,
		Allowance:  any,
	}, nil
}

// Validate implements fmt.Validator
func (u UserGrant) Validate() error {
	if u.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", u.SubspaceID)
	}
	_, err := sdk.AccAddressFromBech32(u.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing granter address")
	}
	_, err = sdk.AccAddressFromBech32(u.Grantee)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing grantee address")
	}
	if u.Grantee == u.Granter {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "cannot self-grant fee authorization")
	}

	f, err := u.GetUnpackedAllowance()
	if err != nil {
		return err
	}
	return f.ValidateBasic()
}

// GetUnpackedAllowance unpacks allowance
func (u UserGrant) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := u.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, fmt.Errorf("failed to unpack allowance")
	}

	return allowance, nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (u UserGrant) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(u.Allowance, &allowance)
}

// --------------------------------------------------------------------------------------------------------------------

// NewGroupGrant is a constructor for the GroupGrant type
func NewGroupGrant(subspaceID uint64, granter string, groupID uint32, feeAllowance feegranttypes.FeeAllowanceI) (GroupGrant, error) {
	msg, ok := feeAllowance.(proto.Message)
	if !ok {
		return GroupGrant{}, sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", feeAllowance)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return GroupGrant{}, err
	}

	return GroupGrant{
		SubspaceID: subspaceID,
		Granter:    granter,
		GroupID:    groupID,
		Allowance:  any,
	}, nil
}

// Validate implements fmt.Validator
func (g GroupGrant) Validate() error {
	if g.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", g.SubspaceID)
	}
	if g.GroupID == 0 {
		return fmt.Errorf("invalid group id: %d", g.GroupID)
	}
	_, err := sdk.AccAddressFromBech32(g.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing granter address")
	}

	f, err := g.GetUnpackedAllowance()
	if err != nil {
		return err
	}
	return f.ValidateBasic()
}

// GetUnpackedAllowance unpacks allowance
func (g GroupGrant) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := g.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, fmt.Errorf("failed to unpack allowance")
	}

	return allowance, nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (a GroupGrant) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(a.Allowance, &allowance)
}
