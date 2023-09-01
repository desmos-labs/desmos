package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
	"github.com/stretchr/testify/require"
)

func TestQueryApplicationLinksRequest_GetQueryPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		request  *types.QueryApplicationLinksRequest
		expected []byte
	}{
		{
			name:     "empty request returns properly",
			request:  types.NewQueryApplicationLinksRequest("", "", "", nil),
			expected: types.ApplicationLinkPrefix,
		},
		{
			name:     "request with all non empty parameters returns properly",
			request:  types.NewQueryApplicationLinksRequest("user", "application", "username", nil),
			expected: types.UserApplicationLinkKey("user", "application", "username"),
		},
		{
			name:     "request without username returns properly",
			request:  types.NewQueryApplicationLinksRequest("user", "application", "", nil),
			expected: types.UserApplicationLinksApplicationPrefix("user", "application"),
		},
		{
			name:     "request without application returns properly",
			request:  types.NewQueryApplicationLinksRequest("user", "", "username", nil),
			expected: types.UserApplicationLinksPrefix("user"),
		},
		{
			name:     "request without user returns properly",
			request:  types.NewQueryApplicationLinksRequest("", "application", "username", nil),
			expected: types.ApplicationLinkPrefix,
		},
		{
			name:     "request with username only returns properly",
			request:  types.NewQueryApplicationLinksRequest("", "", "username", nil),
			expected: types.ApplicationLinkPrefix,
		},
		{
			name:     "request with application only returns properly",
			request:  types.NewQueryApplicationLinksRequest("", "application", "", nil),
			expected: types.ApplicationLinkPrefix,
		},
		{
			name:     "request with user only returns properly",
			request:  types.NewQueryApplicationLinksRequest("user", "", "", nil),
			expected: types.UserApplicationLinksPrefix("user"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			prefix := tc.request.GetQueryPrefix()
			require.Equal(t, tc.expected, prefix)
		})
	}
}

func TestQueryApplicationLinkOwnersRequest_GetQueryPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		request  *types.QueryApplicationLinkOwnersRequest
		expected []byte
	}{
		{
			name:     "empty request returns properly",
			request:  types.NewQueryApplicationLinkOwnersRequest("", "", nil),
			expected: types.ApplicationLinkAppPrefix,
		},
		{
			name:     "request with application name and username returns properly",
			request:  types.NewQueryApplicationLinkOwnersRequest("application", "username", nil),
			expected: types.ApplicationLinkAppUsernameKey("application", "username"),
		},
		{
			name:     "request with application name only returns properly",
			request:  types.NewQueryApplicationLinkOwnersRequest("application", "", nil),
			expected: types.ApplicationLinkAppKey("application"),
		},
		{
			name:     "request with username only returns properly",
			request:  types.NewQueryApplicationLinkOwnersRequest("", "username", nil),
			expected: types.ApplicationLinkAppPrefix,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			prefix := tc.request.GetQueryPrefix()
			require.Equal(t, tc.expected, prefix)
		})
	}
}
