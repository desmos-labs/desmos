package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

// RandomRelationship picks and returns a random relationships from an array
func RandomRelationship(r *rand.Rand, relationships []types.Relationship) types.Relationship {
	idx := r.Intn(len(relationships))
	return relationships[idx]
}

// RandomSubspace returns a random post subspace from the above random subspaces
func RandomSubspace(_ *rand.Rand) string {
	return ""
}

// RandomUserBlock picks and returns a random user block from an array
func RandomUserBlock(r *rand.Rand, userBlocks []types.UserBlock) types.UserBlock {
	idx := r.Intn(len(userBlocks))
	return userBlocks[idx]
}
