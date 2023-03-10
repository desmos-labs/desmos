package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_HasPoll() {
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
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
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
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
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

func (suite *KeeperTestSuite) TestKeeper_GetPoll() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		expFound   bool
		expPoll    *types.Poll
	}{
		{
			name:       "non existing poll returns nil and false",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expFound:   false,
			expPoll:    nil,
		},
		{
			name: "media attachment returns nil and false",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expFound:   false,
			expPoll:    nil,
		},
		{
			name: "poll returns true and correct value",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				)))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expFound:   true,
			expPoll: types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
				nil,
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

			poll, found := suite.k.GetPoll(ctx, tc.subspaceID, tc.postID, tc.pollID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expPoll, poll)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_Tally() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		expResult  *types.PollTallyResults
		check      func(ctx sdk.Context)
	}{
		{
			name:      "not found poll returns null",
			expResult: nil,
		},
		{
			name: "existing poll returns correct results",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
						types.NewProvidedAnswer("No one of the above", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					true,
					false,
					nil,
				)))

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{0, 1}, "cosmos1pmklwgqjqmgc4ynevmtset85uwm0uau90jdtfn"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1zmqjufkg44ngswgf4vmn7evp8k6h07erdyxefd"))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			expResult: types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(0, 1),
				types.NewAnswerResult(1, 2),
				types.NewAnswerResult(2, 0),
			}),
			check: func(ctx sdk.Context) {
				// Make sure all the answers have been deleted
				answers := suite.k.GetPollUserAnswers(ctx, 1, 1, 1)
				suite.Require().Empty(answers)
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

			results := suite.k.Tally(ctx, tc.subspaceID, tc.postID, tc.pollID)
			suite.Require().Equal(tc.expResult, results)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_EndPoll() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		poll  types.Attachment
		check func(ctx sdk.Context)
	}{
		{
			name: "updates the ended poll properly",
			store: func(ctx sdk.Context) {
				attachment := types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
						types.NewProvidedAnswer("No one of the above", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					true,
					false,
					nil,
				))
				suite.k.SaveAttachment(ctx, attachment)
				suite.k.InsertActivePollQueue(ctx, attachment)

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{0, 1}, "cosmos1pmklwgqjqmgc4ynevmtset85uwm0uau90jdtfn"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1zmqjufkg44ngswgf4vmn7evp8k6h07erdyxefd"))
			},
			poll: types.NewAttachment(1, 1, 1, types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
					types.NewProvidedAnswer("No one of the above", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				true,
				false,
				nil,
			)),
			check: func(ctx sdk.Context) {
				poll, found := suite.k.GetPoll(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(0, 1),
					types.NewAnswerResult(1, 2),
					types.NewAnswerResult(2, 0),
				}), poll.FinalTallyResults)

				store := ctx.KVStore(suite.storeKey)
				endTime := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)
				suite.Require().False(store.Has(types.ActivePollQueueKey(1, 1, 1, endTime)))

				suite.Require().False(suite.k.HasUserAnswer(ctx, 1, 1, 1, "cosmos1pmklwgqjqmgc4ynevmtset85uwm0uau90jdtfn"))
				suite.Require().False(suite.k.HasUserAnswer(ctx, 1, 1, 1, "cosmos1zmqjufkg44ngswgf4vmn7evp8k6h07erdyxefd"))
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

			suite.k.EndPoll(ctx, tc.poll)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_SaveUserAnswer() {
	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		answer types.UserAnswer
		check  func(ctx sdk.Context)
	}{
		{
			name:   "non existing answer is stored properly",
			answer: types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw")
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"), stored)
			},
		},
		{
			name: "existing answer is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"))
			},
			answer: types.NewUserAnswer(1, 1, 1, []uint32{1, 2}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw")
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(1, 1, 1, []uint32{1, 2}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"), stored)
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

func (suite *KeeperTestSuite) TestKeeper_HasUserAnswer() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		user       string
		expResult  bool
	}{
		{
			name:       "non existing answer returns false",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
			expResult:  false,
		},
		{
			name: "existing answer returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
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

func (suite *KeeperTestSuite) TestKeeper_GetUserAnswer() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		user       string
		expFound   bool
		expAnswer  types.UserAnswer
	}{
		{
			name:       "not found answer returns false and empty answer",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
			expFound:   false,
			expAnswer:  types.UserAnswer{},
		},
		{
			name: "found answer returns true and correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
			expFound:   true,
			expAnswer:  types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
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

func (suite *KeeperTestSuite) TestKeeper_DeleteUserAnswer() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing answer is deleted properly",
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserAnswer(ctx, 1, 1, 1, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"))
			},
		},
		{
			name: "existing answer is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserAnswer(ctx, 1, 1, 1, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"))
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

			suite.k.DeleteUserAnswer(ctx, tc.subspaceID, tc.postID, tc.pollID, tc.user, false)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
