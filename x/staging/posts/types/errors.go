package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	// ErrInvalidPostID is returned if we cannot parse a post id
	ErrInvalidPostID = sdkerrors.Register(ModuleName, 1, "invalid post id")

	// ErrDuplicatedPost is returned if the post has already been created before
	ErrDuplicatedPost = sdkerrors.Register(ModuleName, 2, "the provided post conflicts with another one with id:")

	// ErrPostNotFound is returned if the post doesn't exist
	ErrPostNotFound = sdkerrors.Register(ModuleName, 3, "post not found")

	// ErrCommentsNotAllowed is returned when comments are not allowed for a post
	ErrCommentsNotAllowed = sdkerrors.Register(ModuleName, 5, "comments not allowed")

	// ErrInvalidSubspace is returned if a post subspace is not valid
	ErrInvalidSubspace = sdkerrors.Register(ModuleName, 7, "invalid subspace")

	// ErrInvalidPost is returned if a post is not valid (message length exceeded, attributes length exceeded, empty fields)
	ErrInvalidPost = sdkerrors.Register(ModuleName, 8, "invalid post")

	// ErrInvalidReactionCode is returned if we cannot validate a reaction short code
	ErrInvalidReactionCode = sdkerrors.Register(ModuleName, 9,
		"invalid reaction shortcode (it must only contains a-z, 0-9, - and _ and must start and end with a ':')")

	// ErrInvalidReaction is returned if a registered reaction is not valid
	ErrInvalidReaction = sdkerrors.Register(ModuleName, 10, "invalid reaction")

	// ErrInvalidPostPoll is returned if a poll is not valid
	ErrInvalidPostPoll = sdkerrors.Register(ModuleName, 11, "invalid poll")

	// ErrPollNotFound is returned when the poll doesn't exist
	ErrPollNotFound = sdkerrors.Register(ModuleName, 14, "poll not found")

	// ErrInvalidPollAnswers is returned when the poll answers are not valid for multiple reasons specified in the error description
	ErrInvalidPollAnswers = sdkerrors.Register(ModuleName, 16, "invalid answers")

	// ErrPollEmptyAnswer is returned if a poll answer is empty
	ErrPollEmptyAnswer = sdkerrors.Register(ModuleName, 18, "answer text must be specified and cannot be empty")

	// ErrInvalidPostAttachment is returned if an attachment is invalid
	ErrInvalidPostAttachment = sdkerrors.Register(ModuleName, 19, "invalid post attachment")

	// ErrEmptyAttachmentTag is returned when an attachment tag is empty
	ErrEmptyAttachmentTag = sdkerrors.Register(ModuleName, 22, "empty attachment tag address")

	// ErrDuplicatedReaction is returned when a reaction has already been added to a post
	ErrDuplicatedReaction = sdkerrors.Register(ModuleName, 23, "reaction already exist")

	// ErrReactionNotFound is returned when a reaction doesn't exist
	ErrReactionNotFound = sdkerrors.Register(ModuleName, 24, "reaction not found")

	// ErrDuplicatedReport is returned when a post has already been reported
	ErrDuplicatedReport = sdkerrors.Register(ModuleName, 25, "report already created")

	// ErrInvalidReport is returned if a report is not valid (invalid type or message)
	ErrInvalidReport = sdkerrors.Register(ModuleName, 26, "invalid report")
)
