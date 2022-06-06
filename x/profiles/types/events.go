package types

// DONTCOVER

const (
	EventTypeProfileSaved              = "save_profile"
	EventTypeProfileDeleted            = "delete_profile"
	EventTypeDTagTransferRequest       = "create_dtag_transfer_request"
	EventTypeDTagTransferAccept        = "accept_dtag_transfer_request"
	EventTypeDTagTransferRefuse        = "refuse_dtag_transfer_request"
	EventTypeDTagTransferCancel        = "cancel_dtag_transfer_request"
	EventTypeLinkChainAccount          = "link_chain_account"
	EventTypeUnlinkChainAccount        = "unlink_chain_account"
	EventTypeSetDefaultExternalAddress = "set_default_external_address"
	EventTypeLinkChainAccountPacket    = "link_chain_account_packet"
	EventTypePacket                    = "receive_profiles_verification_packet"
	EventTypeTimeout                   = "timeout"
	EventTypesApplicationLinkCreated   = "link_application"
	EventTypeApplicationLinkDeleted    = "unlink_application"
	EventTypesApplicationLinkSaved     = "application_link_saved"

	AttributeKeyProfileDTag                 = "profile_dtag"
	AttributeKeyProfileCreator              = "profile_creator"
	AttributeKeyProfileCreationTime         = "profile_creation_time"
	AttributeKeyRequestReceiver             = "request_receiver"
	AttributeKeyRequestSender               = "request_sender"
	AttributeKeyDTagToTrade                 = "dtag_to_trade"
	AttributeKeyNewDTag                     = "new_dtag"
	AttributeKeyChainLinkExternalAddress    = "chain_link_external_address"
	AttributeKeyChainLinkOwner              = "chain_link_owner"
	AttributeKeyChainLinkChainName          = "chain_link_chain_name"
	AttributeKeyChainLinkCreationTime       = "chain_link_creation_time"
	AttributeKeyAckSuccess                  = "success"
	AttributeKeyUser                        = "user"
	AttributeKeyApplicationName             = "application_name"
	AttributeKeyApplicationUsername         = "application_username"
	AttributeKeyApplicationLinkCreationTime = "application_link_creation_time"
	AttributeKeyOracleID                    = "oracle_id"
	AttributeKeyClientID                    = "client_id"
	AttributeKeyRequestID                   = "request_id"
	AttributeKeyResolveStatus               = "resolve_status"
	AttributeKeyAck                         = "acknowledgement"
	AttributeKeyAckError                    = "error"

	AttributeValueCategory = ModuleName
)
