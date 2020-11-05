package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/app"

	"github.com/desmos-labs/desmos/x/relationships/types"

	"github.com/stretchr/testify/require"
)

func TestRelationship_Validate(t *testing.T) {
	tests := []struct {
		name         string
		relationship types.Relationship
		expErr       error
	}{
		{
			name: "Empty creator returns error",
			relationship: types.NewRelationship(
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: fmt.Errorf("invalid creator address: "),
		},
		{
			name: "Empty recipient returns error",
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: fmt.Errorf("invalid recipient address: "),
		},
		{
			name: "Invalid subspace returns error",
			relationship: types.NewRelationship(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
			),
			expErr: fmt.Errorf("subspace must be a valid sha-256"),
		},
		{
			name: "Same creator and recipient return error",
			relationship: types.NewRelationship(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: fmt.Errorf("creator and recipient cannot be the same user"),
		},
		{
			name: "Valid relationship returns no error",
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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

// ___________________________________________________________________________________________________________________

func TestRemoveRelationship(t *testing.T) {
	tests := []struct {
		name     string
		slice    []types.Relationship
		toRemove types.Relationship
		expFound bool
		expSlice []types.Relationship
	}{
		{
			name:     "cannot delete from empty slice",
			slice:    nil,
			toRemove: types.NewRelationship("creator", "recipient_2", "subspace"),
			expFound: false,
			expSlice: nil,
		},
		{
			name: "first relationship is removed correctly",
			slice: []types.Relationship{
				types.NewRelationship("creator", "recipient_1", "subspace"),
				types.NewRelationship("creator", "recipient_2", "subspace"),
				types.NewRelationship("creator", "recipient_3", "subspace"),
			},
			toRemove: types.NewRelationship("creator", "recipient_1", "subspace"),
			expFound: true,
			expSlice: []types.Relationship{
				types.NewRelationship("creator", "recipient_2", "subspace"),
				types.NewRelationship("creator", "recipient_3", "subspace"),
			},
		},
		{
			name: "middle relationship is removed correctly",
			slice: []types.Relationship{
				types.NewRelationship("creator", "recipient_1", "subspace"),
				types.NewRelationship("creator", "recipient_2", "subspace"),
				types.NewRelationship("creator", "recipient_3", "subspace"),
			},
			toRemove: types.NewRelationship("creator", "recipient_2", "subspace"),
			expFound: true,
			expSlice: []types.Relationship{
				types.NewRelationship("creator", "recipient_1", "subspace"),
				types.NewRelationship("creator", "recipient_3", "subspace"),
			},
		},
		{
			name: "last relationship is removed correctly",
			slice: []types.Relationship{
				types.NewRelationship("creator", "recipient_1", "subspace"),
				types.NewRelationship("creator", "recipient_2", "subspace"),
				types.NewRelationship("creator", "recipient_3", "subspace"),
			},
			toRemove: types.NewRelationship("creator", "recipient_3", "subspace"),
			expFound: true,
			expSlice: []types.Relationship{
				types.NewRelationship("creator", "recipient_1", "subspace"),
				types.NewRelationship("creator", "recipient_2", "subspace"),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			slice, found := types.RemoveRelationship(test.slice, test.toRemove)
			require.Equal(t, slice, test.expSlice)
			require.Equal(t, found, test.expFound)
		})
	}
}

func TestRelationshipsMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	relationships := []types.Relationship{
		types.NewRelationship("creator", "recipient_1", "subspace"),
		types.NewRelationship("creator", "recipient_2", "subspace"),
		types.NewRelationship("creator", "recipient_3", "subspace"),
	}
	marshalled := types.MustMarshalRelationships(cdc, relationships)
	unmarshalled := types.MustUnmarshalRelationships(cdc, marshalled)
	require.Equal(t, relationships, unmarshalled)
}

// ___________________________________________________________________________________________________________________

func TestRemoveUserBlock(t *testing.T) {
	tests := []struct {
		name   string
		blocks []types.UserBlock
		data   struct {
			blocker  string
			blocked  string
			subspace string
		}
		expFound bool
		expSlice []types.UserBlock
	}{
		{
			name:   "empty slice does not allow removal",
			blocks: nil,
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expFound: false,
			expSlice: nil,
		},
		{
			name: "first block is removed properly",
			blocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked_1", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_3", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked_1",
				subspace: "subspace",
			},
			expFound: true,
			expSlice: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_3", "reason", "subspace"),
			},
		},
		{
			name: "middle block is removed properly",
			blocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked_1", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_3", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked_2",
				subspace: "subspace",
			},
			expFound: true,
			expSlice: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked_1", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_3", "reason", "subspace"),
			},
		},
		{
			name: "last block is removed properly",
			blocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked_1", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_3", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked_3",
				subspace: "subspace",
			},
			expFound: true,
			expSlice: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked_1", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			slice, found := types.RemoveUserBlock(test.blocks, test.data.blocker, test.data.blocked, test.data.subspace)
			require.Equal(t, test.expSlice, slice)
			require.Equal(t, test.expFound, found)
		})
	}
}

func TestUserBlocksMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	blocks := []types.UserBlock{
		types.NewUserBlock("blocker", "blocked_1", "reason", "subspace"),
		types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
		types.NewUserBlock("blocker", "blocked_3", "reason", "subspace"),
	}
	marshaled := types.MustMarshalUserBlocks(cdc, blocks)
	unmarshalled := types.MustUnmarshalUserBlocks(cdc, marshaled)
	require.Equal(t, blocks, unmarshalled)
}

// ___________________________________________________________________________________________________________________

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
