package types

// DONTCOVER

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

	EventTypeBlockUser   = "block_user"
	EventTypeUnblockUser = "unblock_user"

	AttributeKeyUserBlockBlocker = "blocker"
	AttributeKeyUserBlockBlocked = "blocked"
	AttributeKeyUserBlockReason  = "reason"
	AttributeKeySubspace         = "subspace"

	EventTypeLinkChainAccount   = "link_chain_account"
	EventTypeUnlinkChainAccount = "unlink_chain_account"

	AttributeChainLinkSourceAddress      = "chain_link_account_target"
	AttributeChainLinkDestinationAddress = "chain_link_account_owner"
	AttributeChainLinkSourceChainName    = "chain_link_source_chain_name"
	AttributeChainLinkCreationTime       = "chain_link_creation_time"

	AttributeKeyAckSuccess          = "success"
	EventTypeLinkChainAccountPacket = "link_chain_account_packet"

	EventTypePacket                  = "profiles_verification_packet"
	EventTypeTimeout                 = "timeout"
	EventTypesApplicationLinkCreated = "application_link_created"
	EventTypesApplicationLinkSaved   = "application_link_saved"
	EventTypeApplicationLinkDeleted  = "application_link_deleted"

	AttributeKeyUser                        = "user"
	AttributeKeyApplicationName             = "application_name"
	AttributeKeyApplicationUsername         = "application_username"
	AttributeKeyApplicationLinkCreationTime = "application_link_creation_time"
	AttributeKeyOracleID                    = "oracle_id"
	AttributeKeyClientID                    = "client_id"
	AttributeKeyRequestID                   = "request_id"
	AttributeKeyRequestKey                  = "request_key"
	AttributeKeyResolveStatus               = "resolve_status"
	AttributeKeyAck                         = "acknowledgement"
	AttributeKeyAckError                    = "error"
)
