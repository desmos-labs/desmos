package v0163

import "github.com/cosmos/cosmos-sdk/codec"

// MustUnmarshalRelationships deserializes the given byte array as an array of relationships using
// the provided BinaryMarshaler
func MustUnmarshalRelationships(codec codec.BinaryMarshaler, bz []byte) []Relationship {
	var wrapped Relationships
	codec.MustUnmarshalBinaryBare(bz, &wrapped)
	return wrapped.Relationships
}
