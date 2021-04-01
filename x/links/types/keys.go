package types

const (
	ModuleName = "links"
	StoreKey   = ModuleName
	RouterKey  = ModuleName

	ActionIBCAccountConnection = "ibc_account_connection"

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

	LinksStorePrefix = []byte("link")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
