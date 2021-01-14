package v0150

type GenesisState struct {
	Relationships []Relationship `json:"relationships"`
	Blocks        []UserBlock    `json:"blocks"`
}

func (state GenesisState) FindRelationshipsForUser(user string) []Relationship {
	var relationships []Relationship
	for _, relationship := range state.Relationships {
		if relationship.Creator == user {
			relationships = append(relationships, relationship)
		}
	}
	return relationships
}

type Relationship struct {
	Creator   string `json:"creator,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Subspace  string `json:"subspace,omitempty"`
}

type UserBlock struct {
	Blocker  string `json:"blocker,omitempty"`
	Blocked  string `json:"blocked,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Subspace string `json:"subspace,omitempty"`
}
