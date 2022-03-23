package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewRelationship returns a new relationships with the given recipient and subspace
func NewRelationship(user string, counterparty string, subspaceID uint64) Relationship {
	return Relationship{
		Creator:      user,
		Counterparty: counterparty,
		SubspaceID:   subspaceID,
	}
}

// Validate implement Validator
func (r Relationship) Validate() error {
	_, err := sdk.AccAddressFromBech32(r.Creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %s", r.Creator)
	}

	_, err = sdk.AccAddressFromBech32(r.Counterparty)
	if err != nil {
		return fmt.Errorf("invalid counterparty address: %s", r.Counterparty)
	}

	if r.Creator == r.Counterparty {
		return fmt.Errorf("creator and recipient cannot be the same user")
	}

	return nil
}

// MustMarshalRelationship serializes the given relationship using the provided BinaryCodec
func MustMarshalRelationship(cdc codec.BinaryCodec, relationship Relationship) []byte {
	return cdc.MustMarshal(&relationship)
}

// MustUnmarshalRelationship deserializes the given byte array as a relationship using
// the provided BinaryCodec
func MustUnmarshalRelationship(cdc codec.BinaryCodec, bz []byte) Relationship {
	var relationship Relationship
	cdc.MustUnmarshal(bz, &relationship)
	return relationship
}

// --------------------------------------------------------------------------------------------------------------------

// NewUserBlock returns a new object representing the fact that one user has blocked another one
// for a specific reason on the given subspace.
func NewUserBlock(blocker, blocked string, reason string, subspaceID uint64) UserBlock {
	return UserBlock{
		Blocker:    blocker,
		Blocked:    blocked,
		Reason:     reason,
		SubspaceID: subspaceID,
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

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// MustMarshalUserBlock serializes the given user block using the provided BinaryCodec
func MustMarshalUserBlock(cdc codec.BinaryCodec, userBlock UserBlock) []byte {
	return cdc.MustMarshal(&userBlock)
}

// MustUnmarshalUserBlock deserializes the given byte array as a UserBlock using the provided BinaryCodec
func MustUnmarshalUserBlock(cdc codec.BinaryCodec, bz []byte) UserBlock {
	var block UserBlock
	cdc.MustUnmarshal(bz, &block)
	return block
}
