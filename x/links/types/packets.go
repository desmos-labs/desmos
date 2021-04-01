package types

// ValidateBasic is used for validating the packet
func (p IBCAccountConnectionPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IBCAccountConnectionPacketData) GetBytes() ([]byte, error) {
	var modulePacket LinksPacketData

	modulePacket.Packet = &LinksPacketData_IbcAccountConnectionPacket{&p}

	return modulePacket.Marshal()
}
