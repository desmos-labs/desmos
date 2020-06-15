package v0_8_0

type Profile struct {
	DTag     string         `json:"dtag" yaml:"dtag"`
	Moniker  *string        `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio      *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	Pictures *Pictures      `json:"pictures,omitempty" yaml:"pictures,omitempty"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
}
