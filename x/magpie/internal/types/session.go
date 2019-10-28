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

// SessionId represents a unique session id
type SessionId uint64

func (id SessionId) Valid() bool {
	return id != 0
}

func (id SessionId) Next() SessionId {
	return id + 1
}

func (id SessionId) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func ParseSessionId(value string) (SessionId, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return SessionId(0), err
	}

	return SessionId(intVal), err
}

// ------------------
// --- Session
// ------------------

// Session is a struct of a user session
type Session struct {
	SessionID     SessionId      `json:"id"`
	Owner         sdk.AccAddress `json:"owner"`
	Created       time.Time      `json:"created"`
	Expiry        time.Time      `json:"expiry"`
	Namespace     string         `json:"namespace"`
	ExternalOwner string         `json:"external_owner"`
	Pubkey        string         `json:"pubkey"`
	Signature     string         `json:"signature"`
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
Signature: %s`, s.Owner, s.Created, s.Expiry, s.Namespace, s.ExternalOwner, s.Pubkey, s.Signature))
}
