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
	NewKeeper           = keeper.NewKeeper
	NewHandler          = keeper.NewHandler
	NewQuerier          = keeper.NewQuerier
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
)

type (
	Keeper           = keeper.Keeper
	SessionID        = types.SessionID
	Session          = types.Session
	Sessions         = types.Sessions
	GenesisState     = types.GenesisState
	MsgCreateSession = types.MsgCreateSession
)
