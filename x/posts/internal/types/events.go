package types

// Magpie module event types
const (
	EventTypePostCreated         = "post_created"
	EventTypePostEdited          = "post_edited"
	EventTypeReactionAdded       = "post_reaction_added"
	EventTypePostReactionRemoved = "post_reaction_removed"
	EventTypeAnsweredPoll        = "post_poll_answered"
	EventTypeClosePoll           = "close_poll"

	// Post attributes
	AttributeKeyPostID       = "post_id"
	AttributeKeyPostParentID = "post_parent_id"
	AttributeKeyPostOwner    = "post_owner"
	AttributeKeyPostEditTime = "post_edit_time"

	//Poll attributes
	AttributeKeyPollAnswerer = "poll_answerer"

	// Reaction attributes
	AttributeKeyReactionOwner = "user"
	AttributeKeyReactionValue = "reaction"

	// Generic attributes
	AttributeKeyCreationTime = "creation_time"
)
