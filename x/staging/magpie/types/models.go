package types

import (
	"strconv"
)

// NewSessionID returns a new SessionID value containing the given value
func NewSessionID(value uint64) SessionID {
	return SessionID{Value: value}
}

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

// ___________________________________________________________________________________________________________________

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
