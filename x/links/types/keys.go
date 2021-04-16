package types

const (
	ModuleName = "links"
	StoreKey   = ModuleName
	RouterKey  = ModuleName

	ActionIBCAccountConnection = "ibc_account_connection"
	ActionIBCAccountLink       = "ibc_account_link"

	// Queries
	QuerierRoute = ModuleName
	QueryLink    = "link"

	// IBC keys
	Version = "links-1"
	PortID  = "links"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = []byte("port")

	LinkStorePrefix = []byte("link")
)

func LinkStoreKey(address string) []byte {
	return append(LinkStorePrefix, []byte(address)...)
}
