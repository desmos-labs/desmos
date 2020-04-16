package v040

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName           = "magpie"
	DefaultSessionLength = 240
)

// GenesisState represents the genesis state for the magpie module
type GenesisState struct {
	DefaultSessionLength int64     `json:"default_session_length"`
	Sessions             []Session `json:"sessions"`
}

// NewGenesisState returns a new genesis state from the given data
func NewGenesisState(defaultSessionLength int64, session []Session) GenesisState {
	return GenesisState{DefaultSessionLength: defaultSessionLength, Sessions: session}
}

// SessionID represents a unique session id
type SessionID uint64

// Session is a struct of a user session
type Session struct {
	SessionID     SessionID      `json:"id,string"`              // Id of the session
	Owner         sdk.AccAddress `json:"owner"`                  // Desmos owner of this session
	Created       int64          `json:"creation_time,string"`   // Block height at which the session has been created
	Expiry        int64          `json:"expiration_time,string"` // Block height at which the session will expire
	Namespace     string         `json:"namespace"`              // External chain identifier
	ExternalOwner string         `json:"external_owner"`         // External chain owner address
	PubKey        string         `json:"pub_key"`                // External chain owner public key
	Signature     string         `json:"signature"`              // Session signature
}
