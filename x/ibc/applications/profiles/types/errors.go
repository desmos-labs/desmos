package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrMaxProfilesChannels = sdkerrors.Register(ModuleName, 1, "max profiles channels")
	ErrInvalidVersion      = sdkerrors.Register(ModuleName, 2, "invalid ICS20 version")
)
