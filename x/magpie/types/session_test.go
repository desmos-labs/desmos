package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/magpie/types"
	"github.com/stretchr/testify/require"
)

// ------------------
// --- Session id
// ------------------

func TestSessionID_Valid(t *testing.T) {
	tests := []struct {
		id            types.SessionID
		shouldBeValid bool
	}{
		{
			id:            types.SessionID{Value: 0},
			shouldBeValid: false,
		},
		{
			id:            types.SessionID{Value: 54},
			shouldBeValid: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%s id valid: %t", test.id, test.shouldBeValid), func(t *testing.T) {
			require.Equal(t, test.shouldBeValid, test.id.Valid())
		})
	}
}

func TestSessionID_Next(t *testing.T) {
	tests := []struct {
		id     types.SessionID
		nextID types.SessionID
	}{
		{
			id:     types.SessionID{Value: 0},
			nextID: types.SessionID{Value: 1},
		},
		{
			id:     types.SessionID{Value: 234},
			nextID: types.SessionID{Value: 235},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%s next is %s", test.id, test.nextID), func(t *testing.T) {
			require.Equal(t, test.nextID, test.id.Next())
		})
	}
}

func TestSessionID_String(t *testing.T) {
	tests := []struct {
		id        types.SessionID
		expString string
	}{
		{
			id:        types.SessionID{Value: 0},
			expString: "0",
		},
		{
			id:        types.SessionID{Value: 123123},
			expString: "123123",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("String representation of %s is %s", test.id, test.expString), func(t *testing.T) {
			require.Equal(t, test.expString, test.id.String())
		})
	}
}

func TestParseSessionID(t *testing.T) {
	tests := []struct {
		name   string
		string string
		expID  types.SessionID
		expErr bool
	}{
		{
			name:   "ID 0 is parsed correctly",
			string: "0",
			expID:  types.SessionID{Value: 0},
		},
		{
			name:   "Negative ID returns error",
			string: "-1",
			expID:  types.SessionID{Value: 0},
			expErr: true,
		},
		{
			name:   "Positive ID is parsed correctly",
			string: "54624",
			expID:  types.SessionID{Value: 54624},
		},
		{
			name:   "Invalid string returns error",
			string: "string",
			expID:  types.SessionID{Value: 0},
			expErr: true,
		},
		{
			name:   "Too big number returns error",
			string: "100000000000000000000000000000000000000000000000000000000000",
			expID:  types.SessionID{Value: 0},
			expErr: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			id, err := types.ParseSessionID(test.string)

			require.Equal(t, test.expID, id)
			require.Equal(t, test.expErr, err != nil)
		})
	}
}

// ------------------
// --- Session
// ------------------

func TestSession_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.Session
		second    types.Session
		expEquals bool
	}{
		{
			name: "Different session ID",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 54},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expEquals: false,
		},
		{
			name: "Different owner",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1cl4kjuqz8zrgw9h32v5hrhzulmlf0jcmjaw67c",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expEquals: false,
		},
		{
			name: "Different created",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   65,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expEquals: false,
		},
		{
			name: "Different expiry",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 10,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expEquals: false,
		},
		{
			name: "Different namespace",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expEquals: false,
		},
		{
			name: "Different external owner",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1zck2zu0thlxzg4hh98y3y9rhsd3mju9rdfj2tn",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expEquals: false,
		},
		{
			name: "Different public key",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1zck2zu0thlxzg4hh98y3y9rhsd3mju9rdfj2tn",
				PublicKey:      "cosmospub1addwnpepq2n40mtr2allsj2zd52g0pyhkct5rhj0e4f9n6xwd5jshswaw90av2y92xt",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expEquals: false,
		},
		{
			name: "Different signature",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "=",
			},
			expEquals: false,
		},
		{
			name: "Same data",
			first: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			second: types.Session{
				SessionId:      types.SessionID{Value: 0},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   0,
				ExpirationTime: 0,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equal(test.second))
		})
	}
}

func TestSession_String(t *testing.T) {
	tests := []struct {
		name      string
		session   types.Session
		expString string
	}{
		{
			name:      "Empty session stringed correctly",
			session:   types.Session{},
			expString: `{"id":"0","owner":"","creation_time":"0","expiration_time":"0","namespace":"","external_owner":"","pub_key":"","signature":""}`,
		},
		{
			name: "Complete session stringed correctly",
			session: types.Session{
				SessionId:      types.SessionID{Value: 15},
				Owner:          "cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj",
				CreationTime:   35,
				ExpirationTime: 55,
				Namespace:      "cosmos",
				ExternalOwner:  "cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q",
				PublicKey:      "cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg",
				Signature:      "YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc=",
			},
			expString: `{"id":"15","owner":"cosmos1htw7gatueyhl9at24m62wl3j3kar3q3ass2pkj","creation_time":"35","expiration_time":"55","namespace":"cosmos","external_owner":"cosmos1l5q6tzjpse5p50zg3spef83cd79drahx58f69q","pub_key":"cosmospub1addwnpepqgxp4eye98gy70lwa58uk29rjpdwn033el34wzt2x74pkkqpp5re2gcyypg","signature":"YXNAJTQzMjUzNTRzMzRnMjQyNDR3NTI0emYyYmg0c2EzMjRyaGIuNHM1Z2I1NDFzMWc="}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expString, test.session.String())
		})
	}
}

// ---------------
// --- Sessions
// ---------------

func TestSessions_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.Sessions
		second    types.Sessions
		expEquals bool
	}{
		{
			name:      "Empty lists",
			first:     types.Sessions{},
			second:    types.Sessions{},
			expEquals: true,
		},
		{
			name:      "Empty and non empty list",
			first:     types.Sessions{},
			second:    types.Sessions{types.Session{SessionId: types.SessionID{Value: 10}}},
			expEquals: false,
		},
		{
			name:      "Non empty list, different items",
			first:     types.Sessions{types.Session{SessionId: types.SessionID{Value: 19}}},
			second:    types.Sessions{types.Session{SessionId: types.SessionID{Value: 14}}},
			expEquals: false,
		},
		{
			name: "Non empty lists, same items in different position",
			first: types.Sessions{
				types.Session{SessionId: types.SessionID{Value: 19}},
				types.Session{SessionId: types.SessionID{Value: 45}},
			},
			second: types.Sessions{
				types.Session{SessionId: types.SessionID{Value: 45}},
				types.Session{SessionId: types.SessionID{Value: 19}},
			},
			expEquals: false,
		},
		{
			name: "Non empty lists, same items and same position",
			first: types.Sessions{
				types.Session{SessionId: types.SessionID{Value: 19}},
				types.Session{SessionId: types.SessionID{Value: 45}},
			},
			second: types.Sessions{
				types.Session{SessionId: types.SessionID{Value: 19}},
				types.Session{SessionId: types.SessionID{Value: 45}},
			},
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}
