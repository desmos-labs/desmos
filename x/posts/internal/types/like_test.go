package types_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/magiconair/properties/assert"
)

func TestLike_String(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	like := types.NewLike(1, liker)
	assert.Equal(t, `{"created":"1","owner":"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}`, like.String())
}

func TestLike_Validate(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	tests := []struct {
		name  string
		like  types.Like
		error error
	}{
		{
			name:  "Valid like returns no error",
			like:  types.NewLike(10, liker),
			error: nil,
		},
		{
			name:  "Missing owner returns error",
			like:  types.NewLike(10, nil),
			error: errors.New("invalid like owner: "),
		},
		{
			name:  "Zero creation time returns error",
			like:  types.NewLike(0, liker),
			error: errors.New("invalid like creation block height: 0"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.error, test.like.Validate())
		})
	}
}

func TestLike_Equals(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name          string
		first         types.Like
		second        types.Like
		shouldBeEqual bool
	}{
		{
			name:          "Returns false with different creation time",
			first:         types.NewLike(5, liker),
			second:        types.NewLike(6, liker),
			shouldBeEqual: false,
		},
		{
			name:          "Returns false with different liker",
			first:         types.NewLike(10, liker),
			second:        types.NewLike(10, otherLiker),
			shouldBeEqual: false,
		},
		{
			name:   "Returns true with the same data",
			first:  types.NewLike(10, liker),
			second: types.NewLike(10, liker),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.shouldBeEqual, test.first.Equals(test.second))
		})
	}
}

func TestLikes_AppendIfMissing(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name      string
		likes     types.Likes
		newLike   types.Like
		expLikes  types.Likes
		expAppend bool
	}{
		{
			name:      "New like is appended properly to empty list",
			likes:     types.Likes{},
			newLike:   types.NewLike(10, liker),
			expLikes:  types.Likes{types.NewLike(10, liker)},
			expAppend: true,
		},
		{
			name:      "New like is appended properly to existing list",
			likes:     types.Likes{types.NewLike(1, liker)},
			newLike:   types.NewLike(10, otherLiker),
			expAppend: true,
			expLikes: types.Likes{
				types.NewLike(1, liker),
				types.NewLike(10, otherLiker),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := test.likes.AppendIfMissing(test.newLike)
			assert.Equal(t, test.expLikes, actual)
			assert.Equal(t, test.expAppend, appended)
		})
	}
}

func TestLikes_ContainsOwnerLike(t *testing.T) {
	liker, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherLiker, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name        string
		likes       types.Likes
		owner       sdk.AccAddress
		expContains bool
	}{
		{
			name:        "Non-empty list returns true with valid address",
			likes:       types.Likes{types.NewLike(1, liker)},
			owner:       liker,
			expContains: true,
		},
		{
			name:        "Empty list returns false",
			likes:       types.Likes{},
			owner:       liker,
			expContains: false,
		},
		{
			name:        "Non-empty list returns false with not found address",
			likes:       types.Likes{types.NewLike(1, liker)},
			owner:       otherLiker,
			expContains: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expContains, test.likes.ContainsOwnerLike(test.owner))
		})
	}
}
