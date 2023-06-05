package types

import "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"

const (
	ModuleName   = types.ModuleName
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = types.MemStoreKey

	ActionCreateDenom      = "create_denom"
	ActionMint             = "tf_mint"
	ActionBurn             = "tf_burn"
	ActionSetDenomMetadata = "set_denom_metadata"
	ActionUpdateParams     = "update_params"
)
