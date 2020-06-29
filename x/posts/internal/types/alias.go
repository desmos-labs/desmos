package types

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/posts/internal/types/common"
	"github.com/desmos-labs/desmos/x/posts/internal/types/polls"
	"github.com/desmos-labs/desmos/x/posts/internal/types/reactions"
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
	GetEmojiByShortCodeOrValue = common.GetEmojiByShortCodeOrValue
	NewPostMedia               = common.NewPostMedia
	ValidateURI                = common.ValidateURI
	NewPostMedias              = common.NewPostMedias
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

	// variable aliases
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
)

type (
	OptionalData  = common.OptionalData
	KeyValue      = common.KeyValue
	PostMedia     = common.PostMedia
	PostMedias    = common.PostMedias
	AnswerID      = polls.AnswerID
	PollAnswer    = polls.PollAnswer
	PollAnswers   = polls.PollAnswers
	PollData      = polls.PollData
	UserAnswer    = polls.UserAnswer
	UserAnswers   = polls.UserAnswers
	PostReaction  = reactions.PostReaction
	PostReactions = reactions.PostReactions
	Reaction      = reactions.Reaction
	Reactions     = reactions.Reactions
)
