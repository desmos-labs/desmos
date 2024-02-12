package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/reactions/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetNextRegisteredReactionID() {
	testCases := []struct {
		name                 string
		store                func(ctx sdk.Context)
		subspaceID           uint64
		registeredReactionID uint32
		check                func(ctx sdk.Context)
	}{
		{
			name:                 "non existing next registered reaction id is set properly",
			subspaceID:           1,
			registeredReactionID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetRegisteredReactionIDFromBytes(store.Get(types.NextRegisteredReactionIDStoreKey(1)))
				suite.Require().Equal(uint32(1), stored)
			},
		},
		{
			name: "existing next registered reaction id is overridden properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.NextRegisteredReactionIDStoreKey(1), types.GetRegisteredReactionIDBytes(1))
			},
			subspaceID:           1,
			registeredReactionID: 2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetRegisteredReactionIDFromBytes(store.Get(types.NextRegisteredReactionIDStoreKey(1)))
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

			suite.k.SetNextRegisteredReactionID(ctx, tc.subspaceID, tc.registeredReactionID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasNextRegisteredReactionID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		expResult  bool
	}{
		{
			name:       "non existing next registered reaction id returns false",
			subspaceID: 1,
			expResult:  false,
		},
		{
			name: "existing next registered reaction id returns true",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.NextRegisteredReactionIDStoreKey(1), types.GetRegisteredReactionIDBytes(1))
			},
			subspaceID: 1,
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

			result := suite.k.HasNextRegisteredReactionID(ctx, tc.subspaceID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetNextRegisteredReactionID() {
	testCases := []struct {
		name                    string
		store                   func(ctx sdk.Context)
		subspaceID              uint64
		shouldErr               bool
		expRegisteredReactionID uint32
	}{
		{
			name:       "non existing next registered reaction id returns error",
			subspaceID: 1,
			shouldErr:  true,
		},
		{
			name: "existing next registered reaction id is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextRegisteredReactionID(ctx, 1, 1)
			},
			subspaceID:              1,
			shouldErr:               false,
			expRegisteredReactionID: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			registeredReactionID, err := suite.k.GetNextRegisteredReactionID(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRegisteredReactionID, registeredReactionID)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteNextRegisteredReactionID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing next registered reaction id is deleted properly",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextRegisteredReactionIDStoreKey(1)))
			},
		},
		{
			name: "existing next registered reaction id is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextRegisteredReactionID(ctx, 1, 1)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextRegisteredReactionIDStoreKey(1)))
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

			suite.k.DeleteNextRegisteredReactionID(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_SaveRegisteredReaction() {
	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		reaction types.RegisteredReaction
		check    func(ctx sdk.Context)
	}{
		{
			name: "non existing reaction is stored properly",
			reaction: types.NewRegisteredReaction(
				1,
				1,
				":hello:",
				"https://example.com?image=hello.png",
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetRegisteredReaction(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				), stored)
			},
		},
		{
			name: "existing reaction is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			reaction: types.NewRegisteredReaction(
				1,
				1,
				":hello:",
				"https://example.com?image=hello2.png",
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetRegisteredReaction(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello2.png",
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

			suite.k.SaveRegisteredReaction(ctx, tc.reaction)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasRegisteredReaction() {
	testCases := []struct {
		name                 string
		store                func(ctx sdk.Context)
		subspaceID           uint64
		registeredReactionID uint32
		expResult            bool
	}{
		{
			name:                 "non existing reaction returns false",
			subspaceID:           1,
			registeredReactionID: 1,
			expResult:            false,
		},
		{
			name: "existing reaction returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			subspaceID:           1,
			registeredReactionID: 1,
			expResult:            true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasRegisteredReaction(ctx, tc.subspaceID, tc.registeredReactionID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetRegisteredReaction() {
	testCases := []struct {
		name                  string
		store                 func(ctx sdk.Context)
		subspaceID            uint64
		registeredReactionID  uint32
		expFound              bool
		expRegisteredReaction types.RegisteredReaction
	}{
		{
			name:                  "non existing reaction returns false and empty reaction",
			subspaceID:            1,
			registeredReactionID:  1,
			expFound:              false,
			expRegisteredReaction: types.RegisteredReaction{},
		},
		{
			name: "existing reaction returns true and correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			subspaceID:           1,
			registeredReactionID: 1,
			expFound:             true,
			expRegisteredReaction: types.NewRegisteredReaction(
				1,
				1,
				":hello:",
				"https://example.com?image=hello.png",
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

			reaction, found := suite.k.GetRegisteredReaction(ctx, tc.subspaceID, tc.registeredReactionID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expRegisteredReaction, reaction)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRegisteredReaction() {
	testCases := []struct {
		name                 string
		store                func(ctx sdk.Context)
		subspaceID           uint64
		registeredReactionID uint32
		check                func(ctx sdk.Context)
	}{
		{
			name:                 "non existing reaction is deleted properly",
			subspaceID:           1,
			registeredReactionID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasRegisteredReaction(ctx, 1, 1))
			},
		},
		{
			name: "existing reaction is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			subspaceID:           1,
			registeredReactionID: 1,
			check: func(ctx sdk.Context) {
				// Make sure the reaction is deleted
				suite.Require().False(suite.k.HasRegisteredReaction(ctx, 1, 1))
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

			suite.k.DeleteRegisteredReaction(ctx, tc.subspaceID, tc.registeredReactionID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
