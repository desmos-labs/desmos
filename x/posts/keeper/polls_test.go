package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func (suite *KeeperTestsuite) TestKeeper_HasPoll() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		expResult  bool
	}{
		{
			name: "media attachment returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewMediaAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expResult:  false,
		},
		{
			name: "poll attachment returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewPollAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
				)))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasPoll(ctx, tc.subspaceID, tc.postID, tc.pollID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SaveUserAnswer() {
	user, err := sdk.AccAddressFromBech32("cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw")
	suite.Require().NoError(err)

	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		answer types.UserAnswer
		check  func(ctx sdk.Context)
	}{
		{
			name:   "non existing answer is stored properly",
			answer: types.NewUserAnswer(1, 1, 1, []uint32{1}, user),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, user)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(1, 1, 1, []uint32{1}, user), stored)
			},
		},
		{
			name: "existing answer is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, user))
			},
			answer: types.NewUserAnswer(1, 1, 1, []uint32{1, 2}, user),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, user)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(1, 1, 1, []uint32{1, 2}, user), stored)
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

			suite.k.SaveUserAnswer(ctx, tc.answer)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasUserAnswer() {
	user, err := sdk.AccAddressFromBech32("cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw")
	suite.Require().NoError(err)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		user       sdk.AccAddress
		expResult  bool
	}{
		{
			name:       "non existing answer returns false",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       user,
			expResult:  false,
		},
		{
			name: "existing answer returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, user))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       user,
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasUserAnswer(ctx, tc.subspaceID, tc.postID, tc.pollID, tc.user)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetUserAnswer() {
	user, err := sdk.AccAddressFromBech32("cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw")
	suite.Require().NoError(err)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		user       sdk.AccAddress
		expFound   bool
		expAnswer  types.UserAnswer
	}{
		{
			name:       "not found answer returns false and empty answer",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       user,
			expFound:   false,
			expAnswer:  types.UserAnswer{},
		},
		{
			name: "found answer returns true and correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, user))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       user,
			expFound:   true,
			expAnswer:  types.NewUserAnswer(1, 1, 1, []uint32{1}, user),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			answer, found := suite.k.GetUserAnswer(ctx, tc.subspaceID, tc.postID, tc.pollID, tc.user)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expAnswer, answer)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteUserAnswer() {
	user, err := sdk.AccAddressFromBech32("cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw")
	suite.Require().NoError(err)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		user       sdk.AccAddress
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing answer is deleted properly",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       user,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserAnswer(ctx, 1, 1, 1, user))
			},
		},
		{
			name: "existing answer is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, user))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       user,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserAnswer(ctx, 1, 1, 1, user))
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

			suite.k.DeleteUserAnswer(ctx, tc.subspaceID, tc.postID, tc.pollID, tc.user)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SavePollTallyResults() {
	testCases := []struct {
		name    string
		store   func(ctx sdk.Context)
		results types.PollTallyResults
		check   func(ctx sdk.Context)
	}{
		{
			name: "non existing tally results are stored properly",
			results: types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(1, 10),
			}),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetPollTallyResults(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(1, 10),
				}), stored)
			},
		},
		{
			name: "existing tally results are overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePollTallyResults(ctx, types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(1, 10),
				}))
			},
			results: types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(1, 10),
				types.NewAnswerResult(2, 11),
			}),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetPollTallyResults(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(1, 10),
					types.NewAnswerResult(2, 11),
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

			suite.k.SavePollTallyResults(ctx, tc.results)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasPollTallyResults() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		expResult  bool
	}{
		{
			name:       "non existing tally result return false",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expResult:  false,
		},
		{
			name: "exiting tally result returns true",
			store: func(ctx sdk.Context) {
				suite.k.SavePollTallyResults(ctx, types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(1, 10),
				}))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasPollTallyResults(ctx, tc.subspaceID, tc.postID, tc.pollID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetPollTallyResults() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		expFound   bool
		expResults types.PollTallyResults
	}{
		{
			name:       "non existing tally result return false and empty result",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expFound:   false,
		},
		{
			name: "exiting tally result returns true and correct value",
			store: func(ctx sdk.Context) {
				suite.k.SavePollTallyResults(ctx, types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(1, 10),
				}))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expFound:   true,
			expResults: types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(1, 10),
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			results, found := suite.k.GetPollTallyResults(ctx, tc.subspaceID, tc.postID, tc.pollID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expResults, results)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeletePollTallyResults() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing result is deleted properly",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPollTallyResults(ctx, 1, 1, 1))
			},
		},
		{
			name: "existing result is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePollTallyResults(ctx, types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(1, 10),
				}))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPollTallyResults(ctx, 1, 1, 1))
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

			suite.k.DeletePollTallyResults(ctx, tc.subspaceID, tc.postID, tc.pollID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
