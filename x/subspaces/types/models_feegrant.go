package types

import (
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	proto "github.com/gogo/protobuf/proto"
)

var _ codectypes.UnpackInterfacesMessage = &UserGrant{}
var _ codectypes.UnpackInterfacesMessage = &GroupGrant{}

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

func (a UserGrant) ValidateBasic() error {
	if a.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", a.SubspaceID)
	}
	if strings.TrimSpace(a.Granter) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing granter address")
	}
	if strings.TrimSpace(a.Grantee) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing grantee address")
	}
	if a.Grantee == a.Granter {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "cannot self-grant fee authorization")
	}

	f, err := a.GetUnpackedAllowance()
	if err != nil {
		return err
	}
	return f.ValidateBasic()
}

// GetUnpackedAllowance unpacks allowance
func (a UserGrant) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := a.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
	if !ok {
		return nil, fmt.Errorf("failed to unpack allowance")
	}

	return allowance, nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (a UserGrant) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var allowance feegranttypes.FeeAllowanceI
	return unpacker.UnpackAny(a.Allowance, &allowance)
}

// --------------------------------------------------------------------------------------------------------------------

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

func (a GroupGrant) ValidateBasic() error {
	if a.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", a.SubspaceID)
	}
	if a.GroupID == 0 {
		return fmt.Errorf("invalid group id: %d", a.GroupID)
	}
	if strings.TrimSpace(a.Granter) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing granter address")
	}

	f, err := a.GetUnpackedAllowance()
	if err != nil {
		return err
	}
	return f.ValidateBasic()
}

// GetUnpackedAllowance unpacks allowance
func (a GroupGrant) GetUnpackedAllowance() (feegranttypes.FeeAllowanceI, error) {
	allowance, ok := a.Allowance.GetCachedValue().(feegranttypes.FeeAllowanceI)
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
