package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewRelationship returns a new relationships with the given recipient and subspace
func NewRelationship(creator string, recipient string, subspace string) Relationship {
	return Relationship{
		Creator:   creator,
		Recipient: recipient,
		Subspace:  subspace,
	}
}

// Validate implement Validator
func (r Relationship) Validate() error {
	_, err := sdk.AccAddressFromBech32(r.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %s", r.Creator)
	}

	_, err = sdk.AccAddressFromBech32(r.Recipient)
	if err != nil {
		return fmt.Errorf("invalid recipient address: %s", r.Recipient)
	}

	if r.Creator == r.Recipient {
		return fmt.Errorf("creator and recipient cannot be the same user")
	}

	if !commons.IsValidSubspace(r.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// RemoveRelationship removes the given relationships from the provided relationships slice.
// If the relationship was found, returns the slice with it removed and true.
// Otherwise, returns the original slice and false
func RemoveRelationship(relationships []Relationship, relationship Relationship) ([]Relationship, bool) {
	for index, rel := range relationships {
		if rel.Equal(relationship) {
			return append(relationships[:index], relationships[index+1:]...), true
		}
	}
	return relationships, false
}

// MustMarshalRelationships serializes the given relationships using the provided BinaryMarshaler
func MustMarshalRelationships(cdc codec.BinaryMarshaler, relationships []Relationship) []byte {
	wrapped := Relationships{Relationships: relationships}
	return cdc.MustMarshalBinaryBare(&wrapped)
}

// MustUnmarshalRelationships deserializes the given byte array as an array of relationships using
// the provided BinaryMarshaler
func MustUnmarshalRelationships(codec codec.BinaryMarshaler, bz []byte) []Relationship {
	var wrapped Relationships
	codec.MustUnmarshalBinaryBare(bz, &wrapped)
	return wrapped.Relationships
}

// ___________________________________________________________________________________________________________________

// NewUserBlock returns a new object representing the fact that one user has blocked another one
// for a specific reason on the given subspace.
func NewUserBlock(blocker, blocked string, reason, subspace string) UserBlock {
	return UserBlock{
		Blocker:  blocker,
		Blocked:  blocked,
		Reason:   reason,
		Subspace: subspace,
	}
}

// Validate implements validator
func (ub UserBlock) Validate() error {
	if len(ub.Blocker) == 0 {
		return fmt.Errorf("blocker address cannot be empty")
	}

	if len(ub.Blocked) == 0 {
		return fmt.Errorf("the address of the blocked user cannot be empty")
	}

	if ub.Blocker == ub.Blocked {
		return fmt.Errorf("blocker and blocked addresses cannot be equals")
	}

	if !commons.IsValidSubspace(ub.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256 hash")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// RemoveUserBlock removes the block made from the blocker towards the blocked inside the subspace,
// from the provided slice of blocks.
// If the block is found, returns the new slice with it removed and true.
// If the block is not found, returns the original fl
func RemoveUserBlock(blocks []UserBlock, blocker, blocked, subspace string) ([]UserBlock, bool) {
	for index, ub := range blocks {
		if ub.Blocker == blocker && ub.Blocked == blocked && ub.Subspace == subspace {
			return append(blocks[:index], blocks[index+1:]...), true

		}
	}
	return blocks, false
}

// MustMarshalUserBlocks serializes the given blocks using the provided BinaryMarshaler
func MustMarshalUserBlocks(cdc codec.BinaryMarshaler, block []UserBlock) []byte {
	wrapped := UserBlocks{Blocks: block}
	return cdc.MustMarshalBinaryBare(&wrapped)
}

// MustUnmarshalUserBlocks deserializes the given byte array as an array of blocks using
// the provided BinaryMarshaler
func MustUnmarshalUserBlocks(cdc codec.BinaryMarshaler, bz []byte) []UserBlock {
	var wrapped UserBlocks
	cdc.MustUnmarshalBinaryBare(bz, &wrapped)
	return wrapped.Blocks
}
