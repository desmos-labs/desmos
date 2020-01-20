package types_test

import (
	"fmt"
	"testing"
	"time"

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

// -------------
// --- PostIDs
// -------------

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

// -----------
// --- Post
// -----------

func TestPost_String(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	timeZone, _ := time.LoadLocation("UTC")
	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	post := types.Post{
		PostID:         types.PostID(19),
		ParentID:       types.PostID(1),
		Message:        "My post message",
		Created:        date,
		LastEdited:     date.AddDate(0, 0, 1),
		AllowsComments: true,
		Subspace:       "desmos",
		OptionalData:   map[string]string{},
		Creator:        owner,
	}

	assert.Equal(t,
		`{"id":"19","parent_id":"1","message":"My post message","created":"2020-01-01T12:00:00Z","last_edited":"2020-01-02T12:00:00Z","allows_comments":true,"subspace":"desmos","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":null}`,
		post.String(),
	)
}

func TestPost_Validate(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	timeZone, _ := time.LoadLocation("UTC")
	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	tests := []struct {
		post     types.Post
		expError string
	}{
		{
			post:     types.NewPost(types.PostID(0), types.PostID(0), "Message", true, "Desmos", map[string]string{}, date, owner, medias),
			expError: "invalid post id: 0",
		},
		{
			post:     types.NewPost(types.PostID(1), types.PostID(0), "", true, "Desmos", map[string]string{}, date, nil, medias),
			expError: "invalid post owner: ",
		},
		{
			post:     types.NewPost(types.PostID(1), types.PostID(0), "", true, "Desmos", map[string]string{}, date, owner, medias),
			expError: "post message must be non empty and non blank",
		},
		{
			post:     types.NewPost(types.PostID(1), types.PostID(0), " ", true, "Desmos", map[string]string{}, date, owner, medias),
			expError: "post message must be non empty and non blank",
		},
		{
			post:     types.NewPost(types.PostID(1), types.PostID(0), "Message", true, "Desmos", map[string]string{}, time.Time{}, owner, medias),
			expError: "invalid post creation time: 0001-01-01 00:00:00 +0000 UTC",
		},
		{
			post:     types.Post{PostID: types.PostID(19), Creator: owner, Message: "Message", Subspace: "desmos", Created: date, LastEdited: date.AddDate(0, 0, -1)},
			expError: "invalid post last edit time: 2019-12-31 12:00:00 +0000 UTC",
		},
		{
			post:     types.NewPost(types.PostID(1), types.PostID(0), "Message", true, "", map[string]string{}, date, owner, medias),
			expError: "post subspace must be non empty and non blank",
		},
		{
			post:     types.NewPost(types.PostID(1), types.PostID(0), "Message", true, " ", map[string]string{}, date, owner, medias),
			expError: "post subspace must be non empty and non blank",
		},
		{
			post: types.Post{
				PostID:         types.PostID(1),
				ParentID:       types.PostID(0),
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Created:        time.Now().UTC().Add(time.Hour),
				Creator:        owner,
				Medias:         medias,
			},
			expError: "post creation date cannot be in the future",
		},
		{
			post: types.Post{
				PostID:         types.PostID(1),
				ParentID:       types.PostID(0),
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Created:        time.Now().UTC(),
				LastEdited:     time.Now().UTC().Add(time.Hour),
				Creator:        owner,
				Medias:         medias,
			},
			expError: "post last edit date cannot be in the future",
		},
		{
			post: types.NewPost(
				types.PostID(1),
				types.PostID(0),
				`
				Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque massa felis, aliquam sed ipsum at, 
				mollis pharetra quam. Vestibulum nec nulla ante. Praesent sed dignissim turpis. Curabitur aliquam nunc 
				eu nisi porta, eu gravida purus faucibus. Duis commodo sagittis lacus, vitae luctus enim vulputate a. 
				Nulla tempor eget nunc vitae vulputate. Nulla facilities. Donec sollicitudin odio in arcu efficitur, 
				sit amet vestibulum diam ullamcorper. Ut ac dolor in velit gravida efficitur et et erat volutpat.
				`,
				true,
				"desmos",
				map[string]string{},
				date,
				owner,
				medias,
			),
			expError: "post message cannot be longer than 500 characters",
		},
		{
			post: types.NewPost(
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
				date,
				owner,
				medias,
			),
			expError: "post optional data cannot contain more than 10 key-value pairs",
		},
		{
			post: types.NewPost(
				types.PostID(1),
				types.PostID(0),
				"Message",
				true,
				"desmos",
				map[string]string{
					"key1": `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque euismod, mi at commodo 
							efficitur, quam sapien congue enim, ut porttitor lacus tellus vitae turpis. Vivamus aliquam 
							sem eget neque metus.`,
				},
				date,
				owner,
				medias,
			),
			expError: "post optional data values cannot exceed 200 characters. key1 of post with id 1 is longer than this",
		},
		{
			post: types.NewPost(
				types.PostID(1),
				types.PostID(0),
				"Message",
				true,
				"Desmos",
				map[string]string{},
				date,
				owner,
				types.PostMedias{
					types.PostMedia{
						URI:      "",
						MimeType: "text/plain",
					},
				},
			),
			expError: "uri must be specified and cannot be empty",
		},
		{
			post: types.NewPost(
				types.PostID(1),
				types.PostID(0),
				"Message",
				true,
				"Desmos",
				map[string]string{},
				date,
				owner,
				types.PostMedias{
					types.PostMedia{
						URI:      "https://example.com",
						MimeType: "",
					},
				},
			),
			expError: "mime type must be specified and cannot be empty",
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

	timeZone, _ := time.LoadLocation("UTC")
	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	tests := []struct {
		name      string
		first     types.Post
		second    types.Post
		expEquals bool
	}{
		{
			name: "Different post ID",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(10),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different parent ID",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(10),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different message",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "Another post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different creation time",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date.AddDate(0, 0, 1),
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different last edited",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 2),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different allows comments",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: false,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different subspace",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos-1",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos-2",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different optional data",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData: map[string]string{
					"field": "value",
				},
				Creator: owner,
				Medias:  medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData: map[string]string{
					"field": "other-value",
				},
				Creator: owner,
				Medias:  medias,
			},
			expEquals: false,
		},
		{
			name: "Different owner",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        otherOwner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different medias",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias: types.PostMedias{
					types.PostMedia{
						URI:      "uri2",
						MimeType: "text/plain",
					},
				},
			},
			expEquals: false,
		},
		{
			name: "Same data",
			first: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         types.PostID(19),
				ParentID:       types.PostID(1),
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
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
	timeZone, _ := time.LoadLocation("UTC")
	date := time.Date(2020, 1, 1, 12, 0, 00, 000, timeZone)

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
				types.Post{PostID: types.PostID(0), Created: date, LastEdited: date.AddDate(0, 0, 1)},
			},
			second: types.Posts{
				types.Post{PostID: types.PostID(0), Created: date, LastEdited: date.AddDate(0, 0, 1)},
				types.Post{PostID: types.PostID(1), Created: date, LastEdited: date.AddDate(0, 0, 1)},
			},
			expEquals: false,
		},
		{
			name: "Same lists but in different orders",
			first: types.Posts{
				types.Post{PostID: types.PostID(0), Created: date, LastEdited: date.AddDate(0, 0, 1)},
				types.Post{PostID: types.PostID(1), Created: date, LastEdited: date.AddDate(0, 0, 1)},
			},
			second: types.Posts{
				types.Post{PostID: types.PostID(1), Created: date, LastEdited: date.AddDate(0, 0, 1)},
				types.Post{PostID: types.PostID(0), Created: date, LastEdited: date.AddDate(0, 0, 1)},
			},
			expEquals: false,
		},
		{
			name: "Same lists are equals",
			first: types.Posts{
				types.Post{PostID: types.PostID(0), Created: date, LastEdited: date.AddDate(0, 0, 1)},
				types.Post{PostID: types.PostID(1), Created: date, LastEdited: date.AddDate(0, 0, 1)},
			},
			second: types.Posts{
				types.Post{PostID: types.PostID(0), Created: date, LastEdited: date.AddDate(0, 0, 1)},
				types.Post{PostID: types.PostID(1), Created: date, LastEdited: date.AddDate(0, 0, 1)},
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
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	timeZone, _ := time.LoadLocation("UTC")
	date := time.Date(2020, 1, 1, 12, 0, 00, 000, timeZone)

	posts := types.Posts{
		types.NewPost(types.PostID(1), types.PostID(10), "Post 1", false, "external-ref-1", map[string]string{}, date, owner1, medias),
		types.NewPost(types.PostID(2), types.PostID(10), "Post 2", false, "external-ref-1", map[string]string{}, date, owner2, medias),
	}

	expected := `ID - [Creator] Message
1 - [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns] Post 1
2 - [cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l] Post 2`
	assert.Equal(t, expected, posts.String())
}

// -----------
// --- PostMedias
// -----------

func TestPostMedias_String(t *testing.T) {
	postMedias := types.PostMedias{
		types.PostMedia{
			URI:      "uri",
			MimeType: "text/plain",
		},
		types.PostMedia{
			URI:      "uri",
			MimeType: "application/json",
		},
	}

	actual := postMedias.String()

	assert.Equal(t, `[{"uri":"uri","mime_Type":"text/plain"},{"uri":"uri","mime_Type":"application/json"}]`, actual)
}

func TestPostMedias_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.PostMedias
		second    types.PostMedias
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			expEquals: true,
		},
		{
			name: "different data returns false",
			first: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			expEquals: false,
		},
		{
			name: "different length returns false",
			first: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			expEquals: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

func TestPostMedias_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name        string
		medias      types.PostMedias
		newMedia    types.PostMedia
		expMedias   types.PostMedias
		expAppended bool
	}{
		{
			name: "append a new media and returns true",
			medias: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: types.PostMedia{
				URI:      "uri",
				MimeType: "application/json",
			},
			expMedias: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			expAppended: true,
		},
		{
			name: "not append an existing media and returns false",
			medias: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: types.PostMedia{
				URI:      "uri",
				MimeType: "text/plain",
			},
			expMedias: types.PostMedias{
				types.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			expAppended: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			medias, found := test.medias.AppendIfMissing(test.newMedia)
			assert.Equal(t, test.expMedias, medias)
			assert.Equal(t, test.expAppended, found)
		})
	}
}

