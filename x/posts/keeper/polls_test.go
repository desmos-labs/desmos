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

func (suite *KeeperTestsuite) TestKeeper_GetPoll() {
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

func (suite *KeeperTestsuite) TestKeeper_Tally() {
	firstUser, err := sdk.AccAddressFromBech32("cosmos1pmklwgqjqmgc4ynevmtset85uwm0uau90jdtfn")
	suite.Require().NoError(err)

	secondUser, err := sdk.AccAddressFromBech32("cosmos1zmqjufkg44ngswgf4vmn7evp8k6h07erdyxefd")
	suite.Require().NoError(err)

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

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{0, 1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, secondUser))
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
