package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/reports/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetNextReasonID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reasonID   uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing reason id is set properly",
			subspaceID: 1,
			reasonID:   1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetReasonIDFromBytes(store.Get(types.NextReasonIDStoreKey(1)))
				suite.Require().Equal(uint32(1), stored)
			},
		},
		{
			name: "existing reason id is overridden properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.NextReasonIDStoreKey(1), types.GetReasonIDBytes(1))
			},
			subspaceID: 1,
			reasonID:   2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetReasonIDFromBytes(store.Get(types.NextReasonIDStoreKey(1)))
				suite.Require().Equal(uint32(2), stored)
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

			suite.k.SetNextReasonID(ctx, tc.subspaceID, tc.reasonID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetNextReasonID() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		shouldErr   bool
		expReasonID uint32
	}{
		{
			name:       "non existing reason id returns error",
			subspaceID: 1,
			shouldErr:  true,
		},
		{
			name: "existing reason id is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 1)
			},
			subspaceID:  1,
			shouldErr:   false,
			expReasonID: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			reasonID, err := suite.k.GetNextReasonID(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReasonID, reasonID)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteNextReasonID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing reason id is deleted properly",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextReasonIDStoreKey(1)))
			},
		},
		{
			name: "existing reason id is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReasonID(ctx, 1, 1)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextReasonIDStoreKey(1)))
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

			suite.k.DeleteNextReasonID(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_SaveReason() {
	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		reason types.Reason
		check  func(ctx sdk.Context)
	}{
		{
			name: "non existing reason is stored properly",
			reason: types.NewReason(
				1,
				1,
				"Spam",
				"This content is spam",
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReason(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				), stored)
			},
		},
		{
			name: "existing reason is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			reason: types.NewReason(
				1,
				1,
				"Self harm",
				"This content contains self harm",
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReason(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReason(
					1,
					1,
					"Self harm",
					"This content contains self harm",
				), stored)
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

			suite.k.SaveReason(ctx, tc.reason)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasReason() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reasonID   uint32
		expResult  bool
	}{
		{
			name:       "non existing reason returns false",
			subspaceID: 1,
			reasonID:   1,
			expResult:  false,
		},
		{
			name: "existing reason returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			subspaceID: 1,
			reasonID:   1,
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

			result := suite.k.HasReason(ctx, tc.subspaceID, tc.reasonID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetReason() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reasonID   uint32
		expFound   bool
		expReason  types.Reason
	}{
		{
			name:       "non existing reason returns false and empty reason",
			subspaceID: 1,
			reasonID:   1,
			expFound:   false,
			expReason:  types.Reason{},
		},
		{
			name: "existing reason returns true and correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))
			},
			subspaceID: 1,
			reasonID:   1,
			expFound:   true,
			expReason: types.NewReason(
				1,
				1,
				"Spam",
				"This content is spam",
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

			reason, found := suite.k.GetReason(ctx, tc.subspaceID, tc.reasonID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expReason, reason)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteReason() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		reasonID   uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing reason is deleted properly",
			subspaceID: 1,
			reasonID:   1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasReason(ctx, 1, 1))
			},
		},
		{
			name: "existing reason is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveReason(ctx, types.NewReason(
					1,
					1,
					"Spam",
					"This content is spam",
				))

				suite.k.SaveReport(ctx, types.NewReport(
					1,
					1,
					[]uint32{1},
					"This content is spam",
					types.NewPostTarget(1),
					"cosmos1zkmf50jq4lzvhvp5ekl0sdf2p4g3v9v8edt24z",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			reasonID:   1,
			check: func(ctx sdk.Context) {
				// Make sure the reason is deleted
				suite.Require().False(suite.k.HasReason(ctx, 1, 1))

				// Make sure the associated reports are deleted
				suite.Require().False(suite.k.HasReport(ctx, 1, 1))
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

			suite.k.DeleteReason(ctx, tc.subspaceID, tc.reasonID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
