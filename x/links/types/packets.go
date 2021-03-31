package types

// ValidateBasic is used for validating the packet
func (p IBCLinkPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IBCLinkPacketData) GetBytes() ([]byte, error) {
	var modulePacket LinksPacketData

	modulePacket.Packet = &LinksPacketData_IbcLinkPacket{&p}

	return modulePacket.Marshal()
}
