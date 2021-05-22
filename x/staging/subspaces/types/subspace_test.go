package types_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSubspace_WithName(t *testing.T) {
	sub := types.NewSubspace("123", "name", "", "", true, time.Unix(1, 2))

	sub = sub.WithName("sub")

	assert.Equal(t, "sub", sub.Name)
}

func TestSubspace_WithOwner(t *testing.T) {
	sub := types.NewSubspace("123", "name", "", "", true, time.Unix(1, 2))

	sub = sub.WithOwner("owner")

	assert.Equal(t, "owner", sub.Owner)
}

func TestSubspace_Validate(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	require.NoError(t, err)

	tests := []struct {
		name     string
		subspace types.Subspace
		expError error
	}{
		{
			name:     "Invalid subspace returns error",
			subspace: types.NewSubspace("123", "", "", "", true, time.Time{}),
			expError: fmt.Errorf("invalid subspace id: 123 it must be a valid SHA-256 hash"),
		},
		{
			name: "Invalid name returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"",
				"",
				"",
				true,
				time.Time{},
			),
			expError: fmt.Errorf("subspace name cannot be empty or blank"),
		},
		{
			name: "Invalid owner returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"",
				"",
				true,
				time.Time{},
			),
			expError: fmt.Errorf("invalid subspace owner: "),
		},
		{
			name: "Invalid creator returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"",
				true,
				time.Time{},
			),
			expError: fmt.Errorf("invalid subspace creator: "),
		},
		{
			name: "Invalid creation time returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				true,
				time.Time{},
			),
			expError: fmt.Errorf("invalid subspace creation time: 0001-01-01 00:00:00 +0000 UTC"),
		},
		{
			name: "Invalid admin address returns error",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    date,
				Open:            true,
				Admins:          []string{""},
				BannedUsers:     nil,
				RegisteredUsers: nil,
			},
			expError: fmt.Errorf("invalid subspace admin address"),
		},
		{
			name: "Invalid blocked user address returns error",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    date,
				Open:            true,
				Admins:          nil,
				BannedUsers:     []string{""},
				RegisteredUsers: nil,
			},
			expError: fmt.Errorf("invalid subspace blocked user address"),
		},
		{
			name: "Invalid registered user address returns error",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    date,
				Open:            true,
				Admins:          nil,
				BannedUsers:     nil,
				RegisteredUsers: []string{""},
			},
			expError: fmt.Errorf("invalid subspace registered user address"),
		},
		{
			name: "Valid subspace returns no error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				true,
				date,
			),
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.subspace.Validate()
			require.Equal(t, test.expError, err)
		})
	}
}

func TestUsers_IsPresent(t *testing.T) {
	tests := []struct {
		name    string
		users   []string
		user    string
		expBool bool
	}{
		{
			name:    "User not found returns false",
			users:   []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			user:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expBool: false,
		},
		{
			name:    "User found returns true",
			users:   []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			user:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expBool: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			found := types.IsPresent(test.users, test.user)
			require.Equal(t, test.expBool, found)
		})
	}
}

func TestUsers_RemoveUser(t *testing.T) {
	tests := []struct {
		name     string
		users    []string
		user     string
		expUsers []string
	}{
		{
			name: "User removed correctly",
			users: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			},
		},
		{
			name:  "User not removed",
			users: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			user:  "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			users := types.RemoveUser(test.users, test.user)
			require.Equal(t, test.expUsers, users)
		})
	}
}

func TestUsers_ValidateUsers(t *testing.T) {
	tests := []struct {
		name     string
		users    []string
		userType string
		expErr   error
	}{
		{
			name: "Invalid user returns error",
			users: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"",
			},
			userType: types.Admin,
			expErr:   fmt.Errorf("invalid subspace admin address"),
		},
		{
			name: "Valid users returns no error",
			users: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			userType: types.Admin,
			expErr:   nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := types.ValidateUsers(test.users, test.userType)
			require.Equal(t, test.expErr, err)
		})
	}
}
