package types

// Posts module event types
const (
	EventTypePostCreated         = "post_created"
	EventTypePostEdited          = "post_edited"
	EventTypePostReactionAdded   = "post_reaction_added"
	EventTypePostReactionRemoved = "post_reaction_removed"
	EventTypeAnsweredPoll        = "post_poll_answered"
	EventTypeRegisterReaction    = "reaction_registered"
	EventTypePostReported        = "post_reported"

	// Post attributes
	AttributeKeyPostID           = "post_id"
	AttributeKeyPostParentID     = "post_parent_id"
	AttributeKeyPostOwner        = "post_owner"
	AttributeKeyReportOwner      = "report_owner"
	AttributeKeyPostEditTime     = "post_edit_time"
	AttributeKeyPostCreationTime = "post_creation_time"

	// Poll attributes
	AttributeKeyPollAnswerer = "poll_answerer"

	// PostReaction attributes
	AttributeKeyPostReactionOwner = "reaction_user"
	AttributeKeyPostReactionValue = "reaction_value"
	AttributeKeyReactionShortCode = "reaction_shortcode"

	// Reaction attributes
	AttributeKeyReactionCreator  = "reaction_creator"
	AttributeKeyReactionSubSpace = "reaction_subspace"
)
