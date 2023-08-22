package types

import (
	"encoding/binary"

	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"

	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
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
	ActionEditRegisteredReaction   = "edit_registered_reaction"
	ActionRemoveRegisteredReaction = "remove_registered_reaction"
	ActionSetReactionParams        = "set_reaction_params"

	DoNotModify = "[do-not-modify]"
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
func GetReactionIDBytes(reactionID uint32) (reactionIDBz []byte) {
	reactionIDBz = make([]byte, 4)
	binary.BigEndian.PutUint32(reactionIDBz, reactionID)
	return reactionIDBz
}

// GetReactionIDFromBytes returns reactionID in uint64 format from a byte array
func GetReactionIDFromBytes(bz []byte) (reactionID uint32) {
	return binary.BigEndian.Uint32(bz)
}

// NextSubspaceReactionIDPrefix returns the store prefix used to store all the next reaction ids for the given subspace
func NextSubspaceReactionIDPrefix(subspaceID uint64) []byte {
	return append(NextReactionIDPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// NextReactionIDStoreKey returns the key used to store the next reaction id for the given subspace
func NextReactionIDStoreKey(subspaceID uint64, postID uint64) []byte {
	return append(NextSubspaceReactionIDPrefix(subspaceID), poststypes.GetPostIDBytes(postID)...)
}

// SubspaceReactionsPrefix returns the prefix used to store the reactions for the given subspace
func SubspaceReactionsPrefix(subspaceID uint64) []byte {
	return append(ReactionPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}

// PostReactionsPrefix returns the prefix used to store the reactions for the given post
func PostReactionsPrefix(subspaceID uint64, postID uint64) []byte {
	return append(SubspaceReactionsPrefix(subspaceID), poststypes.GetPostIDBytes(postID)...)
}

// ReactionStoreKey returns the key used to store the reaction with the given id
func ReactionStoreKey(subspaceID uint64, postID uint64, reactionID uint32) []byte {
	return append(PostReactionsPrefix(subspaceID, postID), GetReactionIDBytes(reactionID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// SubspaceReactionsParamsStoreKey returns the key used to store the reactions params for the given subspace id
func SubspaceReactionsParamsStoreKey(subspaceID uint64) []byte {
	return append(ReactionsParamsPrefix, subspacestypes.GetSubspaceIDBytes(subspaceID)...)
}
