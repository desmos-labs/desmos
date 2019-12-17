package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

// -------------
// --- PostID
// -------------

func TestPostID_MarshalJSON(t *testing.T) {
	json := types.ModuleCdc.MustMarshalJSON(types.PostID(10))
	assert.Equal(t, `"10"`, string(json))
}

func TestPostID_UnmarshalJSON(t *testing.T) {
	var id types.PostID
	types.ModuleCdc.MustUnmarshalJSON([]byte(`"10"`), &id)
	assert.Equal(t, types.PostID(10), id)
}

// -----------
// --- Post
// -----------

func TestPost_String(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	post := types.Post{
		PostID:            types.PostID(19),
		ParentID:          types.PostID(1),
		Message:           "My post message",
		Created:           sdk.NewInt(98),
		LastEdited:        sdk.NewInt(105),
		AllowsComments:    true,
		ExternalReference: "My reference",
		Owner:             owner,
	}

	assert.Equal(t,
		`{"id":"19","parent_id":"1","message":"My post message","created":"98","last_edited":"105","allows_comments":true,"external_reference":"My reference","owner":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`,
		post.String(),
	)
}

func TestPost_Validate(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	tests := []struct {
		post     types.Post
		expError string
	}{
		{
			post:     types.Post{PostID: types.PostID(0)},
			expError: "invalid post id: 0",
		},
		{
			post:     types.Post{PostID: types.PostID(19), Owner: nil},
			expError: "invalid post owner: ",
		},
		{
			post:     types.Post{PostID: types.PostID(19), Owner: owner, Message: ""},
			expError: "invalid post message: ",
		},
		{
			post:     types.Post{PostID: types.PostID(19), Owner: owner, Message: "Message", Created: sdk.NewInt(0)},
			expError: "invalid post creation block height: 0",
		},
		{
			post:     types.Post{PostID: types.PostID(19), Owner: owner, Message: "Message", Created: sdk.NewInt(10), LastEdited: sdk.NewInt(9)},
			expError: "invalid post last edit block height: 9",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.expError, func(t *testing.T) {
			assert.Equal(t, test.expError, test.post.Validate().Error())
		})
	}
}

func TestPost_Equals(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	otherOwner, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name      string
		first     types.Post
		second    types.Post
		expEquals bool
	}{
		{
			name: "Different post ID",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(10),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			expEquals: false,
		},
		{
			name: "Different parent ID",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(10),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			expEquals: false,
		},
		{
			name: "Different message",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "Another post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			expEquals: false,
		},
		{
			name: "Different creation time",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(15),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			expEquals: false,
		},
		{
			name: "Different last edited",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(13),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			expEquals: false,
		},
		{
			name: "Different allows comments",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    false,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			expEquals: false,
		},
		{
			name: "Different owner",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             otherOwner,
			},
			expEquals: false,
		},
		{
			name: "Different external reference",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference 1",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference 2",
				Owner:             owner,
			},
			expEquals: false,
		},
		{
			name: "Same data",
			first: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			second: types.Post{
				PostID:            types.PostID(19),
				ParentID:          types.PostID(1),
				Message:           "My post message",
				Created:           sdk.NewInt(98),
				LastEdited:        sdk.NewInt(105),
				AllowsComments:    true,
				ExternalReference: "My reference",
				Owner:             owner,
			},
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

// -----------
// --- Posts
// -----------

func TestPosts_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.Posts
		second    types.Posts
		expEquals bool
	}{
		{
			name:      "Empty lists are equals",
			first:     types.Posts{},
			second:    types.Posts{},
			expEquals: true,
		},
		{
			name: "List of different lengths are not equals",
			first: types.Posts{
				types.Post{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			second: types.Posts{
				types.Post{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.Post{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			expEquals: false,
		},
		{
			name: "Same lists but in different orders",
			first: types.Posts{
				types.Post{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.Post{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			second: types.Posts{
				types.Post{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.Post{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			expEquals: false,
		},
		{
			name: "Same lists are equals",
			first: types.Posts{
				types.Post{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.Post{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			second: types.Posts{
				types.Post{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.Post{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

func TestPosts_String(t *testing.T) {
	owner1, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	owner2, _ := sdk.AccAddressFromBech32("cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l")
	posts := types.Posts{
		types.NewPost(types.PostID(1), types.PostID(10), "Post 1", false, "external-ref-1", 0, owner1),
		types.NewPost(types.PostID(2), types.PostID(10), "Post 2", false, "external-ref-1", 0, owner2),
	}

	expected := `ID - [Creator] Message
1 - [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns] Post 1
2 - [cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l] Post 2`
	assert.Equal(t, expected, posts.String())
}
