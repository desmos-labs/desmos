package simulation

// DONTCOVER

import (
	"math/rand"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

// RandomGenesisAccount picks and returns a random genesis account from an array
func RandomGenesisAccount(r *rand.Rand, accounts []authtypes.GenesisAccount) authtypes.GenesisAccount {
	return accounts[r.Intn(len(accounts))]
}

// RandomRelationship picks and returns a random relationships from an array
func RandomRelationship(r *rand.Rand, relationships []types.Relationship) types.Relationship {
	return relationships[r.Intn(len(relationships))]
}

// RandomUserBlock picks and returns a random user block from an array
func RandomUserBlock(r *rand.Rand, userBlocks []types.UserBlock) types.UserBlock {
	return userBlocks[r.Intn(len(userBlocks))]
}
