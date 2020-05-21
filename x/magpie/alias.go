package magpie

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/magpie/internal/keeper"
	"github.com/desmos-labs/desmos/x/magpie/internal/simulation"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
)

const (
	QuerySessions             = keeper.QuerySessions
	OpWeightMsgCreatePost     = simulation.OpWeightMsgCreatePost
	EventTypeCreateSession    = types.EventTypeCreateSession
	AttributeKeySessionID     = types.AttributeKeySessionID
	AttributeKeyNamespace     = types.AttributeKeyNamespace
	AttributeKeyExternalOwner = types.AttributeKeyExternalOwner
	AttributeKeyExpiry        = types.AttributeKeyExpiry
	AttributeValueCategory    = types.AttributeValueCategory
	ModuleName                = types.ModuleName
	RouterKey                 = types.RouterKey
	StoreKey                  = types.StoreKey
	ActionCreationSession     = types.ActionCreationSession
)

var (
	// functions aliases
	NewHandler               = keeper.NewHandler
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	RandomizedGenState       = simulation.RandomizedGenState
	DecodeStore              = simulation.DecodeStore
	RandomSessionData        = simulation.RandomSessionData
	WeightedOperations       = simulation.WeightedOperations
	SimulateMsgCreateSession = simulation.SimulateMsgCreateSession
	ParseSessionID           = types.ParseSessionID
	NewSession               = types.NewSession
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	RegisterCodec            = types.RegisterCodec
	NewMsgCreateSession      = types.NewMsgCreateSession
	SessionStoreKey          = types.SessionStoreKey

	// variable aliases
	ModuleCdc             = types.ModuleCdc
	SessionLengthKey      = types.SessionLengthKey
	LastSessionIDStoreKey = types.LastSessionIDStoreKey
	SessionStorePrefix    = types.SessionStorePrefix
	RandomNamespaces      = simulation.RandomNamespaces
)

type (
	Keeper           = keeper.Keeper
	SessionData      = simulation.SessionData
	SessionID        = types.SessionID
	Session          = types.Session
	Sessions         = types.Sessions
	GenesisState     = types.GenesisState
	MsgCreateSession = types.MsgCreateSession
)
