package types

const (
	ModuleName   = "ibcprofiles"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName

	ActionIBCAccountConnection = "ibc_account_connection"
	ActionIBCAccountLink       = "ibc_account_link"

	// IBC keys
	Version = "ibcprofiles-1"
	PortID  = "ibcprofiles"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = []byte("port")
)
