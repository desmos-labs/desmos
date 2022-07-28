package simulation

// DONTCOVER

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

// RandomRegisteredReaction returns a random registered reaction from the slice given
func RandomRegisteredReaction(r *rand.Rand, reactions []types.RegisteredReaction) types.RegisteredReaction {
	return reactions[r.Intn(len(reactions))]
}

// GetRandomFreeTextValue returns a random free text value based on the given limit
func GetRandomFreeTextValue(r *rand.Rand, limit uint32) string {
	return simtypes.RandStringOfLength(r, int(limit))
}

// RandomReaction returns a random reaction from the slice given
func RandomReaction(r *rand.Rand, reactions []types.Reaction) types.Reaction {
	return reactions[r.Intn(len(reactions))]
}

// GenerateRandomShorthandCode returns a random shorthand code
func GenerateRandomShorthandCode(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 20)
}

// GenerateRandomDisplayValue returns a random display value
func GenerateRandomDisplayValue(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 40)
}

// GenerateRandomSubspaceReactionsParams returns a randomly reactions params
func GenerateRandomSubspaceReactionsParams(r *rand.Rand, subspaceID uint64) types.SubspaceReactionsParams {
	return types.NewSubspaceReactionsParams(
		subspaceID,
		types.NewRegisteredReactionValueParams(r.Intn(101) < 50),
		types.NewFreeTextValueParams(r.Intn(101) < 50, uint32(r.Intn(100)+2), ""),
	)
}