func TestPostMedias_Validate(t *testing.T) {
	tests := []struct {
		postMedia types.PostMedias
		expErr    string
	}{
		{
			postMedia: types.PostMedias{
				types.PostMedia{
					URI:      "",
					MimeType: "text/plain",
				},
			},
			expErr: "uri must be specified and cannot be empty",
		},

		{
			postMedia: types.PostMedias{
				types.PostMedia{
					URI:      "htt://example.com",
					MimeType: "text/plain",
				},
			},
			expErr: "invalid uri provided",
		},
		{
			postMedia: types.PostMedias{
				types.PostMedia{
					URI:      "https://example.com",
					MimeType: "",
				},
			},
			expErr: "mime type must be specified and cannot be empty",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.expErr, func(t *testing.T) {
			if len(test.expErr) != 0 {
				assert.Equal(t, test.expErr, test.postMedia.Validate().Error())
			} else {
				assert.Nil(t, test.postMedia.Validate())
			}
		})
	}
}

// -----------
// --- PostMedia
// -----------

func TestPostMedia_String(t *testing.T) {
	pm := types.PostMedia{
		URI:      "http://example.com",
		MimeType: "text/plain",
	}

	actual := pm.String()

	assert.Equal(t, `{"uri":"http://example.com","mime_Type":"text/plain"}`, actual)
}

