package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidVersion      = sdkerrors.Register(ModuleName, 1, "invalid version")
	ErrMaxProfilesChannels = sdkerrors.Register(ModuleName, 2, "max profiles channels")

	ErrProfileNotFound = sdkerrors.Register(ModuleName, 10, "profile not found")

	ErrInvalidChainLink   = sdkerrors.Register(ModuleName, 30, "invalid chain link")
	ErrInvalidAddressData = sdkerrors.Register(ModuleName, 31, "invalid address data")
	ErrInvalidProof       = sdkerrors.Register(ModuleName, 32, "invalid proof")
)
