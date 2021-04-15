package v0150

import (
	v0150relationships "github.com/desmos-labs/desmos/x/staging/relationships/types"
)

type GenesisState = v0150relationships.GenesisState

func FindRelationshipsForUser(state GenesisState, user string) []Relationship {
	var relationships []Relationship
	for _, relationship := range state.Relationships {
		if relationship.Creator == user {
			relationships = append(relationships, relationship)
		}
	}
	return relationships
}

type Relationship = v0150relationships.Relationship
type UserBlock = v0150relationships.UserBlock
