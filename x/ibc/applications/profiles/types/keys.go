package types

const (
	// ModuleName defines the IBC profiles name
	ModuleName = "ibc-profiles"

	// StoreKey is the store key string for IBC profiles
	StoreKey = ModuleName

	// RouterKey is the message route for IBC profiles
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC profiles
	QuerierRoute = ModuleName

	// PortID is the default port id that profiles module binds to
	PortID = "transfer"

	EventTypeConnectProfile = "connect_profile"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = []byte{0x01}
)
