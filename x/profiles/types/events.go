package types

const (
	EventTypeProfileSaved   = "profile_saved"
	EventTypeProfileDeleted = "profile_deleted"

	// Relationships events
	EventTypeMonodirectionalRelationshipCreated = "monodirectional_relationship_created"
	EventTypeBidirectionalRelationshipRequested = "bidirectional_relationship_requested"
	EventTypeBidirectionalRelationshipAccepted  = "bidirectional_relationship_accepted"
	EventTypeBidirectionalRelationshipDenied    = "bidirectional_relationship_denied"
	EventTypeRelationshipsDeleted               = "relationships_deleted"

	// Profile attributes
	AttributeProfileDtag         = "profile_dtag"
	AttributeProfileCreator      = "profile_creator"
	AttributeProfileCreationTime = "profile_creation_time"

	// Relationships attributes
	AttributeRelationshipSender   = "relationship_sender"
	AttributeRelationshipReceiver = "relationship_receiver"
	AttributeRelationshipStatus   = "relationship_status"
)
