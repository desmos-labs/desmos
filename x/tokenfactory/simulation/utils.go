package simulation

import (
	"math/rand"
)

// RandomDenom returns a random denom from the given slice
func RandomDenom(r *rand.Rand, denoms []string) string {
	return denoms[r.Intn(len(denoms))]
}
