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
