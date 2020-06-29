package types_test

import (
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types/common"
	"github.com/desmos-labs/desmos/x/posts/internal/types/polls"

	"github.com/stretchr/testify/require"
)

// -------------
// --- PostID
// -------------

func TestPostID_Equals(t *testing.T) {
	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	creationDate := time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone)
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	subspace2 := "ec8202b6f9fb16f9e26b66367afa4e037752f3c09a18cefab426165e06a424b1"

	tests := []struct {
		name    string
		postID  types.PostID
		otherID types.PostID
		expBool bool
	}{
		{
			name:    "Equal IDs returns true",
			postID:  types.ComputeID(creationDate, creator, subspace),
			otherID: types.ComputeID(creationDate, creator, subspace),
			expBool: true,
		},
		{
			name:    "Non Equal IDs returns false",
			postID:  types.ComputeID(creationDate, creator, subspace),
			otherID: types.ComputeID(creationDate, creator, subspace2),
			expBool: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res := test.postID.Equals(test.otherID)
			require.Equal(t, test.expBool, res)
		})
	}

}

func TestPostID_String(t *testing.T) {
	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	creationDate := time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone)

	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	computedID := types.ComputeID(creationDate, creator, subspace)

	require.Equal(t, "f55d90114d81e70399d6330a57081b86ae1bdf928b78a57e88870f64240009ef", computedID.String())
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

// -----------
// --- Post
// -----------

func TestPost_String(t *testing.T) {
	id := types.PostID("dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1")
	id2 := types.PostID("e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	post := types.Post{
		PostID:         id,
		ParentID:       id2,
		Message:        "My post message",
		Created:        date,
		LastEdited:     date.AddDate(0, 0, 1),
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Creator:        owner,
	}

	require.Equal(t,
		`{"id":"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1","parent_id":"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163","message":"My post message","created":"2020-01-01T12:00:00Z","last_edited":"2020-01-02T12:00:00Z","allows_comments":true,"subspace":"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`,
		post.String(),
	)
}

func TestPost_Validate(t *testing.T) {
	id := types.PostID("dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1")
	id2 := types.PostID("e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	medias := common.PostMedias{
		common.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}
	answer := polls.PollAnswer{
		ID:   polls.AnswerID(1),
		Text: "Yes",
	}

	answer2 := polls.PollAnswer{
		ID:   polls.AnswerID(2),
		Text: "No",
	}
	pollData := polls.NewPollData("poll?", time.Now().UTC().Add(time.Hour), polls.PollAnswers{answer, answer2}, true, false, true)

	tests := []struct {
		name     string
		post     types.Post
		expError string
	}{
		{
			name:     "Invalid postID",
			post:     types.NewPost("", "", "Message", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, owner).WithMedias(medias).WithPollData(pollData),
			expError: "invalid postID: ",
		},
		{
			name:     "Invalid post owner",
			post:     types.NewPost(id, id2, "", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, nil).WithMedias(medias).WithPollData(pollData),
			expError: "invalid post owner: ",
		},
		{
			name:     "Empty post message and media",
			post:     types.NewPost(id, id2, "", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, owner).WithPollData(pollData),
			expError: "post message or medias required, they cannot be both empty",
		},
		{
			name:     "Empty post message (blank) and media",
			post:     types.NewPost(id, id2, " ", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, owner).WithPollData(pollData),
			expError: "post message or medias required, they cannot be both empty",
		},
		{
			name:     "Invalid post creation time",
			post:     types.NewPost(id, id2, "Message", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, time.Time{}, owner).WithMedias(medias).WithPollData(pollData),
			expError: "invalid post creation time: 0001-01-01 00:00:00 +0000 UTC",
		},
		{
			name:     "Invalid post last edit time",
			post:     types.Post{PostID: id, Creator: owner, Message: "Message", Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", Created: date, LastEdited: date.AddDate(0, 0, -1)},
			expError: "invalid post last edit time: 2019-12-31 12:00:00 +0000 UTC",
		},
		{
			name:     "Invalid post subspace",
			post:     types.NewPost(id, id2, "Message", true, "", map[string]string{}, date, owner).WithMedias(medias).WithPollData(pollData),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			name:     "Invalid post subspace(blank)",
			post:     types.NewPost(id, id2, "Message", true, " ", map[string]string{}, date, owner).WithMedias(medias).WithPollData(pollData),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			name: "Post creation data in future",
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
			name: "Post last edit date in future",
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
			name: "Post message cannot be longer than 500 characters",
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
			name: "post optional data cannot contain more than 10 key-value",
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
			name: "post optional data values cannot exceed 200 characters",
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
			expError: "post optional data values cannot exceed 200 characters. key1 of post with id dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1 is longer than this",
		},
		{
			name:     "Valid post without poll data",
			post:     types.NewPost(id, "", "Message", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, owner).WithMedias(medias),
			expError: "",
		},
		{
			name:     "Valid post without medias",
			post:     types.NewPost(id, "", "Message", true, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", map[string]string{}, date, owner).WithPollData(pollData),
			expError: "",
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
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	otherOwner, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	medias := common.PostMedias{
		common.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	pollData := polls.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
		polls.NewPollAnswers(
			polls.NewPollAnswer(polls.AnswerID(1), "Yes"),
			polls.NewPollAnswer(polls.AnswerID(2), "No"),
		),
		true,
		false,
		true,
	)

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
				Medias:         common.PostMedias{},
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
				PollData:       &polls.PollData{},
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
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
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
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	owner1, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	owner2, err := sdk.AccAddressFromBech32("cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l")
	require.NoError(t, err)

	medias := common.PostMedias{
		common.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}
	answer := polls.PollAnswer{
		ID:   polls.AnswerID(1),
		Text: "Yes",
	}

	answer2 := polls.PollAnswer{
		ID:   polls.AnswerID(2),
		Text: "No",
	}
	pollData := polls.NewPollData("poll?", time.Now().UTC().Add(time.Hour), polls.PollAnswers{answer, answer2}, true, false, true)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 0, 00, 000, timeZone)

	posts := types.Posts{
		types.NewPost(id, id2, "Post 1", false, "external-ref-1", map[string]string{}, date, owner1).WithMedias(medias).WithPollData(pollData),
		types.NewPost(id, id2, "Post 2", false, "external-ref-1", map[string]string{}, date, owner2).WithMedias(medias).WithPollData(pollData),
	}

	expected := `ID - [Creator] Message
19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af - [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns] Post 1
19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af - [cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l] Post 2`
	require.Equal(t, expected, posts.String())
}
