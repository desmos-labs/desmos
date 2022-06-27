package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name      string
		data      *types.GenesisState
		shouldErr bool
	}{
		{
			name: "invalid subspace data entry returns error",
			data: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspaceDataEntry(0, 0),
			}, nil, nil, nil, nil, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "duplicated subspace data entries return error",
			data: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspaceDataEntry(1, 2),
				types.NewSubspaceDataEntry(1, 3),
			}, nil, nil, nil, nil, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "invalid post returns error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.Post{
					types.NewPost(
						1,
						0,
						0,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				},
				nil,
				nil,
				nil,
				nil,
				types.Params{},
			),
			shouldErr: true,
		},
		{
			name: "duplicated posts return error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.Post{
					types.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						0,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
					types.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						0,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				},
				nil,
				nil,
				nil,
				nil,
				types.Params{},
			),
			shouldErr: true,
		},
		{
			name: "invalid post data entry returns error",
			data: types.NewGenesisState(nil, nil, []types.PostDataEntry{
				types.NewPostDataEntry(0, 1, 1),
			}, nil, nil, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "duplicated post data entry returns error",
			data: types.NewGenesisState(nil, nil, []types.PostDataEntry{
				types.NewPostDataEntry(1, 1, 1),
				types.NewPostDataEntry(1, 1, 1),
			}, nil, nil, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "duplicated attachments return error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.Post{
					types.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						0,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				},
				nil,
				[]types.Attachment{
					types.NewAttachment(1, 1, 1, types.NewMedia(
						"ftp://user:password@example.com/image.png",
						"image/png",
					)),
					types.NewAttachment(1, 1, 1, types.NewMedia(
						"ftp://user:password@example.com/image.png",
						"image/png",
					)),
				},
				nil,
				nil,
				types.Params{},
			),
			shouldErr: true,
		},
		{
			name: "invalid attachment returns error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.Post{
					types.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						0,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				},
				nil,
				[]types.Attachment{
					types.NewAttachment(0, 1, 1, types.NewMedia(
						"ftp://user:password@example.com/image.png",
						"image/png",
					)),
				},
				nil,
				nil,
				types.Params{},
			),
			shouldErr: true,
		},
		{
			name: "invalid poll data returns error",
			data: types.NewGenesisState(nil, nil, nil, nil, []types.ActivePollData{
				types.NewActivePollData(0, 1, 1, time.Now()),
			}, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "duplicated poll data returns error",
			data: types.NewGenesisState(nil, nil, nil, nil, []types.ActivePollData{
				types.NewActivePollData(1, 1, 1, time.Now()),
				types.NewActivePollData(1, 1, 1, time.Now()),
			}, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "invalid user answer returns error",
			data: types.NewGenesisState(nil, nil, nil, nil, nil, []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
			}, types.Params{}),
			shouldErr: true,
		},
		{
			name: "duplicated user answers return error",
			data: types.NewGenesisState(nil, nil, nil, nil, nil, []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
			}, types.Params{}),
			shouldErr: true,
		},
		{
			name:      "default genesis does not error",
			data:      types.DefaultGenesisState(),
			shouldErr: false,
		},
		{
			name: "valid genesis state does not error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.Post{
					types.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						0,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				},
				[]types.PostDataEntry{
					types.NewPostDataEntry(1, 1, 2),
				},
				[]types.Attachment{
					types.NewAttachment(1, 1, 1, types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						nil,
					)),
				},
				nil,
				[]types.UserAnswer{
					types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				},
				types.NewParams(100),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateGenesis(tc.data)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestSubspaceDataEntry_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		entry     types.SubspaceDataEntry
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			entry:     types.NewSubspaceDataEntry(0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid initial post id returns error",
			entry:     types.NewSubspaceDataEntry(1, 0),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			entry:     types.NewSubspaceDataEntry(1, 1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entry.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPostDataEntry_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		entry     types.PostDataEntry
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			entry:     types.NewPostDataEntry(0, 1, 1),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			entry:     types.NewPostDataEntry(1, 0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid attachment id returns error",
			entry:     types.NewPostDataEntry(1, 1, 0),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			entry:     types.NewPostDataEntry(1, 1, 1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entry.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestActivePollData_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		data      types.ActivePollData
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			data:      types.NewActivePollData(0, 1, 1, time.Now()),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			data:      types.NewActivePollData(1, 0, 1, time.Now()),
			shouldErr: true,
		},
		{
			name:      "invalid poll id returns error",
			data:      types.NewActivePollData(1, 1, 0, time.Now()),
			shouldErr: true,
		},
		{
			name:      "invalid end date returns error",
			data:      types.NewActivePollData(1, 1, 1, time.Time{}),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			data:      types.NewActivePollData(1, 1, 1, time.Now()),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.data.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
