package types

const (
	// ModuleName defines the IBC profiles name
	ModuleName = "ibcprofiles"

	// Version defines the current version the IBC profiles module supports
	// TODO: Using our new version for profile packet (new ics?)
	Version = "ics20-1"

	// StoreKey is the store key string for IBC profiles
	StoreKey = ModuleName

	// RouterKey is the message route for IBC profiles
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC profiles
	QuerierRoute = ModuleName

	// PortID is the default port id that profiles module binds to
	PortID = ModuleName
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = []byte{0x01}

	// ConnectionPrefix defines the store prefix to be used when storing connections
	ConnectionPrefix = []byte("connection")

	ConnectionClientIDPrefix = []byte("client_id")
)

func ConnectionKey(connection *Connection) []byte {
	return append(ConnectionPrefix, []byte(connection.User+connection.Application.Name+connection.Application.Username)...)
}

func UserConnectionsPrefix(user string) []byte {
	return append(ConnectionPrefix, []byte(user)...)
}

func ConnectionClientIDKey(clientID string) []byte {
	return append(ConnectionClientIDPrefix, []byte(clientID)...)
}
