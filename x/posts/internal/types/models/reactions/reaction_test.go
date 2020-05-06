package reactions_test

import (
	"errors"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types/models/reactions"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestReaction_String(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	reaction := reactions.NewReaction(
		user,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	require.Equal(t, `{"shortcode":":smile:","value":"https://smile.jpg","subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}`, reaction.String())
}

func TestReaction_Validate(t *testing.T) {
	testOwner, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name     string
		reaction reactions.Reaction
		error    error
	}{
		{
			name: "Valid reaction returns no error (url on value)",
			reaction: reactions.NewReaction(
				testOwner,
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
		{
			name: "Valid reaction returns no error (unicode on value)",
			reaction: reactions.NewReaction(
				testOwner,
				":smile:",
				"U+1F600",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
		{
			name: "Missing creator returns error",
			reaction: reactions.NewReaction(
				nil,
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("invalid reaction creator: "),
		},
		{
			name: "Empty short code returns error",
			reaction: reactions.NewReaction(
				testOwner,
				"",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction short code must be an emoji short code"),
		},
		{
			name: "Invalid short code returns error",
			reaction: reactions.NewReaction(
				testOwner,
				"smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction short code must be an emoji short code"),
		},
		{
			name: "Empty value returns error",
			reaction: reactions.NewReaction(
				testOwner,
				":smile:",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL or an emoji"),
		},
		{
			name: "invalid value returns error (url)",
			reaction: reactions.NewReaction(
				testOwner,
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL or an emoji"),
		},
		{
			name: "invalid value returns error (unicode)",
			reaction: reactions.NewReaction(
				testOwner,
				":smile:",
				"U+1",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL or an emoji"),
		},
		{
			name: "invalid subspace returns no error",
			reaction: reactions.NewReaction(
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
		first         reactions.Reaction
		second        reactions.Reaction
		shouldBeEqual bool
	}{
		{
			name: "Returns false with different user",
			first: reactions.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			second: reactions.NewReaction(otherLiker, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			shouldBeEqual: false,
		},
		{
			name: "Returns true with the same data",
			first: reactions.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			second: reactions.NewReaction(user, ":smile:", "smile.jpg",
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
		reactions   reactions.Reactions
		newReaction reactions.Reaction
		expReaction reactions.Reactions
		expAppend   bool
	}{
		{
			name:      "New reaction is appended properly to empty list",
			reactions: reactions.Reactions{},
			newReaction: reactions.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expReaction: reactions.Reactions{reactions.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
			expAppend: true,
		},
		{
			name: "New reaction is appended properly to existing list",
			reactions: reactions.Reactions{reactions.NewReaction(user, ":smile:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
			newReaction: reactions.NewReaction(user, ":sad:", "smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expAppend: true,
			expReaction: reactions.Reactions{
				reactions.NewReaction(user, ":smile:", "smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				reactions.NewReaction(user, ":sad:", "smile.jpg",
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
