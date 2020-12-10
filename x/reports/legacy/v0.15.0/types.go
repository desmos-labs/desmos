package v0150

type GenesisState struct {
	Reports []Report `json:"reports" yaml:"reports"`
}

type Report struct {
	PostID  string `json:"post_id" yaml:"post_id"`
	Type    string `json:"type,omitempty" yaml:"type"`
	Message string `json:"message,omitempty" yaml:"message"`
	User    string `json:"user,omitempty" yaml:"user"`
}
