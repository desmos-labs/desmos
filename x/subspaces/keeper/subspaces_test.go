package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetSubspaceID() {
	testCases := []struct {
		name  string
		id    uint64
		check func(ctx sdk.Context)
	}{
		{
			name: "zero subspace id",
			id:   0,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().Equal(uint64(0), types.GetSubspaceIDFromBytes(store.Get(types.SubspaceIDKey)))
			},
		},
		{
			name: "non-zero subspace id",
			id:   5,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().Equal(uint64(5), types.GetSubspaceIDFromBytes(store.Get(types.SubspaceIDKey)))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			suite.k.SetSubspaceID(ctx, tc.id)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetSubspaceID() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		expID     uint64
	}{
		{
			name:      "subspace id not set",
			shouldErr: true,
		},
		{
			name: "subspace id set",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.SubspaceIDKey, types.GetSubspaceIDBytes(1))
			},
			shouldErr: false,
			expID:     1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			id, err := suite.k.GetSubspaceID(ctx)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expID, id)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SaveSubspace() {
	testCases := []struct {
		name     string
		store    func(ctx sdk.Context)
		subspace types.Subspace
		check    func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace is stored properly",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				subspace, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), subspace)
			},
		},
		{
			name: "existing subspace is replaced correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspace: types.NewSubspace(
				1,
				"Test subspace with another name and owner",
				"This is a test subspace with a changed description",
				"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				subspace, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), subspace)
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

			suite.k.SaveSubspace(ctx, tc.subspace)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasSubspace() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		expResult  bool
	}{
		{
			name:       "not found subspace returns false",
			subspaceID: 1,
			groupID:    1,
			expResult:  false,
		},
		{
			name: "found subspace returns the correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			subspaceID: 1,
			groupID:    1,
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

			result := suite.k.HasUserGroup(ctx, tc.subspaceID, tc.groupID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetSubspace() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		id          uint64
		expSubspace types.Subspace
		expFound    bool
	}{
		{
			name:     "not found subspace returns false",
			id:       1,
			expFound: false,
		},
		{
			name: "found subspace returns the correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			id:       1,
			expFound: true,
			expSubspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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

			subspace, found := suite.k.GetSubspace(ctx, tc.id)
			suite.Require().Equal(tc.expFound, found)
			if tc.expFound {
				suite.Require().Equal(tc.expSubspace, subspace)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteSubspace() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing subspace is deleted properly",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				found := suite.k.HasSubspace(ctx, 1)
				suite.Require().False(found)
			},
		},
		{
			name: "existing subspace is deleted and all groups and permissions are removed",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionWrite)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				found := suite.k.HasSubspace(ctx, 1)
				suite.Require().False(found)

				groups := suite.k.GetSubspaceGroups(ctx, 1)
				suite.Require().Empty(groups)

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e")
				suite.Require().NoError(err)
				permission := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Equal(types.PermissionNothing, permission)
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

			suite.k.DeleteSubspace(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
