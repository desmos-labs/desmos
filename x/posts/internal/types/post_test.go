package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

// -------------
// --- PostID
// -------------

func TestParsePostID(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	tests := []struct {
		name     string
		value    string
		expID    types.PostID
		expError string
	}{
		{
			name:     "Invalid id returns error",
			value:    "id",
			expID:    types.PostID(id),
			expError: "invalid postID cannot be parsed: id",
		},
		{
			name:     "Empty id returns error",
			value:    "",
			expID:    types.PostID(id),
			expError: "invalid postID cannot be parsed: ",
		},
		{
			name:     "Valid id returns proper value",
			value:    "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			expID:    types.PostID(id),
			expError: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			id, err := types.ParsePostID(test.value)

			if err == nil {
				require.Equal(t, test.expID, id)
			} else {
				require.Equal(t, test.expError, err.Error())
			}
		})
	}
}

// -------------
// --- PostIDs
// -------------

func TestPostIDs_Equals(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := []byte("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	tests := []struct {
		name      string
		first     types.PostIDs
		second    types.PostIDs
		expEquals bool
	}{
		{
			name:      "Different length",
			first:     types.PostIDs{types.PostID(id), types.PostID(id)},
			second:    types.PostIDs{types.PostID(id)},
			expEquals: false,
		},
		{
			name:      "Different order",
			first:     types.PostIDs{types.PostID(id), types.PostID(id2)},
			second:    types.PostIDs{types.PostID(id2), types.PostID(id)},
			expEquals: false,
		},
		{
			name:      "Same length and order",
			first:     types.PostIDs{types.PostID(id), types.PostID(id2)},
			second:    types.PostIDs{types.PostID(id), types.PostID(id2)},
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

func TestPostIDs_AppendIfMissing(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := []byte("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	tests := []struct {
		name      string
		IDs       types.PostIDs
		newID     types.PostID
		expIDs    types.PostIDs
		expEdited bool
	}{
		{
			name:      "AppendIfMissing dont append anything",
			IDs:       types.PostIDs{types.PostID(id)},
			newID:     types.PostID(id),
			expIDs:    types.PostIDs{types.PostID(id)},
			expEdited: false,
		},
		{
			name:      "AppendIfMissing append something",
			IDs:       types.PostIDs{types.PostID(id)},
			newID:     types.PostID(id2),
			expIDs:    types.PostIDs{types.PostID(id), types.PostID(id2)},
			expEdited: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newIDs, edited := test.IDs.AppendIfMissing(test.newID)
			require.Equal(t, test.expIDs, newIDs)
			require.Equal(t, test.expEdited, edited)
		})
	}
}

func TestPostID_MarshalJSON(t *testing.T) {
	creationDate := time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone)
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	computedID := types.ComputeID(creationDate, creator, subspace)
	json := types.ModuleCdc.MustMarshalJSON(computedID)
	stringID := computedID.String()
	require.Equal(t, stringID, string(json))
}

func TestPostID_UnmarshalJSON(t *testing.T) {
	creationDate := time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone)
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	computedID := types.ComputeID(creationDate, creator, subspace)
	stringID := computedID.String()

	tests := []struct {
		name     string
		value    string
		expID    types.PostID
		expError string
	}{
		{
			name:     "Invalid ID returns error",
			value:    "id",
			expID:    nil,
			expError: "invalid postID cannot be parsed: id",
		},
		{
			name:     "Valid id is read properly",
			value:    stringID,
			expID:    computedID,
			expError: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var id types.PostID
			err := types.ModuleCdc.UnmarshalJSON([]byte(test.value), &id)

			if err == nil {
				require.Equal(t, test.expID, id)
			} else {
				require.Equal(t, test.expError, err.Error())
			}
		})
	}
}

// -----------
// --- Post
// -----------

func TestPost_String(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff")
	id2 := []byte("f1b909289cd23188c19da17ae")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	post := types.Post{
		PostID:         types.PostID(id),
		ParentID:       types.PostID(id2),
		Message:        "My post message",
		Created:        date,
		LastEdited:     date.AddDate(0, 0, 1),
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Creator:        owner,
	}

	require.Equal(t,
		`{"id":"3139646530326531303563363861363065343563323839626666","parent_id":"66316239303932383963643233313838633139646131376165","message":"My post message","created":"2020-01-01T12:00:00Z","last_edited":"2020-01-02T12:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`,
		post.String(),
	)
}

