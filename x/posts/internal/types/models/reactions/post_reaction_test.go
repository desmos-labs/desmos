package reactions_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/reactions"
	"github.com/stretchr/testify/require"
)

func TestPostReaction_String(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	reaction := reactions.NewPostReaction("reaction", user)
	require.Equal(t, `{"owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"reaction"}`, reaction.String())
}

func TestPostReaction_Validate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name     string
		reaction reactions.PostReaction
		error    error
	}{
		{
			name:     "Valid reaction returns no error",
			reaction: reactions.NewPostReaction("reaction", user),
			error:    nil,
		},
		{
			name:     "Missing owner returns error",
			reaction: reactions.NewPostReaction("reaction", nil),
			error:    errors.New("invalid reaction owner: "),
		},
		{
			name:     "Missing value returns error",
			reaction: reactions.NewPostReaction("", user),
			error:    errors.New("reaction value cannot be empty or blank"),
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
			first:         reactions.NewPostReaction("reaction", user),
			second:        reactions.NewPostReaction("reaction", otherLiker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns true with the same data",
			first:         reactions.NewPostReaction("reaction", user),
			second:        reactions.NewPostReaction("reaction", user),
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
			newLike:   reactions.NewPostReaction("reaction", user),
			expLikes:  reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			expAppend: true,
		},
		{
			name:      "New reaction is appended properly to existing list",
			reactions: reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			newLike:   reactions.NewPostReaction("reaction", otherLiker),
			expAppend: true,
			expLikes: reactions.PostReactions{
				reactions.NewPostReaction("reaction", user),
				reactions.NewPostReaction("reaction", otherLiker),
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
		value       string
		expContains bool
	}{
		{
			name:        "Non-empty list returns true with valid address",
			reactions:   reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			owner:       user,
			value:       "reaction",
			expContains: true,
		},
		{
			name:        "Empty list returns false",
			reactions:   reactions.PostReactions{},
			owner:       user,
			value:       "reaction",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found address",
			reactions:   reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			owner:       otherLiker,
			value:       "reaction",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found value",
			reactions:   reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			owner:       user,
			value:       "reaction-2",
			expContains: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expContains, test.reactions.ContainsReactionFrom(test.owner, test.value))
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
			reactions: reactions.PostReactions{reactions.NewPostReaction(":+1:", user)},
			owner:     user,
			value:     ":+1:",
			expIndex:  0,
		},
		{
			name:      "Non-empty list returns proper index with valid value (emoji - one code)",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":+1:", user)},
			owner:     user,
			value:     "👍",
			expIndex:  0,
		},
		{
			name:      "Non-empty list returns proper index with valid value (emoji - another code)",
			reactions: reactions.PostReactions{reactions.NewPostReaction(":thumbsup:", user)},
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
			reactions: reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			owner:     otherLiker,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found value",
			reactions: reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			owner:     otherLiker,
			value:     "reaction-2",
			expIndex:  -1,
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
		value     string
		expResult reactions.PostReactions
		expEdited bool
	}{
		{
			name:      "PostReaction is removed from non-empty list",
			reactions: reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			owner:     user,
			value:     "reaction",
			expResult: reactions.PostReactions{},
			expEdited: true,
		},
		{
			name:      "Empty list is not edited",
			reactions: reactions.PostReactions{},
			owner:     user,
			value:     "reaction",
			expResult: reactions.PostReactions{},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found address is not edited",
			reactions: reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			owner:     otherLiker,
			value:     "reaction",
			expResult: reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found value is not edited",
			reactions: reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			owner:     otherLiker,
			value:     "reaction-2",
			expResult: reactions.PostReactions{reactions.NewPostReaction("reaction", user)},
			expEdited: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result, edited := test.reactions.RemoveReaction(test.owner, test.value)
			require.Equal(t, test.expEdited, edited)
			require.Equal(t, test.expResult, result)
		})
	}
}
