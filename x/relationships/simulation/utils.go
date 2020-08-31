package simulation

// DONTCOVER

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RandomRelationship picks and returns a random relationships from an array
func RandomRelationship(r *rand.Rand, relationships []sdk.AccAddress) sdk.AccAddress {
	idx := r.Intn(len(relationships))
	return relationships[idx]
}
