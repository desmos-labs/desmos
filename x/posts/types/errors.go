package types

import errors "cosmossdk.io/errors"

var (
	ErrInvalidGenesis = errors.Register(ModuleName, 1, "invalid genesis state")
	ErrInvalidPost    = errors.Register(ModuleName, 2, "invalid post")
)
