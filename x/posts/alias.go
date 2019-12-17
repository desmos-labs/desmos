package posts

import (
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
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
	NewReaction         = types.NewReaction
	NewPost             = types.NewPost
	ParsePostID         = types.ParsePostID
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// Msgs
	NewMsgCreatePost         = types.NewMsgCreatePost
	NewMsgEditPost           = types.NewMsgEditPost
	NewMsgAddPostReaction    = types.NewMsgAddPostReaction
	NewMsgRemovePostReaction = types.NewMsgRemovePostReaction
)

type (
	// Keeper
	Keeper = keeper.Keeper

	// Types
	PostID       = types.PostID
	PostIDs      = types.PostIDs
	Post         = types.Post
	Posts        = types.Posts
	Reaction     = types.Reaction
	Reactions    = types.Reactions
	GenesisState = types.GenesisState

	// Msgs
	MsgCreatePost         = types.MsgCreatePost
	MsgEditPost           = types.MsgEditPost
	MsgAddPostReaction    = types.MsgAddPostReaction
	MsgRemovePostReaction = types.MsgRemovePostReaction
)
