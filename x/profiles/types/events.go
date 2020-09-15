package types

const (
	EventTypeProfileSaved        = "profile_saved"
	EventTypeProfileDeleted      = "profile_deleted"
	EventTypeDTagTransferRequest = "dtag_transfer_request"
	EventTypeDTagTransferAccept  = "dtag_transfer_accept"

	// Profile attributes
	AttributeProfileDtag         = "profile_dtag"
	AttributeProfileCreator      = "profile_creator"
	AttributeProfileCreationTime = "profile_creation_time"

	// DTag trade attributes
	AttributeCurrentOwner  = "current_owner"
	AttributeReceivingUser = "receiving_user"
	AttributeDTagToTrade   = "dtag_to_trade"
	AttributeNewDTag       = "new_dtag"
)
