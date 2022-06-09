package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetNextReactionID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		reactionID uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing next reaction id is set properly",
			subspaceID: 1,
			postID:     1,
			reactionID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetReactionIDFromBytes(store.Get(types.NextReactionIDStoreKey(1, 1)))
				suite.Require().Equal(uint32(1), stored)
			},
		},
		{
			name: "existing next reaction id is overridden properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.NextReactionIDStoreKey(1, 1), types.GetReactionIDBytes(1))
			},
			subspaceID: 1,
			postID:     1,
			reactionID: 2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetReactionIDFromBytes(store.Get(types.NextReactionIDStoreKey(1, 1)))
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

			suite.k.SetNextReactionID(ctx, tc.subspaceID, tc.postID, tc.reactionID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasNextReactionID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		expResult  bool
	}{
		{
			name:       "non existing next reaction id returns false",
			subspaceID: 1,
			postID:     1,
			expResult:  false,
		},
		{
			name: "existing next reaction id returns true",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.NextReactionIDStoreKey(1, 1), types.GetReactionIDBytes(1))
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

			result := suite.k.HasNextReactionID(ctx, tc.subspaceID, tc.postID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetNextReactionID() {
	testCases := []struct {
		name          string
		store         func(ctx sdk.Context)
		subspaceID    uint64
		postID        uint64
		shouldErr     bool
		expReactionID uint32
	}{
		{
			name:       "non existing next reaction id returns error",
			subspaceID: 1,
			postID:     1,
			shouldErr:  true,
		},
		{
			name: "existing next reaction id is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReactionID(ctx, 1, 1, 1)
			},
			subspaceID:    1,
			postID:        1,
			shouldErr:     false,
			expReactionID: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			reactionID, err := suite.k.GetNextReactionID(ctx, tc.subspaceID, tc.postID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expReactionID, reactionID)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteNextReactionID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing next reaction id is deleted properly",
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextReactionIDStoreKey(1, 1)))
			},
		},
		{
			name: "existing next reaction id is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextReactionID(ctx, 1, 1, 1)
			},
			subspaceID: 1,
			postID:     1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextReactionIDStoreKey(1, 1)))
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

			suite.k.DeleteNextReactionID(ctx, tc.subspaceID, tc.postID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_ValidateReaction() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		reaction  types.Reaction
		shouldErr bool
	}{
		{
			name: "invalid reaction returns error",
			reaction: types.NewReaction(
				0,
				1,
				1,
				types.NewFreeTextValue(""),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "RegisteredReactionValue - not enabled reactions returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(false),
					types.NewFreeTextValueParams(false, 1, ""),
				))
			},
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "RegisteredReactionValue - non exiting registered reaction returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(false, 1, ""),
				))
			},
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "RegisteredReactionValue - valid value returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(false, 1, ""),
				))
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello",
					"https://examplle.com?image=hello.png",
				))
			},
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: false,
		},
		{
			name: "FreeTextValue - not enabled reactions returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(false),
					types.NewFreeTextValueParams(false, 1, ""),
				))
			},
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewFreeTextValue("Wow!"),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "FreeTextValue - too long value length returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(false),
					types.NewFreeTextValueParams(true, 1, ""),
				))
			},
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewFreeTextValue("Wow!"),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "FreeTextValue - not matching regex value returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(false),
					types.NewFreeTextValueParams(true, 10, "[a-z]"),
				))
			},
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewFreeTextValue("🚀"),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "FreeTextValue - valid value returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(false),
					types.NewFreeTextValueParams(true, 10, ""),
				))
			},
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewFreeTextValue("🚀"),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
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

			err := suite.k.ValidateReaction(ctx, tc.reaction)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveReaction() {
	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		reaction types.Reaction
		check    func(ctx sdk.Context)
	}{
		{
			name: "non existing reaction is stored properly",
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReaction(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				), stored)
			},
		},
		{
			name: "existing reaction is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			reaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(2),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReaction(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(2),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
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

			suite.k.SaveReaction(ctx, tc.reaction)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasReaction() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		reactionID uint32
		expResult  bool
	}{
		{
			name:       "non existing reaction returns false",
			subspaceID: 1,
			postID:     1,
			reactionID: 1,
			expResult:  false,
		},
		{
			name: "existing reaction returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			subspaceID: 1,
			postID:     1,
			reactionID: 1,
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

			result := suite.k.HasReaction(ctx, tc.subspaceID, tc.postID, tc.reactionID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetReaction() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		postID      uint64
		reactionID  uint32
		expFound    bool
		expReaction types.Reaction
	}{
		{
			name:        "non existing reaction returns false and empty reaction",
			subspaceID:  1,
			postID:      1,
			reactionID:  1,
			expFound:    false,
			expReaction: types.Reaction{},
		},
		{
			name: "existing reaction returns true and correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			subspaceID: 1,
			postID:     1,
			reactionID: 1,
			expFound:   true,
			expReaction: types.NewReaction(
				1,
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
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

			reaction, found := suite.k.GetReaction(ctx, tc.subspaceID, tc.postID, tc.reactionID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expReaction, reaction)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteReaction() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		reactionID uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing reaction is deleted properly",
			subspaceID: 1,
			postID:     1,
			reactionID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasReaction(ctx, 1, 1, 1))
			},
		},
		{
			name: "existing reaction is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			subspaceID: 1,
			postID:     1,
			reactionID: 1,
			check: func(ctx sdk.Context) {
				// Make sure the reaction is deleted
				suite.Require().False(suite.k.HasReaction(ctx, 1, 1, 1))
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

			suite.k.DeleteReaction(ctx, tc.subspaceID, tc.postID, tc.reactionID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
