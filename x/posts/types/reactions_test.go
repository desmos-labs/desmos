package types_test

import (
	"errors"
	"github.com/desmos-labs/desmos/x/posts/types"
	"testing"

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

func TestReaction_Equals(t *testing.T) {
	tests := []struct {
		name          string
		first         types.RegisteredReaction
		second        types.RegisteredReaction
		shouldBeEqual bool
	}{
		{
			name: "Returns false with different user",
			first: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			second: types.NewRegisteredReaction(
				"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			shouldBeEqual: false,
		},
		{
			name: "Returns true with the same data",
			first: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			second: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
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

// ___________________________________________________________________________________________________________________

func TestReactions_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name        string
		reactions   types.Reactions
		newReaction types.RegisteredReaction
		expReaction types.Reactions
		expAppend   bool
	}{
		{
			name:      "New reaction is appended properly to empty list",
			reactions: types.Reactions{},
			newReaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expReaction: types.Reactions{
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":smile:",
					"smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expAppend: true,
		},
		{
			name: "New reaction is appended properly to existing list",
			reactions: types.Reactions{
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":smile:",
					"smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			newReaction: types.NewRegisteredReaction(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				":sad:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expAppend: true,
			expReaction: types.Reactions{
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":smile:",
					"smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":sad:",
					"smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
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

// ___________________________________________________________________________________________________________________

func TestPostReaction_String(t *testing.T) {
	reaction := types.NewPostReaction(
		":smile:",
		"reaction",
		"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
	)
	require.Equal(t, "[Shortcode] :smile: [Value] reaction [Owner] cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", reaction.String())
}

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

func TestPostReaction_Equals(t *testing.T) {
	tests := []struct {
		name          string
		first         types.PostReaction
		second        types.PostReaction
		shouldBeEqual bool
	}{
		{
			name: "Returns false with different user",
			first: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			second: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			),
			shouldBeEqual: false,
		},
		{
			name: "Returns false with different value",
			first: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			second: types.NewPostReaction(
				":smile:",
				"reactions",
				"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			),
			shouldBeEqual: false,
		},
		{
			name: "Returns false with different shortcode",
			first: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			second: types.NewPostReaction(
				":face:",
				"reaction",
				"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			),
			shouldBeEqual: false,
		},
		{
			name: "Returns true with the same data",
			first: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			second: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
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

// ___________________________________________________________________________________________________________________

func TestPostReactions_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name      string
		reactions types.PostReactions
		newLike   types.PostReaction
		expLikes  types.PostReactions
		expAppend bool
	}{
		{
			name:      "New reaction is appended properly to empty list",
			reactions: types.PostReactions{},
			newLike: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expLikes: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			expAppend: true,
		},
		{
			name: "New reaction is appended properly to existing list",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			newLike: types.NewPostReaction(
				":smile:",
				"reaction",
				"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			),
			expAppend: true,
			expLikes: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := test.reactions.AppendIfMissing(test.newLike)
			require.Equal(t, test.expLikes, actual)
			require.Equal(t, test.expAppend, appended)
		})
	}
}

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
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:       "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode:   ":smile:",
			expContains: true,
		},
		{
			name:        "Empty list returns false",
			reactions:   types.PostReactions{},
			owner:       "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode:   ":smile:",
			expContains: false,
		},
		{
			name: "Non-empty list returns false with not found address",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:       "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			shortcode:   ":smile:",
			expContains: false,
		},
		{
			name: "Non-empty list returns false with not found value",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
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
			reactions: types.PostReactions{
				types.NewPostReaction(
					":+1:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    ":+1:",
			expIndex: 0,
		},
		{
			name: "Non-empty list returns proper index with valid value (emoji - one code)",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":+1:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    "👍",
			expIndex: 0,
		},
		{
			name: "Non-empty list returns proper index with valid value (emoji - another code)",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":thumbsup:",
					"👍",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    "👍",
			expIndex: 0,
		},
		{
			name:      "Empty list returns -1",
			reactions: types.PostReactions{},
			owner:     "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name: "Non-empty list returns -1 with not found address",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:    "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			value:    "reaction",
			expIndex: -1,
		},
		{
			name: "Non-empty list returns -1 with not found value",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:    "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			value:    "reaction-2",
			expIndex: -1,
		},
		{
			name: "Existing reaction search by code",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":reaction:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    ":reaction:",
			expIndex: 0,
		},
		{
			name: "Exiting emoji reaction stored by value search by code",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":fire:",
					"🔥",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    ":fire:",
			expIndex: 0,
		},
		{
			name: "Exiting emoji reaction stored by code search by code",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":fire:",
					"🔥",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:    "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			value:    ":fire:",
			expIndex: 0,
		},
		{
			name: "Exiting emoji reaction stored by code search by value",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":fire:",
					"🔥",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
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
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:     "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode: ":smile:",
			expResult: types.PostReactions{},
			expEdited: true,
		},
		{
			name:      "Empty list is not edited",
			reactions: types.PostReactions{},
			owner:     "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			shortcode: ":smile:",
			expResult: types.PostReactions{},
			expEdited: false,
		},
		{
			name: "Non-empty list with not found address is not edited",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:     "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			shortcode: ":smile:",
			expResult: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			expEdited: false,
		},
		{
			name: "Non-empty list with not found value is not edited",
			reactions: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
			owner:     "cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae",
			shortcode: ":like:",
			expResult: types.PostReactions{
				types.NewPostReaction(
					":smile:",
					"reaction",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				),
			},
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
