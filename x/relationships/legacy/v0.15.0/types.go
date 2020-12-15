package v0150

type GenesisState struct {
	Relationships []Relationship `json:"relationships" yaml:"users_relationships"`
	Blocks        []UserBlock    `json:"blocks" yaml:"users_blocks"`
}

type Relationship struct {
	Creator   string `json:"creator,omitempty" yaml:"creator"`
	Recipient string `json:"recipient,omitempty" yaml:"recipient"`
	Subspace  string `json:"subspace,omitempty" yaml:"subspace"`
}

type UserBlock struct {
	Blocker  string `json:"blocker,omitempty" yaml:"blocker"`
	Blocked  string `json:"blocked,omitempty" yaml:"blocked"`
	Reason   string `json:"reason,omitempty" yaml:"reason"`
	Subspace string `json:"subspace,omitempty" yaml:"subspace"`
}
