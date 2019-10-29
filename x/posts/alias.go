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
	NewMsgLike       = types.NewMsgLike
	NewMsgUnlike     = types.NewMsgUnlike
)

type (
	// Keeper
	Keeper = keeper.Keeper

	// Types
	Post   = types.Post
	PostID = types.PostID
	Like   = types.Like
	LikeID = types.LikeID

	// Msgs
	MsgCreatePost = types.MsgCreatePost
	MsgEditPost   = types.MsgEditPost
	MsgLike       = types.MsgLike
	MsgUnlike     = types.MsgUnlike
)
