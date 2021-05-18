package types

const (
	EventTypeProfileSaved        = "profile_saved"
	EventTypeProfileDeleted      = "profile_deleted"
	EventTypeDTagTransferRequest = "dtag_transfer_request"
	EventTypeDTagTransferAccept  = "dtag_transfer_accept"
	EventTypeDTagTransferRefuse  = "dtag_transfer_refuse"
	EventTypeDTagTransferCancel  = "dtag_transfer_cancel"

	AttributeProfileDTag         = "profile_dtag"
	AttributeProfileCreator      = "profile_creator"
	AttributeProfileCreationTime = "profile_creation_time"

	AttributeRequestReceiver = "request_receiver"
	AttributeRequestSender   = "request_sender"
	AttributeDTagToTrade     = "dtag_to_trade"
	AttributeNewDTag         = "new_dtag"

	EventTypeRelationshipCreated  = "relationship_created"
	EventTypeRelationshipsDeleted = "relationships_deleted"

	AttributeRelationshipSender   = "sender"
	AttributeRelationshipReceiver = "receiver"
	AttributeRelationshipSubspace = "subspace"

	EventTypeBlockUser   = "block_user"
	EventTypeUnblockUser = "unblock_user"

	AttributeKeyUserBlockBlocker = "blocker"
	AttributeKeyUserBlockBlocked = "blocked"
	AttributeKeyUserBlockReason  = "reason"
	AttributeKeySubspace         = "subspace"
)
