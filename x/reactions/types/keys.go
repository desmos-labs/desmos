package types

import (
	"encoding/binary"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// DONTCOVER

const (
	ModuleName   = "reactions"
	RouterKey    = ModuleName
	StoreKey     = ModuleName
	QuerierRoute = ModuleName

	ActionAddReaction              = "add_reaction"
	ActionRemoveReaction           = "remove_reaction"
	ActionAddRegisteredReaction    = "add_registered_reaction"
	ActionRemoveRegisteredReaction = "remove_registered_reaction"
	ActionSetReactionParams        = "set_reaction_params"
)

var (
	NextRegisteredReactionIDPrefix = []byte{0x01}
	RegisteredReactionPrefix       = []byte{0x02}

	NextReactionIDPrefix = []byte{0x10}
	ReactionPrefix       = []byte{0x11}

	ReactionsParamsPrefix = []byte{0x20}
)

// GetRegisteredReactionIDBytes returns the byte representation of the registeredReactionID
func GetRegisteredReactionIDBytes(registeredReactionID uint32) (reactionIDBz []byte) {
	reactionIDBz = make([]byte, 4)
	binary.BigEndian.PutUint32(reactionIDBz, registeredReactionID)
	return reactionIDBz
}

// GetRegisteredReactionIDFromBytes returns registeredReactionID in uint32 format from a byte array
func GetRegisteredReactionIDFromBytes(bz []byte) (registeredReactionID uint32) {
	return binary.BigEndian.Uint32(bz)
}

// NextRegisteredReactionIDStoreKey returns the key used to store the next registered reaction id for the given subspace
func NextRegisteredReactionIDStoreKey(subspaceID uint64) []byte {
	return append(NextRegisteredReactionIDPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// SubspaceRegisteredReactionsPrefix returns the prefix used to store the registered reactions for the given subspace
func SubspaceRegisteredReactionsPrefix(subspaceID uint64) []byte {
	return append(RegisteredReactionPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// RegisteredReactionStoreKey returns the key used to store the registered reaction with the given id
func RegisteredReactionStoreKey(subspaceID uint64, registeredReactionID uint32) []byte {
	return append(SubspaceRegisteredReactionsPrefix(subspaceID), GetRegisteredReactionIDBytes(registeredReactionID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// GetReactionIDBytes returns the byte representation f the reactionID
func GetReactionIDBytes(reactionID uint64) (reactionIDBz []byte) {
	reactionIDBz = make([]byte, 8)
	binary.BigEndian.PutUint64(reactionIDBz, reactionID)
	return reactionIDBz
}

// GetReactionIDFromBytes returns reactionID in uint64 format from a byte array
func GetReactionIDFromBytes(bz []byte) (reactionID uint64) {
	return binary.BigEndian.Uint64(bz)
}

// NextReactionIDStoreKey returns the key used to store the next reaction id for the given subspace
func NextReactionIDStoreKey(subspaceID uint64) []byte {
	return append(NextReactionIDPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// SubspaceReactionsPrefix returns the prefix used to store the registered reactions for the given subspace
func SubspaceReactionsPrefix(subspaceID uint64) []byte {
	return append(ReactionPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// ReactionStoreKey returns the key used to store the registered reaction with the given id
func ReactionStoreKey(subspaceID uint64, reactionID uint64) []byte {
	return append(SubspaceReactionsPrefix(subspaceID), GetReactionIDBytes(reactionID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// ReactionsParamsStoreKey returns the key used to store the reactions params for the given subspace id
func ReactionsParamsStoreKey(subspaceID uint64) []byte {
	return append(ReactionsParamsPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}
