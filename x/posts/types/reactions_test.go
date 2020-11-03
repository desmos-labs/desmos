package types_test

import (
	"errors"
	"github.com/desmos-labs/desmos/x/posts/types"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestReaction_Validate(t *testing.T) {
	testOwner, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name     string
		reaction types.RegisteredReaction
		error    error
	}{
		{
			name: "Valid reaction returns no error (url on value)",
			reaction: types.NewReaction(
				testOwner,
				":smile-jpg:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
		{
			name: "Invalid reaction returns error (unicode on value)",
			reaction: types.NewReaction(
				testOwner,
				":smiles:",
				"U+1F600",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
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
			//nolint - errcheck
			error: errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
		},
		{
			name: "Invalid short code returns error",
			reaction: types.NewReaction(
				testOwner,
				"smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			//nolint - errcheck
			error: errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
		},
		{
			name: "Empty value returns error",
			reaction: types.NewReaction(
				testOwner,
				":smile:",
				"",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "invalid value returns error (url)",
			reaction: types.NewReaction(
				testOwner,
				":smile:",
				"smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
		},
		{
			name: "invalid value returns error (unicode)",
			reaction: types.NewReaction(
				testOwner,
				":smile:",
				"U+1",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: errors.New("reaction value should be a URL"),
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
		first         types.RegisteredReaction
		second        types.RegisteredReaction
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

// ___________________________________________________________________________________________________________________

func TestReactions_AppendIfMissing(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

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

// ___________________________________________________________________________________________________________________

func TestPostReaction_String(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	reaction := types.NewPostReaction(":smile:", "reaction", user)
	require.Equal(t, "[Shortcode] :smile: [Value] reaction [Owner] cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", reaction.String())
}

func TestPostReaction_Validate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	//nolint - errcheck
	tests := []struct {
		name     string
		reaction types.PostReaction
		error    error
	}{
		{
			name:     "Valid reaction returns no error",
			reaction: types.NewPostReaction(":smile:", "reaction", user),
			error:    nil,
		},
		{
			name:     "Missing owner returns error",
			reaction: types.NewPostReaction(":smile:", "reaction", nil),
			error:    errors.New("invalid reaction owner: "),
		},
		{
			name:     "Missing value returns error",
			reaction: types.NewPostReaction(":smile:", "", user),
			error:    errors.New("reaction value cannot be empty or blank"),
		},
		{
			name:     "Invalid shortcode returns error",
			reaction: types.NewPostReaction("invalid", "reaction", user),
			error:    errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a ':'"),
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
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name          string
		first         types.PostReaction
		second        types.PostReaction
		shouldBeEqual bool
	}{
		{
			name:          "Returns false with different user",
			first:         types.NewPostReaction(":smile:", "reaction", user),
			second:        types.NewPostReaction(":smile:", "reaction", otherLiker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns false with different value",
			first:         types.NewPostReaction(":smile:", "reaction", user),
			second:        types.NewPostReaction(":smile:", "reactions", otherLiker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns false with different shortcode",
			first:         types.NewPostReaction(":smile:", "reaction", user),
			second:        types.NewPostReaction(":face:", "reaction", otherLiker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns true with the same data",
			first:         types.NewPostReaction(":smile:", "reaction", user),
			second:        types.NewPostReaction(":smile:", "reaction", user),
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
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

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
			newLike:   types.NewPostReaction(":smile:", "reaction", user),
			expLikes:  types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			expAppend: true,
		},
		{
			name:      "New reaction is appended properly to existing list",
			reactions: types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			newLike:   types.NewPostReaction(":smile:", "reaction", otherLiker),
			expAppend: true,
			expLikes: types.PostReactions{
				types.NewPostReaction(":smile:", "reaction", user),
				types.NewPostReaction(":smile:", "reaction", otherLiker),
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
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name        string
		reactions   types.PostReactions
		owner       sdk.AccAddress
		shortcode   string
		expContains bool
	}{
		{
			name:        "Non-empty list returns true with valid address",
			reactions:   types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			owner:       user,
			shortcode:   ":smile:",
			expContains: true,
		},
		{
			name:        "Empty list returns false",
			reactions:   types.PostReactions{},
			owner:       user,
			shortcode:   ":smile:",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found address",
			reactions:   types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			owner:       otherLiker,
			shortcode:   ":smile:",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found value",
			reactions:   types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			owner:       user,
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
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name      string
		reactions types.PostReactions
		owner     sdk.AccAddress
		value     string
		expIndex  int
	}{
		{
			name:      "Non-empty list returns proper index with valid value (shortcode)",
			reactions: types.PostReactions{types.NewPostReaction(":+1:", "👍", user)},
			owner:     user,
			value:     ":+1:",
			expIndex:  0,
		},
		{
			name:      "Non-empty list returns proper index with valid value (emoji - one code)",
			reactions: types.PostReactions{types.NewPostReaction(":+1:", "👍", user)},
			owner:     user,
			value:     "👍",
			expIndex:  0,
		},
		{
			name:      "Non-empty list returns proper index with valid value (emoji - another code)",
			reactions: types.PostReactions{types.NewPostReaction(":thumbsup:", "👍", user)},
			owner:     user,
			value:     "👍",
			expIndex:  0,
		},
		{
			name:      "Empty list returns -1",
			reactions: types.PostReactions{},
			owner:     user,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found address",
			reactions: types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			owner:     otherLiker,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found value",
			reactions: types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			owner:     otherLiker,
			value:     "reaction-2",
			expIndex:  -1,
		},
		{
			name:      "Existing reaction search by code",
			reactions: types.PostReactions{types.NewPostReaction(":reaction:", "reaction", user)},
			owner:     user,
			value:     ":reaction:",
			expIndex:  0,
		},
		{
			name:      "Exiting emoji reaction stored by value search by code",
			reactions: types.PostReactions{types.NewPostReaction(":fire:", "🔥", user)},
			owner:     user,
			value:     ":fire:",
			expIndex:  0,
		},
		{
			name:      "Exiting emoji reaction stored by code search by code",
			reactions: types.PostReactions{types.NewPostReaction(":fire:", "🔥", user)},
			owner:     user,
			value:     ":fire:",
			expIndex:  0,
		},
		{
			name:      "Exiting emoji reaction stored by code search by value",
			reactions: types.PostReactions{types.NewPostReaction(":fire:", "🔥", user)},
			owner:     user,
			value:     "🔥",
			expIndex:  0,
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
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name      string
		reactions types.PostReactions
		owner     sdk.AccAddress
		shortcode string
		expResult types.PostReactions
		expEdited bool
	}{
		{
			name:      "PostReaction is removed from non-empty list",
			reactions: types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			owner:     user,
			shortcode: ":smile:",
			expResult: types.PostReactions{},
			expEdited: true,
		},
		{
			name:      "Empty list is not edited",
			reactions: types.PostReactions{},
			owner:     user,
			shortcode: ":smile:",
			expResult: types.PostReactions{},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found address is not edited",
			reactions: types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			owner:     otherLiker,
			shortcode: ":smile:",
			expResult: types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found value is not edited",
			reactions: types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
			owner:     otherLiker,
			shortcode: ":like:",
			expResult: types.PostReactions{types.NewPostReaction(":smile:", "reaction", user)},
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
