package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

// -------------
// --- PostID
// -------------

func TestPostID_Next(t *testing.T) {
	tests := []struct {
		id       types.PostID
		expected types.PostID
	}{
		{
			id:       types.PostID(0),
			expected: types.PostID(1),
		},
		{
			id:       types.PostID(1123123),
			expected: types.PostID(1123124),
		},
	}

	for index, test := range tests {
		test := test
		t.Run(fmt.Sprintf("Test index: %d", index), func(t *testing.T) {
			assert.Equal(t, test.expected, test.id.Next())
		})
	}
}

func TestPostID_MarshalJSON(t *testing.T) {
	json := types.ModuleCdc.MustMarshalJSON(types.PostID(10))
	assert.Equal(t, `"10"`, string(json))
}

func TestPostID_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expID    types.PostID
		expError string
	}{
		{
			name:     "Invalid ID returns error",
			value:    "id",
			expID:    types.PostID(0),
			expError: "invalid character 'i' looking for beginning of value",
		},
		{
			name:     "Valid id is read properly",
			value:    `"10"`,
			expID:    types.PostID(10),
			expError: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var id types.PostID
			err := types.ModuleCdc.UnmarshalJSON([]byte(test.value), &id)

			if err == nil {
				assert.Equal(t, test.expID, id)
			} else {
				assert.Equal(t, test.expError, err.Error())
			}
		})
	}
}

func TestParsePostID(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expID    types.PostID
		expError string
	}{
		{
			name:     "Invalid id returns error",
			value:    "id",
			expID:    types.PostID(0),
			expError: "strconv.ParseUint: parsing \"id\": invalid syntax",
		},
		{
			name:     "Valid id returns proper value",
			value:    "10",
			expID:    types.PostID(10),
			expError: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			id, err := types.ParsePostID(test.value)

			if err == nil {
				assert.Equal(t, test.expID, id)
			} else {
				assert.Equal(t, test.expError, err.Error())
			}
		})
	}
}

func TestPostIDs_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.PostIDs
		second    types.PostIDs
		expEquals bool
	}{
		{
			name:      "Different length",
			first:     types.PostIDs{types.PostID(1), types.PostID(0)},
			second:    types.PostIDs{types.PostID(1)},
			expEquals: false,
		},
		{
			name:      "Different order",
			first:     types.PostIDs{types.PostID(0), types.PostID(1)},
			second:    types.PostIDs{types.PostID(1), types.PostID(0)},
			expEquals: false,
		},
		{
			name:      "Same length and order",
			first:     types.PostIDs{types.PostID(0), types.PostID(1)},
			second:    types.PostIDs{types.PostID(0), types.PostID(1)},
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

func TestPostIDs_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name      string
		IDs       types.PostIDs
		newID     types.PostID
		expIDs    types.PostIDs
		expEdited bool
	}{
		{
			name:      "AppendIfMissing dont append anything",
			IDs:       types.PostIDs{types.PostID(1)},
			newID:     types.PostID(1),
			expIDs:    types.PostIDs{types.PostID(1)},
			expEdited: false,
		},
		{
			name:      "AppendIfMissing append something",
			IDs:       types.PostIDs{types.PostID(1)},
			newID:     types.PostID(2),
			expIDs:    types.PostIDs{types.PostID(1), types.PostID(2)},
			expEdited: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newIDs, edited := test.IDs.AppendIfMissing(test.newID)
			assert.Equal(t, test.expIDs, newIDs)
			assert.Equal(t, test.expEdited, edited)
		})
	}
}

func TestPost_GetID(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.GetID()

	assert.Equal(t, types.PostID(2), actual)
}

func TestPost_GetParentID(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.GetParentID()

	assert.Equal(t, types.PostID(0), actual)
}

func TestPost_SetMessage(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.SetMessage("edited media post")

	assert.Equal(t, "edited media post", actual.GetMessage())
}

func TestPost_GetMessage(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.GetMessage()

	assert.Equal(t, "media Post", actual)

}

func TestPost_CreationTime(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.CreationTime()

	assert.Equal(t, sdk.NewInt(0), actual)
}

func TestPost_SetEditTime(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.SetEditTime(sdk.NewInt(1))

	assert.Equal(t, sdk.NewInt(1), actual.GetEditTime())
}

func TestPost_GetEditTime(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.GetEditTime()

	assert.Equal(t, sdk.NewInt(0), actual)
}

func TestPost_CanComment(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.CanComment()

	assert.Equal(t, false, actual)
}

func TestPost_GetSubspace(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.GetSubspace()

	assert.Equal(t, "desmos", actual)
}

func TestPost_GetOptionalData(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{"key1": "value1"}, 0, owner)

	actual := post.GetOptionalData()

	assert.Equal(t, map[string]string{"key1": "value1"}, actual)
}

func TestPost_Owner(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	var post = types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, 0, owner)

	actual := post.Owner()

	assert.Equal(t, owner, actual)
}

// -----------
// --- TextPost
// -----------

func TestPost_String(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	post := types.TextPost{
		PostID:         types.PostID(19),
		ParentID:       types.PostID(1),
		Message:        "My post message",
		Created:        sdk.NewInt(98),
		LastEdited:     sdk.NewInt(105),
		AllowsComments: true,
		Subspace:       "desmos",
		OptionalData:   map[string]string{},
		Creator:        owner,
	}

	assert.Equal(t,
		`{"id":"19","parent_id":"1","message":"My post message","created":"98","last_edited":"105","allows_comments":true,"subspace":"desmos","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`,
		post.String(),
	)
}

