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
	NewKeeper                = keeper.NewKeeper
	NewHandler               = keeper.NewHandler
	NewQuerier               = keeper.NewQuerier
	ModuleCdc                = types.ModuleCdc
	RegisterCodec            = types.RegisterCodec
	ParsePostID              = types.ParsePostID
	NewReaction              = types.NewReaction
	NewPost                  = types.NewPost
	NewPostMedia             = types.NewPostMedia
	NewPostMedias            = types.NewPostMedias
	NewPollData              = types.NewPollData
	NewPollAnswer            = types.NewPollAnswer
	NewPollAnswers           = types.NewPollAnswers
	NewUserAnswer            = types.NewUserAnswer
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	DecodeStore              = simulation.DecodeStore
	RandomizedGenState       = simulation.RandomizedGenState
	WeightedOperations       = simulation.WeightedOperations
	NewMsgCreatePost         = types.NewMsgCreatePost
	NewMsgEditPost           = types.NewMsgEditPost
	NewMsgAddPostReaction    = types.NewMsgAddPostReaction
	NewMsgRemovePostReaction = types.NewMsgRemovePostReaction
)

type (
	Keeper                = keeper.Keeper
	PostID                = types.PostID
	PostIDs               = types.PostIDs
	Post                  = types.Post
	Posts                 = types.Posts
	PostMedia             = types.PostMedia
	PostMedias            = types.PostMedias
	AnswerID              = types.AnswerID
	UserAnswer            = types.UserAnswer
	UserAnswers           = types.UserAnswers
	Reaction              = types.Reaction
	Reactions             = types.Reactions
	PostQueryResponse     = types.PostQueryResponse
	GenesisState          = types.GenesisState
	MsgCreatePost         = types.MsgCreatePost
	MsgEditPost           = types.MsgEditPost
	MsgAddPostReaction    = types.MsgAddPostReaction
	MsgRemovePostReaction = types.MsgRemovePostReaction
	MsgAnswerPoll         = types.MsgAnswerPoll
	QueryPostsParams      = types.QueryPostsParams
)
