package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/testutil"
	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveRelationship() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		relationship types.Relationship
		shouldErr    bool
		check        func(ctx sdk.Context)
	}{
		{
			name: "existent relationship returns error",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Counterparty)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				0,
			),
			shouldErr: true,
		},
		{
			name: "relationship added correctly",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))
			},
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				0,
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				expected := []types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
						0,
					),
				}
				suite.Require().Equal(expected, suite.k.GetAllRelationships(ctx))
			},
		},
		{
			name: "relationship added correctly (different subspace)",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Counterparty)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				1,
			),
			shouldErr: false,
		},
		{
			name: "relationship added correctly (different receiver)",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Counterparty)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
				1,
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

			err := suite.k.SaveRelationship(ctx, tc.relationship)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllRelationships() {
	testCases := []struct {
		name             string
		store            func(ctx sdk.Context)
		expRelationships []types.Relationship
	}{
		{
			name: "non empty relationships slice is returned properly",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Counterparty)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))

				relationship = types.NewRelationship(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Counterparty)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					0,
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				),
			},
		},
		{
			name:             "empty relationships slice is returned properly",
			expRelationships: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			relationships := suite.k.GetAllRelationships(ctx)
			suite.Require().Equal(tc.expRelationships, relationships)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	testCases := []struct {
		name             string
		store            func(ctx sdk.Context)
		user             string
		expRelationships []types.Relationship
	}{
		{
			name: "non empty relationships slice is returned properly",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Counterparty)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			user: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				),
			},
		},
		{
			name:             "empty relationships slice is returned properly",
			expRelationships: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			relationships := suite.k.GetUserRelationships(ctx, tc.user)
			suite.Require().Equal(tc.expRelationships, relationships)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveRelationship() {
	testCases := []struct {
		name                 string
		store                func(ctx sdk.Context)
		relationshipToDelete types.Relationship
		check                func(ctx sdk.Context)
	}{
		{
			name: "deleting an existing relationship does not error",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Counterparty)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))

				relationship = types.NewRelationship(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Counterparty)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			relationshipToDelete: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				0,
			),
			check: func(ctx sdk.Context) {
				expected := []types.Relationship{
					types.NewRelationship(
						"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
						"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
						0,
					),
				}
				suite.Require().Equal(expected, suite.k.GetAllRelationships(ctx))
			},
		},
		{
			name: "deleting a non existing relationship returns an error",
			relationshipToDelete: types.NewRelationship(
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

			suite.k.RemoveRelationship(ctx, tc.relationshipToDelete)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