func TestPost_Validate(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	tests := []struct {
		post     types.TextPost
		expError string
	}{
		{
			post:     types.NewTextPost(types.PostID(0), types.PostID(0), "Message", true, "Desmos", map[string]string{}, 10, owner),
			expError: "invalid post id: 0",
		},
		{
			post:     types.NewTextPost(types.PostID(1), types.PostID(0), "", true, "Desmos", map[string]string{}, 10, nil),
			expError: "invalid post owner: ",
		},
		{
			post:     types.NewTextPost(types.PostID(1), types.PostID(0), "", true, "Desmos", map[string]string{}, 10, owner),
			expError: "post message must be non empty and non blank",
		},
		{
			post:     types.NewTextPost(types.PostID(1), types.PostID(0), " ", true, "Desmos", map[string]string{}, 10, owner),
			expError: "post message must be non empty and non blank",
		},
		{
			post:     types.NewTextPost(types.PostID(1), types.PostID(0), "Message", true, "Desmos", map[string]string{}, 0, owner),
			expError: "invalid post creation block height: 0",
		},
		{
			post:     types.TextPost{PostID: types.PostID(19), Creator: owner, Message: "Message", Subspace: "desmos", Created: sdk.NewInt(10), LastEdited: sdk.NewInt(9)},
			expError: "invalid post last edit block height: 9",
		},
		{
			post:     types.NewTextPost(types.PostID(1), types.PostID(0), "Message", true, "", map[string]string{}, 1, owner),
			expError: "post subspace must be non empty and non blank",
		},
		{
			post:     types.NewTextPost(types.PostID(1), types.PostID(0), "Message", true, " ", map[string]string{}, 1, owner),
			expError: "post subspace must be non empty and non blank",
		},
		{
			post: types.NewTextPost(
				types.PostID(1),
				types.PostID(0),
				"Message",
				true,
				"desmos",
				map[string]string{
					"key1":  "value",
					"key2":  "value",
					"key3":  "value",
					"key4":  "value",
					"key5":  "value",
					"key6":  "value",
					"key7":  "value",
					"key8":  "value",
					"key9":  "value",
					"key10": "value",
					"key11": "value",
				},
				1,
				owner,
			),
			expError: "post optional data cannot contain more than 10 key-value pairs",
		},
		{
			post: types.NewTextPost(
				types.PostID(1),
				types.PostID(0),
				"Message",
				true,
				"desmos",
				map[string]string{
					"key1": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque euismod, mi at commodo efficitur, quam sapien congue enim, ut porttitor lacus tellus vitae turpis. Vivamus aliquam sem eget neque metus.",
				},
				1,
				owner,
			),
			expError: "post optional data values cannot exceed 200 characters. key1 of post with id 1 is longer than this",
		},
		{
			post:     types.NewTextPost(types.PostID(1), types.PostID(0), "Message", true, "Desmos", map[string]string{}, 1, owner),
			expError: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.expError, func(t *testing.T) {
			if len(test.expError) != 0 {
				assert.Equal(t, test.expError, test.post.Validate().Error())
			} else {
				assert.Nil(t, test.post.Validate())
			}
		})
	}
}

func TestPost_Equals(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	otherOwner, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name      string
		first     types.TextPost
		second    types.TextPost
		expEquals bool
	}{
		{
			name: "Different post ID",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(10),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			expEquals: false,
		},
		{
			name: "Different parent ID",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(10),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			expEquals: false,
		},
		{
			name: "Different message",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "Another post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			expEquals: false,
		},
		{
			name: "Different creation time",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(15),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			expEquals: false,
		},
		{
			name: "Different last edited",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(13),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			expEquals: false,
		},
		{
			name: "Different allows comments",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: false,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			expEquals: false,
		},
		{
			name: "Different subspace",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos-1",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos-2",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			expEquals: false,
		},
		{
			name: "Different optional data",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData: map[string]string{
					"field": "value",
				},
				Creator: owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData: map[string]string{
					"field": "other-value",
				},
				Creator: owner,
			},
			expEquals: false,
		},
		{
			name: "Different owner",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        otherOwner,
			},
			expEquals: false,
		},
		{
			name: "Same data",
			first: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
			},
			second: types.TextPost{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        sdk.NewInt(98),
				LastEdited:     sdk.NewInt(105),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
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
// --- TextPosts
// -----------
func TestPosts_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.TextPosts
		second    types.TextPosts
		expEquals bool
	}{
		{
			name:      "Empty lists are equals",
			first:     types.TextPosts{},
			second:    types.TextPosts{},
			expEquals: true,
		},
		{
			name: "List of different lengths are not equals",
			first: types.TextPosts{
				types.TextPost{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			second: types.TextPosts{
				types.TextPost{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.TextPost{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			expEquals: false,
		},
		{
			name: "Same lists but in different orders",
			first: types.TextPosts{
				types.TextPost{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.TextPost{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			second: types.TextPosts{
				types.TextPost{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.TextPost{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			expEquals: false,
		},
		{
			name: "Same lists are equals",
			first: types.TextPosts{
				types.TextPost{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.TextPost{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
			},
			second: types.TextPosts{
				types.TextPost{PostID: types.PostID(0), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
				types.TextPost{PostID: types.PostID(1), Created: sdk.ZeroInt(), LastEdited: sdk.ZeroInt()},
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
	posts := types.TextPosts{
		types.NewTextPost(types.PostID(1), types.PostID(10), "TextPost 1", false, "external-ref-1", map[string]string{}, 0, owner1),
		types.NewTextPost(types.PostID(2), types.PostID(10), "TextPost 2", false, "external-ref-1", map[string]string{}, 0, owner2),
	}

	expected := `ID - [Creator] Message
1 - [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns] TextPost 1
2 - [cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l] TextPost 2`
	assert.Equal(t, expected, posts.String())
}
