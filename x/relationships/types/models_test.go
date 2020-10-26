package types_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewUserBlock(t *testing.T) {
	userBlock := types.NewUserBlock(
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		"idk",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	actual := types.NewUserBlock(
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		"idk",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	require.Equal(t, userBlock, actual)
}

func TestUserBlock_String(t *testing.T) {
	userBlock := types.NewUserBlock(
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		"idk",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	require.Equal(t,
		"User Block: [Blocker] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 [Blocked] cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 [Reason] idk [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		userBlock.String(),
	)
}

func TestUserBlock_Validate(t *testing.T) {
	tests := []struct {
		name      string
		userBlock types.UserBlock
		expError  error
	}{
		{
			name: "empty blocker address returns error",
			userBlock: types.NewUserBlock(
				"",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expError: fmt.Errorf("blocker address cannot be empty"),
		},
		{
			name: "empty blocked address returns error",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expError: fmt.Errorf("the address of the blocked user cannot be empty"),
		},
		{
			name: "equals blocker and blocked addresses returns error",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expError: fmt.Errorf("blocker and blocked addresses cannot be equals"),
		},
		{
			name: "invalid subspace returns error",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"reason",
				"yeah",
			),
			expError: fmt.Errorf("subspace must be a valid sha-256 hash"),
		},
		{
			name: "correct user block returns no error",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expError, test.userBlock.Validate())
		})
	}
}

func TestUserBlock_Equals(t *testing.T) {
	tests := []struct {
		name           string
		userBlock      types.UserBlock
		otherUserBlock types.UserBlock
		expBool        bool
	}{
		{
			name: "Equals user block returns true",
			userBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			otherUserBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expBool: true,
		},
		{
			name: "Non equals user block returns false",
			userBlock: types.NewUserBlock(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			otherUserBlock: types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.userBlock.Equal(test.otherUserBlock))
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestNewRelationship(t *testing.T) {
	actual := types.NewRelationship(
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	require.Equal(t, actual.Recipient, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.Equal(t, actual.Subspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
}

func TestRelationship_String(t *testing.T) {
	relationship := types.NewRelationship(
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	require.Equal(t,
		"Relationship:[Recipient] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		relationship.String(),
	)
}

func TestRelationship_Validate(t *testing.T) {
	tests := []struct {
		name         string
		relationship types.Relationship
		expErr       error
	}{
		{
			name:         "Empty recipient returns error",
			relationship: types.NewRelationship("", ""),
			expErr:       fmt.Errorf("recipient can't be empty"),
		},
		{
			name:         "Invalid subspace returns error",
			relationship: types.NewRelationship("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", ""),
			expErr:       fmt.Errorf("subspace must be a valid sha-256"),
		},
		{
			name: "Valid relationship returns no error",
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.relationship.Validate())
		})
	}
}

func TestRelationship_Equals(t *testing.T) {
	tests := []struct {
		name         string
		relationship types.Relationship
		otherRel     types.Relationship
		expBool      bool
	}{
		{
			name: "Equals relationships returns true",
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			otherRel: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expBool: true,
		},
		{
			name: "Non equals relationships returns false",
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			otherRel: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"1234",
			),
			expBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.relationship.Equal(test.otherRel))
		})
	}
}
