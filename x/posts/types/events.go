package types

// Posts module event types
const (
	EventTypeCreatePost               = "create_post"
	EventTypeEditPost                 = "edit_post"
	EventTypeDeletePost               = "delete_post"
	EventTypeAddPostAttachment        = "add_post_attachment"
	EventTypeRemovePostAttachment     = "remove_post_attachment"
	EventTypeAnswerPoll               = "answer_poll"
	EventTypeTallyPoll                = "tally_poll"
	EventTypeMovePost                 = "move_post"
	EventTypeRequestPostOwnerTransfer = "request_post_owner_transfer"
	EventTypeCancelPostOwnerTransfer  = "cancel_post_owner_transfer"
	EventTypeAcceptPostOwnerTransfer  = "accept_post_owner_transfer"
	EventTypeRefusePostOwnerTransfer  = "refuse_post_owner_transfer"

	AttributeKeySubspaceID    = "subspace_id"
	AttributeKeySectionID     = "section_id"
	AttributeKeyPostID        = "post_id"
	AttributeKeyAuthor        = "author"
	AttributeKeyCreationTime  = "creation_date"
	AttributeKeyLastEditTime  = "last_edit_date"
	AttributeKeyAttachmentID  = "attachment_id"
	AttributeKeyPollID        = "poll_id"
	AttributeKeyNewSubspaceID = "new_subspace_id"
	AttributeKeyNewPostID     = "new_post_id"
	AttributeKeySender        = "sender"
	AttributeKeyReceiver      = "receiver"
)
