package reactions_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/types/models/reactions"
	"github.com/stretchr/testify/require"
)

func TestPostReaction_String(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	reaction := reactions.NewPostReaction(":smile:", "reaction", user)
	require.Equal(t, `{"shortcode":":smile:","value":"reaction","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}`, reaction.String())
}

func TestPostReaction_Validate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	//nolint - errcheck
	tests := []struct {
		name     string
		reaction reactions.PostReaction
		error    error
	}{
		{
			name:     "Valid reaction returns no error",
			reaction: reactions.NewPostReaction(":smile:", "reaction", user),
			error:    nil,
		},
		{
			name:     "Missing owner returns error",
			reaction: reactions.NewPostReaction(":smile:", "reaction", nil),
			error:    errors.New("invalid reaction owner: "),
		},
		{
			name:     "Missing value returns error",
			reaction: reactions.NewPostReaction(":smile:", "", user),
			error:    errors.New("reaction value cannot be empty or blank"),
		},
		{
			name:     "Invalid shortcode returns error",
			reaction: reactions.NewPostReaction("invalid", "reaction", user),
			error:    errors.New("the specified shortcode is not valid. To be valid it must only contains a-z, 0-9, - and _ and must start and end with a :"),
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
		first         reactions.PostReaction
		second        reactions.PostReaction
		shouldBeEqual bool
	}{
		{
			name:          "Returns false with different user",
			first:         reactions.NewPostReaction(":smile:", "reaction", user),
			second:        reactions.NewPostReaction(":smile:", "reaction", otherLiker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns false with different value",
			first:         reactions.NewPostReaction(":smile:", "reaction", user),
			second:        reactions.NewPostReaction(":smile:", "reactions", otherLiker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns false with different shortcode",
			first:         reactions.NewPostReaction(":smile:", "reaction", user),
			second:        reactions.NewPostReaction(":face:", "reaction", otherLiker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns true with the same data",
			first:         reactions.NewPostReaction(":smile:", "reaction", user),
			second:        reactions.NewPostReaction(":smile:", "reaction", user),
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

func TestPostReactions_AppendIfMissing(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name      string
		reactions reactions.PostReactions
		newLike   reactions.PostReaction
		expLikes  reactions.PostReactions
		expAppend bool
	}{
		{
			name:      "New reaction is appended properly to empty list",
			reactions: reactions.PostReactions{},
			newLike:   reactions.NewPostReaction(":smile:", "reaction", user),
			expLikes:  reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			expAppend: true,
		},
		{
			name:      "New reaction is appended properly to existing list",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			newLike:   reactions.NewPostReaction(":smile:", "reaction", otherLiker),
			expAppend: true,
			expLikes: reactions.PostReactions{
				reactions.NewPostReaction(":smile:", "reaction", user),
				reactions.NewPostReaction(":smile:", "reaction", otherLiker),
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
		reactions   reactions.PostReactions
		owner       sdk.AccAddress
		shortcode   string
		expContains bool
	}{
		{
			name:        "Non-empty list returns true with valid address",
			reactions:   reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			owner:       user,
			shortcode:   ":smile:",
			expContains: true,
		},
		{
			name:        "Empty list returns false",
			reactions:   reactions.PostReactions{},
			owner:       user,
			shortcode:   ":smile:",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found address",
			reactions:   reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			owner:       otherLiker,
			shortcode:   ":smile:",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found value",
			reactions:   reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
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
		reactions reactions.PostReactions
		owner     sdk.AccAddress
		value     string
		expIndex  int
	}{
		{
			name:      "Non-empty list returns proper index with valid value (shortcode)",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":+1:", "👍", user)},
			owner:     user,
			value:     ":+1:",
			expIndex:  0,
		},
		{
			name:      "Non-empty list returns proper index with valid value (emoji - one code)",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":+1:", "👍", user)},
			owner:     user,
			value:     "👍",
			expIndex:  0,
		},
		{
			name:      "Non-empty list returns proper index with valid value (emoji - another code)",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":thumbsup:", "👍", user)},
			owner:     user,
			value:     "👍",
			expIndex:  0,
		},
		{
			name:      "Empty list returns -1",
			reactions: reactions.PostReactions{},
			owner:     user,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found address",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			owner:     otherLiker,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found value",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			owner:     otherLiker,
			value:     "reaction-2",
			expIndex:  -1,
		},
		{
			name:      "Existing reaction search by code",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":reaction:", "reaction", user)},
			owner:     user,
			value:     ":reaction:",
			expIndex:  0,
		},
		{
			name:      "Exiting emoji reaction stored by value search by code",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":fire:", "🔥", user)},
			owner:     user,
			value:     ":fire:",
			expIndex:  0,
		},
		{
			name:      "Exiting emoji reaction stored by code search by code",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":fire:", "🔥", user)},
			owner:     user,
			value:     ":fire:",
			expIndex:  0,
		},
		{
			name:      "Exiting emoji reaction stored by code search by value",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":fire:", "🔥", user)},
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
		reactions reactions.PostReactions
		owner     sdk.AccAddress
		shortcode string
		expResult reactions.PostReactions
		expEdited bool
	}{
		{
			name:      "PostReaction is removed from non-empty list",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			owner:     user,
			shortcode: ":smile:",
			expResult: reactions.PostReactions{},
			expEdited: true,
		},
		{
			name:      "Empty list is not edited",
			reactions: reactions.PostReactions{},
			owner:     user,
			shortcode: ":smile:",
			expResult: reactions.PostReactions{},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found address is not edited",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			owner:     otherLiker,
			shortcode: ":smile:",
			expResult: reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found value is not edited",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
			owner:     otherLiker,
			shortcode: ":like:",
			expResult: reactions.PostReactions{reactions.NewPostReaction(":smile:", "reaction", user)},
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
