package types

import errors "cosmossdk.io/errors"

var (
	// ErrPermissionDenied is returned if a user cannot perform a specific action inside a subspace
	ErrPermissionDenied = errors.Register(ModuleName, 1, "permissions denied for user")
	ErrInvalidGenesis   = errors.Register(ModuleName, 2, "invalid genesis state")
)
