package keeper_test

import (
	"time"

	relationshipstypes "github.com/desmos-labs/desmos/v3/x/relationships/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetNextPostID() {
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

func (suite *KeeperTestsuite) TestKeeper_GetNextPostID() {
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

func (suite *KeeperTestsuite) TestKeeper_DeleteNextPostID() {
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

func (suite *KeeperTestsuite) TestKeeper_ValidatePostReference() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		postAuthor  string
		subspaceID  uint64
		referenceID uint64
		shouldErr   bool
	}{
		{
			name:        "non existing referenced post returns error",
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "blocked post author returns error",
			store: func(ctx sdk.Context) {
				suite.rk.SaveUserBlock(ctx, relationshipstypes.NewUserBlock(
					"cosmos1fvnkn5yjhdc6sxwlph8e98udw8nsly0w9yznrk",
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"",
					1,
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External id",
					"This is a long post text to make sure tags are valid",
					"cosmos1fvnkn5yjhdc6sxwlph8e98udw8nsly0w9yznrk",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "valid post reference returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External id",
					"This is a long post text to make sure tags are valid",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.ValidatePostReference(ctx, tc.postAuthor, tc.subspaceID, tc.referenceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_ValidatePostReply() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		postAuthor  string
		subspaceID  uint64
		referenceID uint64
		shouldErr   bool
	}{
		{
			name:        "reply post not found returns error",
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_FOLLOWERS and not follower returns error",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"This is a test post",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					types.REPLY_SETTING_FOLLOWERS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_FOLLOWERS and follower returns no error",
			store: func(ctx sdk.Context) {
				suite.rk.SaveRelationship(ctx, relationshipstypes.NewRelationship(
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"This is a test post",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					types.REPLY_SETTING_FOLLOWERS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   false,
		},
		{
			name: "REPLY_SETTING_MUTUAL and not mutual returns error",
			store: func(ctx sdk.Context) {
				suite.rk.SaveRelationship(ctx, relationshipstypes.NewRelationship(
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"This is a test post",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					types.REPLY_SETTING_MENTIONS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_MUTUAL and mutual returns no error",
			store: func(ctx sdk.Context) {
				suite.rk.SaveRelationship(ctx, relationshipstypes.NewRelationship(
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
				))

				suite.rk.SaveRelationship(ctx, relationshipstypes.NewRelationship(
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					1,
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"This is a test post",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					types.REPLY_SETTING_MENTIONS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_MENTIONS and no mention returns error",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"This is a test post",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					types.REPLY_SETTING_MENTIONS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   true,
		},
		{
			name: "REPLY_SETTING_MENTIONS and mention returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
					"This is a test post",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					types.NewEntities(nil, []types.Tag{
						types.NewTag(0, 44, "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g"),
					}, nil),
					nil,
					types.REPLY_SETTING_MENTIONS,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			postAuthor:  "cosmos1t457f629cc3ykftepjejgzxv0vmz5dw2gn940g",
			subspaceID:  1,
			referenceID: 1,
			shouldErr:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.ValidatePostReply(ctx, tc.postAuthor, tc.subspaceID, tc.referenceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_ValidatePost() {
	testCases := []struct {
		name      string
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
				2,
				"External id",
				"Text",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 3, "tag"),
					},
					[]types.Tag{
						types.NewTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
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
				2,
				"External id",
				"Text",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 3, "tag"),
					},
					[]types.Tag{
						types.NewTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				nil,
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
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
				2,
				"External id",
				"Text",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				0,
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 3, "tag"),
					},
					[]types.Tag{
						types.NewTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
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
				"External id",
				"Text",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				0,
				nil,
				nil,
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "valid post returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External id",
					"This is a long post text to make sure tags are valid",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			post: types.NewPost(
				1,
				2,
				"External id",
				"This is a long post text to make sure tags are valid",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 3, "tag"),
					},
					[]types.Tag{
						types.NewTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]types.PostReference{
					types.NewPostReference(types.TYPE_REPLIED_TO, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
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

func (suite *KeeperTestsuite) TestKeeper_SavePost() {
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
				2,
				"External id",
				"Text",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 3, "tag"),
					},
					[]types.Tag{
						types.NewTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			check: func(ctx sdk.Context) {
				// Check the post exists
				stored, found := suite.k.GetPost(ctx, 1, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPost(
					1,
					2,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					types.NewEntities(
						[]types.Tag{
							types.NewTag(1, 3, "tag"),
						},
						[]types.Tag{
							types.NewTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]types.PostReference{
						types.NewPostReference(types.TYPE_QUOTED, 1),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				), stored)

				// Make sure the attachment it is generated properly
				store := ctx.KVStore(suite.storeKey)
				attachmentID := types.GetAttachmentIDFromBytes(store.Get(types.NextAttachmentIDStoreKey(1, 2)))
				suite.Require().Equal(uint32(1), attachmentID)
			},
		},
		{
			name: "existing post is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					2,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					types.NewEntities(
						[]types.Tag{
							types.NewTag(1, 3, "tag"),
						},
						[]types.Tag{
							types.NewTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]types.PostReference{
						types.NewPostReference(types.TYPE_QUOTED, 1),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			post: types.NewPost(
				1,
				2,
				"External id",
				"This is a new text",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 3, "tag"),
					},
					[]types.Tag{
						types.NewTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			check: func(ctx sdk.Context) {
				// Make sure the post is saved properly
				stored, found := suite.k.GetPost(ctx, 1, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPost(
					1,
					2,
					"External id",
					"This is a new text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					types.NewEntities(
						[]types.Tag{
							types.NewTag(1, 3, "tag"),
						},
						[]types.Tag{
							types.NewTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]types.PostReference{
						types.NewPostReference(types.TYPE_QUOTED, 1),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				), stored)

				// Make sure the attachment it is generated properly
				store := ctx.KVStore(suite.storeKey)
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

func (suite *KeeperTestsuite) TestKeeper_HasPost() {
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
					2,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					types.NewEntities(
						[]types.Tag{
							types.NewTag(1, 3, "tag"),
						},
						[]types.Tag{
							types.NewTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]types.PostReference{
						types.NewPostReference(types.TYPE_QUOTED, 1),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
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

func (suite *KeeperTestsuite) TestKeeper_GetPost() {
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
					2,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					types.NewEntities(
						[]types.Tag{
							types.NewTag(1, 3, "tag"),
						},
						[]types.Tag{
							types.NewTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]types.PostReference{
						types.NewPostReference(types.TYPE_QUOTED, 1),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			subspaceID: 1,
			postID:     2,
			expFound:   true,
			expPost: types.NewPost(
				1,
				2,
				"External id",
				"Text",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 3, "tag"),
					},
					[]types.Tag{
						types.NewTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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

			post, found := suite.k.GetPost(ctx, tc.subspaceID, tc.postID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expPost, post)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeletePost() {
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
					2,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					types.NewEntities(
						[]types.Tag{
							types.NewTag(1, 3, "tag"),
						},
						[]types.Tag{
							types.NewTag(4, 6, "tag"),
						},
						[]types.Url{
							types.NewURL(7, 9, "URL", "Display URL"),
						},
					),
					[]types.PostReference{
						types.NewPostReference(types.TYPE_QUOTED, 1),
					},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
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
				suite.Require().False(suite.k.HasPost(ctx, 1, 2))
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
