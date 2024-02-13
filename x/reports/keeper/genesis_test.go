package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v7/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	testCases := []struct {
		name       string
		setup      func()
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "subspaces data entries are exported properly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
					subspacestypes.NewSubspace(
						2,
						"Another test subspace",
						"This is another test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				}
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool)) {
						for _, subspace := range subspaces {
							fn(subspace)
						}
					})
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 1)
				suite.k.SetNextReportID(ctx, 1, 2)

				suite.k.SetNextReasonID(ctx, 2, 3)
				suite.k.SetNextReportID(ctx, 2, 4)

				suite.k.SetParams(ctx, types.Params{})
			},
			expGenesis: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspacesDataEntry(1, 1, 2),
				types.NewSubspacesDataEntry(2, 3, 4),
			}, nil, nil, types.Params{}),
		},
		{
			name: "reasons are exported properly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{}
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool)) {
						for _, subspace := range subspaces {
							fn(subspace)
						}
					})
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SetParams(ctx, types.Params{})
			},
			expGenesis: types.NewGenesisState(nil, []types.Reason{
				types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				),
			}, nil, types.Params{}),
		},
		{
			name: "reports are exported properly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{}
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool)) {
						for _, subspace := range subspaces {
							fn(subspace)
						}
					})
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SetParams(ctx, types.Params{})
			},
			expGenesis: types.NewGenesisState(nil, nil, []types.Report{
				types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewUserTarget("cosmos1pjffdtweghpyxru9alssyqtdkq8mn6sepgstgm"),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			}, types.Params{}),
		},
		{
			name: "params are exported properly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{}
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool)) {
						for _, subspace := range subspaces {
							fn(subspace)
						}
					})
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams([]types.StandardReason{
					types.NewStandardReason(1, "Spam", "This content is spam"),
				}))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, types.NewParams([]types.StandardReason{
				types.NewStandardReason(1, "Spam", "This content is spam"),
			})),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			genesis := suite.k.ExportGenesis(ctx)
			suite.Require().Equal(tc.expGenesis, genesis)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_InitGenesis() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		data  types.GenesisState
		check func(ctx sdk.Context)
	}{
		{
			name: "subspaces data are initialized properly",
			data: types.GenesisState{
				SubspacesData: []types.SubspaceDataEntry{
					types.NewSubspacesDataEntry(1, 2, 3),
				},
			},
			check: func(ctx sdk.Context) {
				nextReasonID, err := suite.k.GetNextReasonID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), nextReasonID)

				nextReportID, err := suite.k.GetNextReportID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(3), nextReportID)
			},
		},
		{
			name: "reasons are imported properly",
			data: types.GenesisState{
				Reasons: []types.Reason{
					types.NewReason(
						2,
						1,
						"Spam",
						"This content is spam",
					),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReason(ctx, 2, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReason(
					2,
					1,
					"Spam",
					"This content is spam",
				), stored)
			},
		},
		{
			name: "reports are imported properly",
			data: types.GenesisState{
				Reports: []types.Report{
					types.NewReport(
						1,
						1,
						[]uint32{1},
						"This content is spam",
						types.NewPostTarget(1),
						"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReport(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewPostTarget(1),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), stored)
			},
		},
		{
			name: "params are initialized properly",
			data: types.GenesisState{
				Params: types.NewParams([]types.StandardReason{
					types.NewStandardReason(1, "Spam", "This content is spam"),
				}),
			},
			check: func(ctx sdk.Context) {
				stored := suite.k.GetParams(ctx)
				suite.Require().Equal(types.NewParams([]types.StandardReason{
					types.NewStandardReason(1, "Spam", "This content is spam"),
				}), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.k.InitGenesis(ctx, tc.data)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
