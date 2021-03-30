package types

// NewLink returns a new Link object
func NewLink(source string, destination string) Link {
	return Link{
		Source:      source,
		Destination: destination,
	}
}
