package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/testutil"
	"github.com/desmos-labs/desmos/x/profiles/types"
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
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"subspace",
			),
			shouldErr: true,
		},
		{
			name: "relationship added correctly",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))
			},
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"subspace",
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				expected := []types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
						"subspace",
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
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"subspace_2",
			),
			shouldErr: false,
		},
		{
			name: "relationship added correctly (different receiver)",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			relationship: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
				"subspace_2",
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
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))

				relationship = types.NewRelationship(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"subspace",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"subspace",
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
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			user: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"subspace",
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
		shouldErr            bool
		check                func(ctx sdk.Context)
	}{
		{
			name: "deleting an existing relationship does not error",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))

				relationship = types.NewRelationship(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					"subspace",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			relationshipToDelete: types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				"subspace",
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				expected := []types.Relationship{
					types.NewRelationship(
						"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
						"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
						"subspace",
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
				"subspace",
			),
			shouldErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.RemoveRelationship(ctx, tc.relationshipToDelete)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteSubspaceUserRelationships() {
	ctx, _ := suite.ctx.CacheContext()

	// Init relationships
	relationships := []types.Relationship{
		types.NewRelationship(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewRelationship(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewRelationship(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}
	suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
	for _, rel := range relationships {
		suite.Require().NoError(suite.k.SaveRelationship(ctx, rel))
	}

	suite.k.DeleteSubspaceUserRelationships(
		ctx,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	)

	// Check the result
	suite.Require().Equal(
		[]types.Relationship{
			types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
				"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
		}, suite.k.GetAllRelationships(ctx))
}
