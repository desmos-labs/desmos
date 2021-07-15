package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrInvalidVersion is returned if a wrong version of IBC channel is used
	ErrInvalidVersion = sdkerrors.Register(ModuleName, 1, "invalid version")

	// ErrMaxProfilesChannels is returned if the channel sequence exceed the max allowed
	ErrMaxProfilesChannels = sdkerrors.Register(ModuleName, 2, "max profiles channels")

	// ErrProfileNotFound is returned when a profile doesn't exist
	ErrProfileNotFound = sdkerrors.Register(ModuleName, 3, "profile not found")

	// ErrProfileAlreadyCreated is returned when a profile has already been created
	ErrProfileAlreadyCreated = sdkerrors.Register(ModuleName, 4, "profile already created")

	// ErrInvalidNickname is returned when a profile's nickname doesn't respect the nickname parameters
	ErrInvalidNickname = sdkerrors.Register(ModuleName, 5, "invalid profile nickname")

	// ErrInvalidDTag is returned when a profile's DTag is empty or doesn't respect the DTag parameters
	ErrInvalidDTag = sdkerrors.Register(ModuleName, 6, "invalid profile DTag")

	// ErrInvalidBio is returned when a profile's Bio doesn't respect the Bio parameters
	ErrInvalidBio = sdkerrors.Register(ModuleName, 7, "invalid profile biography")

	// ErrInvalidDTagRequest is returned when a DTagRequest is not valid (doesn't exist or same sender/receiver)
	ErrInvalidDTagRequest = sdkerrors.Register(ModuleName, 8, "invalid DTag transfer request")

	// ErrBlockEqualUsers is returned when the users of a block are the same
	ErrBlockEqualUsers = sdkerrors.Register(ModuleName, 9, "blocker and blocked cannot be the same user")

	// ErrBlockUserAlreadyBlocked is returned if the blocked user has already been blocked
	ErrBlockUserAlreadyBlocked = sdkerrors.Register(ModuleName, 10, "the user has already been blocked")

	// ErrBlockNotFound is returned when a block doesn't exist
	ErrBlockNotFound = sdkerrors.Register(ModuleName, 11, "block not found")

	// ErrBlockedByUser is returned if a user has been blocked by the user
	ErrBlockedByUser = sdkerrors.Register(ModuleName, 12, "blocked by the user")

	// ErrRelationshipEqualUsers is returned if the two users of a relationship are the same
	ErrRelationshipEqualUsers = sdkerrors.Register(ModuleName, 13, "creator and recipient cannot be the same user")

	// ErrRelationshipNotFound is returned when a relationship doesn't exist
	ErrRelationshipNotFound = sdkerrors.Register(ModuleName, 14, "relationship not found")

	// ErrRelationshipAlreadyCreated is returned when a relationship already exist
	ErrRelationshipAlreadyCreated = sdkerrors.Register(ModuleName, 15, "relationship already created")

	// ErrAppLinkEmptyName is returned when an app link application name is empty
	ErrAppLinkEmptyName = sdkerrors.Register(ModuleName, 16, "application cannot be empty or blank")

	// ErrAppLinkEmptyUsername is returned when an app link username is empty
	ErrAppLinkEmptyUsername = sdkerrors.Register(ModuleName, 17, "username cannot be empty or blank")

	// ErrInvalidPacketData is returned when an IBC packed is invalid
	ErrInvalidPacketData = sdkerrors.Register(ModuleName, 31, "invalid packet data type")

	// ErrChainLinkNotFound is returned when a chain link doesn't exist
	ErrChainLinkNotFound = sdkerrors.Register(ModuleName, 32, "chain link not found")

	// ErrInvalidChainLink is returned when a chain link is not valid
	ErrInvalidChainLink = sdkerrors.Register(ModuleName, 33, "invalid chain link")

	// ErrDuplicatedChainLink is returned when a chain link is duplicated
	ErrDuplicatedChainLink = sdkerrors.Register(ModuleName, 34, "chain link already exists")

	// ErrChainLinkInvalidAddressData is returned when an address data is invalid
	ErrChainLinkInvalidAddressData = sdkerrors.Register(ModuleName, 35, "invalid address data")

	// ErrChainLinkInvalidProof is returned when the the chain link proof is not valid
	ErrChainLinkInvalidProof = sdkerrors.Register(ModuleName, 36, "invalid proof")
)

const (
	ErrIBCTimeout         = "ibc connection timeout"
	ErrRequestExpired     = "oracle request expired"
	ErrRequestFailed      = "oracle request failed"
	ErrInvalidSignature   = "invalid signature"
	ErrInvalidAppUsername = "invalid application username"
)
