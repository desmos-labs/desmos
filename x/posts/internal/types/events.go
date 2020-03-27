package types

// Magpie module event types
const (
	EventTypePostCreated         = "post_created"
	EventTypePostEdited          = "post_edited"
	EventTypePostReactionAdded   = "post_reaction_added"
	EventTypePostReactionRemoved = "post_reaction_removed"
	EventTypeAnsweredPoll        = "post_poll_answered"
	EventTypeClosePoll           = "post_poll_closed"
	EventTypeRegisterReaction    = "reaction_registered"

	// Post attributes
	AttributeKeyPostID       = "post_id"
	AttributeKeyPostParentID = "post_parent_id"
	AttributeKeyPostOwner    = "post_owner"
	AttributeKeyPostEditTime = "post_edit_time"

	// Poll attributes
	AttributeKeyPollAnswerer = "poll_answerer"

	// PostReaction attributes
	AttributeKeyPostReactionOwner = "user"
	AttributeKeyPostReactionValue = "reaction"

	// Reaction attributes
	AttributeKeyReactionCreator   = "creator"
	AttributeKeyReactionShortCode = "short_code"
	AttributeKeyReactionSubSpace  = "subspace"

	// Generic attributes
	AttributeKeyCreationTime = "creation_time"
)
