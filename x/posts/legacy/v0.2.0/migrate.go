package v0_2_0

import (
	v120docs "github.com/desmos-labs/desmos/x/posts/legacy/v0.1.0"
)

// Migrate accepts exported genesis state from v0.1.0 and migrates it to v0.2.0
// genesis state. This migration changes the data that are saved for each post
// moving the external reference into the arbitrary data map and adding the new
// subspace field
func Migrate(oldGenState v120docs.GenesisState) GenesisState {
	// TODO: Implement the migration
	return GenesisState{}
}
