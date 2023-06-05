package types

// DONTCOVER

import (
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

var (
	// PermissionManageSubspaceTokens allows users to manage subspace tokens
	PermissionManageSubspaceTokens = subspacestypes.RegisterPermission("manage custom subspace tokens")
)
