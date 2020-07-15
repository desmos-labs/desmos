package types

const (
	ModuleName = "magpie"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreationSession = "create_session"
)

var (
	SessionLengthKey      = []byte("default_session_length")
	LastSessionIDStoreKey = []byte("last_session_id")
	SessionStorePrefix    = []byte("session")
)

// SessionStoreKey turns a session id to a key used to store a session into the sessions store
//nolint: interfacer
func SessionStoreKey(id SessionID) []byte {
	return append(SessionStorePrefix, []byte(id.String())...)
}
