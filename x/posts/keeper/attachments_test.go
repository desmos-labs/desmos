package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetAttachmentID() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		subspaceID   uint64
		postID       uint64
		attachmentID uint32
		check        func(ctx sdk.Context)
	}{
		{
			name:         "non exiting attachment id is set properly",
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				attachmentID := types.GetAttachmentIDFromBytes(store.Get(types.AttachmentIDStoreKey(1, 1)))
				suite.Require().Equal(uint32(1), attachmentID)
			},
		},
		{
			name: "existing attachment id is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SetAttachmentID(ctx, 1, 1, 1)
			},
			subspaceID:   1,
			postID:       1,
			attachmentID: 2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				attachmentID := types.GetAttachmentIDFromBytes(store.Get(types.AttachmentIDStoreKey(1, 1)))
				suite.Require().Equal(uint32(2), attachmentID)
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

			suite.k.SetAttachmentID(ctx, tc.subspaceID, tc.postID, tc.attachmentID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasAttachmentID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		expResult  bool
	}{
		{
			name:       "not found attachment id returns false",
			subspaceID: 1,
			postID:     1,
			expResult:  false,
		},
		{
			name: "found attachment id returns true",
			store: func(ctx sdk.Context) {
				suite.k.SetAttachmentID(ctx, 1, 1, 1)
			},
			subspaceID: 1,
			postID:     1,
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

			result := suite.k.HasAttachmentID(ctx, tc.subspaceID, tc.postID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetAttachmentID() {
	testCases := []struct {
		name            string
		store           func(ctx sdk.Context)
		subspaceID      uint64
		postID          uint64
		shouldErr       bool
		expAttachmentID uint32
	}{
		{
			name:       "non existing attachment id returns error",
			subspaceID: 1,
			postID:     1,
			shouldErr:  true,
		},
		{
			name: "exiting attachment id returns correct value",
			store: func(ctx sdk.Context) {
				suite.k.SetAttachmentID(ctx, 1, 1, 1)
			},
			subspaceID:      1,
			postID:          1,
			shouldErr:       false,
			expAttachmentID: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			attachmentID, err := suite.k.GetAttachmentID(ctx, tc.subspaceID, tc.postID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expAttachmentID, attachmentID)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SaveAttachment() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		attachment types.Attachment
		check      func(ctx sdk.Context)
	}{
		{
			name: "non existing attachment is stored properly",
			attachment: types.NewAttachment(1, 1, 1, types.NewMedia(
				"ftp://user:password@example.com/image.png",
				"image/png",
			)),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetAttachment(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)), stored)
			},
		},
		{
			name: "existing attachment is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			attachment: types.NewAttachment(1, 1, 1, types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			)),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetAttachment(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
				)), stored)
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

			suite.k.SaveAttachment(ctx, tc.attachment)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasAttachment() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		subspaceID   uint64
		postID       uint64
		attachmentID uint32
		expResult    bool
	}{
		{
			name:         "not found attachment returns false",
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			expResult:    false,
		},
		{
			name: "found attachment returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			expResult:    true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasAttachment(ctx, tc.subspaceID, tc.postID, tc.attachmentID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetAttachment() {
	testCases := []struct {
		name          string
		store         func(ctx sdk.Context)
		subspaceID    uint64
		postID        uint64
		attachmentID  uint32
		expFound      bool
		expAttachment types.Attachment
	}{
		{
			name:          "non existing attachment returns false and empty attachment",
			subspaceID:    1,
			postID:        1,
			attachmentID:  1,
			expFound:      false,
			expAttachment: types.Attachment{},
		},
		{
			name: "existing attachment returns true and the correct value",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			expFound:     true,
			expAttachment: types.NewAttachment(1, 1, 1, types.NewMedia(
				"ftp://user:password@example.com/image.png",
				"image/png",
			)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			attachment, found := suite.k.GetAttachment(ctx, tc.subspaceID, tc.postID, tc.attachmentID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expAttachment, attachment)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteAttachment() {
	user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
	suite.Require().NoError(err)

	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		subspaceID   uint64
		postID       uint64
		attachmentID uint32
		check        func(ctx sdk.Context)
	}{
		{
			name:         "non existing attachment is deleted properly",
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 1, 1))
			},
		},
		{
			name: "exiting media attachment is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 1, 1))
			},
		},
		{
			name: "existing poll attachment is deleted along with all the data",
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
				)))

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{1},
					user,
				))

				suite.k.SavePollTallyResults(ctx, types.NewPollTallyResults(
					1,
					1,
					1,
					[]types.PollTallyResults_AnswerResult{
						types.NewAnswerResult(1, 10),
					}),
				)
			},
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 1, 1))

				suite.Require().Empty(suite.k.GetUserAnswers(ctx, 1, 1, 1))
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

			suite.k.DeleteAttachment(ctx, tc.subspaceID, tc.postID, tc.attachmentID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
