package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func TestSplitActivePollQueueKey(t *testing.T) {
	testCases := []struct {
		name          string
		key           []byte
		shouldErr     bool
		expSubspaceID uint64
		expPostID     uint64
		expoPollID    uint32
		expDate       time.Time
	}{
		{
			name:      "invalid key returns error",
			key:       []byte{},
			shouldErr: true,
		},
		{
			name:          "valid key returns proper values",
			key:           types.ActivePollQueueKey(1, 2, 3, time.Date(2020, 1, 1, 12, 00, 00, 00, time.UTC)),
			expSubspaceID: 1,
			expPostID:     2,
			expoPollID:    3,
			expDate:       time.Date(2020, 1, 1, 12, 00, 00, 00, time.UTC),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldErr {
				require.Panics(t, func() { types.SplitActivePollQueueKey(tc.key) })
			} else {
				subspaceID, postID, pollID, date := types.SplitActivePollQueueKey(tc.key)
				require.Equal(t, tc.expSubspaceID, subspaceID)
				require.Equal(t, tc.expPostID, postID)
				require.Equal(t, tc.expoPollID, pollID)
				require.Equal(t, tc.expDate, date)
			}
		})
	}
}

func TestGetPollIDFromBytes(t *testing.T) {
	testCases := []struct {
		name          string
		key           []byte
		shouldErr     bool
		expSubspaceID uint64
		expPostID     uint64
		expPollID     uint32
	}{
		{
			name:      "invalid key length returns error",
			key:       []byte{},
			shouldErr: true,
		},
		{
			name:          "valid key length returns proper values",
			key:           types.GetPollIDBytes(1, 1, 1),
			shouldErr:     false,
			expSubspaceID: 1,
			expPostID:     1,
			expPollID:     1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldErr {
				require.Panics(t, func() { types.GetPollIDFromBytes(tc.key) })
			} else {
				subspaceID, postID, pollID := types.GetPollIDFromBytes(tc.key)
				require.Equal(t, tc.expSubspaceID, subspaceID)
				require.Equal(t, tc.expPostID, postID)
				require.Equal(t, tc.expPollID, pollID)
			}
		})
	}

}
