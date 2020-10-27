package types

import (
	"strconv"
)

// ------------------
// --- Session id
// ------------------

// Valid returns true if and only if this id can be considered
// valid to be stored on the chain
func (id SessionID) Valid() bool {
	return id.Value != 0
}

// Next returns the next id to this one
func (id SessionID) Next() SessionID {
	return SessionID{Value: id.Value + 1}
}

// ParseSessionID take the given string value and parses it returning a SessionID.
// If the given value is not valid, returns an error instead
func ParseSessionID(value string) (SessionID, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return SessionID{Value: 0}, err
	}

	return SessionID{Value: intVal}, err
}

// ------------------
// --- Session
// ------------------

// NewSession return a new session containing the given parameters
func NewSession(
	id SessionID, owner string, created, expiry uint64,
	namespace, externalOwner, pubKey, signature string,
) Session {
	return Session{
		SessionId:      id,
		Owner:          owner,
		CreationTime:   created,
		ExpirationTime: expiry,
		Namespace:      namespace,
		ExternalOwner:  externalOwner,
		PublicKey:      pubKey,
		Signature:      signature,
	}
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
		if !s.Equal(other[index]) {
			return false
		}
	}

	return true
}
