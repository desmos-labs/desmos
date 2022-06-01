package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetNextSectionID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing next section id is set properly",
			subspaceID: 1,
			sectionID:  1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetSectionIDFromBytes(store.Get(types.NextSectionIDStoreKey(1)))
				suite.Require().Equal(uint32(1), stored)
			},
		},
		{
			name: "existing next section id is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextSectionID(ctx, 1, 1)
			},
			subspaceID: 1,
			sectionID:  2,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				stored := types.GetSectionIDFromBytes(store.Get(types.NextSectionIDStoreKey(1)))
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

			suite.k.SetNextSectionID(ctx, tc.subspaceID, tc.sectionID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasNextSectionID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		expResult  bool
	}{
		{
			name:       "non existing next section id returns false",
			subspaceID: 1,
			expResult:  false,
		},
		{
			name: "existing next section id returns true",
			store: func(ctx sdk.Context) {
				suite.k.SetNextSectionID(ctx, 1, 1)
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

			result := suite.k.HasNextSectionID(ctx, tc.subspaceID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetNextSectionID() {
	testCases := []struct {
		name             string
		store            func(ctx sdk.Context)
		subspaceID       uint64
		shouldErr        bool
		expNextSectionID uint32
	}{
		{
			name:       "non existing next section id returns error",
			subspaceID: 1,
			shouldErr:  true,
		},
		{
			name: "existing next section id is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextSectionID(ctx, 1, 1)
			},
			subspaceID:       1,
			shouldErr:        false,
			expNextSectionID: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			nextSectionID, err := suite.k.GetNextSectionID(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expNextSectionID, nextSectionID)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteNextSectionID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing next section id is deleted properly",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextSectionIDStoreKey(1)))
			},
		},
		{
			name: "existing next section is is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextSectionID(ctx, 1, 1)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.NextSectionIDStoreKey(1)))
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

			suite.k.DeleteNextSectionID(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SaveSection() {
	testCases := []struct {
		name    string
		store   func(ctx sdk.Context)
		section types.Section
		check   func(ctx sdk.Context)
	}{
		{
			name: "non existing section is stored properly",
			section: types.NewSection(
				1,
				1,
				0,
				"Test section",
				"This is a test section",
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetSection(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				), stored)
			},
		},
		{
			name: "existing section is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				))
			},
			section: types.NewSection(
				1,
				1,
				1,
				"Edited section",
				"This is an edited section",
			),
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetSection(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSection(
					1,
					1,
					1,
					"Edited section",
					"This is an edited section",
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

			suite.k.SaveSection(ctx, tc.section)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_IsSectionPathValid() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		expResult  bool
	}{
		{
			name: "section with circular path returns false",
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

				// Create the following subspaces sections
				//     Root
				//
				//      A
				//     / \
				//    B - C
				suite.k.SaveSection(ctx, types.DefaultSection(1))
				suite.k.SaveSection(ctx, types.NewSection(1, 1, 3, "A", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 2, 1, "B", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 3, 2, "C", ""))
			},
			subspaceID: 1,
			sectionID:  3,
			expResult:  false,
		},
		{
			name: "section with relative path returns false",
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

				// Create the following subspaces sections
				//     Root
				//
				//    A - B
				suite.k.SaveSection(ctx, types.DefaultSection(1))
				suite.k.SaveSection(ctx, types.NewSection(1, 1, 2, "A", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 2, 1, "B", ""))
			},
			subspaceID: 1,
			sectionID:  2,
			expResult:  false,
		},
		{
			name: "section with valid path returns true",
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

				// Create the following subspaces sections
				//     Root
				//     /   \
				//    A     B
				//    |
				//    C
				suite.k.SaveSection(ctx, types.DefaultSection(1))
				suite.k.SaveSection(ctx, types.NewSection(1, 1, types.RootSectionID, "A", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 2, types.RootSectionID, "B", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 3, 1, "C", ""))
			},
			subspaceID: 1,
			sectionID:  3,
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

			result := suite.k.IsSectionPathValid(ctx, tc.subspaceID, tc.sectionID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasSection() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		expResult  bool
	}{
		{
			name:       "non existing section returns false",
			subspaceID: 1,
			sectionID:  1,
			expResult:  false,
		},
		{
			name: "existing section returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					1,
					"Edited section",
					"This is an edited section",
				))
			},
			subspaceID: 1,
			sectionID:  1,
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

			result := suite.k.HasSection(ctx, tc.subspaceID, tc.sectionID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetSection() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		expSection types.Section
		expFound   bool
	}{
		{
			name:       "non existing section returns empty section and false",
			subspaceID: 1,
			sectionID:  1,
			expSection: types.Section{},
			expFound:   false,
		},
		{
			name: "existing section returns correct data and true",
			store: func(ctx sdk.Context) {
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					1,
					"Edited section",
					"This is an edited section",
				))
			},
			subspaceID: 1,
			sectionID:  1,
			expSection: types.NewSection(
				1,
				1,
				1,
				"Edited section",
				"This is an edited section",
			),
			expFound: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			section, found := suite.k.GetSection(ctx, tc.subspaceID, tc.sectionID)
			suite.Require().Equal(tc.expSection, section)
			suite.Require().Equal(tc.expFound, found)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteSection() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing section is deleted properly",
			subspaceID: 1,
			sectionID:  1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasSection(ctx, 1, 1))
			},
		},
		{
			name: "existing section is deleted properly along all the related data",
			store: func(ctx sdk.Context) {
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test edited section",
				))

				// User groups
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				// Permissions
				suite.k.SetUserPermissions(ctx, 1, 1, "cosmos1p7vudy57pw08w6plujlpqpuqea2hkqusfq5zjc", types.PermissionManageSections)

				// Children section
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					2,
					1,
					"Child section",
					"This is a child section",
				))
			},
			subspaceID: 1,
			sectionID:  1,
			check: func(ctx sdk.Context) {
				// Make sure the section and its children have been deleted
				sections := suite.k.GetSubspaceSections(ctx, 1)
				suite.Require().Empty(sections)

				// Make sure user groups have been deleted
				groups := suite.k.GetSectionUserGroups(ctx, 1, 1)
				suite.Require().Empty(groups)

				// Make sure all user group members have been deleted
				permissions := suite.k.GetSectionUserPermissions(ctx, 1, 1)
				suite.Require().Empty(permissions)
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

			suite.k.DeleteSection(ctx, tc.subspaceID, tc.sectionID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
