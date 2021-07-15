package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	// ErrInvalidPostID is returned if we cannot parse a post id
	ErrInvalidPostID = sdkerrors.Register(ModuleName, 1, "invalid post id")

	// ErrPostAlreadyCreated is returned if the post has already been created before
	ErrPostAlreadyCreated = sdkerrors.Register(ModuleName, 2, "the provided post conflicts with another one with id:")

	// ErrPostNotFound is returned if the post doesn't exist
	ErrPostNotFound = sdkerrors.Register(ModuleName, 3, "post not found")

	// ErrInvalidEmptyFields is returned when both message and attachments or poll are empty
	ErrInvalidEmptyFields = sdkerrors.Register(ModuleName, 4, "post message, attachments or poll are required and cannot be all blank or empty")

	// ErrCommentsNotAllowed is returned when comments are not allowed for a post
	ErrCommentsNotAllowed = sdkerrors.Register(ModuleName, 5, "comments not allowed")

	// ErrInvalidEditDate is returned when a post's edit date is set before the creation date
	ErrInvalidEditDate = sdkerrors.Register(ModuleName, 6, "edit date cannot be before creation date")

	// ErrInvalidSubspace is returned if a post subspace is not valid
	ErrInvalidSubspace = sdkerrors.Register(ModuleName, 7, "invalid subspace")

	// ErrMessageLengthExceeded is returned if the post message exceed the max length possible
	ErrMessageLengthExceeded = sdkerrors.Register(ModuleName, 8, "message max length exceeded")

	// ErrMaxAdditionalAttributesNumberExceeded is returned if the additional attributes are more than the max possible
	ErrMaxAdditionalAttributesNumberExceeded = sdkerrors.Register(ModuleName, 9, "additional attributes max number exceeded")

	// ErrAdditionalAttributeValLenExceeded is returned if the additional attribute value length exceed the max length possible
	ErrAdditionalAttributeValLenExceeded = sdkerrors.Register(ModuleName, 10, "additional attribute value length exceeded")

	// ErrInvalidReactionCode is returned if we cannot validate a reaction short code
	ErrInvalidReactionCode = sdkerrors.Register(ModuleName, 11,
		"invalid reaction shortcode (it must only contains a-z, 0-9, - and _ and must start and end with a ':')")

	// ErrReactionCodeAlreadyExist is returned if a reaction shortcode already exist
	ErrReactionCodeAlreadyExist = sdkerrors.Register(ModuleName, 12, "reaction shortcode already exists")

	// ErrReactionInvalidValue is returned if a reaction value is not valid
	ErrReactionInvalidValue = sdkerrors.Register(ModuleName, 13, "reaction value should be a valid URI")

	// ErrPollEmptyQuestion is returned when a poll question is empty or blank
	ErrPollEmptyQuestion = sdkerrors.Register(ModuleName, 14, "missing poll question")

	// ErrPollEndDate is returned when the poll end date is invalid
	ErrPollEndDate = sdkerrors.Register(ModuleName, 15, "invalid poll end date")

	// ErrPollNotFound is returned when the poll doesn't exist
	ErrPollNotFound = sdkerrors.Register(ModuleName, 16, "poll not found")

	// ErrPollClosed is returned if the poll has already been closed
	ErrPollClosed = sdkerrors.Register(ModuleName, 17, "poll closed")

	// ErrPollInvalidAnswers is returned when the poll answers are not valid for multiple reasons specified in the error description
	ErrPollInvalidAnswers = sdkerrors.Register(ModuleName, 18, "invalid answers")

	// ErrPollInvalidAnswersMinNumber is returned when the number of the possible poll answers is not reached
	ErrPollInvalidAnswersMinNumber = sdkerrors.Register(ModuleName, 19, "poll answers must be at least two")

	// ErrPollEmptyAnswer is returned if a poll answer is empty
	ErrPollEmptyAnswer = sdkerrors.Register(ModuleName, 20, "answer text must be specified and cannot be empty")

	// ErrPollUnregisteredAnswer is returned if a provided answer for a poll is not one of the provided
	ErrPollUnregisteredAnswer = sdkerrors.Register(ModuleName, 21, "unregistered answer")

	// ErrInvalidAttachmentURI is returned when an attachment URI is not valid
	ErrInvalidAttachmentURI = sdkerrors.Register(ModuleName, 22, "invalid URI provided")

	// ErrEmptyAttachmentMimeType is returned if the attachment type is empty
	ErrEmptyAttachmentMimeType = sdkerrors.Register(ModuleName, 23, "mime type must be specified and cannot be empty")

	// ErrEmptyAttachmentTag is returned when an attachment tag is empty
	ErrEmptyAttachmentTag = sdkerrors.Register(ModuleName, 24, "empty attachment tag address")

	// ErrReactionAlreadyAdded is returned when a reaction has already been added to a post
	ErrReactionAlreadyAdded = sdkerrors.Register(ModuleName, 26, "reaction already added")

	// ErrReactionNotFound is returned when a reaction doesn't exist
	ErrReactionNotFound = sdkerrors.Register(ModuleName, 27, "reaction not found")

	// ErrReportAlreadyCreated is returned when a post has already been reported
	ErrReportAlreadyCreated = sdkerrors.Register(ModuleName, 28, "report already created")

	// ErrReportNotFound is returned when a report doesn't exist
	ErrReportNotFound = sdkerrors.Register(ModuleName, 29, "report not found")

	// ErrReportEmptyType is returned if the report's type is empty
	ErrReportEmptyType = sdkerrors.Register(ModuleName, 30, "report type should not be empty")

	// ErrReportEmptyMessage is returned if the report's message is empty
	ErrReportEmptyMessage = sdkerrors.Register(ModuleName, 31, "report message cannot be empty")
)
