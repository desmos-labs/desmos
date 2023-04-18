package errors

import errors "cosmossdk.io/errors"

// RootCodespace is the codespace for all errors defined in this package
const RootCodespace = "desmos"

var (
	ErrInvalidURI = errors.Register(RootCodespace, 1, "invalid uri")
)
