package types

const (
	ModuleName = "links"
	StoreKey   = ModuleName
	RouterKey  = ModuleName

	ActionIBCLink = "ibc_link"

	// Queries
	QuerierRoute = ModuleName
	QueryLink    = "link"

	// IBC keys
	MemStoreKey = "mem_capability"
	Version     = "links-1"
	PortID      = "links"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("links-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
