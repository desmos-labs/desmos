package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"
	"github.com/desmos-labs/desmos/v4/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestKeeper_AfterSubspaceSaved() {
	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		subspace subspacestypes.Subspace
		check    func(ctx sdk.Context)
	}{
		{
			name: "next registered reaction id is saved properly",
			subspace: subspacestypes.NewSubspace(1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextRegisteredReactionID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(1), stored)
			},
		},
		{
			name: "next registered reaction id not overridden",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextRegisteredReactionID(ctx, 1, 2)
			},
			subspace: subspacestypes.NewSubspace(1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextRegisteredReactionID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), stored)
			},
		},
		{
			name: "reactions params are saved properly",
			subspace: subspacestypes.NewSubspace(1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetSubspaceReactionsParams(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(types.DefaultReactionsParams(1), stored)
			},
		},
		{
			name: "reactions params are not overridden",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 1000, "[a-z]"),
				))
			},
			subspace: subspacestypes.NewSubspace(1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetSubspaceReactionsParams(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 1000, "[a-z]"),
				), stored)
			},
		},
	}

	suite.sk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.sk.SaveSubspace(ctx, tc.subspace)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_AfterSubspaceDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "next registered reaction id is deleted properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextRegisteredReactionID(ctx, 1, 2)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasNextRegisteredReactionID(ctx, 1))
			},
		},
		{
			name: "registered reactions are deleted properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))

				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					2,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasRegisteredReaction(ctx, 1, 1))
				suite.Require().False(suite.k.HasRegisteredReaction(ctx, 1, 2))
			},
		},
		{
			name: "reactions params are deleted properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 1000, "[a-z]"),
				))
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasSubspaceReactionsParams(ctx, 1))
			},
		},
	}

	suite.sk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.sk.DeleteSubspace(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_AfterPostSaved() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		post  poststypes.Post
		check func(ctx sdk.Context)
	}{
		{
			name: "next reaction id is set properly",
			post: poststypes.NewPost(
				1,
				0,
				1,
				"External ID",
				"This is a text",
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				1,
				nil,
				nil,
				poststypes.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextReactionID(ctx, 1, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(1), stored)
			},
		},
		{
			name: "next reaction id is not overridden",
			store: func(ctx sdk.Context) {
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
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextReactionID(ctx, 1, 1, 2)
			},
			post: poststypes.NewPost(
				1,
				0,
				1,
				"External ID",
				"This is a text",
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				1,
				nil,
				nil,
				poststypes.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextReactionID(ctx, 1, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), stored)
			},
		},
	}

	suite.pk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.pk.SavePost(ctx, tc.post)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_AfterPostDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "next reaction id is deleted properly",
			store: func(ctx sdk.Context) {
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
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextReactionID(ctx, 1, 1, 1)
			},
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasNextReactionID(ctx, 1, 1))
			},
		},
		{
			name: "reactions are deleted properly",
			store: func(ctx sdk.Context) {
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
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))

				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					2,
					types.NewRegisteredReactionValue(2),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasReaction(ctx, 1, 1, 1))
				suite.Require().False(suite.k.HasReaction(ctx, 1, 1, 2))
			},
		},
	}

	suite.pk.SetHooks(suite.k.Hooks())

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.pk.DeletePost(ctx, tc.subspaceID, tc.postID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
