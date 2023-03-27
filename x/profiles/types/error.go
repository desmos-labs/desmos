package types

import errors "cosmossdk.io/errors"

// DONTCOVER

var (
	ErrInvalidVersion      = errors.Register(ModuleName, 1, "invalid version")
	ErrMaxProfilesChannels = errors.Register(ModuleName, 2, "max profiles channels")

	ErrProfileNotFound = errors.Register(ModuleName, 10, "profile not found")

	ErrInvalidPacketData   = errors.Register(ModuleName, 31, "invalid packet data type")
	ErrInvalidChainLink    = errors.Register(ModuleName, 35, "invalid chain link")
	ErrDuplicatedChainLink = errors.Register(ModuleName, 36, "chain link already exists")
	ErrInvalidAddressData  = errors.Register(ModuleName, 37, "invalid address data")
	ErrInvalidProof        = errors.Register(ModuleName, 38, "invalid proof")
)

const (
	ErrIBCTimeout         = "ibc connection timeout"
	ErrRequestExpired     = "oracle request expired"
	ErrRequestFailed      = "oracle request failed"
	ErrInvalidSignature   = "invalid signature"
	ErrInvalidAppUsername = "invalid application username"
)
