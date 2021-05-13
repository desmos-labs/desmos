package types

const (
	// IBC events
	EventTypeTimeout                    = "timeout"
	AttributeKeyAckSuccess              = "success"
	AttributeKeyAck                     = "acknowledgement"
	AttributeKeyAckError                = "error"
	EventTypeIBCAccountConnectionPacket = "ibc_account_connection_packet"
	EventTypeIBCAccountLinkPacket       = "ibc_account_link_packet"
)
