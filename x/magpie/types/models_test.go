package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

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
		t.Run(fmt.Sprintf("%d id valid: %t", test.id, test.shouldBeValid), func(t *testing.T) {
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
		t.Run(fmt.Sprintf("%d next is %d", test.id, test.nextID), func(t *testing.T) {
			require.Equal(t, test.nextID, test.id.Next())
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

// ___________________________________________________________________________________________________________________

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
