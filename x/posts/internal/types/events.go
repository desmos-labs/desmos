package types

// Magpie module event types
const (
	EventTypePostCreated         = "post_created"
	EventTypePostEdited          = "post_edited"
	EventTypeReactionAdded       = "post_reaction_added"
	EventTypePostReactionRemoved = "post_reaction_removed"

	// Post attributes
	AttributeKeyPostID       = "post_id"
	AttributeKeyPostParentID = "post_parent_id"
	AttributeKeyPostOwner    = "post_owner"
	AttributeKeyPostEditTime = "post_edit_time"

	// Reaction attributes
	AttributeKeyReactionOwner = "user"
	AttributeKeyReactionValue = "reaction"

	// Generic attributes
	AttributeKeyCreationTime = "creation_time"
)
