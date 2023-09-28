package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v6/x/reactions/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Reactions() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		request      *types.QueryReactionsRequest
		shouldErr    bool
		expReactions []types.Reaction
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryReactionsRequest(0, 1, "", nil),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			request:   types.NewQueryReactionsRequest(1, 0, "", nil),
			shouldErr: true,
		},
		{
			name: "request without user returns properly",
			store: func(ctx sdk.Context) {
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
			request: types.NewQueryReactionsRequest(1, 1, "", &query.PageRequest{
				Limit:  1,
				Offset: 1,
			}),
			shouldErr: false,
			expReactions: []types.Reaction{
				types.NewReaction(
					1,
					1,
					2,
					types.NewRegisteredReactionValue(2),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				),
			},
		},
		{
			name: "request with user returns properly",
			store: func(ctx sdk.Context) {
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
					"cosmos14z8mn9ywhqu84alr5grxuljwj87jyz0zpxnlxy",
				))
			},
			request:   types.NewQueryReactionsRequest(1, 1, "cosmos14z8mn9ywhqu84alr5grxuljwj87jyz0zpxnlxy", nil),
			shouldErr: false,
			expReactions: []types.Reaction{
				types.NewReaction(
					1,
					1,
					2,
					types.NewRegisteredReactionValue(2),
					"cosmos14z8mn9ywhqu84alr5grxuljwj87jyz0zpxnlxy",
				),
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

			res, err := suite.k.Reactions(ctx, tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReactions, res.Reactions)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Reaction() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		request     *types.QueryReactionRequest
		shouldErr   bool
		expReaction types.Reaction
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryReactionRequest(0, 1, 1),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			request:   types.NewQueryReactionRequest(1, 0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid reaction id returns error",
			request:   types.NewQueryReactionRequest(1, 1, 0),
			shouldErr: true,
		},
		{
			name: "valid request returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			request:   types.NewQueryReactionRequest(1, 1, 1),
			shouldErr: false,
			expReaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.Reaction(ctx, tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReaction, res.Reaction)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_RegisteredReactions() {
	testCases := []struct {
		name                   string
		store                  func(ctx sdk.Context)
		request                *types.QueryRegisteredReactionsRequest
		shouldErr              bool
		expRegisteredReactions []types.RegisteredReaction
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryRegisteredReactionsRequest(0, nil),
			shouldErr: true,
		},
		{
			name: "request with pagination works properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))

				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					2,
					":wave:",
					"https://example.com?image=wave.png",
				))
			},
			request: types.NewQueryRegisteredReactionsRequest(1, &query.PageRequest{
				Limit:  1,
				Offset: 1,
			}),
			shouldErr: false,
			expRegisteredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					1,
					2,
					":wave:",
					"https://example.com?image=wave.png",
				),
			},
		},
		{
			name: "request without pagination works properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))

				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					2,
					":wave:",
					"https://example.com?image=wave.png",
				))
			},
			request:   types.NewQueryRegisteredReactionsRequest(1, nil),
			shouldErr: false,
			expRegisteredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				),
				types.NewRegisteredReaction(
					1,
					2,
					":wave:",
					"https://example.com?image=wave.png",
				),
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

			res, err := suite.k.RegisteredReactions(ctx, tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRegisteredReactions, res.RegisteredReactions)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_RegisteredReaction() {
	testCases := []struct {
		name                  string
		store                 func(ctx sdk.Context)
		request               *types.QueryRegisteredReactionRequest
		shouldErr             bool
		expRegisteredReaction types.RegisteredReaction
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryRegisteredReactionRequest(0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid reaction id returns error",
			request:   types.NewQueryRegisteredReactionRequest(1, 0),
			shouldErr: true,
		},
		{
			name: "valid request returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			request:   types.NewQueryRegisteredReactionRequest(1, 1),
			shouldErr: false,
			expRegisteredReaction: types.NewRegisteredReaction(
				1,
				1,
				":hello:",
				"https://example.com?image=hello.png",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.RegisteredReaction(ctx, tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRegisteredReaction, res.RegisteredReaction)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_ReactionsParams() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryReactionsParamsRequest
		shouldErr bool
		expParams types.SubspaceReactionsParams
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryReactionsParamsRequest(0),
			shouldErr: true,
		},
		{
			name:      "not stored params return error",
			request:   types.NewQueryReactionsParamsRequest(1),
			shouldErr: true,
		},
		{
			name: "valid request returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 1000, "[a-z]"),
				))
			},
			request:   types.NewQueryReactionsParamsRequest(1),
			shouldErr: false,
			expParams: types.NewSubspaceReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 1000, "[a-z]"),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.ReactionsParams(ctx, tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expParams, res.Params)
			}
		})
	}
}
