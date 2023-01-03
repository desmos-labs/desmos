package types

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	proto "github.com/gogo/protobuf/proto"
)

var _ codectypes.UnpackInterfacesMessage = &Grant{}

// GrantTarget represents a generic grant target
type GrantTarget interface {
	proto.Message

	isGrantTarget()
	Validate() error
}

// NewUserTarget is a constructor for the UserTarget type
func NewUserTarget(user string) *UserTarget {
	return &UserTarget{
		User: user,
	}
}

// isGrantTarget implements GrantTarget
func (t *UserTarget) isGrantTarget() {}

// isGrantTarget implements GrantTarget
func (t *UserTarget) Validate() error {
	_, err := sdk.AccAddressFromBech32(t.User)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid grantee address")
	}
	return err
}

// --------------------------------------------------------------------------------------------------------------------

// NewGroupTarget is a constructor for the GroupTarget type
func NewGroupTarget(groupID uint32) *GroupTarget {
	return &GroupTarget{
		GroupID: groupID,
	}
}

// isGrantTarget implements GrantTarget
func (t *GroupTarget) isGrantTarget() {}

// isGrantTarget implements GrantTarget
func (t *GroupTarget) Validate() error {
	if t.GroupID == 0 {
		return fmt.Errorf("invalid group id: %d", t.GroupID)
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewGrant is a constructor for the Grant type
func NewGrant(subspaceID uint64, granter string, target GrantTarget, feeAllowance feegranttypes.FeeAllowanceI) (Grant, error) {
	msg, ok := feeAllowance.(proto.Message)
	if !ok {
		return Grant{}, sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", feeAllowance)
	}

	allowanceAny, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return Grant{}, err
	}

	targetAny, err := codectypes.NewAnyWithValue(target)
	if err != nil {
		return Grant{}, err
	}

	return Grant{
		SubspaceID: subspaceID,
		Granter:    granter,
		Target:     targetAny,
		Allowance:  allowanceAny,
	}, nil
}

// Validate implements fmt.Validator
func (g Grant) Validate() error {
	if g.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", g.SubspaceID)
	}
	_, err := sdk.AccAddressFromBech32(g.Granter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid granter address")
	}

	target := g.Target.GetCachedValue().(GrantTarget)
	err = target.Validate()
	if err != nil {
		return err
	}

	if u, ok := target.(*UserTarget); ok {
		if u.User == g.Granter {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "cannot self-grant fee authorization")
		}
	}

	f, err := g.GetUnpackedAllowance()
	if err != nil {
		return err
	}
	return f.ValidateBasic()
}

// GetUnpackedAllowance unpacks allowance
func (u Grant) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := u.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, fmt.Errorf("failed to unpack allowance")
	}

	return allowance, nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (u Grant) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var target GrantTarget
	err := unpacker.UnpackAny(u.Target, &target)
	if err != nil {
		return err
	}
	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(u.Allowance, &allowance)
}
