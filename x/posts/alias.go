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
	NewKeeper           = keeper.NewKeeper
	NewHandler          = keeper.NewHandler
	NewQuerier          = keeper.NewQuerier
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
	ParsePostID         = types.ParsePostID
	NewReaction         = types.NewReaction
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	DecodeStore         = simulation.DecodeStore
	RandomizedGenState  = simulation.RandomizedGenState
	WeightedOperations  = simulation.WeightedOperations
)

type (
	Keeper                = keeper.Keeper
	PostID                = types.PostID
	PostIDs               = types.PostIDs
	Post                  = types.Post
	Posts                 = types.Posts
	PostMedia             = types.PostMedia
	PostMedias            = types.PostMedias
	Reaction              = types.Reaction
	Reactions             = types.Reactions
	PostQueryResponse     = types.PostQueryResponse
	GenesisState          = types.GenesisState
	MsgCreatePost         = types.MsgCreatePost
	MsgEditPost           = types.MsgEditPost
	MsgAddPostReaction    = types.MsgAddPostReaction
	MsgRemovePostReaction = types.MsgRemovePostReaction
	QueryPostsParams      = types.QueryPostsParams
)
