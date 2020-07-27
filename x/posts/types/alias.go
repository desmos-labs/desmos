package types

// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/posts/types/models"
	"github.com/desmos-labs/desmos/x/posts/types/models/common"
	"github.com/desmos-labs/desmos/x/posts/types/models/polls"
	"github.com/desmos-labs/desmos/x/posts/types/models/reactions"
	"github.com/desmos-labs/desmos/x/posts/types/msgs"
)

const (
	ModuleName               = common.ModuleName
	RouterKey                = common.RouterKey
	StoreKey                 = common.StoreKey
	ActionCreatePost         = common.ActionCreatePost
	ActionEditPost           = common.ActionEditPost
	ActionAnswerPoll         = common.ActionAnswerPoll
	ActionAddPostReaction    = common.ActionAddPostReaction
	ActionRemovePostReaction = common.ActionRemovePostReaction
	ActionRegisterReaction   = common.ActionRegisterReaction
	QuerierRoute             = common.QuerierRoute
	QueryPost                = common.QueryPost
	QueryPosts               = common.QueryPosts
	QueryPollAnswers         = common.QueryPollAnswers
	QueryRegisteredReactions = common.QueryRegisteredReactions
	QueryParams              = common.QueryParams
	PostSortByCreationDate   = common.PostSortByCreationDate
	PostSortByID             = common.PostSortByID
	PostSortOrderAscending   = common.PostSortOrderAscending
	PostSortOrderDescending  = common.PostSortOrderDescending
)

var (
	// functions aliases
	RegisterModelsCodec        = models.RegisterModelsCodec
	ParsePostID                = models.ParsePostID
	NewPost                    = models.NewPost
	NewPostResponse            = models.NewPostResponse
	PostStoreKey               = models.PostStoreKey
	PostIDStoreKey             = models.PostIndexedIDStoreKey
	PostCommentsStoreKey       = models.PostCommentsStoreKey
	PostReactionsStoreKey      = models.PostReactionsStoreKey
	ReactionsStoreKey          = models.ReactionsStoreKey
	PollAnswersStoreKey        = models.PollAnswersStoreKey
	IsValidPostID              = common.IsValidPostID
	IsValidSubspace            = common.IsValidSubspace
	IsValidReactionCode        = common.IsValidReactionCode
	GetEmojiByShortCodeOrValue = common.GetEmojiByShortCodeOrValue
	NewAttachment              = common.NewAttachment
	NewAttachments             = common.NewAttachments
	NewPollData                = polls.NewPollData
	ArePollDataEquals          = polls.ArePollDataEquals
	NewUserAnswer              = polls.NewUserAnswer
	NewUserAnswers             = polls.NewUserAnswers
	ParseAnswerID              = polls.ParseAnswerID
	NewPollAnswer              = polls.NewPollAnswer
	NewPollAnswers             = polls.NewPollAnswers
	NewReaction                = reactions.NewReaction
	IsEmoji                    = reactions.IsEmoji
	NewReactions               = reactions.NewReactions
	NewPostReaction            = reactions.NewPostReaction
	NewPostReactions           = reactions.NewPostReactions
	NewMsgAddPostReaction      = msgs.NewMsgAddPostReaction
	NewMsgRemovePostReaction   = msgs.NewMsgRemovePostReaction
	NewMsgAnswerPoll           = msgs.NewMsgAnswerPoll
	NewMsgCreatePost           = msgs.NewMsgCreatePost
	NewMsgEditPost             = msgs.NewMsgEditPost
	NewMsgRegisterReaction     = msgs.NewMsgRegisterReaction
	RegisterMessagesCodec      = msgs.RegisterMessagesCodec

	// variable aliases
	ModuleAddress            = common.ModuleAddress
	PostStorePrefix          = common.PostStorePrefix
	PostIDStorePrefix        = common.PostIndexedIDStorePrefix
	PostTotalNumberPrefix    = common.PostTotalNumberPrefix
	PostCommentsStorePrefix  = common.PostCommentsStorePrefix
	PostReactionsStorePrefix = common.PostReactionsStorePrefix
	ReactionsStorePrefix     = common.ReactionsStorePrefix
	PollAnswersStorePrefix   = common.PollAnswersStorePrefix
	MsgsCodec                = msgs.MsgsCodec
	ModelsCdc                = models.ModelsCdc
)

type (
	OptionalData             = common.OptionalData
	KeyValue                 = common.KeyValue
	Attachment               = common.Attachment
	Attachments              = common.Attachments
	PollData                 = polls.PollData
	UserAnswer               = polls.UserAnswer
	UserAnswers              = polls.UserAnswers
	AnswerID                 = polls.AnswerID
	PollAnswer               = polls.PollAnswer
	PollAnswers              = polls.PollAnswers
	Reaction                 = reactions.Reaction
	Reactions                = reactions.Reactions
	PostReaction             = reactions.PostReaction
	PostReactions            = reactions.PostReactions
	MsgAddPostReaction       = msgs.MsgAddPostReaction
	MsgRemovePostReaction    = msgs.MsgRemovePostReaction
	MsgAnswerPoll            = msgs.MsgAnswerPoll
	MsgCreatePost            = msgs.MsgCreatePost
	MsgEditPost              = msgs.MsgEditPost
	MsgRegisterReaction      = msgs.MsgRegisterReaction
	PostID                   = models.PostID
	PostIDs                  = models.PostIDs
	Post                     = models.Post
	Posts                    = models.Posts
	PostQueryResponse        = models.PostQueryResponse
	PollAnswersQueryResponse = models.PollAnswersQueryResponse
)
