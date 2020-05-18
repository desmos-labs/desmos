package posts

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/simulation"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/common"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/polls"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/reactions"
	"github.com/desmos-labs/desmos/x/posts/internal/types/msgs"
)

const (
	OpWeightMsgCreatePost           = simulation.OpWeightMsgCreatePost
	OpWeightMsgEditPost             = simulation.OpWeightMsgEditPost
	OpWeightMsgAddReaction          = simulation.OpWeightMsgAddReaction
	OpWeightMsgRemoveReaction       = simulation.OpWeightMsgRemoveReaction
	OpWeightMsgAnswerPoll           = simulation.OpWeightMsgAnswerPoll
	OpWeightMsgRegisterReaction     = simulation.OpWeightMsgRegisterReaction
	DefaultGasValue                 = simulation.DefaultGasValue
	EventTypePostCreated            = types.EventTypePostCreated
	EventTypePostEdited             = types.EventTypePostEdited
	EventTypePostReactionAdded      = types.EventTypePostReactionAdded
	EventTypePostReactionRemoved    = types.EventTypePostReactionRemoved
	EventTypeAnsweredPoll           = types.EventTypeAnsweredPoll
	EventTypeClosePoll              = types.EventTypeClosePoll
	EventTypeRegisterReaction       = types.EventTypeRegisterReaction
	AttributeKeyPostID              = types.AttributeKeyPostID
	AttributeKeyPostParentID        = types.AttributeKeyPostParentID
	AttributeKeyPostOwner           = types.AttributeKeyPostOwner
	AttributeKeyPostEditTime        = types.AttributeKeyPostEditTime
	AttributeKeyPostCreationTime    = types.AttributeKeyPostCreationTime
	AttributeKeyPollAnswerer        = types.AttributeKeyPollAnswerer
	AttributeKeyPostReactionOwner   = types.AttributeKeyPostReactionOwner
	AttributeKeyPostReactionValue   = types.AttributeKeyPostReactionValue
	AttributeKeyReactionShortCode   = types.AttributeKeyReactionShortCode
	AttributeKeyReactionCreator     = types.AttributeKeyReactionCreator
	AttributeKeyReactionSubSpace    = types.AttributeKeyReactionSubSpace
	ModuleName                      = common.ModuleName
	RouterKey                       = common.RouterKey
	StoreKey                        = common.StoreKey
	MaxPostMessageLength            = common.MaxPostMessageLength
	MaxOptionalDataFieldsNumber     = common.MaxOptionalDataFieldsNumber
	MaxOptionalDataFieldValueLength = common.MaxOptionalDataFieldValueLength
	ActionCreatePost                = common.ActionCreatePost
	ActionEditPost                  = common.ActionEditPost
	ActionAnswerPoll                = common.ActionAnswerPoll
	ActionAddPostReaction           = common.ActionAddPostReaction
	ActionRemovePostReaction        = common.ActionRemovePostReaction
	ActionRegisterReaction          = common.ActionRegisterReaction
	QuerierRoute                    = common.QuerierRoute
	QueryPost                       = common.QueryPost
	QueryPosts                      = common.QueryPosts
	QueryPollAnswers                = common.QueryPollAnswers
	QueryRegisteredReactions        = common.QueryRegisteredReactions
	PostSortByCreationDate          = common.PostSortByCreationDate
	PostSortByID                    = common.PostSortByID
	PostSortOrderAscending          = common.PostSortOrderAscending
	PostSortOrderDescending         = common.PostSortOrderDescending
)

