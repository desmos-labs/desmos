package types

// Posts module event types
const (
	EventTypeCreatePost           = "create_post"
	EventTypeEditPost             = "edit_post"
	EventTypeAddPostAttachment    = "add_post_attachment"
	EventTypeRemovePostAttachment = "add_post_attachment"
	EventTypeAnswerPoll           = "answer_poll"

	AttributeValueCategory   = ModuleName
	AttributeKeySubspaceID   = "subspace_id"
	AttributeKeyPostID       = "post_id"
	AttributeKeyAuthor       = "author"
	AttributeKeyCreationTime = "creation_date"
	AttributeKeyLastEditTime = "last_edit_date"
	AttributeKeyAttachmentID = "attachment_id"
	AttributeKeyPollID       = "poll_id"
)
