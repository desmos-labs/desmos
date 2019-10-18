package magpie

import (
	"github.com/kwunyeung/desmos/x/magpie/internal/keeper"
	"github.com/kwunyeung/desmos/x/magpie/internal/types"
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
	NewLike    = types.NewLike
	NewPost    = types.NewPost
	NewSession = types.NewSession

	// Msgs
	NewMsgCreatePost = types.NewMsgCreatePost
	NewMsgEditPost   = types.NewMsgEditPost
	NewMsgLike       = types.NewMsgLike
	NewMsgUnlike     = types.NewMsgUnlike
	NewMsgSession    = types.NewMsgCreateSession
)

type (
	// Keeper
	Keeper = keeper.Keeper

	// Types
	Post    = types.Post
	Like    = types.Like
	Session = types.Session

	// Msgs
	MsgCreatePost    = types.MsgCreatePost
	MsgEditPost      = types.MsgEditPost
	MsgLike          = types.MsgLike
	MsgUnlike        = types.MsgUnlike
	MsgCreateSession = types.MsgCreateSession

	// Queries
	QueryResPost    = keeper.QueryResPost
	QueryResLike    = keeper.QueryResLike
	QueryResSession = keeper.QueryResSession
)
