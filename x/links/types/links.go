package types

// NewLink returns a new Link object
func NewLink(sourceAddress string, destinationAddress string) Link {
	return Link{
		SourceAddress:      sourceAddress,
		DestinationAddress: destinationAddress,
	}
}
