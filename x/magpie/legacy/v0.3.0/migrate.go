package v0_3_0

import (
	v020magpie "github.com/desmos-labs/desmos/x/magpie/legacy/v0.2.0"
)

// Migrate accepts exported genesis state from v0.2.0 and migrates it to v0.3.0
// genesis state. This migration adds the default session length to the new genesis state.
func Migrate(oldGenState v020magpie.GenesisState) GenesisState {
	return GenesisState{
		Sessions:             migrateSessions(oldGenState.Sessions),
		DefaultSessionLength: DefaultSessionLength,
	}
}

func migrateSessions(oldSessions []v020magpie.Session) []Session {
	sessions := make([]Session, len(oldSessions))
	for index, session := range oldSessions {
		sessions[index] = Session{
			SessionID:     SessionID(session.SessionID),
			Owner:         session.Owner,
			Created:       session.Created,
			Expiry:        session.Expiry,
			Namespace:     session.Namespace,
			ExternalOwner: session.ExternalOwner,
			PubKey:        session.PubKey,
			Signature:     session.Signature,
		}
	}
	return sessions
}
