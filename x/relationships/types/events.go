package types

// DONTCOVER

const (
	EventTypeCreatedRelationship = "created_relationship"
	EventTypeDeletedRelationship = "deleted_relationship"
	EventTypeBlockedUser         = "blocked_user"
	EventTypeUnblockedUser       = "unblocked_user"

	AttributeRelationshipCreator      = "creator"
	AttributeRelationshipCounterparty = "counterparty"
	AttributeKeyUserBlockBlocker      = "blocker"
	AttributeKeyUserBlockBlocked      = "blocked"
)
