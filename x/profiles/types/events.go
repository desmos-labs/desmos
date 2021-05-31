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

	EventTypeLinkChainAccount   = "link_chain_account"
	EventTypeUnlinkChainAccount = "unlink_chain_account"

	AttributeChainLinkAccountTarget   = "chain_link_account_target"
	AttributeChainLinkAccountOwner    = "chain_link_account_owner"
	AttributeChainLinkSourceChainName = "chain_link_source_chain_name"
	AttributeChainLinkCreated         = "chain_link_created"

	// IBC events
	EventTypeTimeout                = "timeout"
	AttributeKeyAckSuccess          = "success"
	AttributeKeyAck                 = "acknowledgement"
	AttributeKeyAckError            = "error"
	EventTypeLinkChainAccountPacket = "link_chain_account_packet"
)
