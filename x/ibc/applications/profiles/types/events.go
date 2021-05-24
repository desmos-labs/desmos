package types

const (
	EventTypePacket          = "profiles_verification_packet"
	EventTypeTimeout         = "timeout"
	EventTypeConnectProfile  = "connect_profile"
	EventTypeConnectionSaved = "connection_saved"

	AttributeKeyUser                = "user"
	AttributeKeyApplicationName     = "application_name"
	AttributeKeyApplicationUsername = "application_username"
	AttributeKeyOracleID            = "oracle_id"
	AttributeKeyClientID            = "client_id"
	AttributeKeyRequestID           = "request_id"
	AttributeKeyRequestKey          = "request_key"
	AttributeKeyResolveStatus       = "resolve_status"
	AttributeKeyAckSuccess          = "success"
	AttributeKeyAck                 = "acknowledgement"
	AttributeKeyAckError            = "error"
)
