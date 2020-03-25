package profile

import (
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	DecodeStore         = simulation.DecodeStore
	RandomizedGenState  = simulation.RandomizedGenState
	WeightedOperations  = simulation.WeightedOperations
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	NewMsgCreateAccount = types.NewMsgCreateAccount
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewHandler          = keeper.NewHandler
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier

	ModuleCdc = types.ModuleCdc
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
