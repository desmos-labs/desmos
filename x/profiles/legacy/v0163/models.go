package v0163

import "github.com/cosmos/cosmos-sdk/codec"

// MustUnmarshalRelationships deserializes the given byte array as an array of Relationship using
// the provided BinaryMarshaler
func MustUnmarshalRelationships(codec codec.BinaryMarshaler, bz []byte) []Relationship {
	var wrapped Relationships
	codec.MustUnmarshalBinaryBare(bz, &wrapped)
	return wrapped.Relationships
}

// MustUnmarshalUserBlocks deserializes the given byte array as an array of UserBlock using
// the provided BinaryMarshaler
func MustUnmarshalUserBlocks(codec codec.BinaryMarshaler, bz []byte) []UserBlock {
	var wrapped UserBlocks
	codec.MustUnmarshalBinaryBare(bz, &wrapped)
	return wrapped.Blocks
}

// MustUnmarshalDTagTransferRequests deserializes the given byte array as an array of DTagTransferRequest using
// the provided BinaryMarshaler
func MustUnmarshalDTagTransferRequests(codec codec.BinaryMarshaler, bz []byte) []DTagTransferRequest {
	var wrapped DTagTransferRequests
	codec.MustUnmarshalBinaryBare(bz, &wrapped)
	return wrapped.Requests
}
