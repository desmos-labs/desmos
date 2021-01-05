package types_test

import (
	"errors"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/types"

	"github.com/stretchr/testify/require"
)

func TestReaction_Validate(t *testing.T) {
	tests := []struct {
		name     string
		reaction types.RegisteredReaction
		error    error
	}{
		{
			name: "Valid reaction returns no error (url on value)",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile-jpg:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
		{
			name: "Invalid reaction returns error (unicode on value)",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smiles:",
				"U+1F600",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "Missing creator returns error",
			reaction: types.NewRegisteredReaction(
				"",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("invalid reaction creator: "),
		},
		{
			name: "Empty short code returns error",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			//nolint - errcheck
			error: errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
		},
		{
			name: "Invalid short code returns error",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			//nolint - errcheck
			error: errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
		},
		{
			name: "Empty value returns error",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "invalid value returns error (url)",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "invalid value returns error (unicode)",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"U+1",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "invalid subspace returns no error",
			reaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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

// ___________________________________________________________________________________________________________________

func TestPostReaction_Validate(t *testing.T) {
	tests := []struct {
		name     string
		reaction types.PostReaction
		error    error
	}{
		{
			name: "Valid reaction returns no error",
			reaction: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			error: nil,
		},
		{
			name: "Missing owner returns error",
			reaction: types.NewPostReaction(
				":smile:",
				"reaction",
				"",
			),
			error: errors.New("invalid reaction owner: "),
		},
		{
			name: "Missing value returns error",
			reaction: types.NewPostReaction(
				":smile:",
				"",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			error: errors.New("reaction value cannot be empty or blank"),
		},
		{
			name: "Invalid shortcode returns error",
			reaction: types.NewPostReaction(
				"invalid",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			error: errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.error, test.reaction.Validate())
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestPostReactions_ContainsOwnerLike(t *testing.T) {
	tests := []struct {
		name        string
		reactions   types.PostReactions
		owner       string
		shortcode   string
		expContains bool
	}{
		{
			name: "Non-empty list returns true with valid address",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:       "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode:   ":smile:",
			expContains: true,
		},
		{
			name:        "Empty list returns false",
			reactions:   types.NewPostReactions(),
			owner:       "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode:   ":smile:",
			expContains: false,
		},
		{
			name: "Non-empty list returns false with not found address",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:       "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			shortcode:   ":smile:",
			expContains: false,
		},
		{
			name: "Non-empty list returns false with not found value",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:       "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode:   ":like:",
			expContains: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expContains, test.reactions.ContainsReactionFrom(test.owner, test.shortcode))
		})
	}
}

func TestPostReactions_IndexOfByUserAndValue(t *testing.T) {
	tests := []struct {
		name      string
		reactions types.PostReactions
		owner     string
		value     string
		expIndex  int
	}{
		{
			name: "Non-empty list returns proper index with valid value (shortcode)",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":+1:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    ":+1:",
			expIndex: 0,
		},
		{
			name: "Non-empty list returns proper index with valid value (emoji - one code)",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":+1:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    "👍",
			expIndex: 0,
		},
		{
			name: "Non-empty list returns proper index with valid value (emoji - another code)",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":thumbsup:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    "👍",
			expIndex: 0,
		},
		{
			name:      "Empty list returns -1",
			reactions: types.NewPostReactions(),
			owner:     "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name: "Non-empty list returns -1 with not found address",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			value:    "reaction",
			expIndex: -1,
		},
		{
			name: "Non-empty list returns -1 with not found value",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			value:    "reaction-2",
			expIndex: -1,
		},
		{
			name: "Existing reaction search by code",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":reaction:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    ":reaction:",
			expIndex: 0,
		},
		{
			name: "Exiting emoji reaction stored by value search by code",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":fire:",
					"🔥",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    ":fire:",
			expIndex: 0,
		},
		{
			name: "Exiting emoji reaction stored by code search by code",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":fire:",
					"🔥",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    ":fire:",
			expIndex: 0,
		},
		{
			name: "Exiting emoji reaction stored by code search by value",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":fire:",
					"🔥",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    "🔥",
			expIndex: 0,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expIndex, test.reactions.IndexOfByUserAndValue(test.owner, test.value))
		})
	}
}

func TestPostReactions_RemoveReaction(t *testing.T) {
	tests := []struct {
		name      string
		reactions types.PostReactions
		owner     string
		shortcode string
		expResult types.PostReactions
		expEdited bool
	}{
		{
			name: "PostReaction is removed from non-empty list",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:     "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode: ":smile:",
			expResult: types.NewPostReactions([]types.PostReaction{}...),
			expEdited: true,
		},
		{
			name:      "Empty list is not edited",
			reactions: types.NewPostReactions(),
			owner:     "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode: ":smile:",
			expResult: types.NewPostReactions(),
			expEdited: false,
		},
		{
			name: "Non-empty list with not found address is not edited",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:     "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			shortcode: ":smile:",
			expResult: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			expEdited: false,
		},
		{
			name: "Non-empty list with not found value is not edited",
			reactions: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			owner:     "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			shortcode: ":like:",
			expResult: types.NewPostReactions(
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			),
			expEdited: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result, edited := test.reactions.RemoveReaction(test.owner, test.shortcode)
			require.Equal(t, test.expEdited, edited)
			require.Equal(t, test.expResult, result)
		})
	}
}
