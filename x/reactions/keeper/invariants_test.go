package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v5/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v5/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestValidSubspacesInvariant() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non existing next registered reaction id breaks invariant",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
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
			name: "non existing reactions params break invariant",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
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
			name: "valid data does not break invairiant",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
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
				suite.k.SetNextRegisteredReactionID(ctx, 1, 1)
				suite.k.SaveSubspaceReactionsParams(ctx, types.DefaultReactionsParams(1))
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

func (suite *KeeperTestSuite) TestValidRegisteredReactionsInvariant() {
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
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			expBroken: true,
		},
		{
			name: "non existing next registered reaction id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			expBroken: true,
		},
		{
			name: "invalid next registered reaction id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextRegisteredReactionID(ctx, 1, 1)

				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			expBroken: true,
		},
		{
			name: "invalid registered reaction breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextRegisteredReactionID(ctx, 1, 2)

				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					"",
					"https://example.com?image=hello.png",
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
				suite.k.SetNextRegisteredReactionID(ctx, 1, 2)

				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
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

			_, broken := keeper.ValidRegisteredReactionsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidReactionsInvariant() {
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

				suite.pk.EXPECT().
					HasPost(gomock.Any(), uint64(1), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			expBroken: true,
		},
		{
			name: "non existing post breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					HasPost(gomock.Any(), uint64(1), uint64(1)).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			expBroken: true,
		},
		{
			name: "non existing next reaction id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					HasPost(gomock.Any(), uint64(1), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			expBroken: true,
		},
		{
			name: "invalid next reaction id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					HasPost(gomock.Any(), uint64(1), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReactionID(ctx, 1, 1, 1)

				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			expBroken: true,
		},
		{
			name: "invalid registered reaction breaks invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)

				suite.pk.EXPECT().
					HasPost(gomock.Any(), uint64(1), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReactionID(ctx, 1, 1, 2)

				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(0),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
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
				suite.k.SetNextReactionID(ctx, 1, 1, 2)

				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
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

			_, broken := keeper.ValidReactionsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidReactionsParamsInvariant() {
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
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 1000, "[a-z]"),
				))
			},
			expBroken: true,
		},
		{
			name: "invalid params break invariant",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 1000, ".*{1,2}"),
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
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 1000, ""),
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

			_, broken := keeper.ValidReactionsParamsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}
