package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ------------------
// --- Session id
// ------------------

// SessionID represents a unique session id
type SessionID uint64

// Valid returns true if and only if this id can be considered
// valid to be stored on the chain
func (id SessionID) Valid() bool {
	return id != 0
}

// Next returns the next id to this one
func (id SessionID) Next() SessionID {
	return id + 1
}

// String implements fmt.Stringer
func (id SessionID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// ParseSessionID take the given string value and parses it returning a SessionID.
// If the given value is not valid, returns an error instead
func ParseSessionID(value string) (SessionID, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return SessionID(0), err
	}

	return SessionID(intVal), err
}

// ------------------
// --- Session
// ------------------

// Session is a struct of a user session
type Session struct {
	SessionID     SessionID      `json:"id"`             // Id of the session
	Owner         sdk.AccAddress `json:"owner"`          // Desmos owner of this session
	Created       time.Time      `json:"created"`        // Creation time
	Expiry        time.Time      `json:"expiry"`         // Expiration time
	Namespace     string         `json:"namespace"`      // Bech32 HRP of the external_owner field
	ExternalOwner string         `json:"external_owner"` // External chain owner address
	PubKey        string         `json:"pub_key"`        // External chain owner public key
	Signature     string         `json:"signature"`      // Session signature
}

// NewSession return an empty Session
func NewSession() Session {
	return Session{}
}

// implement fmt.Stringer
func (s Session) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Created: %s
Expiry: %s
Namespace: %s
External Owner: %s
Pubkey: %s
Signature: %s`, s.Owner, s.Created, s.Expiry, s.Namespace, s.ExternalOwner, s.PubKey, s.Signature))
}