func TestPost_Validate(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff")
	id2 := []byte("f1b909289cd23188c19da17ae")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}
	answer := types.PollAnswer{
		ID:   types.AnswerID(1),
		Text: "Yes",
	}

	answer2 := types.PollAnswer{
		ID:   types.AnswerID(2),
		Text: "No",
	}
	pollData := types.NewPollData("poll?", time.Now().UTC().Add(time.Hour), types.PollAnswers{answer, answer2}, true, false, true)

	tests := []struct {
		post     types.Post
		expError string
	}{
		{
			post:     types.NewPost(types.PostID(nil), types.PostID(nil), "Message", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, owner).WithMedias(medias).WithPollData(pollData),
			expError: "invalid post id: ",
		},
		{
			post:     types.NewPost(id, id2, "", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, nil).WithMedias(medias).WithPollData(pollData),
			expError: "invalid post owner: ",
		},
		{
			post:     types.NewPost(id, id2, "", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, owner).WithPollData(pollData),
			expError: "post message or medias required, they cannot be both empty",
		},
		{
			post:     types.NewPost(id, id2, " ", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, owner).WithPollData(pollData),
			expError: "post message or medias required, they cannot be both empty",
		},
		{
			post:     types.NewPost(id, id2, "Message", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, time.Time{}, owner).WithMedias(medias).WithPollData(pollData),
			expError: "invalid post creation time: 0001-01-01 00:00:00 +0000 UTC",
		},
		{
			post:     types.Post{PostID: id, Creator: owner, Message: "Message", Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", Created: date, LastEdited: date.AddDate(0, 0, -1)},
			expError: "invalid post last edit time: 2019-12-31 12:00:00 +0000 UTC",
		},
		{
			post:     types.NewPost(id, id2, "Message", true, "", map[string]string{}, date, owner).WithMedias(medias).WithPollData(pollData),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			post:     types.NewPost(id, id2, "Message", true, " ", map[string]string{}, date, owner).WithMedias(medias).WithPollData(pollData),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			post: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Created:        time.Now().UTC().Add(time.Hour),
				Creator:        owner,
				Medias:         medias,
			},
			expError: "post creation date cannot be in the future",
		},
		{
			post: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
				id,
				id2,
				`
				Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque massa felis, aliquam sed ipsum at, 
				mollis pharetra quam. Vestibulum nec nulla ante. Praesent sed dignissim turpis. Curabitur aliquam nunc 
				eu nisi porta, eu gravida purus faucibus. Duis commodo sagittis lacus, vitae luctus enim vulputate a. 
				Nulla tempor eget nunc vitae vulputate. Nulla facilities. Donec sollicitudin odio in arcu efficitur, 
				sit amet vestibulum diam ullamcorper. Ut ac dolor in velit gravida efficitur et et erat volutpat.
				`,
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				date,
				owner,
			).WithMedias(medias).WithPollData(pollData),
			expError: "post message cannot be longer than 500 characters",
		},
		{
			post: types.NewPost(
				id,
				id2,
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
			).WithMedias(medias).WithPollData(pollData),
			expError: "post optional data cannot contain more than 10 key-value pairs",
		},
		{
			post: types.NewPost(
				id,
				id2,
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{
					"key1": `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque euismod, mi at commodo 
							efficitur, quam sapien congue enim, ut porttitor lacus tellus vitae turpis. Vivamus aliquam 
							sem eget neque metus.`,
				},
				date,
				owner,
			).WithMedias(medias).WithPollData(pollData),
			expError: "post optional data values cannot exceed 200 characters. key1 of post with id 3139646530326531303563363861363065343563323839626666 is longer than this",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.expError, func(t *testing.T) {
			if len(test.expError) != 0 {
				require.Equal(t, test.expError, test.post.Validate().Error())
			} else {
				require.Nil(t, test.post.Validate())
			}
		})
	}
}

