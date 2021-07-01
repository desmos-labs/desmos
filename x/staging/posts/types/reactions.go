package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	subspacestypes "github.com/desmos-labs/desmos/x/staging/subspaces/types"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewRegisteredReaction returns a new RegisteredReaction
func NewRegisteredReaction(creator string, shortCode, value, subspace string) RegisteredReaction {
	return RegisteredReaction{
		ShortCode: shortCode,
		Value:     value,
		Subspace:  subspace,
		Creator:   creator,
	}
}

// Validate implements validator
func (reaction RegisteredReaction) Validate() error {
	if reaction.Creator == "" {
		return fmt.Errorf("invalid reaction creator: %s", reaction.Creator)
	}

	if !IsValidReactionCode(reaction.ShortCode) {
		return fmt.Errorf("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'")
	}

	if !commons.IsURIValid(reaction.Value) {
		return fmt.Errorf("reaction value should be a URL")
	}

	if !subspacestypes.IsValidSubspace(reaction.Subspace) {
		return fmt.Errorf("reaction subspace must be a valid sha-256 hash")
	}

	if _, found := GetEmojiByShortCodeOrValue(reaction.ShortCode); found {
		return fmt.Errorf("reaction has emoji shortcode: %s", reaction.ShortCode)
	}

	return nil
}

// MustMarshalRegisteredReaction serializes the given registered reaction using the provided BinaryMarshaler
func MustMarshalRegisteredReaction(cdc codec.BinaryMarshaler, reaction RegisteredReaction) []byte {
	return cdc.MustMarshalBinaryBare(&reaction)
}

// MustUnmarshalRegisteredReaction deserializes the given byte array as a registered reaction using
// the provided BinaryMarshaler
func MustUnmarshalRegisteredReaction(cdc codec.BinaryMarshaler, bz []byte) RegisteredReaction {
	var reaction RegisteredReaction
	cdc.MustUnmarshalBinaryBare(bz, &reaction)
	return reaction
}

// ___________________________________________________________________________________________________________________

// NewPostReaction returns a new PostReaction
func NewPostReaction(postID, shortcode, value, owner string) PostReaction {
	return PostReaction{
		PostID:    postID,
		ShortCode: shortcode,
		Value:     value,
		Owner:     owner,
	}
}

// Validate implements validator
func (reaction PostReaction) Validate() error {
	if !IsValidPostID(reaction.PostID) {
		return fmt.Errorf("invalid post id: %s", reaction.PostID)
	}

	if reaction.Owner == "" {
		return fmt.Errorf("invalid reaction owner: %s", reaction.Owner)
	}

	if len(strings.TrimSpace(reaction.Value)) == 0 {
		return fmt.Errorf("reaction value cannot be empty or blank")
	}

	if !IsValidReactionCode(reaction.ShortCode) {
		return fmt.Errorf("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'")
	}

	return nil
}

// MustMarshalPostReaction serializes the given post reaction using the provided BinaryMarshaler
func MustMarshalPostReaction(cdc codec.BinaryMarshaler, reaction PostReaction) []byte {
	return cdc.MustMarshalBinaryBare(&reaction)
}

// MustUnmarshalPostReaction deserializes the given byte array as a post reaction using
// the provided BinaryMarshaler
func MustUnmarshalPostReaction(cdc codec.BinaryMarshaler, bz []byte) PostReaction {
	var reaction PostReaction
	cdc.MustUnmarshalBinaryBare(bz, &reaction)
	return reaction
}
