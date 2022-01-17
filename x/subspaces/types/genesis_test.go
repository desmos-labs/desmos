package types_test

//
//import (
//	"testing"
//	"time"
//
//	types2 "github.com/desmos-labs/desmos/v2/x/subspaces/types"
//
//	"github.com/stretchr/testify/require"
//)
//
//func TestValidateGenesis(t *testing.T) {
//	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
//	require.NoError(t, err)
//
//	tests := []struct {
//		name      string
//		genesis   *types2.GenesisState
//		shouldErr bool
//	}{
//		{
//			name:      "Default genesis does not error",
//			genesis:   types2.DefaultGenesisState(),
//			shouldErr: false,
//		},
//		{
//			name: "Genesis with invalid subspaces returns error",
//			genesis: types2.NewGenesisState(
//				[]types2.Subspace{
//					types2.NewSubspace(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						"",
//						"",
//						"",
//						types2.SubspaceTypeOpen,
//						time.Time{},
//					),
//				},
//				nil,
//				nil,
//				nil,
//			),
//			shouldErr: true,
//		},
//		{
//			name: "Genesis with duplicated subspaces returns error",
//			genesis: types2.NewGenesisState(
//				[]types2.Subspace{
//					types2.NewSubspace(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						"name",
//						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//						types2.SubspaceTypeOpen,
//						date,
//					),
//					types2.NewSubspace(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						"name",
//						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//						types2.SubspaceTypeOpen,
//						date,
//					),
//				},
//				nil,
//				nil,
//				nil,
//			),
//			shouldErr: true,
//		},
//		{
//			name: "Genesis with duplicated admins entry returns error",
//			genesis: types2.NewGenesisState(
//				nil,
//				[]types2.UsersEntry{
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
//					),
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn"},
//					),
//				},
//				nil,
//				nil,
//			),
//			shouldErr: true,
//		},
//		{
//			name: "Genesis with duplicated admins returns error",
//			genesis: types2.NewGenesisState(
//				nil,
//				[]types2.UsersEntry{
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{
//							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//						},
//					),
//				},
//				nil,
//				nil,
//			),
//			shouldErr: true,
//		},
//		{
//			name: "Genesis with duplicated registered users entry returns error",
//			genesis: types2.NewGenesisState(
//				nil,
//				nil,
//				[]types2.UsersEntry{
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
//					),
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn"},
//					),
//				},
//				nil,
//			),
//			shouldErr: true,
//		},
//		{
//			name: "Genesis with duplicated registered users returns error",
//			genesis: types2.NewGenesisState(
//				nil,
//				nil,
//				[]types2.UsersEntry{
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{
//							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//						},
//					),
//				},
//				nil,
//			),
//			shouldErr: true,
//		},
//		{
//			name: "Genesis with duplicated banned users entry returns error",
//			genesis: types2.NewGenesisState(
//				nil,
//				nil,
//				nil,
//				[]types2.UsersEntry{
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
//					),
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn"},
//					),
//				},
//			),
//			shouldErr: true,
//		},
//		{
//			name: "Genesis with duplicated banned users returns error",
//			genesis: types2.NewGenesisState(
//				nil,
//				nil,
//				nil,
//				[]types2.UsersEntry{
//					types2.NewUsersEntry(
//						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
//						[]string{
//							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
//						},
//					),
//				},
//			),
//			shouldErr: true,
//		},
//	}
//
//	for _, test := range tests {
//		test := test
//		t.Run(test.name, func(t *testing.T) {
//			err = types2.ValidateGenesis(test.genesis)
//			if test.shouldErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
