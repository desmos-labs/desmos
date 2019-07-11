package magpie

import (
	"github.com/kwunyeung/desmos/x/magpie/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgCreatePost = types.NewMsgCreatePost
	NewMsgEditPost   = types.NewMsgEditPost
	NewMsgLike       = types.NewMsgLike
	NewMsgUnlike     = types.NewMsgUnlike
	NewLike          = types.NewLike
	NewPost          = types.NewPost
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	MsgCreatePost = types.MsgCreatePost
	MsgEditPost   = types.MsgEditPost
	MsgLike       = types.MsgLike
	MsgUnlike     = types.MsgUnlike
	Post          = types.Post
	Like          = types.Like
	QueryResPost  = types.QueryResPost
	QueryResLike  = types.QueryResLike
)
