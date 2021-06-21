package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidVersion      = sdkerrors.Register(ModuleName, 1, "invalid version")
	ErrMaxProfilesChannels = sdkerrors.Register(ModuleName, 2, "max profiles channels")

	ErrProfileNotFound = sdkerrors.Register(ModuleName, 10, "profile not found")

	ErrInvalidPacketData   = sdkerrors.Register(ModuleName, 31, "invalid packet data type")
	ErrInvalidChainLink    = sdkerrors.Register(ModuleName, 35, "invalid chain link")
	ErrDuplicatedChainLink = sdkerrors.Register(ModuleName, 36, "chain link already exists")
	ErrInvalidAddressData  = sdkerrors.Register(ModuleName, 37, "invalid address data")
	ErrInvalidProof        = sdkerrors.Register(ModuleName, 38, "invalid proof")
)

const (
	ErrIBCTimeout         = "ibc connection timeout"
	ErrRequestExpired     = "oracle request expired"
	ErrRequestFailed      = "oracle request failed"
	ErrInvalidSignature   = "invalid signature"
	ErrInvalidAppUsername = "invalid application username"
)
