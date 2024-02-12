package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveRelationship() {
	testCases := []struct {
		name         string
		relationship types.Relationship
		check        func(ctx sdk.Context)
	}{
		{
			name: "relationship saved correctly",
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				0,
			),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasRelationship(
					ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			suite.k.SaveRelationship(ctx, tc.relationship)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasRelationship() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		creator      string
		counterparty string
		subspace     uint64
		expResult    bool
	}{
		{
			name:         "non existing relationship returns false",
			creator:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			counterparty: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace:     1,
			expResult:    false,
		},
		{
			name: "existing relationship returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					1,
				))
			},
			creator:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			counterparty: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspace:     1,
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

			result := suite.k.HasRelationship(ctx, tc.creator, tc.counterparty, tc.subspace)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetRelationship() {
	testCases := []struct {
		name            string
		store           func(ctx sdk.Context)
		user            string
		counterparty    string
		subspaceID      uint64
		expFound        bool
		expRelationship types.Relationship
	}{
		{
			name:         "non existing relationship returns false",
			user:         "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			counterparty: "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			subspaceID:   0,
			expFound:     false,
		},
		{
			name: "existing relationship is returned correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				))
			},
			user:         "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			counterparty: "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			subspaceID:   0,
			expFound:     true,
			expRelationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				0,
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

			relationship, found := suite.k.GetRelationship(ctx, tc.user, tc.counterparty, tc.subspaceID)
			suite.Require().Equal(tc.expFound, found)
			if tc.expFound {
				suite.Require().Equal(tc.expRelationship, relationship)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRelationship() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		user         string
		counterparty string
		subspaceID   uint64
		check        func(ctx sdk.Context)
	}{
		{
			name: "deleting a relationship works properly",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				)
				suite.k.SaveRelationship(ctx, relationship)

				relationship = types.NewRelationship(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					0,
				)
				suite.k.SaveRelationship(ctx, relationship)
			},
			user:         "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			counterparty: "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			subspaceID:   0,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasRelationship(ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				))
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

			suite.k.DeleteRelationship(ctx, tc.user, tc.counterparty, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
