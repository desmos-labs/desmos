package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
	"github.com/stretchr/testify/require"
)

func TestQueryChainLinksRequest_GetQueryPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		request  *types.QueryChainLinksRequest
		expected []byte
	}{
		{
			name:     "empty request returns properly",
			request:  types.NewQueryChainLinksRequest("", "", "", nil),
			expected: types.ChainLinksPrefix,
		},
		{
			name:     "request with all non empty parameters returns properly",
			request:  types.NewQueryChainLinksRequest("user", "chain", "target", nil),
			expected: types.ChainLinksStoreKey("user", "chain", "target"),
		},
		{
			name:     "request without target returns properly",
			request:  types.NewQueryChainLinksRequest("user", "chain", "", nil),
			expected: types.UserChainLinksChainPrefix("user", "chain"),
		},
		{
			name:     "request without chain returns properly",
			request:  types.NewQueryChainLinksRequest("user", "", "target", nil),
			expected: types.UserChainLinksPrefix("user"),
		},
		{
			name:     "request without user returns properly",
			request:  types.NewQueryChainLinksRequest("", "chain", "target", nil),
			expected: types.ChainLinksPrefix,
		},
		{
			name:     "request with username only returns properly",
			request:  types.NewQueryChainLinksRequest("", "", "target", nil),
			expected: types.ChainLinksPrefix,
		},
		{
			name:     "request with chain only returns properly",
			request:  types.NewQueryChainLinksRequest("", "chain", "", nil),
			expected: types.ChainLinksPrefix,
		},
		{
			name:     "request with user only returns properly",
			request:  types.NewQueryChainLinksRequest("user", "", "", nil),
			expected: types.UserChainLinksPrefix("user"),
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

func TestQueryChainLinkOwnersRequest_GetQueryPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		request  *types.QueryChainLinkOwnersRequest
		expected []byte
	}{
		{
			name:     "empty request returns properly",
			request:  types.NewQueryChainLinkOwnersRequest("", "", nil),
			expected: types.ChainLinkChainPrefix,
		},
		{
			name:     "request with application name and username returns properly",
			request:  types.NewQueryChainLinkOwnersRequest("chain", "username", nil),
			expected: types.ChainLinkChainAddressKey("chain", "username"),
		},
		{
			name:     "request with chain only returns properly",
			request:  types.NewQueryChainLinkOwnersRequest("chain", "", nil),
			expected: types.ChainLinkChainKey("chain"),
		},
		{
			name:     "request with username only returns properly",
			request:  types.NewQueryChainLinkOwnersRequest("", "username", nil),
			expected: types.ChainLinkChainPrefix,
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

func TestQueryDefaultExternalAddressesRequest_GetQueryPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		request  *types.QueryDefaultExternalAddressesRequest
		expected []byte
	}{
		{
			name:     "empty request returns properly",
			request:  types.NewQueryDefaultExternalAddressesRequest("", "", nil),
			expected: types.DefaultExternalAddressPrefix,
		},
		{
			name:     "request with owner and chain returns properly",
			request:  types.NewQueryDefaultExternalAddressesRequest("owner", "chain", nil),
			expected: types.DefaultExternalAddressKey("owner", "chain"),
		},
		{
			name:     "request with owner only returns properly",
			request:  types.NewQueryDefaultExternalAddressesRequest("owner", "", nil),
			expected: types.OwnerDefaultExternalAddressPrefix("owner"),
		},
		{
			name:     "request with chain only returns properly",
			request:  types.NewQueryDefaultExternalAddressesRequest("", "chain", nil),
			expected: types.DefaultExternalAddressPrefix,
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
