package types

import (
	"fmt"

	subspacestypes "github.com/desmos-labs/desmos/x/staging/subspaces/types"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	if !subspacestypes.IsValidSubspace(r.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256")
	}

	return nil
}

// MustMarshalRelationship serializes the given relationship using the provided BinaryMarshaler
func MustMarshalRelationship(cdc codec.BinaryMarshaler, relationship Relationship) []byte {
	return cdc.MustMarshalBinaryBare(&relationship)
}

// MustUnmarshalRelationship deserializes the given byte array as a relationship using
// the provided BinaryMarshaler
func MustUnmarshalRelationship(cdc codec.BinaryMarshaler, bz []byte) Relationship {
	var relationship Relationship
	cdc.MustUnmarshalBinaryBare(bz, &relationship)
	return relationship
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

	if !subspacestypes.IsValidSubspace(ub.Subspace) {
		return fmt.Errorf("subspace must be a valid sha-256 hash")
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// MustUnmarshalUserBlock deserializes the given byte array as a UserBlock using the provided BinaryMarshaler
func MustUnmarshalUserBlock(cdc codec.BinaryMarshaler, bz []byte) UserBlock {
	var block UserBlock
	cdc.MustUnmarshalBinaryBare(bz, &block)
	return block
}
