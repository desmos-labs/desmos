package types

const (
	EventTypeProfileSaved            = "profile_saved"
	EventTypeProfileDeleted          = "profile_deleted"
	EventTypeDTagTransferRequest     = "dtag_transfer_request"
	EventTypeDTagTransferReqAccepted = "accepted_transfer_request"

	// Profile attributes
	AttributeProfileDtag         = "profile_dtag"
	AttributeProfileCreator      = "profile_creator"
	AttributeProfileCreationTime = "profile_creation_time"

	// DTag trade attributes
	AttributeCurrentOwner  = "current_owner"
	AttributeReceivingUser = "receiving_user"
	AttributeNewDTag       = "new_DTag"
)
