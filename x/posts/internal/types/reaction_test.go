package types_test

import (
	"errors"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestReaction_String(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	reaction := types.NewReaction(
		user,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	require.Equal(t, `{"ShortCode":":smile:","Value":"https://smile.jpg","Subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","Creator":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}`, reaction.String())
}

func TestReaction_Validate(t *testing.T) {

	tests := []struct {
		name     string
		reaction types.Reaction
		error    error
	}{
		{
			name: "Valid reaction returns no error (url on value)",
			reaction: types.NewReaction(
				testOwner,
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
		/*{
			name: "Valid reaction returns no error (unicode on value)",
			reaction: types.NewReaction(
				testOwner,
				":smile:",
				"U+1F600",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},*/
		{
			name: "Missing creator returns error",
			reaction: types.NewReaction(
				nil,
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("invalid reaction creator: "),
		},
		{
			name: "Empty short code returns error",
			reaction: types.NewReaction(
				testOwner,
				"",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction short code cannot be empty or blank"),
		},
		{
			name: "Invalid short code returns error",
			reaction: types.NewReaction(
				testOwner,
				"smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction short code must be an emoji short code"),
		},
		{
			name: "Empty value returns error",
			reaction: types.NewReaction(
				testOwner,
				":smile:",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value cannot be empty or blank"),
		},
		{
			name: "invalid value returns error",
			reaction: types.NewReaction(
				testOwner,
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL or an emoji unicode"),
		},
		{
			name: "invalid subspace returns no error",
			reaction: types.NewReaction(
				testOwner,
				":smile:",
				"https://smile.jpg",
				"1234",
			),
			error: errors.New("reaction subspace must be a valid sha-256 hash"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.error, test.reaction.Validate())
		})
	}
}

func TestReaction_Equals(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name          string
		first         types.Reaction
		second        types.Reaction
		shouldBeEqual bool
	}{
		{
			name: "Returns false with different user",
			first: types.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			second: types.NewReaction(otherLiker, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			shouldBeEqual: false,
		},
		{
			name: "Returns true with the same data",
			first: types.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			second: types.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			shouldBeEqual: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.shouldBeEqual, test.first.Equals(test.second))
		})
	}
}

func TestReactions_AppendIfMissing(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name        string
		reactions   types.Reactions
		newReaction types.Reaction
		expReaction types.Reactions
		expAppend   bool
	}{
		{
			name:      "New reaction is appended properly to empty list",
			reactions: types.Reactions{},
			newReaction: types.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expReaction: types.Reactions{types.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
			expAppend: true,
		},
		{
			name: "New reaction is appended properly to existing list",
			reactions: types.Reactions{types.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
			newReaction: types.NewReaction(user, ":sad:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expAppend: true,
			expReaction: types.Reactions{
				types.NewReaction(user, ":smile:", "smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewReaction(user, ":sad:", "smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := test.reactions.AppendIfMissing(test.newReaction)
			require.Equal(t, test.expReaction, actual)
			require.Equal(t, test.expAppend, appended)
		})
	}
}
