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
	IsValidPostID              = common.IsValidPostID
	IsValidReactionCode        = common.IsValidReactionCode
	GetEmojiByShortCodeOrValue = common.GetEmojiByShortCodeOrValue
	NewAttachment              = common.NewAttachment
	NewAttachments             = common.NewAttachments
	ParseAnswerID              = polls.ParseAnswerID
	NewPollAnswer              = polls.NewPollAnswer
	NewPollAnswers             = polls.NewPollAnswers
	NewPollData                = polls.NewPollData
	ArePollDataEquals          = polls.ArePollDataEquals
	NewUserAnswer              = polls.NewUserAnswer
	NewUserAnswers             = polls.NewUserAnswers
	NewPostReaction            = reactions.NewPostReaction
	NewPostReactions           = reactions.NewPostReactions
	NewReaction                = reactions.NewReaction
	IsEmoji                    = reactions.IsEmoji
	NewReactions               = reactions.NewReactions
	NewMsgAnswerPoll           = msgs.NewMsgAnswerPoll
	NewMsgCreatePost           = msgs.NewMsgCreatePost
	NewMsgEditPost             = msgs.NewMsgEditPost
	NewMsgRegisterReaction     = msgs.NewMsgRegisterReaction
	RegisterMessagesCodec      = msgs.RegisterMessagesCodec
	NewMsgAddPostReaction      = msgs.NewMsgAddPostReaction
	NewMsgRemovePostReaction   = msgs.NewMsgRemovePostReaction
	ParsePostID                = models.ParsePostID
	NewPost                    = models.NewPost
	NewPostResponse            = models.NewPostResponse
	PostStoreKey               = models.PostStoreKey
	PostIndexedIDStoreKey      = models.PostIndexedIDStoreKey
	PostCommentsStoreKey       = models.PostCommentsStoreKey
	PostReactionsStoreKey      = models.PostReactionsStoreKey
	ReactionsStoreKey          = models.ReactionsStoreKey
	PollAnswersStoreKey        = models.PollAnswersStoreKey
	RegisterModelsCodec        = models.RegisterModelsCodec

	// variable aliases
	ModuleAddress            = common.ModuleAddress
	PostStorePrefix          = common.PostStorePrefix
	PostIndexedIDStorePrefix = common.PostIndexedIDStorePrefix
	PostTotalNumberPrefix    = common.PostTotalNumberPrefix
	PostCommentsStorePrefix  = common.PostCommentsStorePrefix
	PostReactionsStorePrefix = common.PostReactionsStorePrefix
	ReactionsStorePrefix     = common.ReactionsStorePrefix
	PollAnswersStorePrefix   = common.PollAnswersStorePrefix
	MsgsCodec                = msgs.MsgsCodec
	ModelsCdc                = models.ModelsCdc
)

type (
	PostID                   = models.PostID
	PostIDs                  = models.PostIDs
	Post                     = models.Post
	Posts                    = models.Posts
	PostQueryResponse        = models.PostQueryResponse
	PollAnswersQueryResponse = models.PollAnswersQueryResponse
	OptionalData             = common.OptionalData
	KeyValue                 = common.KeyValue
	Attachment               = common.Attachment
	Attachments              = common.Attachments
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
	MsgAnswerPoll            = msgs.MsgAnswerPoll
	MsgCreatePost            = msgs.MsgCreatePost
	MsgEditPost              = msgs.MsgEditPost
	MsgRegisterReaction      = msgs.MsgRegisterReaction
	MsgAddPostReaction       = msgs.MsgAddPostReaction
	MsgRemovePostReaction    = msgs.MsgRemovePostReaction
)
