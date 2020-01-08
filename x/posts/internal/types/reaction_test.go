package types_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/magiconair/properties/assert"
)

func TestReaction_String(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	reaction := types.NewReaction("reaction", 1, user)
	assert.Equal(t, `{"created":"1","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4","value":"reaction"}`, reaction.String())
}

func TestReaction_Validate(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	tests := []struct {
		name     string
		reaction types.Reaction
		error    error
	}{
		{
			name:     "Valid reaction returns no error",
			reaction: types.NewReaction("reaction", 10, user),
			error:    nil,
		},
		{
			name:     "Missing owner returns error",
			reaction: types.NewReaction("reaction", 10, nil),
			error:    errors.New("invalid reaction owner: "),
		},
		{
			name:     "Zero creation time returns error",
			reaction: types.NewReaction("reaction", 0, user),
			error:    errors.New("invalid reaction creation block height: 0"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.error, test.reaction.Validate())
		})
	}
}

func TestReaction_Equals(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name          string
		first         types.Reaction
		second        types.Reaction
		shouldBeEqual bool
	}{
		{
			name:          "Returns false with different creation time",
			first:         types.NewReaction("reaction", 5, user),
			second:        types.NewReaction("reaction", 6, user),
			shouldBeEqual: false,
		},
		{
			name:          "Returns false with different user",
			first:         types.NewReaction("reaction", 10, user),
			second:        types.NewReaction("reaction", 10, otherLiker),
			shouldBeEqual: false,
		},
		{
			name:   "Returns true with the same data",
			first:  types.NewReaction("reaction", 10, user),
			second: types.NewReaction("reaction", 10, user),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.shouldBeEqual, test.first.Equals(test.second))
		})
	}
}

func TestReactions_AppendIfMissing(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name      string
		reactions types.Reactions
		newLike   types.Reaction
		expLikes  types.Reactions
		expAppend bool
	}{
		{
			name:      "New reaction is appended properly to empty list",
			reactions: types.Reactions{},
			newLike:   types.NewReaction("reaction", 10, user),
			expLikes:  types.Reactions{types.NewReaction("reaction", 10, user)},
			expAppend: true,
		},
		{
			name:      "New reaction is appended properly to existing list",
			reactions: types.Reactions{types.NewReaction("reaction", 1, user)},
			newLike:   types.NewReaction("reaction", 10, otherLiker),
			expAppend: true,
			expLikes: types.Reactions{
				types.NewReaction("reaction", 1, user),
				types.NewReaction("reaction", 10, otherLiker),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := test.reactions.AppendIfMissing(test.newLike)
			assert.Equal(t, test.expLikes, actual)
			assert.Equal(t, test.expAppend, appended)
		})
	}
}

func TestReactions_ContainsOwnerLike(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name        string
		reactions   types.Reactions
		owner       sdk.AccAddress
		value       string
		expContains bool
	}{
		{
			name:        "Non-empty list returns true with valid address",
			reactions:   types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:       user,
			value:       "reaction",
			expContains: true,
		},
		{
			name:        "Empty list returns false",
			reactions:   types.Reactions{},
			owner:       user,
			value:       "reaction",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found address",
			reactions:   types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:       otherLiker,
			value:       "reaction",
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found value",
			reactions:   types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:       user,
			value:       "reaction-2",
			expContains: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expContains, test.reactions.ContainsReactionFrom(test.owner, test.value))
		})
	}
}

func TestReactions_IndexOfByUserAndValue(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name      string
		reactions types.Reactions
		owner     sdk.AccAddress
		value     string
		expIndex  int
	}{
		{
			name:      "Non-empty list returns proper index with valid value",
			reactions: types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:     user,
			value:     "reaction",
			expIndex:  0,
		},
		{
			name:      "Empty list returns -1",
			reactions: types.Reactions{},
			owner:     user,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found address",
			reactions: types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:     otherLiker,
			value:     "reaction",
			expIndex:  -1,
		},
		{
			name:      "Non-empty list returns -1 with not found value",
			reactions: types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:     otherLiker,
			value:     "reaction-2",
			expIndex:  -1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expIndex, test.reactions.IndexOfByUserAndValue(test.owner, test.value))
		})
	}
}

func TestReactions_RemoveReaction(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name      string
		reactions types.Reactions
		owner     sdk.AccAddress
		value     string
		expResult types.Reactions
		expEdited bool
	}{
		{
			name:      "Reaction is removed from non-empty list",
			reactions: types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:     user,
			value:     "reaction",
			expResult: types.Reactions{},
			expEdited: true,
		},
		{
			name:      "Empty list is not edited",
			reactions: types.Reactions{},
			owner:     user,
			value:     "reaction",
			expResult: types.Reactions{},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found address is not edited",
			reactions: types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:     otherLiker,
			value:     "reaction",
			expResult: types.Reactions{types.NewReaction("reaction", 1, user)},
			expEdited: false,
		},
		{
			name:      "Non-empty list with not found value is not edited",
			reactions: types.Reactions{types.NewReaction("reaction", 1, user)},
			owner:     otherLiker,
			value:     "reaction-2",
			expResult: types.Reactions{types.NewReaction("reaction", 1, user)},
			expEdited: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result, edited := test.reactions.RemoveReaction(test.owner, test.value)
			assert.Equal(t, test.expEdited, edited)
			assert.Equal(t, test.expResult, result)
		})
	}
}
