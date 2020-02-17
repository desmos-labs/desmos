package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

// -----------
// --- PostMedias
// -----------

func TestPostMedias_String(t *testing.T) {
	postMedias := types.PostMedias{
		types.PostMedia{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
		types.PostMedia{
			URI:      "https://another.com",
			MimeType: "application/json",
		},
	}

	actual := postMedias.String()

	expected := "Medias - [URI] [Mime-Type]\n[https://uri.com] text/plain \n[https://another.com] application/json"

	require.Equal(t, expected, actual)
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
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
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

func TestPostMedia_String(t *testing.T) {
	pm := types.PostMedia{
		URI:      "http://example.com",
		MimeType: "text/plain",
	}

	actual := pm.String()
	expected := "Media - URI - [http://example.com] ; Mime-Type - [text/plain] \n"

	require.Equal(t, expected, actual)
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
			require.Equal(t, test.expErr, types.ValidateURI(test.uri))
		})
	}
}
