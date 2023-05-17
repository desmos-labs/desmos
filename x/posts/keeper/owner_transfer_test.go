package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SavePostOwnerTransferRequest() {
	testCases := []struct {
		name    string
		store   func(ctx sdk.Context)
		request types.PostOwnerTransferRequest
		check   func(ctx sdk.Context)
	}{
		{
			name: "non existing request is stored properly",
			request: types.NewPostOwnerTransferRequest(
				1,
				1,
				"receiver",
				"sender",
			),
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				var request types.PostOwnerTransferRequest
				suite.cdc.MustUnmarshal(
					store.Get(types.PostOwnerTransferRequestStoreKey(1, 1)),
					&request,
				)

				suite.Require().Equal(types.NewPostOwnerTransferRequest(
					1,
					1,
					"receiver",
					"sender",
				), request)
			},
		},
		{
			name: "existing request is overwritten properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"other_receiver",
					"sender",
				))
			},
			request: types.NewPostOwnerTransferRequest(
				1,
				1,
				"receiver",
				"sender",
			),
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				var request types.PostOwnerTransferRequest
				suite.cdc.MustUnmarshal(
					store.Get(types.PostOwnerTransferRequestStoreKey(1, 1)),
					&request,
				)

				suite.Require().Equal(types.NewPostOwnerTransferRequest(
					1,
					1,
					"receiver",
					"sender",
				), request)
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

			suite.k.SavePostOwnerTransferRequest(ctx, tc.request)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasPostOwnerTransferRequest() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		expResult  bool
	}{
		{
			name:       "not found request returns false",
			subspaceID: 1,
			postID:     1,
			expResult:  false,
		},
		{
			name: "found request returns true",
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"receiver",
					"sender",
				))
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

			found := suite.k.HasPostOwnerTransferRequest(ctx, tc.subspaceID, tc.postID)
			suite.Require().Equal(tc.expResult, found)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeletePostOwnerTransferRequest() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "non existing request is deleted properly",
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.PostOwnerTransferRequestStoreKey(1, 1)))
			},
		},
		{
			name: "existing request is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
					1,
					1,
					"receiver",
					"sender",
				))
			},
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.PostOwnerTransferRequestStoreKey(1, 1)))
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

			suite.k.DeletePostOwnerTransferRequest(ctx, tc.subspaceID, tc.postID)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
