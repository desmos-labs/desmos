package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/subspaces/types"
)

func TestSubspace_WithName(t *testing.T) {
	tests := []struct {
		name        string
		subspace    types.Subspace
		expSubspace types.Subspace
		newName     string
	}{
		{
			name: "Do not modify does not modify the name",
			subspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			newName: types.DoNotModify,
		},
		{
			name: "Name edited correctly",
			subspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"newName",
				"",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			newName: "newName",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			subspace := test.subspace.WithName(test.newName)
			require.Equal(t, test.expSubspace, subspace)
		})
	}
}

func TestSubspace_WithOwner(t *testing.T) {
	tests := []struct {
		name        string
		subspace    types.Subspace
		expSubspace types.Subspace
		newOwner    string
	}{
		{
			name: "empty owner does not modify the owner field",
			subspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			newOwner: "",
		},
		{
			name: "new owner modify the owner field",
			subspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"newOwner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			newOwner: "newOwner",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			subspace := test.subspace.WithOwner(test.newOwner)
			require.Equal(t, test.expSubspace, subspace)
		})
	}
}

func TestSubspace_WithSubspaceType(t *testing.T) {
	tests := []struct {
		name        string
		subspace    types.Subspace
		expSubspace types.Subspace
		subType     types.SubspaceType
	}{
		{
			name: "unspecified subspace type does not modify the owner field",
			subspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			subType: types.SubspaceTypeUnspecified,
		},
		{
			name: "new subspace type modify the type field",
			subspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"owner",
				"",
				types.SubspaceTypeClosed,
				time.Unix(1, 2),
			),
			subType: types.SubspaceTypeClosed,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			subspace := test.subspace.WithSubspaceType(test.subType)
			require.Equal(t, test.expSubspace, subspace)
		})
	}
}

func TestSubspace_WithDescription(t *testing.T) {
	tests := []struct {
		name           string
		subspace       types.Subspace
		expSubspace    types.Subspace
		newDescription string
	}{
		{
			name: "Do not modify does not modify the description field",
			subspace: types.NewSubspace(
				"123",
				"name",
				"desc",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"desc",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			newDescription: types.DoNotModify,
		},
		{
			name: "new description modify the description field",
			subspace: types.NewSubspace(
				"123",
				"name",
				"",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"newDescr",
				"",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			newDescription: "newDescr",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			subspace := test.subspace.WithDescription(test.newDescription)
			require.Equal(t, test.expSubspace, subspace)
		})
	}
}

func TestSubspace_WithLogo(t *testing.T) {
	tests := []struct {
		name        string
		subspace    types.Subspace
		expSubspace types.Subspace
		newLogo     string
	}{
		{
			name: "Do not modify does not modify the logo field",
			subspace: types.NewSubspace(
				"123",
				"name",
				"desc",
				"logo",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"desc",
				"logo",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			newLogo: types.DoNotModify,
		},
		{
			name: "new logo modify the logo field",
			subspace: types.NewSubspace(
				"123",
				"name",
				"desc",
				"logo",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			expSubspace: types.NewSubspace(
				"123",
				"name",
				"desc",
				"newLogo",
				"owner",
				"",
				types.SubspaceTypeOpen,
				time.Unix(1, 2),
			),
			newLogo: "newLogo",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			subspace := test.subspace.WithLogo(test.newLogo)
			require.Equal(t, test.expSubspace, subspace)
		})
	}
}

func TestSubspace_Validate(t *testing.T) {
	date := time.Date(2050, 01, 01, 15, 15, 00, 000, time.UTC)
	tests := []struct {
		name     string
		subspace types.Subspace
		expError bool
	}{
		{
			name: "Invalid ID returns error",
			subspace: types.NewSubspace(
				"123",
				"",
				"",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid name returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"",
				"",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid owner returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"",
				"",
				"",
				"",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid creator returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"",
				"",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid creation time returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"",
				"",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types.SubspaceTypeOpen,
				time.Time{},
			),
			expError: true,
		},
		{
			name: "Invalid subspace logo returns error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"",
				"h",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types.SubspaceTypeOpen,
				date,
			),
			expError: true,
		},
		{
			name: "Valid subspace returns no error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"",
				"https://shorturl.at/adnX3",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types.SubspaceTypeOpen,
				date,
			),
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.subspace.Validate()
			if test.expError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_IsValidSubspaceType(t *testing.T) {
	tests := []struct {
		name     string
		subType  types.SubspaceType
		expValid bool
	}{
		{
			name:     "valid open type returns true",
			subType:  types.SubspaceTypeOpen,
			expValid: true,
		},
		{
			name:     "valid close type returns true",
			subType:  types.SubspaceTypeClosed,
			expValid: true,
		},
		{
			name:     "invalid type returns false",
			subType:  types.SubspaceTypeUnspecified,
			expValid: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expValid, types.IsValidSubspaceType(test.subType))
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
			name:       "Valid open subspace Type",
			subType:    "open",
			expSubType: types.SubspaceTypeOpen.String(),
		},
		{
			name:       "Valid closed subspace type",
			subType:    "closed",
			expSubType: types.SubspaceTypeClosed.String(),
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
			expSubType: types.SubspaceTypeUnspecified,
			expError:   fmt.Errorf("'invalid' is not a valid subspace type"),
		},
		{
			name:       "Valid subspace type",
			subType:    types.SubspaceTypeOpen.String(),
			expSubType: types.SubspaceTypeOpen,
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
