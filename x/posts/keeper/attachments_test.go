package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetNextAttachmentID() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		subspaceID   uint64
		postID       uint64
		attachmentID uint32
		check        func(ctx sdk.Context)
	}{
		{
			name:         "non existing attachment id is set properly",
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				attachmentID := types.GetAttachmentIDFromBytes(store.Get(types.NextAttachmentIDStoreKey(1, 1)))
				suite.Require().Equal(uint32(1), attachmentID)
			},
		},
		{
			name: "existing attachment id is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextAttachmentID(ctx, 1, 1, 1)
			},
			subspaceID:   1,
			postID:       1,
			attachmentID: 2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				attachmentID := types.GetAttachmentIDFromBytes(store.Get(types.NextAttachmentIDStoreKey(1, 1)))
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

			suite.k.SetNextAttachmentID(ctx, tc.subspaceID, tc.postID, tc.attachmentID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasNextAttachmentID() {
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
				suite.k.SetNextAttachmentID(ctx, 1, 1, 1)
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

			result := suite.k.HasNextAttachmentID(ctx, tc.subspaceID, tc.postID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetNextAttachmentID() {
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
			name: "existing attachment id returns correct value",
			store: func(ctx sdk.Context) {
				suite.k.SetNextAttachmentID(ctx, 1, 1, 1)
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

			attachmentID, err := suite.k.GetNextAttachmentID(ctx, tc.subspaceID, tc.postID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expAttachmentID, attachmentID)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteNextAttachmentID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing attachment id is deleted properly",
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextAttachmentIDStoreKey(1, 1)))
			},
		},
		{
			name: "existing post id is deleted properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.NextAttachmentIDStoreKey(1, 1), types.GetAttachmentIDBytes(1))
			},
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextAttachmentIDStoreKey(1, 1)))
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

			suite.k.DeleteNextAttachmentID(ctx, tc.subspaceID, tc.postID)
			if tc.check != nil {
				tc.check(ctx)
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
				time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
				types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(0, 1),
					types.NewAnswerResult(2, 5),
				}),
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
					time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
						types.NewAnswerResult(0, 1),
						types.NewAnswerResult(2, 5),
					}),
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
					nil,
				)))

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{1},
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 1, 1))

				suite.Require().Empty(suite.k.GetPollUserAnswers(ctx, 1, 1, 1))
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
