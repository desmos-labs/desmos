package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/stretchr/testify/require"
)

func TestSubspace_WithName(t *testing.T) {
	sub := types.NewSubspace("123", "name", "", "", types.Open, time.Unix(1, 2))

	sub = sub.WithName("sub")

	require.Equal(t, "sub", sub.Name)
}

func TestSubspace_WithOwner(t *testing.T) {
	sub := types.NewSubspace("123", "name", "", "", types.Open, time.Unix(1, 2))

	sub = sub.WithOwner("owner")

	require.Equal(t, "owner", sub.Owner)
}

func TestSubspace_WithSubspaceType(t *testing.T) {
	sub := types.NewSubspace("123", "name", "", "", types.Open, time.Unix(1, 2))

	sub = sub.WithSubspaceType(types.Close)

	require.Equal(t, types.Close, sub.Type)
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
			subspace: types.NewSubspace("123", "", "", "", types.Open, time.Time{}),
			expError: fmt.Errorf("invalid subspace id: 123 it must be a valid SHA-256 hash"),
		},
		{
			name: "Invalid name returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"",
				"",
				"",
				types.Open,
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
				types.Open,
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
				types.Open,
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
				types.Open,
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
				Type:            types.Open,
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
				Type:            types.Open,
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
				Type:            types.Open,
				Admins:          nil,
				BannedUsers:     nil,
				RegisteredUsers: []string{""},
			},
			expError: fmt.Errorf("invalid subspace registered user address"),
		},
		{
			name: "Invalid subspace types returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types.Unspecified,
				date,
			),
			expError: fmt.Errorf("invalid subspace type: %s", types.Unspecified),
		},
		{
			name: "Valid subspace returns no error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types.Open,
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

func TestSubspace_IsAdmin(t *testing.T) {
	tests := []struct {
		name     string
		user     string
		subspace types.Subspace
		expBool  bool
	}{
		{
			name: "user is an admin",
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    time.Time{},
				Type:            types.Open,
				Admins:          []string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				BannedUsers:     nil,
				RegisteredUsers: nil,
			},
			expBool: true,
		},
		{
			name: "user is not an admin",
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    time.Time{},
				Type:            types.Open,
				Admins:          nil,
				BannedUsers:     nil,
				RegisteredUsers: nil,
			},
			expBool: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res := test.subspace.IsAdmin(test.user)
			require.Equal(t, test.expBool, res)
		})
	}
}

func TestSubspace_IsBanned(t *testing.T) {
	tests := []struct {
		name     string
		user     string
		subspace types.Subspace
		expBool  bool
	}{
		{
			name: "user is banned",
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    time.Time{},
				Type:            types.Open,
				Admins:          nil,
				BannedUsers:     []string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				RegisteredUsers: nil,
			},
			expBool: true,
		},
		{
			name: "user is not banned",
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    time.Time{},
				Type:            types.Open,
				Admins:          nil,
				BannedUsers:     nil,
				RegisteredUsers: nil,
			},
			expBool: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res := test.subspace.IsBanned(test.user)
			require.Equal(t, test.expBool, res)
		})
	}
}

func TestSubspace_IsRegistered(t *testing.T) {
	tests := []struct {
		name     string
		user     string
		subspace types.Subspace
		expBool  bool
	}{
		{
			name: "user is registered",
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    time.Time{},
				Type:            types.Open,
				Admins:          nil,
				BannedUsers:     nil,
				RegisteredUsers: []string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
			},
			expBool: true,
		},
		{
			name: "user is not registered",
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace: types.Subspace{
				ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:            "test",
				Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime:    time.Time{},
				Type:            types.Open,
				Admins:          nil,
				BannedUsers:     nil,
				RegisteredUsers: nil,
			},
			expBool: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res := test.subspace.IsRegistered(test.user)
			require.Equal(t, test.expBool, res)
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

func Test_IsValidSubspaceType(t *testing.T) {
	tests := []struct {
		name    string
		subType types.SubspaceType
		expBool bool
	}{
		{
			name:    "valid open type returns true",
			subType: types.Open,
			expBool: true,
		},
		{
			name:    "valid close type returns true",
			subType: types.Close,
			expBool: true,
		},
		{
			name:    "invalid type returns false",
			subType: types.Unspecified,
			expBool: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, types.IsValidSubspaceType(test.subType))
		})
	}
}

func Test_NormalizeSubspaceType(t *testing.T) {
	tests := []struct {
		name       string
		subType    string
		expSubType string
	}{
		{
			name:       "Valid Open subspace Type",
			subType:    "open",
			expSubType: types.Open.String(),
		},
		{
			name:       "Valid Close subspace type",
			subType:    "Close",
			expSubType: types.Close.String(),
		},
		{
			name:       "Invalid subspace type",
			subType:    "Invalid",
			expSubType: "Invalid",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {

			subspaceType := types.NormalizeSubspaceType(test.subType)
			require.Equal(t, test.expSubType, subspaceType)
		})
	}
}

func Test_SubspaceTypeFromString(t *testing.T) {
	tests := []struct {
		name       string
		subType    string
		expSubType types.SubspaceType
		expError   error
	}{
		{
			name:       "Invalid subspace type",
			subType:    "invalid",
			expSubType: types.Unspecified,
			expError:   fmt.Errorf("'invalid' is not a valid subspace type"),
		},
		{
			name:       "Valid subspace type",
			subType:    types.Open.String(),
			expSubType: types.Open,
			expError:   nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res, err := types.SubspaceTypeFromString(test.subType)
			require.Equal(t, test.expError, err)
			require.Equal(t, test.expSubType, res)
		})
	}
}
