package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -----------
// --- MediaPost
// -----------

func TestMediaPost_String(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)

	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)
	assert.Equal(t,
		`{"id":"2","parent_id":"0","message":"media Post","created":"2020-01-01T15:15:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":false,"subspace":"desmos","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"uri":"uri","provider":"provider","mime_Type":"text/plain"}]}`,
		mp.String(),
	)
}

func TestMediaPost_GetID(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.GetID()

	assert.Equal(t, types.PostID(2), actual)
}

func TestMediaPost_GetParentID(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.GetParentID()

	assert.Equal(t, types.PostID(0), actual)
}

func TestMediaPost_SetMessage(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.SetMessage("edited media post")

	assert.Equal(t, "edited media post", actual.GetMessage())
}

func TestMediaPost_GetMessage(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.GetMessage()

	assert.Equal(t, "media Post", actual)

}

func TestMediaPost_CreationTime(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.CreationTime()

	assert.Equal(t, testPostCreationDate, actual)
}

func TestMediaPost_SetEditTime(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.SetEditTime(time.Date(2020, 1, 2, 15, 15, 00, 000, timeZone))

	assert.Equal(t, time.Date(2020, 1, 2, 15, 15, 00, 000, timeZone), actual.GetEditTime())
}

func TestMediaPost_GetEditTime(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.GetEditTime()

	assert.Equal(t, time.Time{}, actual)
}

func TestMediaPost_CanComment(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.CanComment()

	assert.Equal(t, false, actual)
}

func TestMediaPost_GetSubspace(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.GetSubspace()

	assert.Equal(t, "desmos", actual)
}

func TestMediaPost_GetOptionalData(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{"key1": "value1"}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.GetOptionalData()

	assert.Equal(t, map[string]string{"key1": "value1"}, actual)
}

func TestMediaPost_Owner(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{"key1": "value1"}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)

	actual := mp.Owner()

	assert.Equal(t, owner, actual)
}

func TestMediaPost_Validate(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	tests := []struct {
		post   types.MediaPost
		expErr string
	}{
		{
			post: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			expErr: "media provider must be specified and cannot be empty",
		},
		{
			post: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "",
						MimeType: "text/plain",
					},
				},
			),
			expErr: "uri must be specified and cannot be empty",
		},
		{
			post: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "htt://example.com",
						MimeType: "text/plain",
					},
				},
			),
			expErr: "invalid uri provided",
		},
		{
			post: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "",
					},
				},
			),
			expErr: "mime type must be specified and cannot be empty",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.expErr, func(t *testing.T) {
			if len(test.expErr) != 0 {
				assert.Equal(t, test.expErr, test.post.Validate().Error())
			} else {
				assert.Nil(t, test.post.Validate())
			}
		})
	}
}

func TestMediaPost_Equals(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	otherOwner, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)

	tests := []struct {
		name      string
		first     types.MediaPost
		second    types.MediaPost
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			second: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			expEquals: true,
		},
		{
			name: "Different owner returns false",
			first: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			second: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, otherOwner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			expEquals: false,
		},
		{
			name: "Different provider returns false",
			first: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			second: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider2",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			expEquals: false,
		},
		{
			name: "Different URI returns false",
			first: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			second: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, otherOwner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://another.com",
						MimeType: "text/plain",
					},
				},
			),
			expEquals: false,
		},
		{
			name: "Different mime type returns false",
			first: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			),
			second: types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, otherOwner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "application/json",
					},
				},
			),
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

func TestMediaPost_MarshalJSON(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)
	json := types.ModuleCdc.MustMarshalJSON(mp)
	assert.Equal(t, `{"type":"desmos/MediaPost","value":{"id":"2","parent_id":"0","message":"media Post","created":"2020-01-01T15:15:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":false,"subspace":"desmos","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"uri":"uri","provider":"provider","mime_Type":"text/plain"}]}}`, string(json))
}

func TestMediaPost_UnmarshalJSON(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mp = types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", nil, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)
	tests := []struct {
		name         string
		value        string
		expMediaPost types.MediaPost
		expError     string
	}{
		{
			name:         "Valid media post is ready properly",
			value:        `{"type":"desmos/MediaPost","value":{"id":"2","parent_id":"0","message":"media Post","created":"2020-01-01T15:15:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":false,"subspace":"desmos","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"uri":"uri","provider":"provider","mime_Type":"text/plain"}]}}`,
			expMediaPost: mp,
			expError:     "",
		},
		{
			name:         "Invalid media post returns error",
			value:        `{"post":{"id":"2","parent_id":"0","message":"media Post","created":"0","last_edited":"0","allows_comments":false,"subspace":"desmos","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},"medias":[{"provider":"ipfs","uri":"uri","mime_Type":"text/plain"}]}`,
			expMediaPost: mp,
			expError:     "JSON encoding of interfaces require non-empty type field.",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var mediaPost types.MediaPost
			err := types.ModuleCdc.UnmarshalJSON([]byte(test.value), &mediaPost)

			if err == nil {
				assert.Equal(t, test.expMediaPost, mediaPost)
			} else {
				assert.Equal(t, test.expError, err.Error())
			}
		})
	}
}

// -----------
// --- MediaPosts
// -----------

