package types

// Posts module event types
const (
	EventTypeCreatePost           = "create_post"
	EventTypeEditPost             = "edit_post"
	EventTypeDeletePost           = "delete_post"
	EventTypeAddPostAttachment    = "add_post_attachment"
	EventTypeRemovePostAttachment = "remove_post_attachment"
	EventTypeAnswerPoll           = "answer_poll"
	EventTypeTallyPoll            = "tally_poll"
	EventTypeChangePostOwner      = "change_post_owner"

	AttributeValueCategory   = ModuleName
	AttributeKeySubspaceID   = "subspace_id"
	AttributeKeySectionID    = "section_id"
	AttributeKeyPostID       = "post_id"
	AttributeKeyAuthor       = "author"
	AttributeKeyCreationTime = "creation_date"
	AttributeKeyLastEditTime = "last_edit_date"
	AttributeKeyAttachmentID = "attachment_id"
	AttributeKeyPollID       = "poll_id"
	AttributeKeyOwner        = "owner"
	AttributeKeyNewOwner     = "new_owner"
)
