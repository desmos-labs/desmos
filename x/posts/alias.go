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
	NewLike = types.NewLike
	NewPost = types.NewPost

	// Msgs
	NewMsgCreatePost = types.NewMsgCreatePost
	NewMsgEditPost   = types.NewMsgEditPost
	NewMsgLikePost   = types.NewMsgLikePost
	NewMsgUnlikePost = types.NewMsgUnlikePost
)

type (
	// Keeper
	Keeper = keeper.Keeper

	// Types
	PostID = types.PostID
	Post   = types.Post
	Like   = types.Like
	Likes  = types.Likes

	// Msgs
	MsgCreatePost = types.MsgCreatePost
	MsgEditPost   = types.MsgEditPost
	MsgLikePost   = types.MsgLikePost
	MsgUnlikePost = types.MsgUnlikePost
)
