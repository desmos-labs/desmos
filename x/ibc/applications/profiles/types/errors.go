package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrMaxProfilesChannels = sdkerrors.Register(ModuleName, 1, "max profiles channels")
	ErrInvalidVersion      = sdkerrors.Register(ModuleName, 2, "invalid ICS20 version")
)

const (
	ErrIBCTimeout         = "ibc connection timeout"
	ErrRequestExpired     = "oracle request expired"
	ErrRequestFailed      = "oracle request failed"
	ErrInvalidSignature   = "invalid signature"
	ErrInvalidAppUsername = "invalid application username"
)
