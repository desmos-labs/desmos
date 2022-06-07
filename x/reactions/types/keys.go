package types

// DONTCOVER

const (
	ModuleName   = "reactions"
	RouterKey    = ModuleName
	StoreKey     = ModuleName
	QuerierRoute = ModuleName

	ActionAddReaction              = "add_reaction"
	ActionRemoveReaction           = "remove_reaction"
	ActionAddRegisteredReaction    = "add_registered_reaction"
	ActionRemoveRegisteredReaction = "remove_registered_reaction"
	ActionSetReactionParams        = "set_reaction_params"
)
