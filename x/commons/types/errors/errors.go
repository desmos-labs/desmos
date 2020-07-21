package errors

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// RootCodespace is the codespace for all errors defined in this package
const RootCodespace = "desmos"

var (
	ErrInvalidURI = sdkerrors.Register(RootCodespace, 1, "invalid uri")
)
