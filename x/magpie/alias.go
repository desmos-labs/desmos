package magpie

import (
	"github.com/desmos-labs/desmos/x/magpie/internal/keeper"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	// Keeper methods
	NewKeeper  = keeper.NewKeeper
	NewHandler = keeper.NewHandler
	NewQuerier = keeper.NewQuerier

	// Codec
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	// Types
	NewSession = types.NewSession

	// Msgs
	NewMsgSession = types.NewMsgCreateSession
)

type (
	// Keeper
	Keeper = keeper.Keeper

	// Types
	SessionID = types.SessionID
	Session   = types.Session
	Sessions  = types.Sessions

	// Msgs
	MsgCreateSession = types.MsgCreateSession
)
