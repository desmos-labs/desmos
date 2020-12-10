package v0130

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "reports"
)

type GenesisState struct {
	Reports map[string][]Report `json:"reports" yaml:"reports"`
}

type Report struct {
	Type    string         `json:"type" yaml:"type"`       // Identifies the type of the reports
	Message string         `json:"message" yaml:"message"` // Contains the user message
	User    sdk.AccAddress `json:"user" yaml:"user"`       // Identifies the reporting user
}
