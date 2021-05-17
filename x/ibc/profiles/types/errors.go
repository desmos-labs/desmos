package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ibc/profiles module sentinel errors
var (
	ErrInvalidVersion         = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrMaxIBCProfilesChannels = sdkerrors.Register(ModuleName, 1, "max ibcprofiles channels")
)
