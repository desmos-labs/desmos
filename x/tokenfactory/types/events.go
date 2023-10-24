package types

// DONTCOVER

//nolint:gosec // This is a false positive
const (
	EventTypeCreateDenom      = "create_denom"
	EventTypeMint             = "tf_mint"
	EventTypeBurn             = "tf_burn"
	EventTypeSetDenomMetadata = "set_denom_metadata"

	AttributeValueCategory   = ModuleName
	AttributeKeySubspaceID   = "subspace_id"
	AttributeAmount          = "amount"
	AttributeCreator         = "creator"
	AttributeSubdenom        = "subdenom"
	AttributeNewTokenDenom   = "new_token_denom"
	AttributeMintToAddress   = "mint_to_address"
	AttributeBurnFromAddress = "burn_from_address"
	AttributeDenom           = "denom"
	AttributeDenomMetadata   = "denom_metadata"
)
