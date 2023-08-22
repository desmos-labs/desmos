package keeper_test

import (
	"time"

	"github.com/golang/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/reports/keeper"
	"github.com/desmos-labs/desmos/v6/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestValidSubspacesInvariant() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing next reason id breaks invariant",
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
				}
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool)) {
						for _, subspace := range subspaces {
							fn(subspace)
						}
					})
			},
			expBroken: true,
		},
		{
			name: "non existing next report id breaks invariant",
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

			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
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
				suite.k.SetNextReportID(ctx, 1, 1)
			},
			expBroken: false,
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

			_, broken := keeper.ValidSubspacesInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidReasonsInvariant() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing subspace breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			expBroken: true,
		},
		{
			name: "non existing next reason id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			expBroken: true,
		},
		{
			name: "invalid reason id compared to next reason id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 1)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			expBroken: true,
		},
		{
			name: "invalid reason breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 2)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"",
					"This content is spam",
				))
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 2)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			expBroken: false,
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

			_, broken := keeper.ValidReasonsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidReportsInvariant() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "missing subspace breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: true,
		},
		{
			name: "missing reason breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: true,
		},
		{
			name: "missing next report id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: true,
		},
		{
			name: "invalid report id compared to next report id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 1)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: true,
		},
		{
			name: "missing post breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					HasPost(gomock.Any(), uint64(1), uint64(1)).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 2)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: true,
		},
		{
			name: "invalid report breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 2)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewUserTarget("cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79"),
					"",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					HasPost(gomock.Any(), uint64(1), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReportID(ctx, 1, 2)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This user is spamming",
					types.NewPostTarget(1),
					"cosmos1ggzk8tnte9lmzgpvyzzdtmwmn6rjlct4spmjjd",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: false,
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

			_, broken := keeper.ValidReportsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}