func TestPostMedia_Validate(t *testing.T) {
	tests := []struct {
		postMedia types.PostMedia
		expErr    string
	}{
		{
			postMedia: types.PostMedia{
				URI:      "",
				MimeType: "text/plain",
			},
			expErr: "uri must be specified and cannot be empty",
		},
		{
			postMedia: types.PostMedia{
				URI:      "htt://example.com",
				MimeType: "text/plain",
			},
			expErr: "invalid uri provided",
		},
		{
			postMedia: types.PostMedia{
				URI:      "https://example.com",
				MimeType: "",
			},
			expErr: "mime type must be specified and cannot be empty",
		},
		{
			postMedia: types.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			expErr: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.expErr, func(t *testing.T) {
			if len(test.expErr) != 0 {
				assert.Equal(t, test.expErr, test.postMedia.Validate().Error())
			} else {
				assert.Nil(t, test.postMedia.Validate())
			}
		})
	}
}

func TestPostMedia_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.PostMedia
		second    types.PostMedia
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: types.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: types.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			expEquals: true,
		},
		{
			name: "Different URI returns false",
			first: types.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: types.PostMedia{
				URI:      "https://another.com",
				MimeType: "text/plain",
			},
			expEquals: false,
		},
		{
			name: "Different mime type returns false",
			first: types.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: types.PostMedia{
				URI:      "https://example.com",
				MimeType: "application/json",
			},
			expEquals: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

func TestPostMedia_ParseURI(t *testing.T) {
	tests := []struct {
		uri    string
		expErr error
	}{
		{
			uri:    "http://error.com",
			expErr: fmt.Errorf("invalid uri provided"),
		},
		{
			uri:    "http://",
			expErr: fmt.Errorf("invalid uri provided"),
		},
		{
			uri:    "error.com",
			expErr: fmt.Errorf("invalid uri provided"),
		},
		{
			uri:    ".com",
			expErr: fmt.Errorf("invalid uri provided"),
		},
		{
			uri:    "ttps://",
			expErr: fmt.Errorf("invalid uri provided"),
		},
		{
			uri:    "ps://site.com",
			expErr: fmt.Errorf("invalid uri provided"),
		},
		{
			uri:    "https://",
			expErr: fmt.Errorf("invalid uri provided"),
		},
		{
			uri:    "https://example.com",
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.uri, func(t *testing.T) {
			assert.Equal(t, test.expErr, types.ParseURI(test.uri))
		})
	}
}
