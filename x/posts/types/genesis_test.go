package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
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
			}, nil, nil, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "duplicated subspace data entries return error",
			data: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspaceDataEntry(1, 2),
				types.NewSubspaceDataEntry(1, 3),
			}, nil, nil, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "invalid initial post id returns error",
			data: types.NewGenesisState(nil, []types.GenesisPost{
				types.NewGenesisPost(0, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				)),
			}, nil, nil, types.Params{}),
			shouldErr: true,
		},
		{
			name: "invalid genesis post returns error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.GenesisPost{
					types.NewGenesisPost(1, types.NewPost(
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
					)),
				},
				nil,
				nil,
				types.Params{},
			),
			shouldErr: true,
		},
		{
			name: "duplicated genesis posts return error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.GenesisPost{
					types.NewGenesisPost(1, types.NewPost(
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
					)),
					types.NewGenesisPost(3, types.NewPost(
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
					)),
				},
				nil,
				nil,
				types.Params{},
			),
			shouldErr: true,
		},
		{
			name: "duplicated attachments return error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.GenesisPost{
					types.NewGenesisPost(1, types.NewPost(
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
					)),
				},
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
				types.Params{},
			),
			shouldErr: true,
		},
		{
			name: "invalid initial attachment id returns error",
			data: types.NewGenesisState(
				[]types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
				[]types.GenesisPost{
					types.NewGenesisPost(0, types.NewPost(
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
					)),
				},
				[]types.Attachment{
					types.NewAttachment(1, 1, 1, types.NewMedia(
						"ftp://user:password@example.com/image.png",
						"image/png",
					)),
				},
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
				[]types.GenesisPost{
					types.NewGenesisPost(1, types.NewPost(
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
					)),
				},
				[]types.Attachment{
					types.NewAttachment(0, 1, 1, types.NewMedia(
						"ftp://user:password@example.com/image.png",
						"image/png",
					)),
				},
				nil,
				types.Params{},
			),
			shouldErr: true,
		},

		{
			name: "invalid user answer returns error",
			data: types.NewGenesisState(nil, nil, nil, []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
			}, types.Params{}),
			shouldErr: true,
		},
		{
			name: "duplicated user answers return error",
			data: types.NewGenesisState(nil, nil, nil, []types.UserAnswer{
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
				[]types.GenesisPost{
					types.NewGenesisPost(2, types.NewPost(
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
					)),
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
