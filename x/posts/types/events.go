package types

// Posts module event types
const (
	EventTypeCreatedPost                = "created_post"
	EventTypeEditedPost                 = "edited_post"
	EventTypeDeletedPost                = "deleted_post"
	EventTypeAddedPostAttachment        = "added_post_attachment"
	EventTypeRemovedPostAttachment      = "removed_post_attachment"
	EventTypeAnsweredPoll               = "answered_poll"
	EventTypeTalliedPoll                = "tallied_poll"
	EventTypeMovedPost                  = "moved_post"
	EventTypeRequestedPostOwnerTransfer = "requested_post_owner_transfer"
	EventTypeCanceledPostOwnerTransfer  = "canceled_post_owner_transfer"
	EventTypeAcceptedPostOwnerTransfer  = "accepted_post_owner_transfer"
	EventTypeRefusedPostOwnerTransfer   = "refused_post_owner_transfer"

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
