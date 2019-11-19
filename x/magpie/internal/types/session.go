package types

import (
	"encoding/json"
	"strconv"

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
	SessionID     SessionID      `json:"id,string"`              // Id of the session
	Owner         sdk.AccAddress `json:"owner"`                  // Desmos owner of this session
	Created       int64          `json:"creation_time,string"`   // Block height at which the session has been created
	Expiry        int64          `json:"expiration_time,string"` // Block height at which the session will expire
	Namespace     string         `json:"namespace"`              // External chain identifier
	ExternalOwner string         `json:"external_owner"`         // External chain owner address
	PubKey        string         `json:"pub_key"`                // External chain owner public key
	Signature     string         `json:"signature"`              // Session signature
}

// NewSession return an empty Session
func NewSession() Session {
	return Session{}
}

// Equals returns true iff s and other contain the same data
func (s Session) Equals(other Session) bool {
	return s.SessionID == other.SessionID &&
		s.Owner.Equals(other.Owner) &&
		s.Created == other.Created &&
		s.Expiry == other.Expiry &&
		s.Namespace == other.Namespace &&
		s.ExternalOwner == other.ExternalOwner &&
		s.PubKey == other.PubKey &&
		s.Signature == other.Signature
}

// implement fmt.Stringer
func (s Session) String() string {
	bytes, err := json.Marshal(&s)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// ---------------
// --- Sessions
// ---------------

// Sessions represents a slice of Session objects
type Sessions []Session

// Equals returns true if and only if slice contains the
// same session objects as the other slice, false otherwise
func (slice Sessions) Equals(other Sessions) bool {
	if len(slice) != len(other) {
		return false
	}

	for index, s := range slice {
		if !s.Equals(other[index]) {
			return false
		}
	}

	return true
}
