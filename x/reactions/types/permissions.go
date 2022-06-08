package types

// DONTCOVER

import (
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

var (
	PermissionsReact                    = subspacestypes.Permission(0b111111111111)
	PermissionManageRegisteredReactions = subspacestypes.Permission(0b111111111111)
	PermissionManageReactionParams      = subspacestypes.Permission(0b111111111111)
)
