package keeper_test

import (
	"time"

	"github.com/golang/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetNextPostID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing post id is set properly",
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetPostIDFromBytes(store.Get(types.NextPostIDStoreKey(1)))
				suite.Require().Equal(uint64(1), stored)
			},
		},
		{
			name: "existing post id is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)
			},
			subspaceID: 1,
			postID:     2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetPostIDFromBytes(store.Get(types.NextPostIDStoreKey(1)))
				suite.Require().Equal(uint64(2), stored)
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

			suite.k.SetNextPostID(ctx, tc.subspaceID, tc.postID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestKeeper_GetNextPostID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		shouldErr  bool
		expPostID  uint64
	}{
		{
			name:       "not found post id returns error",
			subspaceID: 1,
			shouldErr:  true,
		},
		{
			name: "found post id returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)
			},
			subspaceID: 1,
			shouldErr:  false,
			expPostID:  1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			postID, err := suite.k.GetNextPostID(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expPostID, postID)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteNextPostID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing post id is deleted properly",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextPostIDStoreKey(1)))
			},
		},
		{
			name: "existing post id is deleted properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.NextPostIDStoreKey(1), types.GetPostIDBytes(1))
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextPostIDStoreKey(1)))
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

			suite.k.DeleteNextPostID(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_ValidatePostReference() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		postOwner   string
		subspaceID  uint64
		referenceID uint64
		shouldErr   bool
	}{
		{
			name:        "non existing referenced post returns error",
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "blocked post author returns error",
			setup: func() {
				suite.rk.EXPECT().HasUserBlocked(gomock.Any(),
					"cosmos1fvnkn5yjhdc6sxwlph8e98udw8nsly0w9yznrk",
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"This is a long post text to make sure tags are valid",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1fvnkn5yjhdc6sxwlph8e98udw8nsly0w9yznrk",
				))
			},
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "valid post reference returns no error",
			setup: func() {
				suite.rk.EXPECT().HasUserBlocked(gomock.Any(),
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					uint64(1)).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"This is a long post text to make sure tags are valid",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   false,
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

			err := suite.k.ValidatePostReference(ctx, tc.postOwner, tc.subspaceID, tc.referenceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_ValidatePostReply() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		postOwner   string
		subspaceID  uint64
		referenceID uint64
		shouldErr   bool
	}{
		{
			name:        "reply post not found returns error",
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_FOLLOWERS and not follower returns error",
			setup: func() {
				suite.rk.EXPECT().HasRelationship(gomock.Any(),
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					uint64(1)).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"This is a test post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_FOLLOWERS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_FOLLOWERS and follower returns no error",
			setup: func() {
				suite.rk.EXPECT().HasRelationship(gomock.Any(),
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"This is a test post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_FOLLOWERS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   false,
		},
		{
			name: "REPLY_SETTING_MUTUAL and not mutual returns error",
			setup: func() {
				suite.rk.EXPECT().HasRelationship(gomock.Any(),
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					uint64(1)).Return(true)
				suite.rk.EXPECT().HasRelationship(gomock.Any(),
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					uint64(1)).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"This is a test post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_MUTUAL,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_MUTUAL and mutual returns no error",
			setup: func() {
				suite.rk.EXPECT().HasRelationship(gomock.Any(),
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					uint64(1)).Return(true)
				suite.rk.EXPECT().HasRelationship(gomock.Any(),
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"This is a test post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_MUTUAL,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   false,
		},
		{
			name: "REPLY_SETTING_MENTIONS and no mention returns error",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"This is a test post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_MENTIONS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_MENTIONS and mention returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"This is a test post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					types.NewEntities(nil, []types.TextTag{
						types.NewTextTag(0, 44, "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g"),
					}, nil),
					nil,
					nil,
					types.REPLY_SETTING_MENTIONS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			postOwner:   "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   false,
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

			err := suite.k.ValidatePostReply(ctx, tc.postOwner, tc.subspaceID, tc.referenceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_ValidatePost() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		post      types.Post
		shouldErr bool
	}{
		{
			name: "invalid text length returns error",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(1))
			},
			post: types.NewPost(
				1,
				0,
				2,
				"External id",
				"Text",
				"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				1,
				nil,
				nil,
				nil,
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "post with invalid conversation id returns error",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			post: types.NewPost(
				1,
				0,
				2,
				"External id",
				"Text",
				"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				1,
				nil,
				nil,
				nil,
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "post with invalid reference returns error",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			post: types.NewPost(
				1,
				0,
				2,
				"External id",
				"Text",
				"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				0,
				nil,
				nil,
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "invalid post returns error",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(1))
			},
			post: types.NewPost(
				1,
				0,
				0,
				"External id",
				"Text",
				"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				0,
				nil,
				nil,
				nil,
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: true,
		},
		{
			name: "valid post returns no error",
			setup: func() {
				suite.rk.EXPECT().HasUserBlocked(gomock.Any(),
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					uint64(1)).
					Return(false).
					AnyTimes()
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"This is a long post text to make sure tags are valid",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			post: types.NewPost(
				1,
				0,
				2,
				"External id",
				"This is a long post text to make sure tags are valid",
				"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				1,
				types.NewEntities(
					[]types.TextTag{
						types.NewTextTag(1, 3, "tag"),
					},
					[]types.TextTag{
						types.NewTextTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]string{"generic"},
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_REPLY, 1, 0),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			shouldErr: false,
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

			err := suite.k.ValidatePost(ctx, tc.post)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SavePost() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		post  types.Post
		check func(ctx sdk.Context)
	}{
		{
			name: "non existing post is saved properly",
			post: types.NewPost(
				1,
				0,
				2,
				"External id",
				"Text",
				"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				1,
				types.NewEntities(
					[]types.TextTag{
						types.NewTextTag(1, 3, "tag"),
					},
					[]types.TextTag{
						types.NewTextTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]string{"generic"},
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			check: func(ctx sdk.Context) {
				// Check the post exists
				stored, found := suite.k.GetPost(ctx, 1, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPost(
					1,
					0,
					2,
					"External id",
					"Text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					types.NewEntities(
						[]types.TextTag{
							types.NewTextTag(1, 3, "tag"),
						},
						[]types.TextTag{
							types.NewTextTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]string{"generic"},
					[]types.PostReference{
						types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				), stored)

				store := ctx.KVStore(suite.storeKey)

				// Make sure the section reference is properly set
				suite.Require().True(store.Has(types.PostSectionStoreKey(1, 0, 2)))

				// Make sure the attachment it is generated properly
				attachmentID := types.GetAttachmentIDFromBytes(store.Get(types.NextAttachmentIDStoreKey(1, 2)))
				suite.Require().Equal(uint32(1), attachmentID)
			},
		},
		{
			name: "existing post is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"External id",
					"Text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					types.NewEntities(
						[]types.TextTag{
							types.NewTextTag(1, 3, "tag"),
						},
						[]types.TextTag{
							types.NewTextTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]string{"generic"},
					[]types.PostReference{
						types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			post: types.NewPost(
				1,
				0,
				2,
				"External id",
				"This is a new text",
				"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				1,
				types.NewEntities(
					[]types.TextTag{
						types.NewTextTag(1, 3, "tag"),
					},
					[]types.TextTag{
						types.NewTextTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]string{"generic"},
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
			),
			check: func(ctx sdk.Context) {
				// Make sure the post is saved properly
				stored, found := suite.k.GetPost(ctx, 1, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPost(
					1,
					0,
					2,
					"External id",
					"This is a new text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					types.NewEntities(
						[]types.TextTag{
							types.NewTextTag(1, 3, "tag"),
						},
						[]types.TextTag{
							types.NewTextTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]string{"generic"},
					[]types.PostReference{
						types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				), stored)

				store := ctx.KVStore(suite.storeKey)

				// Make sure the section reference is stored
				suite.Require().True(store.Has(types.PostSectionStoreKey(1, 0, 2)))

				// Make sure the attachment it is generated properly
				attachmentID := types.GetAttachmentIDFromBytes(store.Get(types.NextAttachmentIDStoreKey(1, 2)))
				suite.Require().Equal(uint32(1), attachmentID)
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

			suite.k.SavePost(ctx, tc.post)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasPost() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		expResult  bool
	}{
		{
			name:       "non existing post returns false",
			subspaceID: 1,
			postID:     1,
			expResult:  false,
		},
		{
			name: "existing post returns true",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"External id",
					"Text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			subspaceID: 1,
			postID:     2,
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

			result := suite.k.HasPost(ctx, tc.subspaceID, tc.postID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetPost() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		expFound   bool
		expPost    types.Post
	}{
		{
			name:       "non existing post returns false and empty post",
			subspaceID: 1,
			postID:     1,
			expFound:   false,
			expPost:    types.Post{},
		},
		{
			name: "existing post returns correct value and true",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"External id",
					"Text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					types.NewEntities(
						[]types.TextTag{
							types.NewTextTag(1, 3, "tag"),
						},
						[]types.TextTag{
							types.NewTextTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]string{"generic"},
					[]types.PostReference{
						types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))
			},
			subspaceID: 1,
			postID:     2,
			expFound:   true,
			expPost: types.NewPost(
				1,
				0,
				2,
				"External id",
				"Text",
				"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
				1,
				types.NewEntities(
					[]types.TextTag{
						types.NewTextTag(1, 3, "tag"),
					},
					[]types.TextTag{
						types.NewTextTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]string{"generic"},
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
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

			post, found := suite.k.GetPost(ctx, tc.subspaceID, tc.postID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expPost, post)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeletePost() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing post is deleted properly",
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPost(ctx, 1, 1))
			},
		},
		{
			name: "existing post is deleted along with attachments",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"External id",
					"Text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				))

				suite.k.SetNextAttachmentID(ctx, 1, 2, 2)

				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 2, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			subspaceID: 1,
			postID:     2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(suite.k.HasPost(ctx, 1, 2))
				suite.Require().False(store.Has(types.PostSectionStoreKey(1, 0, 1)))
				suite.Require().False(suite.k.HasNextAttachmentID(ctx, 1, 2))
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 2, 1))
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

			suite.k.DeletePost(ctx, tc.subspaceID, tc.postID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