var (
	// functions aliases
	NewPostResponse                  = models.NewPostResponse
	RegisterModelsCodec              = models.RegisterModelsCodec
	PostStoreKey                     = models.PostStoreKey
	PostCommentsStoreKey             = models.PostCommentsStoreKey
	PostReactionsStoreKey            = models.PostReactionsStoreKey
	ReactionsStoreKey                = models.ReactionsStoreKey
	PollAnswersStoreKey              = models.PollAnswersStoreKey
	ComputeID                        = models.ComputeID
	ParsePostID                      = models.ParsePostID
	NewPost                          = models.NewPost
	NewPostMedia                     = common.NewPostMedia
	ValidateURI                      = common.ValidateURI
	NewPostMedias                    = common.NewPostMedias
	ParseAnswerID                    = polls.ParseAnswerID
	NewPollAnswer                    = polls.NewPollAnswer
	NewPollAnswers                   = polls.NewPollAnswers
	NewPollData                      = polls.NewPollData
	ArePollDataEquals                = polls.ArePollDataEquals
	NewUserAnswer                    = polls.NewUserAnswer
	NewUserAnswers                   = polls.NewUserAnswers
	NewPostReaction                  = reactions.NewPostReaction
	NewPostReactions                 = reactions.NewPostReactions
	NewReaction                      = reactions.NewReaction
	IsEmoji                          = reactions.IsEmoji
	NewReactions                     = reactions.NewReactions
	RegisterMessagesCodec            = msgs.RegisterMessagesCodec
	NewMsgCreatePost                 = msgs.NewMsgCreatePost
	NewMsgEditPost                   = msgs.NewMsgEditPost
	NewMsgAnswerPoll                 = msgs.NewMsgAnswerPoll
	NewMsgAddPostReaction            = msgs.NewMsgAddPostReaction
	NewMsgRemovePostReaction         = msgs.NewMsgRemovePostReaction
	NewMsgRegisterReaction           = msgs.NewMsgRegisterReaction
	NewHandler                       = keeper.NewHandler
	RegisterInvariants               = keeper.RegisterInvariants
	AllInvariants                    = keeper.AllInvariants
	ValidPostsInvariant              = keeper.ValidPostsInvariant
	ValidCommentsDateInvariant       = keeper.ValidCommentsDateInvariant
	ValidPostForReactionsInvariant   = keeper.ValidPostForReactionsInvariant
	ValidPollForPollAnswersInvariant = keeper.ValidPollForPollAnswersInvariant
	NewKeeper                        = keeper.NewKeeper
	NewQuerier                       = keeper.NewQuerier
	DecodeStore                      = simulation.DecodeStore
	RandomizedGenState               = simulation.RandomizedGenState
	WeightedOperations               = simulation.WeightedOperations
	SimulateMsgAnswerToPoll          = simulation.SimulateMsgAnswerToPoll
	SimulateMsgCreatePost            = simulation.SimulateMsgCreatePost
	SimulateMsgEditPost              = simulation.SimulateMsgEditPost
	SimulateMsgAddPostReaction       = simulation.SimulateMsgAddPostReaction
	SimulateMsgRemovePostReaction    = simulation.SimulateMsgRemovePostReaction
	SimulateMsgRegisterReaction      = simulation.SimulateMsgRegisterReaction
	RandomPost                       = simulation.RandomPost
	RandomPostData                   = simulation.RandomPostData
	RandomPostReactionData           = simulation.RandomPostReactionData
	RandomPostReactionValue          = simulation.RandomPostReactionValue
	RandomPostID                     = simulation.RandomPostID
	RandomMessage                    = simulation.RandomMessage
	RandomSubspace                   = simulation.RandomSubspace
	RandomHashtag                    = simulation.RandomHashtag
	RandomMedias                     = simulation.RandomMedias
	RandomPollData                   = simulation.RandomPollData
	GetAccount                       = simulation.GetAccount
	RandomReactionValue              = simulation.RandomReactionValue
	RandomReactionShortCode          = simulation.RandomReactionShortCode
	RandomReactionData               = simulation.RandomReactionData
	RegisteredReactionsData          = simulation.RegisteredReactionsData
	RegisterCodec                    = types.RegisterCodec
	NewGenesisState                  = types.NewGenesisState
	DefaultGenesisState              = types.DefaultGenesisState
	ValidateGenesis                  = types.ValidateGenesis
	DefaultQueryPostsParams          = types.DefaultQueryPostsParams

	// variable aliases
	RandomMimeTypes          = simulation.RandomMimeTypes
	RandomHosts              = simulation.RandomHosts
	ModuleCdc                = types.ModuleCdc
	ModelsCdc                = models.ModelsCdc
	Sha256RegEx              = common.Sha256RegEx
	HashtagRegEx             = common.HashtagRegEx
	ShortCodeRegEx           = common.ShortCodeRegEx
	URIRegEx                 = common.URIRegEx
	ModuleAddress            = common.ModuleAddress
	PostStorePrefix          = common.PostStorePrefix
	PostCommentsStorePrefix  = common.PostCommentsStorePrefix
	PostReactionsStorePrefix = common.PostReactionsStorePrefix
	ReactionsStorePrefix     = common.ReactionsStorePrefix
	PollAnswersStorePrefix   = common.PollAnswersStorePrefix
	MsgsCodec                = msgs.MsgsCodec
)

type (
	PostData                 = simulation.PostData
	PostReactionData         = simulation.PostReactionData
	ReactionData             = simulation.ReactionData
	GenesisState             = types.GenesisState
	QueryPostsParams         = types.QueryPostsParams
	PostQueryResponse        = models.PostQueryResponse
	PollAnswersQueryResponse = models.PollAnswersQueryResponse
	PostID                   = models.PostID
	PostIDs                  = models.PostIDs
	Post                     = models.Post
	Posts                    = models.Posts
	PostMedia                = common.PostMedia
	PostMedias               = common.PostMedias
	OptionalData             = common.OptionalData
	KeyValue                 = common.KeyValue
	AnswerID                 = polls.AnswerID
	PollAnswer               = polls.PollAnswer
	PollAnswers              = polls.PollAnswers
	PollData                 = polls.PollData
	UserAnswer               = polls.UserAnswer
	UserAnswers              = polls.UserAnswers
	PostReaction             = reactions.PostReaction
	PostReactions            = reactions.PostReactions
	Reaction                 = reactions.Reaction
	Reactions                = reactions.Reactions
	MsgCreatePost            = msgs.MsgCreatePost
	MsgEditPost              = msgs.MsgEditPost
	MsgAnswerPoll            = msgs.MsgAnswerPoll
	MsgAddPostReaction       = msgs.MsgAddPostReaction
	MsgRemovePostReaction    = msgs.MsgRemovePostReaction
	MsgRegisterReaction      = msgs.MsgRegisterReaction
	Keeper                   = keeper.Keeper
)
