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
	NewMsgSession    = types.NewMsgCreateSession
	NewLike          = types.NewLike
	NewPost          = types.NewPost
	NewSession       = types.NewSession
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	MsgCreatePost    = types.MsgCreatePost
	MsgEditPost      = types.MsgEditPost
	MsgLike          = types.MsgLike
	MsgUnlike        = types.MsgUnlike
	MsgCreateSession = types.MsgCreateSession
	Post             = types.Post
	Like             = types.Like
	Session          = types.Session
	QueryResPost     = types.QueryResPost
	QueryResLike     = types.QueryResLike
	QueryResSession  = types.QueryResSession
)
