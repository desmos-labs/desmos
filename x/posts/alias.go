package posts

import (
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
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
	NewReaction         = types.NewReaction
	NewTextPost         = types.NewTextPost
	NewMediaPost        = types.NewMediaPost
	ParsePostID         = types.ParsePostID
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// Msgs
	NewMsgCreateTextPost     = types.NewMsgCreateTextPost
	NewMsgCreateMediaPost    = types.NewMsgCreateMediaPost
	NewMsgEditPost           = types.NewMsgEditPost
	NewMsgAddPostReaction    = types.NewMsgAddPostReaction
	NewMsgRemovePostReaction = types.NewMsgRemovePostReaction

	// Queries
	NewQueryPostsParams = types.NewQueryPostsParams
)

type (
	// Keeper
	Keeper = keeper.Keeper

	// Types
	PostID       = types.PostID
	PostIDs      = types.PostIDs
	Post         = types.Post
	Posts        = types.Posts
	TextPost     = types.TextPost
	TextPosts    = types.TextPosts
	MediaPost    = types.MediaPost
	MediaPosts   = types.MediaPosts
	PostMedia    = types.PostMedia
	PostMedias   = types.PostMedias
	Reaction     = types.Reaction
	Reactions    = types.Reactions
	GenesisState = types.GenesisState

	// Msgs
	MsgCreatePost         = types.MsgCreatePost
	MsgCreateTextPost     = types.MsgCreateTextPost
	MsgCreateMediaPost    = types.MsgCreateMediaPost
	MsgEditPost           = types.MsgEditPost
	MsgAddPostReaction    = types.MsgAddPostReaction
	MsgRemovePostReaction = types.MsgRemovePostReaction

	// Queries
	QueryPostsParams = types.QueryPostsParams
)
