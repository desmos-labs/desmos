package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/posts/types"
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
				request, _ := suite.k.GetPostOwnerTransferRequest(ctx, 1, 1)
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
				request, _ := suite.k.GetPostOwnerTransferRequest(ctx, 1, 1)
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

func (suite *KeeperTestSuite) TestKeeper_GetPostOwnerTransferRequest() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		expFound   bool
		expRequest types.PostOwnerTransferRequest
	}{
		{
			name:       "non existing request returns false and empty request",
			subspaceID: 1,
			postID:     1,
			expFound:   false,
			expRequest: types.PostOwnerTransferRequest{},
		},
		{
			name: "existing request returns correct value and true",
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
			expFound:   true,
			expRequest: types.NewPostOwnerTransferRequest(
				1,
				1,
				"receiver",
				"sender",
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

			request, found := suite.k.GetPostOwnerTransferRequest(ctx, tc.subspaceID, tc.postID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expRequest, request)
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
