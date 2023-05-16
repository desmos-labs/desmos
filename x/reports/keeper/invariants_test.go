package keeper_test

import (
	"time"

	poststypes "github.com/desmos-labs/desmos/v5/x/posts/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/reports/keeper"
	"github.com/desmos-labs/desmos/v5/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestValidSubspacesInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing next reason id breaks invariant",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: true,
		},
		{
			name: "non existing next report id breaks invariant",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReasonID(ctx, 1, 1)

			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing subspace breaks invariant",
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "missing subspace breaks invariant",
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
			name: "missing next report id breaks invariant",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextReportID(ctx, 1, 2)

				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
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
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidReportsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}
