package types

import "strings"

const (
	ModuleName   = "tokenfactory"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName

	ActionCreateDenom      = "create_denom"
	ActionMint             = "tf_mint"
	ActionBurn             = "tf_burn"
	ActionSetDenomMetadata = "set_denom_metadata"
	ActionUpdateParams     = "update_params"
)

// KeySeparator is used to combine parts of the keys in the store
const KeySeparator = "|"

var (
	DenomAuthorityMetadataKey = "authoritymetadata"
	DenomsPrefixKey           = "denoms"
	CreatorPrefixKey          = "creator"
	ParamsPrefixKey           = "params"
)

// GetDenomPrefixStore returns the store prefix where all the data associated with a specific denom
// is stored
func GetDenomPrefixStore(denom string) []byte {
	return []byte(strings.Join([]string{DenomsPrefixKey, denom, ""}, KeySeparator))
}

// GetCreatorsPrefix returns the store prefix where the list of the denoms created by a specific
// creator are stored
func GetCreatorPrefix(creator string) []byte {
	return []byte(strings.Join([]string{CreatorPrefixKey, creator, ""}, KeySeparator))
}

// GetCreatorsPrefix returns the store prefix where a list of all creator addresses are stored
func GetCreatorsPrefix() []byte {
	return []byte(strings.Join([]string{CreatorPrefixKey, ""}, KeySeparator))
}
