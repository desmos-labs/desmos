package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidVersion      = sdkerrors.Register(ModuleName, 1, "invalid version")
	ErrMaxProfilesChannels = sdkerrors.Register(ModuleName, 2, "max profiles channels")

	ErrProfileNotFound       = sdkerrors.Register(ModuleName, 3, "profile not found")
	ErrProfileAlreadyCreated = sdkerrors.Register(ModuleName, 4, "profile already created")

	ErrInvalidNickname = sdkerrors.Register(ModuleName, 5, "invalid profile nickname")
	ErrInvalidDTag     = sdkerrors.Register(ModuleName, 6, "invalid profile DTag")
	ErrInvalidBio      = sdkerrors.Register(ModuleName, 7, "invalid profile biography")

	ErrInvalidDTagRequest = sdkerrors.Register(ModuleName, 8, "invalid DTag transfer request")

	ErrInvalidBlock        = sdkerrors.Register(ModuleName, 9, "invalid block")
	ErrBlockAlreadyCreated = sdkerrors.Register(ModuleName, 10, "block already created")
	ErrBlockNotFound       = sdkerrors.Register(ModuleName, 11, "block not found")
	ErrBlockedByUser       = sdkerrors.Register(ModuleName, 12, "blocked by the user")

	ErrInvalidRelationship        = sdkerrors.Register(ModuleName, 13, "invalid relationship")
	ErrRelationshipNotFound       = sdkerrors.Register(ModuleName, 14, "relationship not found")
	ErrRelationshipAlreadyCreated = sdkerrors.Register(ModuleName, 15, "relationship already created")

	ErrChainLinkNotFound = sdkerrors.Register(ModuleName, 16, "chain link not found")

	ErrInvalidAppLink = sdkerrors.Register(ModuleName, 17, "invalid app link")

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