func TestMediaPosts_Equals(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	otherOwner, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)

	tests := []struct {
		name      string
		first     types.MediaPosts
		second    types.MediaPosts
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			)},
			second: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			)},
			expEquals: true,
		},
		{
			name: "Different owner returns false",
			first: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			)},
			second: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, otherOwner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			)},
			expEquals: false,
		},
		{
			name: "Different provider returns false",
			first: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			)},
			second: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider2",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			)},
			expEquals: false,
		},
		{
			name: "Different URI returns false",
			first: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			)},
			second: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://different.com",
						MimeType: "text/plain",
					},
				},
			)},
			expEquals: false,
		},
		{
			name: "Different mime type returns false",
			first: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "text/plain",
					},
				},
			)},
			second: types.MediaPosts{types.NewMediaPost(
				types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "https://example.com",
						MimeType: "application/json",
					},
				},
			)},
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

func TestMediaPosts_String(t *testing.T) {
	owner, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	testPostCreationDate := time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone)
	var mps = types.MediaPosts{types.NewMediaPost(
		types.NewTextPost(types.PostID(2), types.PostID(0), "media Post", false, "desmos", map[string]string{}, testPostCreationDate, owner),
		types.PostMedias{
			types.PostMedia{
				Provider: "ipfs",
				URI:      "uri",
				MimeType: "text/plain",
			},
		},
	)}

	assert.Equal(t,
		`[{"id":"2","parent_id":"0","message":"media Post","created":"2020-01-01T15:15:00Z","last_edited":"0001-01-01T00:00:00Z","allows_comments":false,"subspace":"desmos","creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","medias":[{"uri":"uri","provider":"ipfs","mime_Type":"text/plain"}]}]`,
		mps.String(),
	)
}

// -----------
// --- PostMedias
// -----------

func TestPostMedias_String(t *testing.T) {
	postMedias := types.PostMedias{
		types.PostMedia{
			Provider: "ipfs",
			URI:      "uri",
			MimeType: "text/plain",
		},
		types.PostMedia{
			Provider: "dropbox",
			URI:      "uri",
			MimeType: "application/json",
		},
	}

	actual := postMedias.String()

	assert.Equal(t, `[{"uri":"uri","provider":"ipfs","mime_Type":"text/plain"},{"uri":"uri","provider":"dropbox","mime_Type":"application/json"}]`, actual)
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
					Provider: "ipfs",
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					Provider: "dropbox",
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: types.PostMedias{
				types.PostMedia{
					Provider: "ipfs",
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					Provider: "dropbox",
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
					Provider: "ipfs",
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					Provider: "dropbox",
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: types.PostMedias{
				types.PostMedia{
					Provider: "dropbox",
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					Provider: "ipfs",
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
					Provider: "ipfs",
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					Provider: "dropbox",
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: types.PostMedias{
				types.PostMedia{
					Provider: "dropbox",
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
					Provider: "ipfs",
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: types.PostMedia{
				Provider: "dropbox",
				URI:      "uri",
				MimeType: "application/json",
			},
			expMedias: types.PostMedias{
				types.PostMedia{
					Provider: "ipfs",
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.PostMedia{
					Provider: "dropbox",
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
					Provider: "ipfs",
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: types.PostMedia{
				Provider: "ipfs",
				URI:      "uri",
				MimeType: "text/plain",
			},
			expMedias: types.PostMedias{
				types.PostMedia{
					Provider: "ipfs",
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

// -----------
// --- PostMedia
// -----------

func TestPostMedia_String(t *testing.T) {
	pm := types.PostMedia{
		Provider: "provider",
		URI:      "http://example.com",
		MimeType: "text/plain",
	}

	actual := pm.String()

	assert.Equal(t, `{"uri":"http://example.com","provider":"provider","mime_Type":"text/plain"}`, actual)
}

func TestPostMedia_Validate(t *testing.T) {
	tests := []struct {
		postMedia types.PostMedia
		expErr    string
	}{
		{
			postMedia: types.PostMedia{
				Provider: "",
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			expErr: "media provider must be specified and cannot be empty",
		},
		{
			postMedia: types.PostMedia{
				Provider: "provider",
				URI:      "",
				MimeType: "text/plain",
			},
			expErr: "uri must be specified and cannot be empty",
		},
		{
			postMedia: types.PostMedia{
				Provider: "provider",
				URI:      "htt://example.com",
				MimeType: "text/plain",
			},
			expErr: "invalid uri provided",
		},
		{
			postMedia: types.PostMedia{
				Provider: "provider",
				URI:      "https://example.com",
				MimeType: "",
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
				Provider: "provider",
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: types.PostMedia{
				Provider: "provider",
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			expEquals: true,
		},
		{
			name: "Different provider returns false",
			first: types.PostMedia{
				Provider: "provider",
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: types.PostMedia{
				Provider: "provider2",
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			expEquals: false,
		},
		{
			name: "Different URI returns false",
			first: types.PostMedia{
				Provider: "provider",
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: types.PostMedia{
				Provider: "provider",
				URI:      "https://another.com",
				MimeType: "text/plain",
			},
			expEquals: false,
		},
		{
			name: "Different mime type returns false",
			first: types.PostMedia{
				Provider: "provider",
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: types.PostMedia{
				Provider: "provider",
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
