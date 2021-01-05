package v0150

import (
	v0130relationships "github.com/desmos-labs/desmos/x/relationships/legacy/v0.13.0"
)

// Migrate accepts exported genesis state from v0.13.0 and migrates it to v0.15.0
// genesis state.
func Migrate(oldGenState v0130relationships.GenesisState) GenesisState {
	return GenesisState{
		Relationships: ConvertRelationships(oldGenState.UsersRelationships),
		Blocks:        ConvertUserBlocks(oldGenState.UsersBlocks),
	}
}

func ConvertRelationships(oldRelationshipsMap map[string][]v0130relationships.Relationship) []Relationship {
	var relationships []Relationship
	for creator, oldRelationships := range oldRelationshipsMap {
		for _, oldRelationship := range oldRelationships {
			relationship := Relationship{
				Creator:   creator,
				Recipient: oldRelationship.Recipient.String(),
				Subspace:  oldRelationship.Subspace,
			}
			relationships = append(relationships, relationship)
		}
	}

	return relationships
}

func ConvertUserBlocks(oldUserBlocks []v0130relationships.UserBlock) []UserBlock {
	userBlocks := make([]UserBlock, len(oldUserBlocks))

	for index, oldUserBlock := range oldUserBlocks {
		userBlock := UserBlock{
			Blocker:  oldUserBlock.Blocker.String(),
			Blocked:  oldUserBlock.Blocked.String(),
			Reason:   oldUserBlock.Reason,
			Subspace: oldUserBlock.Subspace,
		}

		userBlocks[index] = userBlock
	}

	return userBlocks
}
