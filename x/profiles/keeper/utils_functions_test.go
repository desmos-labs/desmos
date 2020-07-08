package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_IterateProfile() {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	creator2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)

	creator3, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)

	creator4, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	suite.NoError(err)

	timeZone, err := time.LoadLocation("UTC")
	suite.NoError(err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profiles := types.Profiles{
		types.NewProfile("first", creator, date),
		types.NewProfile("second", creator2, date),
		types.NewProfile("not", creator3, date),
		types.NewProfile("third", creator4, date),
	}

	expProfiles := types.Profiles{
		profiles[0],
		profiles[1],
		profiles[3],
	}

	for _, profile := range profiles {
		err := suite.keeper.SaveProfile(suite.ctx, profile)
		suite.NoError(err)
	}

	var validProfiles types.Profiles
	suite.keeper.IterateProfiles(suite.ctx, func(_ int64, profile types.Profile) (stop bool) {
		if profile.DTag == "not" {
			return false
		}
		validProfiles = append(validProfiles, profile)
		return false
	})

	suite.Len(expProfiles, len(validProfiles))

	for _, profile := range validProfiles {
		suite.Contains(expProfiles, profile)
	}

}
