package posts

import (
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
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
	ParsePostID         = types.ParsePostID
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// Simulation
	DecodeStore        = simulation.DecodeStore
	RandomizedGenState = simulation.RandomizedGenState
	WeightedOperations = simulation.WeightedOperations
)

type (
	// Keeper
	Keeper = keeper.Keeper

	// Types
	PostID       = types.PostID
	PostIDs      = types.PostIDs
	Post         = types.Post
	Posts        = types.Posts
	PostMedia    = types.PostMedia
	PostMedias   = types.PostMedias
	Reaction     = types.Reaction
	Reactions    = types.Reactions
	GenesisState = types.GenesisState

	// Msgs
	MsgCreatePost         = types.MsgCreatePost
	MsgEditPost           = types.MsgEditPost
	MsgAddPostReaction    = types.MsgAddPostReaction
	MsgRemovePostReaction = types.MsgRemovePostReaction

	// Queries
	QueryPostsParams = types.QueryPostsParams
)
