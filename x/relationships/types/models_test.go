package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/v4/app"

	"github.com/desmos-labs/desmos/v4/x/relationships/types"

	"github.com/stretchr/testify/require"
)

func TestRelationship_Validate(t *testing.T) {
	testCases := []struct {
		name         string
		relationship types.Relationship
		shouldErr    bool
	}{
		{
			name: "empty creator returns error",
			relationship: types.NewRelationship(
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				0,
			),
			shouldErr: true,
		},
		{
			name: "empty recipient returns error",
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				0,
			),
			shouldErr: true,
		},
		{
			name: "same creator and recipient return error",
			relationship: types.NewRelationship(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				0,
			),
			shouldErr: true,
		},
		{
			name: "valid relationship returns no error",
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				0,
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.relationship.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRelationshipMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	relationship := types.NewRelationship("creator", "recipient_1", 1)
	marshalled := types.MustMarshalRelationship(cdc, relationship)
	unmarshalled := types.MustUnmarshalRelationship(cdc, marshalled)
	require.Equal(t, relationship, unmarshalled)
}

// --------------------------------------------------------------------------------------------------------------------

func TestUserBlock_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		userBlock types.UserBlock
		shouldErr bool
	}{
		{
			name: "empty blocker address returns error",
			userBlock: types.NewUserBlock(
				"",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"reason",
				0,
			),
			shouldErr: true,
		},
		{
			name: "empty blocked address returns error",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				"reason",
				0,
			),
			shouldErr: true,
		},
		{
			name: "equals blocker and blocked addresses returns error",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"reason",
				0,
			),
			shouldErr: true,
		},
		{
			name: "correct user block returns no error",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"reason",
				0,
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.userBlock.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
