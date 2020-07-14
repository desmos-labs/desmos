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
			name:        "Invariants not violated",
			profile:     types.NewProfile("dtag", suite.testData.postOwner, suite.testData.profile.CreationDate),
			expResponse: "Every invariant condition is fulfilled correctly",
			expBool:     true,
		},
		{
			name:        "ValidProfile invariant violated",
			profile:     types.NewProfile("", suite.testData.postOwner, suite.testData.profile.CreationDate),
			expResponse: "profiles: invalid profiles invariant\nThe following list contains invalid profiles:\n Invalid profiles:\n[DTag]: , [Creator]: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47\n\n",
			expBool:     true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() //reset

			err := suite.keeper.SaveProfile(suite.ctx, test.profile)
			suite.NoError(err)

			res, stop := keeper.AllInvariants(suite.keeper)(suite.ctx)

			suite.Equal(test.expResponse, res)
			suite.Equal(test.expBool, stop)
		})
	}

}
