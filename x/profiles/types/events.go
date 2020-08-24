package types

const (
	EventTypeProfileSaved   = "profile_saved"
	EventTypeProfileDeleted = "profile_deleted"

	// Relationships events
	EventTypeMonoDirectionalRelationshipCreated = "relationship_created"
	EventTypeRelationshipsDeleted               = "relationships_deleted"

	// Profile attributes
	AttributeProfileDtag         = "profile_dtag"
	AttributeProfileCreator      = "profile_creator"
	AttributeProfileCreationTime = "profile_creation_time"

	// Relationships attributes
	AttributeRelationshipSender   = "relationship_sender"
	AttributeRelationshipReceiver = "relationship_receiver"
)
