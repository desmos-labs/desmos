package simulation

// DONTCOVER

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

// RandomRelationship picks and returns a random relationships from an array
func RandomRelationship(r *rand.Rand, relationships []sdk.AccAddress) sdk.AccAddress {
	idx := r.Intn(len(relationships))
	return relationships[idx]
}

// RandomUserBlock picks and returns a random user block from an array
func RandomUserBlock(r *rand.Rand, userBlocks []types.UserBlock) types.UserBlock {
	idx := r.Intn(len(userBlocks))
	return userBlocks[idx]
}
