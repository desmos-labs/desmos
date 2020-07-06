package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/internal/keeper"
	"github.com/desmos-labs/desmos/x/profiles/internal/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	timeZone, err := time.LoadLocation("UTC")
	suite.NoError(err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	tests := []struct {
		name        string
		profile     types.Profile
		expResponse string
		expBool     bool
	}{
		{
			name:        "Invariants not violated",
			profile:     types.NewProfile("dtag", owner, date),
			expResponse: "Every invariant condition is fulfilled correctly",
			expBool:     true,
		},
		{
			name:        "ValidProfile invariant violated",
			profile:     types.NewProfile("", owner, date),
			expResponse: "profiles: invalid profiles invariant\nThe following list contains invalid profiles:\n Invalid profiles:\n[DTag]: , [Creator]: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\n\n",
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
