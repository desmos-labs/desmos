package common_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/posts/types/models/common"
)

// -----------
// --- Attachments
// -----------

func TestPostMedias_String(t *testing.T) {
	var tag, err = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	var tag2, err2 = sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	require.NoError(t, err2)

	postMedias := common.Attachments{
		common.Attachment{
			URI:      "https://uri.com",
			MimeType: "text/plain",
			Tags:     []sdk.AccAddress{tag, tag2},
		},
		common.Attachment{
			URI:      "https://another.com",
			MimeType: "application/json",
			Tags:     []sdk.AccAddress{tag},
		},
	}

	actual := postMedias.String()

	expected := "[URI] [Mime-Type] [Tags]\n[https://uri.com] [text/plain] [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns,\ncosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h,\n] \n[https://another.com] [application/json] [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns,\n]"

	require.Equal(t, expected, actual)
}

func TestPostMedias_Equals(t *testing.T) {
	var tag, err = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	var tag2, err2 = sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	require.NoError(t, err2)

	tests := []struct {
		name      string
		first     common.Attachments
		second    common.Attachments
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
					Tags:     []sdk.AccAddress{tag, tag2},
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
					Tags:     []sdk.AccAddress{tag},
				},
			},
			second: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
					Tags:     []sdk.AccAddress{tag, tag2},
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
					Tags:     []sdk.AccAddress{tag},
				},
			},
			expEquals: true,
		},
		{
			name: "different data returns false",
			first: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			second: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			expEquals: false,
		},
		{
			name: "different length returns false",
			first: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			expEquals: false,
		},
		{
			name: "different tags length returns false",
			first: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
					Tags:     []sdk.AccAddress{tag, tag2},
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
					Tags:     []sdk.AccAddress{tag},
				},
			},
			second: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
					Tags:     []sdk.AccAddress{tag},
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
					Tags:     []sdk.AccAddress{tag},
				},
			},
			expEquals: false,
		},
		{
			name: "different tags returns false",
			first: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
					Tags:     []sdk.AccAddress{tag2},
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
					Tags:     []sdk.AccAddress{tag},
				},
			},
			second: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
					Tags:     []sdk.AccAddress{tag},
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
					Tags:     []sdk.AccAddress{tag2},
				},
			},
			expEquals: false,
		},
		{
			name: "nil tags returns true",
			first: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
					Tags:     nil,
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
					Tags:     nil,
				},
			},
			second: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
					Tags:     nil,
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
					Tags:     nil,
				},
			},
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

func TestPostMedias_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name        string
		medias      common.Attachments
		newMedia    common.Attachment
		expMedias   common.Attachments
		expAppended bool
	}{
		{
			name: "append a new attachment and returns true",
			medias: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: common.Attachment{
				URI:      "uri",
				MimeType: "application/json",
			},
			expMedias: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
				common.Attachment{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
		},
		{
			name: "not append an existing attachment and returns false",
			medias: common.Attachments{
				common.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: common.Attachment{
				URI:      "uri",
				MimeType: "text/plain",
			},
			expMedias: common.Attachments{
				common.Attachment{
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
		postMedia common.Attachments
		expErr    string
	}{
		{
			postMedia: common.Attachments{
				common.Attachment{
					URI:      "",
					MimeType: "text/plain",
				},
			},
			expErr: "invalid uri provided",
		},

		{
			postMedia: common.Attachments{
				common.Attachment{
					URI:      "htt://example.com",
					MimeType: "text/plain",
				},
			},
			expErr: "invalid uri provided",
		},
		{
			postMedia: common.Attachments{
				common.Attachment{
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
// --- Attachment
// -----------

func TestPostMedia_Validate(t *testing.T) {
	var tag, err = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name      string
		postMedia common.Attachment
		expErr    string
	}{
		{
			name: "Empty URI",
			postMedia: common.Attachment{
				URI:      "",
				MimeType: "text/plain",
			},
			expErr: "invalid uri provided",
		},
		{
			name: "Invalid URI",
			postMedia: common.Attachment{
				URI:      "htt://example.com",
				MimeType: "text/plain",
			},
			expErr: "invalid uri provided",
		},
		{
			name: "Empty mime type",
			postMedia: common.Attachment{
				URI:      "https://example.com",
				MimeType: "",
			},
			expErr: "mime type must be specified and cannot be empty",
		},
		{
			name: "Invalid Tags",
			postMedia: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
				Tags:     []sdk.AccAddress{{}},
			},
			expErr: "invalid empty tag address: ",
		},
		{
			name: "No errors attachment (with tags)",
			postMedia: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
				Tags:     []sdk.AccAddress{tag},
			},
			expErr: "",
		},
		{
			name: "No errors attachment (without tags)",
			postMedia: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
				Tags:     nil,
			},
			expErr: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if len(test.expErr) != 0 {
				require.Equal(t, test.expErr, test.postMedia.Validate().Error())
			} else {
				require.Nil(t, test.postMedia.Validate())
			}
		})
	}
}

func TestPostMedia_Equals(t *testing.T) {
	var tag, err = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name      string
		first     common.Attachment
		second    common.Attachment
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
				Tags:     []sdk.AccAddress{tag},
			},
			second: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
				Tags:     []sdk.AccAddress{tag},
			},
			expEquals: true,
		},
		{
			name: "Different URI returns false",
			first: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: common.Attachment{
				URI:      "https://another.com",
				MimeType: "text/plain",
			},
			expEquals: false,
		},
		{
			name: "Different mime type returns false",
			first: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: common.Attachment{
				URI:      "https://example.com",
				MimeType: "application/json",
			},
			expEquals: false,
		},
		{
			name: "Different tags returns false",
			first: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
				Tags:     []sdk.AccAddress{tag},
			},
			second: common.Attachment{
				URI:      "https://example.com",
				MimeType: "text/plain",
				Tags:     []sdk.AccAddress{},
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
