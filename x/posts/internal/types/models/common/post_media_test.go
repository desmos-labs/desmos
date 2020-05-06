package common_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types/models/common"
	"github.com/stretchr/testify/require"
)

// -----------
// --- PostMedias
// -----------

func TestPostMedias_String(t *testing.T) {
	postMedias := common.PostMedias{
		common.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
		common.PostMedia{
			URI:      "https://another.com",
			MimeType: "application/json",
		},
	}

	actual := postMedias.String()

	expected := "URI [Mime-Type]\n[https://uri.com] text/plain \n[https://another.com] application/json"

	require.Equal(t, expected, actual)
}

func TestPostMedias_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     common.PostMedias
		second    common.PostMedias
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				common.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				common.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			expEquals: true,
		},
		{
			name: "different data returns false",
			first: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				common.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			expEquals: false,
		},
		{
			name: "different length returns false",
			first: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				common.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: common.PostMedias{
				common.PostMedia{
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
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

func TestPostMedias_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name        string
		medias      common.PostMedias
		newMedia    common.PostMedia
		expMedias   common.PostMedias
		expAppended bool
	}{
		{
			name: "append a new media and returns true",
			medias: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: common.PostMedia{
				URI:      "uri",
				MimeType: "application/json",
			},
			expMedias: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				common.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
		},
		{
			name: "not append an existing media and returns false",
			medias: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: common.PostMedia{
				URI:      "uri",
				MimeType: "text/plain",
			},
			expMedias: common.PostMedias{
				common.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			medias := test.medias.AppendIfMissing(test.newMedia)
			require.Equal(t, test.expMedias, medias)
		})
	}
}

func TestPostMedias_Validate(t *testing.T) {
	tests := []struct {
		postMedia common.PostMedias
		expErr    string
	}{
		{
			postMedia: common.PostMedias{
				common.PostMedia{
					URI:      "",
					MimeType: "text/plain",
				},
			},
			expErr: "invalid uri provided",
		},

		{
			postMedia: common.PostMedias{
				common.PostMedia{
					URI:      "htt://example.com",
					MimeType: "text/plain",
				},
			},
			expErr: "invalid uri provided",
		},
		{
			postMedia: common.PostMedias{
				common.PostMedia{
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
				require.Equal(t, test.expErr, test.postMedia.Validate().Error())
			} else {
				require.Nil(t, test.postMedia.Validate())
			}
		})
	}
}

// -----------
// --- PostMedia
// -----------

func TestPostMedia_Validate(t *testing.T) {
	tests := []struct {
		postMedia common.PostMedia
		expErr    string
	}{
		{
			postMedia: common.PostMedia{
				URI:      "",
				MimeType: "text/plain",
			},
			expErr: "invalid uri provided",
		},
		{
			postMedia: common.PostMedia{
				URI:      "htt://example.com",
				MimeType: "text/plain",
			},
			expErr: "invalid uri provided",
		},
		{
			postMedia: common.PostMedia{
				URI:      "https://example.com",
				MimeType: "",
			},
			expErr: "mime type must be specified and cannot be empty",
		},
		{
			postMedia: common.PostMedia{
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
				require.Equal(t, test.expErr, test.postMedia.Validate().Error())
			} else {
				require.Nil(t, test.postMedia.Validate())
			}
		})
	}
}

func TestPostMedia_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     common.PostMedia
		second    common.PostMedia
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: common.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: common.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			expEquals: true,
		},
		{
			name: "Different URI returns false",
			first: common.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: common.PostMedia{
				URI:      "https://another.com",
				MimeType: "text/plain",
			},
			expEquals: false,
		},
		{
			name: "Different mime type returns false",
			first: common.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: common.PostMedia{
				URI:      "https://example.com",
				MimeType: "application/json",
			},
			expEquals: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
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
			expErr: nil,
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
			require.Equal(t, test.expErr, common.ValidateURI(test.uri))
		})
	}
}
