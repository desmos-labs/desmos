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
	ModuleName                = types.ModuleName
	RouterKey                 = types.RouterKey
	StoreKey                  = types.StoreKey
	ActionCreationSession     = types.ActionCreationSession
	EventTypeCreateSession    = types.EventTypeCreateSession
	AttributeKeySessionID     = types.AttributeKeySessionID
	AttributeKeyNamespace     = types.AttributeKeyNamespace
	AttributeKeyExternalOwner = types.AttributeKeyExternalOwner
	AttributeKeyExpiry        = types.AttributeKeyExpiry
	AttributeValueCategory    = types.AttributeValueCategory
)

var (
	// functions aliases
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	NewHandler               = keeper.NewHandler
	RandomizedGenState       = simulation.RandomizedGenState
	DecodeStore              = simulation.DecodeStore
	RandomSessionData        = simulation.RandomSessionData
	WeightedOperations       = simulation.WeightedOperations
	SimulateMsgCreateSession = simulation.SimulateMsgCreateSession
	RegisterCodec            = types.RegisterCodec
	NewMsgCreateSession      = types.NewMsgCreateSession
	SessionStoreKey          = types.SessionStoreKey
	ParseSessionID           = types.ParseSessionID
	NewSession               = types.NewSession
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis

	// variable aliases
	RandomNamespaces      = simulation.RandomNamespaces
	ModuleCdc             = types.ModuleCdc
	SessionLengthKey      = types.SessionLengthKey
	LastSessionIDStoreKey = types.LastSessionIDStoreKey
	SessionStorePrefix    = types.SessionStorePrefix
)

type (
	SessionData      = simulation.SessionData
	MsgCreateSession = types.MsgCreateSession
	SessionID        = types.SessionID
	Session          = types.Session
	Sessions         = types.Sessions
	GenesisState     = types.GenesisState
	Keeper           = keeper.Keeper
)
