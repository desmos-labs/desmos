package types

import "cosmossdk.io/errors"

// DONTCOVER

var (
	ErrInvalidGenesis = errors.Register(ModuleName, 1, "invalid genesis state")
)
