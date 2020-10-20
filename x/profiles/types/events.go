package types

const (
	EventTypeProfileSaved        = "profile_saved"
	EventTypeProfileDeleted      = "profile_deleted"
	EventTypeDTagTransferRequest = "dtag_transfer_request"
	EventTypeDTagTransferAccept  = "dtag_transfer_accept"
	EventTypeDTagTransferRefuse  = "dtag_transfer_refuse"
	EventTypeDTagTransferCancel  = "dtag_transfer_cancel"

	// Profile attributes
	AttributeProfileDtag         = "profile_dtag"
	AttributeProfileCreator      = "profile_creator"
	AttributeProfileCreationTime = "profile_creation_time"

	// DTag trade attributes
	AttributeRequestReceiver = "request_receiver"
	AttributeRequestSender   = "request_sender"
	AttributeDTagToTrade     = "dtag_to_trade"
	AttributeNewDTag         = "new_dtag"
)
