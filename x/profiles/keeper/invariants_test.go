package keeper_test

import (
	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	tests := []struct {
		name        string
		profile     types.Profile
		expResponse string
		expBool     bool
	}{
		{
			name: "Invariants not violated",
			profile: types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures("", ""),
				suite.testData.profile.CreationDate,
				suite.testData.user,
			),
			expResponse: "Every invariant condition is fulfilled correctly",
			expBool:     true,
		},
		{
			name: "ValidProfile invariant violated",
			profile: types.NewProfile(
				"",
				"",
				"",
				types.NewPictures("", ""),
				suite.testData.profile.CreationDate,
				suite.testData.user,
			),
			expResponse: "profiles: invalid profiles invariant\nThe following list contains invalid profiles:\n Invalid profiles:\n[DTag]: , [Creator]: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\n\n",
			expBool:     true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() //reset

			err := suite.k.StoreProfile(suite.ctx, test.profile)
			suite.Require().NoError(err)

			res, stop := keeper.AllInvariants(suite.k)(suite.ctx)

			suite.Require().Equal(test.expResponse, res)
			suite.Require().Equal(test.expBool, stop)
		})
	}
}
