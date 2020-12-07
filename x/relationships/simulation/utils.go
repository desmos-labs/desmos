package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

var (
	subspaces = []string{
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88",
		"3d59f7548e1af2151b64135003ce63c0a484c26b9b8b166a7b1c1805ec34b00a",
		"ec8202b6f9fb16f9e26b66367afa4e037752f3c09a18cefab426165e06a424b1",
		"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		"3f40462915a3e6026a4d790127b95ded4d870f6ab18d9af2fcbc454168255237",
	}
)

// RandomRelationship picks and returns a random relationships from an array
func RandomRelationship(r *rand.Rand, relationships []types.Relationship) types.Relationship {
	idx := r.Intn(len(relationships))
	return relationships[idx]
}

// RandomSubspace returns a random post subspace from the above random subspaces
func RandomSubspace(r *rand.Rand) string {
	idx := r.Intn(len(subspaces))
	return subspaces[idx]
}

// RandomUserBlock picks and returns a random user block from an array
func RandomUserBlock(r *rand.Rand, userBlocks []types.UserBlock) types.UserBlock {
	idx := r.Intn(len(userBlocks))
	return userBlocks[idx]
}
