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

	// ApplicationLinkPrefix defines the store prefix to be used when storing connections
	ApplicationLinkPrefix         = []byte("application_link")
	ApplicationLinkClientIDPrefix = []byte("client_id")
)

func ConnectionKey(link *ApplicationLink) []byte {
	return append(ApplicationLinkPrefix, []byte(link.User+link.Application.Name+link.Application.Username)...)
}

func UserApplicationsLinksPrefix(user string) []byte {
	return append(ApplicationLinkPrefix, []byte(user)...)
}

func ApplicationLinkClientIDKey(clientID string) []byte {
	return append(ApplicationLinkClientIDPrefix, []byte(clientID)...)
}