func TestPost_Equals(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := []byte("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	otherOwner, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	answer := types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: types.AnswerID(2), Text: "No"}
	pollData := types.NewPollData("poll?", pollEndDate, types.PollAnswers{answer, answer2}, true, false, true)

	tests := []struct {
		name      string
		first     types.Post
		second    types.Post
		expEquals bool
	}{
		{
			name: "Different post ID",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         id2,
				ParentID:       id,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different parent ID",
			first: types.Post{
				PostID:         id,
				ParentID:       id,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different message",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "Another post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different creation time",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date.AddDate(0, 0, 1),
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different last edited",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 2),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different allows comments",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: false,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different subspace",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
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
				PostID:         id,
				ParentID:       id2,
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
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: map[string]string{
					"field": "value",
				},
				Creator: owner,
				Medias:  medias,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        otherOwner,
				Medias:         medias,
			},
			expEquals: false,
		},
		{
			name: "Different medias",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        otherOwner,
				Medias:         types.PostMedias{},
			},
			expEquals: false,
		},
		{
			name: "Different polls",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
				Medias:         medias,
				PollData:       nil,
			},
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        otherOwner,
				Medias:         medias,
				PollData:       &types.PollData{},
			},
			expEquals: false,
		},
		{
			name: "Equals posts",
			first: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
			}.WithMedias(medias).WithPollData(pollData),
			second: types.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Creator:        owner,
			}.WithMedias(medias).WithPollData(pollData),
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

func TestPost_GetPostHashtags(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := []byte("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)

	tests := []struct {
		name        string
		post        types.Post
		expHashtags []string
	}{
		{
			name: "Hashtags in message extracted correctly (spaced hashtags)",
			post: types.NewPost(id,
				id2,
				"Post with #test #desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				date,
				owner,
			),
			expHashtags: []string{"test", "desmos"},
		},
		{
			name: "Hashtags in message extracted correctly (non-spaced hashtags)",
			post: types.NewPost(id,
				id2,
				"Post with #test#desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				date,
				owner,
			),
			expHashtags: []string{},
		},
		{
			name: "Hashtags in message extracted correctly (underscore separated hashtags)",
			post: types.NewPost(id,
				id2,
				"Post with #test_#desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				date,
				owner,
			),
			expHashtags: []string{},
		},
		{
			name: "Hashtags in message extracted correctly (only number hashtag)",
			post: types.NewPost(id,
				id2,
				"Post with #101112",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				date,
				owner,
			),
			expHashtags: []string{},
		},
		{
			name: "No hashtags in message",
			post: types.NewPost(id,
				id2,
				"Post with no hashtag",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				date,
				owner,
			),
			expHashtags: []string{},
		},
		{
			name: "No same hashtags inside string array",
			post: types.NewPost(id,
				id2,
				"Post with double #hashtag #hashtag",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				date,
				owner,
			),
			expHashtags: []string{"hashtag"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			hashtags := test.post.GetPostHashtags()
			require.Equal(t, test.expHashtags, hashtags)
		})
	}
}

// -----------
// --- Posts
// -----------
func TestPosts_String(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff")
	id2 := []byte("f1b909289cd23188c19da17ae")
	owner1, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	owner2, err := sdk.AccAddressFromBech32("cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l")
	require.NoError(t, err)

	medias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}
	answer := types.PollAnswer{
		ID:   types.AnswerID(1),
		Text: "Yes",
	}

	answer2 := types.PollAnswer{
		ID:   types.AnswerID(2),
		Text: "No",
	}
	pollData := types.NewPollData("poll?", time.Now().UTC().Add(time.Hour), types.PollAnswers{answer, answer2}, true, false, true)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 0, 00, 000, timeZone)

	posts := types.Posts{
		types.NewPost(id, id2, "Post 1", false, "external-ref-1", map[string]string{}, date, owner1).WithMedias(medias).WithPollData(pollData),
		types.NewPost(id, id2, "Post 2", false, "external-ref-1", map[string]string{}, date, owner2).WithMedias(medias).WithPollData(pollData),
	}

	expected := `ID - [Creator] Message
3139646530326531303563363861363065343563323839626666 - [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns] Post 1
3139646530326531303563363861363065343563323839626666 - [cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l] Post 2`
	require.Equal(t, expected, posts.String())
}
